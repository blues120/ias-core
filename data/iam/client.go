package iam

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	goHttp "net/http"
	"regexp"
	"strings"

	kratosErrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"

	"github.com/blues120/ias-core/conf"
	"github.com/blues120/ias-core/pkg/openapi"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	sessionLoginKey   = "$default#$login"
	sessionUserIdKey  = "$user#userId"
	queryParameterWID = "workspaceId"

	highestPrivilegeAction = "c_.*" // 所有权限，需要在 iam 服务定义中配置该动作 (re2)c_.*

	orgRoleNote = "组织架构" // 组织架构类角色特有标记

	iamSessionCookiekey = "iam_tgc"
)

var (
	SessionNotFoundError  = kratosErrors.New(410001, "", "登录状态已过期, 请重新登录")
	SessionNotLoginError  = kratosErrors.New(410002, "", "登录状态已过期, 请重新登录")
	UnauthorizedError     = kratosErrors.New(410003, "", "用户未授权")
	OrgRolesNotFoundError = kratosErrors.New(410004, "", "用户未配置组织架构角色，请联系工作区管理员分配")
)

func newRedisClient(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})

	// Enable tracing instrumentation.
	if err := redisotel.InstrumentTracing(rdb); err != nil {
		panic(err)
	}

	// Enable metrics instrumentation.
	if err := redisotel.InstrumentMetrics(rdb); err != nil {
		panic(err)
	}

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return rdb
}

type Client struct {
	conf *conf.IAM
	api  *openapi.OpenClient
	rdb  *redis.Client
	log  *log.Helper
}

func NewClient(c *conf.IAM, logger log.Logger) *Client {
	helper := log.NewHelper(logger)
	api := openapi.NewOpenClient(c.IamHost, c.AppId, c.AppSecret, c.Ac, "", helper)
	var rdb *redis.Client
	if c.Enable && c.Redis != nil {
		rdb = newRedisClient(&conf.Data{Redis: c.Redis}) // 注意：使用 data.NewRedisClient 会循环引用
	}
	return &Client{conf: c, api: api, rdb: rdb, log: helper}
}

// 从 http request 中获取用户信息
func (c *Client) GetUserInfo(r *http.Request) (*UserInfo, error) {
	session, err := r.Cookie(iamSessionCookiekey)
	if err != nil || session.Value == "" {
		return nil, SessionNotFoundError
	}
	kvs, err := c.rdb.HGetAll(context.Background(), getSessionRedisKey(session.Value)).Result()
	if err != nil {
		return nil, SessionNotFoundError
	}
	if login, ok := kvs[sessionLoginKey]; !ok || login != "true" {
		return nil, SessionNotLoginError
	}
	uid := kvs[sessionUserIdKey]
	wid := r.URL.Query().Get(queryParameterWID)
	return &UserInfo{
		TenantID: wid,
		UserID:   uid,
	}, nil
}

type checkPermissionResp struct {
	Data struct {
		Enable string `json:"enable,omitempty"`
		At     string `json:"at,omitempty"`
		On     string `json:"on,omitempty"`
		Do     string `json:"do,omitempty"`
		Who    string `json:"who,omitempty"`
	} `json:"data,omitempty"`
}

// 检查用户接口权限，目前仅过滤工作区权限
func (c *Client) CheckPermission(info *UserInfo, opts ...openapi.Option) error {
	query := fmt.Sprintf("/iam/gw/v1/privilege/CheckSingle?workspaceId=%s&who=%s&on=%s&do=%s", info.TenantID, info.UserID, info.TenantID, c.conf.PrivilegeAction)
	resp, err := c.api.Get(query, opts...)
	if err != nil {
		c.log.Errorf("[IamClient] failed to checkPrivilege, err: %s", err)
		return err
	}
	var permission checkPermissionResp
	err = json.Unmarshal([]byte(resp), &permission)
	if err != nil {
		c.log.Errorf("[IamClient] failed to unmarshal, err: %s", err)
		return err
	}
	if permission.Data.Enable != "true" {
		c.log.Warnf("[IamClient] no privilege, resp: %s", resp)
		return UnauthorizedError
	}
	return nil
}

// 使用 iam 自身设置的 session，以 h:sess 开头；之前 ecx 服务封装的 o:session$ 有问题，不再使用
func getSessionRedisKey(session string) string {
	return fmt.Sprintf("h:sess$%s", session)
}

