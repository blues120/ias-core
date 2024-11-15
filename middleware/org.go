package middleware

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/redis/go-redis/v9"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/iam"
)

func findOrgListByNames(uc *biz.OrganizationUsecase, names []string) (uint32, []string, error) {
	orgs, err := uc.List(context.Background(), &biz.OrganizationListFilter{
		Names: names,
	})
	if err != nil {
		return 0, nil, err
	}

	var ids []string
	for i := range orgs {
		ids = append(ids, fmt.Sprintf("%d", orgs[i].ID))
	}
	sort.Strings(ids)

	var defaultOrg uint32
	if len(orgs) > 0 {
		defaultOrg = orgs[0].ID
	}

	return defaultOrg, ids, nil
}

func toStrArr(arr []uint32) []string {
	var ret []string
	for i := range arr {
		ret = append(ret, fmt.Sprintf("%d", arr[i]))
	}
	return ret
}

func isUserOrgCacheExpired(dbOrgs []string, cacheOrg uint32) bool {
	cacheOrgStr := fmt.Sprintf("%d", cacheOrg)
	for i := range dbOrgs {
		if dbOrgs[i] == cacheOrgStr {
			return false
		}
	}
	return true
}

type OrgMiddleware middleware.Middleware

// 组织架构中间件，在 iam 鉴权通过后执行，即初始化时需要放到 iam 中间件之后
func NewOrgMiddleware(uc *biz.OrganizationUsecase, iamCli *iam.Client, logger log.Logger) OrgMiddleware {
	helper := log.NewHelper(logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			// 如果已经设置了 org, 则忽略下面的操作
			if _, ok := iam.GetOrgInfo(ctx); ok {
				return handler(ctx, req)
			}

			// 查询 iam 中间件上一步设置的用户信息
			uinfo, ok := iam.GetUserInfo(ctx)
			if !ok {
				helper.Error("failed to get user info from ctx")
				return nil, iam.UnauthorizedError
			}

			// 查询用户的组织架构角色名
			roles, err := iamCli.GetOrgRoles(uinfo)
			if err != nil {
				helper.Errorf("failed to get user roles: %s", err)
				return nil, err
			}

			if len(roles) > 0 {
				// 查询用户所属的组织架构
				defaultOrg, orgs, err := findOrgListByNames(uc, roles)
				if err != nil {
					helper.Errorf("failed to get user org list: %s", err)
					return nil, err
				}

				// 查询用户当前层级
				var currentOrg uint32
				currentOrg, err = uc.GetCurrentUserOrg(context.Background(), uinfo.UserID)
				if err != nil {
					if err == redis.Nil {
						if len(orgs) > 0 {
							currentOrg = defaultOrg
						}
					} else {
						helper.Errorf("failed to get user current org: %s", err)
						return nil, err
					}
				}

				// 处理缓存和数据库不一致的情况
				if isUserOrgCacheExpired(orgs, currentOrg) {
					if err := uc.DeleteUserOrgCache(ctx, uinfo.UserID); err != nil { // 清除之前可能残存的组织架构缓存
						helper.Errorf("failed to clear unused org cache for %s: %s", uinfo.UserID, err)
					}
					if len(orgs) > 0 {
						currentOrg = defaultOrg
					}
				}

				// 查询对当前层级具有权限的所有层级列表
				if currentOrg > 0 {
					accessOrgList, err := uc.FindAccessOrgListById(context.Background(), currentOrg)
					if err != nil {
						helper.Errorf("failed to get user access org list: %s", err)
						return nil, err
					}

					var orgInfo = iam.OrgInfo{
						OrgList:       strings.Join(orgs, ","),                                         // e.g. "2,5"
						CurrentOrg:    fmt.Sprintf("#%d#", currentOrg),                                 // e.g. "#2#"
						AccessOrgList: fmt.Sprintf("#%s#", strings.Join(toStrArr(accessOrgList), "#")), // e.g. "#1#2#"
					}

					// 组织架构信息写入上下文
					ctx = iam.SetOrgInfo(ctx, &orgInfo)

					helper.Debugf("user %s orgList: %s, currentOrg: %s, accessOrgList: %s", uinfo.UserID, orgInfo.OrgList, orgInfo.CurrentOrg, orgInfo.AccessOrgList)
				} else {
					helper.Debugf("user %s does not have org info", uinfo.UserID)
				}
			} else {
				helper.Debugf("user %s does not have org info", uinfo.UserID)
				if err := uc.DeleteUserOrgCache(ctx, uinfo.UserID); err != nil { // 清除之前可能残存的组织架构缓存
					helper.Errorf("failed to clear unused org cache for %s: %s", uinfo.UserID, err)
				}
			}

			return handler(ctx, req)
		}
	}
}
