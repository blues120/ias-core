package signature

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	kratosErrors "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/data/iam"
)

var (
	AuthFieldsAbsentError   = kratosErrors.New(410005, "", "鉴权字段缺失，需要 x-ctyunai-now、x-ctyunai-ak、x-ctyunai-signature、x-ctyunai-org (除组织架构接口)")
	AppIdNotFoundError      = kratosErrors.New(410006, "", "找不到 appId 信息或 appId 已禁用")
	UserOrgNotMatchError    = kratosErrors.New(410007, "", "用户和组织架构不匹配")
	UserNotAssignedOrgError = kratosErrors.New(410008, "", "用户未分配组织架构")
	UserOrgConvertError     = kratosErrors.New(410009, "", "x-ctyunai-org 格式转换错误")
	VerifySignatureError    = kratosErrors.New(410010, "", "验签错误，请检查加签算法")
	InvalidRequestError     = kratosErrors.New(410011, "", "鉴权失败，无效的请求")
	TimestampInvalidError   = kratosErrors.New(410012, "", "无效的时间戳")
)

const (
	TimestampDiffSeconds = 10
)

// 以下接口不检查组织架构字段，还未获取组织架构
var orgIdWhiteListAPIs = map[string]struct{}{
	"/api.organization.v1.OrganizationService/GetUserOrganization": struct{}{}, // 获取用户组织架构列表
}

// 是否跳过组织架构信息检查
func shouldSkipOrgIdCheck(op string) bool {
	if _, ok := orgIdWhiteListAPIs[op]; ok {
		return true
	}
	return false
}

// 校验用户是否分配了该组织架构
func checkUserOrgId(ctx context.Context, orgId string, orgUc *biz.OrganizationUsecase, uinfo *iam.UserInfo, helper *log.Helper) error {
	orgs, err := orgUc.GetUserOrgListById(ctx, uinfo)
	if err != nil {
		return UserNotAssignedOrgError
	}

	for i := range orgs {
		if fmt.Sprintf("%d", orgs[i]) == orgId {
			return nil
		}
	}

	return UserOrgNotMatchError
}

// 校验用户组织架构 id
func verifyOrgId(ctx context.Context, orgId string, orgUc *biz.OrganizationUsecase, uinfo *iam.UserInfo, helper *log.Helper) (uint32, []uint32, error) {
	if err := checkUserOrgId(ctx, orgId, orgUc, uinfo, helper); err != nil {
		return 0, nil, err
	}
	orgIdInt, err := strconv.Atoi(orgId)
	if err != nil {
		helper.Errorf("failed to convert org id: %s", err)
		return 0, nil, UserOrgConvertError
	}

	accessOrgList, err := orgUc.FindAccessOrgListById(context.Background(), uint32(orgIdInt))
	if err != nil {
		helper.Errorf("failed to get user access org list: %s", err)
		return 0, nil, err
	}

	return uint32(orgIdInt), accessOrgList, nil
}

func toStrArr(arr []uint32) []string {
	var ret []string
	for i := range arr {
		ret = append(ret, fmt.Sprintf("%d", arr[i]))
	}
	return ret
}

// 对请求参数 encode 并按 key 排序
func prepareEncodedUri(u *url.URL) string {
	return fmt.Sprintf("%s?%s", u.Path, u.Query().Encode())
}

func abs(s int64) int64 {
	if s < 0 {
		return -s
	}
	return s
}

// 校验时间戳，和服务器时间相差不能超过 10s
func isTimestampValid(ts string) error {
	tsInt, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return err
	}
	if abs(time.Now().UnixMilli()-tsInt) > TimestampDiffSeconds*1000 {
		return TimestampInvalidError
	}
	return nil
}

func DoSignatureV2Auth(ctx context.Context, req interface{}, client *iam.Client, orgUc *biz.OrganizationUsecase, nextHandler middleware.Handler, helper *log.Helper, opts ...Option) (reply interface{}, err error) {
	options := NewOptions()
	for _, o := range opts {
		o.Apply(&options)
	}

	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return nil, InvalidRequestError
	}

	// 是否跳过组织架构信息校验
	skipOrgCheck := shouldSkipOrgIdCheck(tr.Operation())

	ht, ok := tr.(*http.Transport)
	if !ok {
		return nil, InvalidRequestError
	}

	r := ht.Request()

	// 从请求头中获取签名信息
	timestamp := r.Header.Get(options.H.now)
	appId := r.Header.Get(options.H.app)
	signature := r.Header.Get(options.H.signature)
	orgIdStr := r.Header.Get(options.H.orgId)
	if timestamp == "" || appId == "" || signature == "" || (!skipOrgCheck && orgIdStr == "") {
		helper.Debugf("verifySignature params, timestamp: %s, appId: %s, signature: %s, org: %s", timestamp, appId, signature, orgIdStr)
		return nil, AuthFieldsAbsentError
	}

	// 校验时间戳
	if err := isTimestampValid(timestamp); err != nil {
		helper.Errorf("invalid timestap: %s", timestamp)
		return nil, TimestampInvalidError
	}

	// 获取 sk 信息
	sk, err := client.GetUserSecretKey(appId)
	if err != nil || !sk.Valid() {
		helper.Errorf("failed to get secret key for %s: %s", appId, err)
		return nil, AppIdNotFoundError
	}

	// 获取用户信息
	uinfo := &iam.UserInfo{
		TenantID: sk.WorkspaceId,
		UserID:   sk.ClientId,
	}

	// 验签
	requestUri := prepareEncodedUri(r.URL)
	if err := verifySignature(requestUri, timestamp, appId, sk.SecretKey, signature); err != nil {
		helper.Debugf("verifySignature failed: %s, uri: %s, timestamp: %s, appId: %s, appSecret: %s, signature: %s", err, requestUri, timestamp, appId, sk.SecretKey, signature)
		return nil, VerifySignatureError
	}

	// 校验组织架构信息
	if !skipOrgCheck {
		orgId, accessOrgList, err := verifyOrgId(ctx, orgIdStr, orgUc, uinfo, helper)
		if err != nil {
			helper.Errorf("failed to get org id: %s", err)
			return nil, err
		}

		// 组织架构信息写入上下文
		ctx = iam.SetOrgInfo(ctx, &iam.OrgInfo{
			CurrentOrg:    fmt.Sprintf("#%d#", orgId),
			AccessOrgList: fmt.Sprintf("#%s#", strings.Join(toStrArr(accessOrgList), "#")), // e.g. "#1#2#"
		})
	}

	// 用户信息写入上下文
	ctx = iam.SetUserInfo(ctx, uinfo)

	return nextHandler(ctx, req)
}

/*
OpenAPI 验签中间件，用于第三方系统调用鉴权（V1 版本用于盒子调用鉴权）
1. ak,sk 通过 IAM 个人信息页申请
2. 可指定组织架构信息，实现数据隔离
*/
func NewSignatureMiddlewareV2(client *iam.Client, orgUc *biz.OrganizationUsecase, logger log.Logger, opts ...Option) middleware.Middleware {
	helper := log.NewHelper(logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			return DoSignatureV2Auth(ctx, req, client, orgUc, handler, helper, opts...)
		}
	}
}