type MenuItem struct {
	Renderer     string      `json:"renderer,omitempty"`
	Plist        string      `json:"plist,omitempty"`
	Enable       string      `json:"enable,omitempty"`
	Name         string      `json:"name,omitempty"`
	Icon         string      `json:"icon,omitempty"`
	DisplayOrder string      `json:"displayOrder,omitempty"`
	Id           string      `json:"id,omitempty"`
	State        string      `json:"state,omitempty"` // 是否上线，在 iam 配置，如果子菜单上线，则认为父菜单也一定上线。上线不等于有权限
	Ucode        string      `json:"ucode,omitempty"`
	Items        []*MenuItem `json:"items,omitempty"`
	ParentId     string      `json:"parentId,omitempty"`
}

type getMenusResp struct {
	Duration string `json:"duration,omitempty"`
	Reason   string `json:"reason,omitempty"`
	Code     string `json:"code,omitempty"`
	Serial   string `json:"serial,omitempty"`
	Host     string `json:"host,omitempty"`
	Data     struct {
		Domain string      `json:"domain,omitempty"`
		Name   string      `json:"name,omitempty"`
		Items  []*MenuItem `json:"items,omitempty"`
	} `json:"data,omitempty"`
}

// 查询某个工作区的菜单列表
func (c *Client) GetMenus(info *UserInfo, domain string, depth int, opts ...openapi.Option) ([]*MenuItem, error) {
	if depth == 0 {
		depth = 5 // depth 先写 5，暂时用不到这么多层级
	}
	if domain == "" {
		domain = "ias.main"
	}
	query := fmt.Sprintf("/iam/api/v1/workspace/menu/GetTree?workspaceId=%s&userId=%s&domain=%s&depth=%d&parent=0", info.TenantID, info.UserID, domain, depth)
	resp, err := c.api.Get(query, opts...)
	if err != nil {
		c.log.Errorf("[IamClient] failed to get menus, err: %s", err)
		return nil, err
	}
	var menusResp getMenusResp
	err = json.Unmarshal([]byte(resp), &menusResp)
	if err != nil {
		c.log.Errorf("[IamClient] failed to unmarshal, err: %s", err)
		return nil, err
	}

	return menusResp.Data.Items, nil
}

type Policy struct {
	Mode     string `json:"mode,omitempty"`
	Resource string `json:"resource,omitempty"`
	Action   string `json:"action,omitempty"`
	Id       string `json:"id,omitempty"`
}

type Role struct {
	Note       string    `json:"note,omitempty"`
	At         string    `json:"at,omitempty"`
	AllowFirst bool      `json:"allowFirst,omitempty"`
	Name       string    `json:"name,omitempty"`
	Id         string    `json:"id,omitempty"`
	Policies   []*Policy `json:"policies,omitempty"`
}

type getPrivilegesResp struct {
	Duration string `json:"duration,omitempty"`
	Reason   string `json:"reason,omitempty"`
	Code     string `json:"code,omitempty"`
	Serial   string `json:"serial,omitempty"`
	Host     string `json:"host,omitempty"`
	Data     struct {
		Roles []*Role `json:"roles,omitempty"`
	} `json:"data,omitempty"`
}

// 查询某个用户的角色列表
func (c *Client) GetPrivileges(info *UserInfo, opts ...openapi.Option) ([]*Role, error) {
	query := fmt.Sprintf("/iam/gw/v1/privilege/List?workspaceId=%s&who=%s", info.TenantID, info.UserID)
	resp, err := c.api.Get(query, opts...)
	if err != nil {
		c.log.Errorf("[IamClient] failed to get privileges, err: %s", err)
		return nil, err
	}
	var privilegesResp getPrivilegesResp
	err = json.Unmarshal([]byte(resp), &privilegesResp)
	if err != nil {
		c.log.Errorf("[IamClient] failed to unmarshal, err: %s", err)
		return nil, err
	}
	return privilegesResp.Data.Roles, nil
}

// 访问资源的权限，原始格式如 w_test_view@enterprise|user|subuser|iam|osp@0
type Plist struct {
	PolicyAction string   // 用户看到资源必须具备的权限动作
	Users        []string // 能够看到资源的用户类型，多个用|分隔
	ServiceId    string   // 服务 id，0 等同于当前服务 id
}

func parseMenuPlist(plist string) (*Plist, error) {
	parts := strings.Split(plist, "@")
	if len(parts) != 3 {
		return nil, errors.New("invalid plist")
	}
	users := strings.Split(parts[1], "|")
	return &Plist{
		PolicyAction: parts[0],
		Users:        users,
		ServiceId:    parts[2],
	}, nil
}

