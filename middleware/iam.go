package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/iam"
)

type IAMMiddleware middleware.Middleware

// 以下接口为前端定时调用，非人为操作，不进行续期，可能有安全风险
var renewSessionWhiteListAPIs = map[string]struct{}{
	"GET:/api/v1/camera": struct{}{}, // 摄像机列表接口
	"GET:/api/v1/task":   struct{}{}, // 任务列表接口
}

func shouldSkipSessionRenew(r *http.Request) bool {
	if _, ok := renewSessionWhiteListAPIs[fmt.Sprintf("%s:%s", strings.ToUpper(r.Method), r.URL.Path)]; ok {
		return true
	}
	return false
}

// session 续期，iam session 默认有效期 30分钟，刷新页面会更新，现在增加后端接口刷新逻辑
func renewSession(cli *iam.Client, r *http.Request) {
	if shouldSkipSessionRenew(r) {
		return
	}

	// 调用 iam current 接口可实现续期，暂时忽略返回
	cli.GetUserCurrentAuthInfo(r)
}

func DoIAMAuth(ctx context.Context, req interface{}, client *iam.Client, nextHandler middleware.Handler, helper *log.Helper) (reply interface{}, err error) {
	request, ok := http.RequestFromServerContext(ctx)
	if !ok {
		helper.Debug("failed to get transport from ctx")
		return nil, iam.UnauthorizedError
	}
	// 获取用户信息
	info, err := client.GetUserInfo(request)
	if err != nil {
		helper.Warnf("failed to get user info from request: %s", err)
		return nil, err
	}
	// 检查用户权限
	if err := client.CheckPermission(info); err != nil {
		helper.Warnf("check permission failed: %s", err)
		return nil, err
	}
	// 用户信息放入上下文
	ctx = iam.SetUserInfo(ctx, info)

	helper.Infof("user %s has logged in and has access to workspace %s", info.UserID, info.TenantID)

	// 前端接口请求过来更新 session 有效期
	go renewSession(client, request)

	return nextHandler(ctx, req)

}

func NewIAMMiddleware(client *iam.Client, logger log.Logger) IAMMiddleware {
	helper := log.NewHelper(logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			return DoIAMAuth(ctx, req, client, handler, helper)
		}
	}
}