func hasHighestPrivilege(actionMap map[string]struct{}) bool {
	if _, ok := actionMap[highestPrivilegeAction]; ok {
		return true
	}
	return false
}

func isReMatched(reMatchActionMap map[string]struct{}, action string) bool {
	for pattern := range reMatchActionMap {
		if matched, err := regexp.MatchString(pattern, action); err == nil && matched {
			return true
		}
	}
	return false
}

/*
递归查找菜单
1. 如果父菜单和子菜单都配置了，则父菜单及配置的子菜单可见
2. 如果父菜单配置了，子菜单没配置，则父菜单和它的全部子菜单都可见
3. 如果父菜单没配置，子菜单配置了，则父菜单及配置的子菜单可见
*/
func (c *Client) getEnableMenus(menus []*MenuItem, exactMatchActionMap map[string]struct{}, reMatchActionMap map[string]struct{}) []*MenuItem {
	var items []*MenuItem
	for i := range menus {
		if menus[i].State != "online" && len(menus[i].Items) == 0 { // 子菜单上线，则父菜单也上线，无论父菜单 state 是什么
			continue
		}
		plist, err := parseMenuPlist(menus[i].Plist)
		if err != nil {
			c.log.Errorf("failed to parse menu plist %s: %s", menus[i].Plist, err)
			continue
		}
		var currentMenu *MenuItem
		if _, ok := exactMatchActionMap[plist.PolicyAction]; ok || isReMatched(reMatchActionMap, plist.PolicyAction) || hasHighestPrivilege(reMatchActionMap) {
			currentMenu = &MenuItem{
				Name:     menus[i].Name,
				Id:       menus[i].Id,
				Ucode:    menus[i].Ucode,
				ParentId: menus[i].ParentId,
			}
		}
		childrenMenus := c.getEnableMenus(menus[i].Items, exactMatchActionMap, reMatchActionMap)
		if len(childrenMenus) > 0 {
			if currentMenu == nil {
				currentMenu = &MenuItem{
					Name:     menus[i].Name,
					Id:       menus[i].Id,
					Ucode:    menus[i].Ucode,
					ParentId: menus[i].ParentId,
					Items:    childrenMenus,
				}
			} else {
				currentMenu.Items = childrenMenus
			}
		} else {
			if currentMenu != nil {
				for j := range menus[i].Items {
					if menus[i].Items[j].State != "offline" {
						currentMenu.Items = append(currentMenu.Items, &MenuItem{
							Name:     menus[i].Items[j].Name,
							Id:       menus[i].Items[j].Id,
							Ucode:    menus[i].Items[j].Ucode,
							ParentId: menus[i].Items[j].ParentId,
						})
					}
				}
			}
		}
		if currentMenu != nil {
			items = append(items, currentMenu)
		}
	}
	return items
}

// 查询某个用户授权的菜单列表
func (c *Client) GetUserEnabledMenus(info *UserInfo, domain string, depth int, opts ...openapi.Option) ([]*MenuItem, error) {
	menus, err := c.GetMenus(info, domain, depth, opts...)
	if err != nil {
		c.log.Errorf("failed to get menus: %s", err)
		return nil, err
	}

	roles, err := c.GetPrivileges(info, opts...)
	if err != nil {
		c.log.Errorf("failed to get roles: %s", err)
		return nil, err
	}

	var exactMatchActionMap = make(map[string]struct{}) // 精确匹配的 action，格式如 "(exact)c_view"
	var reMatchActionMap = make(map[string]struct{})    // 正则匹配的 action，格式如 "(re2)c_*"
	for i := range roles {
		for j := range roles[i].Policies {
			if strings.HasPrefix(roles[i].Policies[j].Action, "(exact)") {
				exactMatchActionMap[strings.TrimPrefix(roles[i].Policies[j].Action, "(exact)")] = struct{}{}
			} else if strings.HasPrefix(roles[i].Policies[j].Action, "(re2)") {
				reMatchActionMap[strings.TrimPrefix(roles[i].Policies[j].Action, "(re2)")] = struct{}{}
			} else { // 默认当做精确匹配
				exactMatchActionMap[strings.TrimPrefix(roles[i].Policies[j].Action, "(exact)")] = struct{}{}
			}
		}
	}

	return c.getEnableMenus(menus, exactMatchActionMap, reMatchActionMap), nil
}

// 获取用户组织架构角色名列表
func (c *Client) GetOrgRoles(info *UserInfo) ([]string, error) {
	roles, err := c.GetPrivileges(info)
	if err != nil {
		c.log.Errorf("failed to get privileges: %s", err)
		return nil, err
	}
	var orgRoles []string
	for i := range roles {
		if roles[i].Note == orgRoleNote {
			parts := strings.Split(roles[i].Name, "-") // e.g. "中石油省公司-管理员"
			if len(parts) > 0 {
				orgRoles = append(orgRoles, parts[0])
			}
		}
	}
	return orgRoles, nil
}

type UserCurrentAuthInfoResp struct {
	Duration string `json:"duration,omitempty"`
	Reason   string `json:"reason,omitempty"`
	Code     string `json:"code,omitempty"`
	Serial   string `json:"serial,omitempty"`
	Host     string `json:"host,omitempty"`
	Data     struct {
		IsLoggedIn bool   `json:"isLoggedIn,omitempty"`
		Id         string `json:"id,omitempty"`
		Property   struct {
			LoginId     string `json:"loginId,omitempty"`
			LoginTime   string `json:"loginTime,omitempty"`
			Name        string `json:"name,omitempty"`
			Mobile      string `json:"mobile,omitempty"`
			AvatarPath  string `json:"avatarPath,omitempty"`
			UserType    string `json:"userType,omitempty"`
			UserId      string `json:"userId,omitempty"`
			Email       string `json:"email,omitempty"`
			LastLoginIp string `json:"lastLoginIp,omitempty"`
		} `json:"property,omitempty"`
	} `json:"data,omitempty"`
}

// 获取当前用户登录信息
func (c *Client) GetUserCurrentAuthInfo(r *http.Request) (*UserCurrentAuthInfoResp, error) {
	session, err := r.Cookie(iamSessionCookiekey)
	if err != nil || session.Value == "" {
		return nil, SessionNotFoundError
	}

	query := fmt.Sprintf("%s/iam/gw/auth/Current", c.conf.IamHost)
	req, err := goHttp.NewRequest("GET", query, nil)
	if err != nil {
		c.log.Errorf("failed to new request: %s", err)
		return nil, err
	}
	req.Header.Add("Cookie", fmt.Sprintf("%s=%s", iamSessionCookiekey, session.Value))
	req.Header.Add("Referer", c.conf.IamHost)

	client := &goHttp.Client{}
	tr := &goHttp.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.Transport = tr

	res, err := client.Do(req)
	if err != nil {
		c.log.Errorf("failed to fire request: %s", err)
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		c.log.Errorf("failed to read response body: %s", err)
		return nil, err
	}

	var authInfo UserCurrentAuthInfoResp
	if err = json.Unmarshal(resBody, &authInfo); err != nil {
		c.log.Errorf("failed to unmarshal response: %s", err)
		return nil, err
	}

	return &authInfo, nil
}

type UserSecretKey struct {
	Enable      string `json:"enable,omitempty"`
	WorkspaceId string `json:"workspaceId,omitempty"`
	Ucode       string `json:"ucode,omitempty"`
	Name        string `json:"name,omitempty"`
	CtyunAcctId string `json:"ctyunAcctId,omitempty"`
	CtyunUserId string `json:"ctyunUserId,omitempty"`
	Note        string `json:"note,omitempty"`
	ClientId    string `json:"clientId,omitempty"`
	SecretKey   string `json:"secretKey,omitempty"`
	Verifier    string `json:"verifier,omitempty"`
}

func (k *UserSecretKey) Valid() bool {
	if k.Enable == "false" || k.WorkspaceId == "" || k.ClientId == "" || k.SecretKey == "" {
		return false
	}
	return true
}

type UserSecretKeyResp struct {
	Duration string         `json:"duration,omitempty"`
	Reason   string         `json:"reason,omitempty"`
	Code     string         `json:"code,omitempty"`
	Serial   string         `json:"serial,omitempty"`
	Host     string         `json:"host,omitempty"`
	Data     *UserSecretKey `json:"data,omitempty"`
}

// 获取用户级别密钥, keyId 即 access_key
func (c *Client) GetUserSecretKey(keyId string, opts ...openapi.Option) (*UserSecretKey, error) {
	query := fmt.Sprintf("/iam/gw/v1/client/Get?keyId=%s", keyId)
	resp, err := c.api.Get(query, opts...)
	if err != nil {
		c.log.Errorf("[IamClient] failed to get user secret key, err: %s", err)
		return nil, err
	}
	var userSecretKeyResp UserSecretKeyResp
	err = json.Unmarshal([]byte(resp), &userSecretKeyResp)
	if err != nil {
		c.log.Errorf("[IamClient] failed to unmarshal, err: %s", err)
		return nil, err
	}
	return userSecretKeyResp.Data, nil
}
