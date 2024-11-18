package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/data/iam"
	"github.com/blues120/ias-core/middleware/signature"
)

type AuthMiddleware middleware.Middleware

// NewAuthMiddleware 统一鉴权中间件
func NewAuthMiddleware(uc *biz.OrganizationUsecase, iamCli *iam.Client, logger log.Logger, opts ...signature.Option) AuthMiddleware {
	helper := log.NewHelper(logger)
	options := signature.NewOptions()
	for _, o := range opts {
		o.Apply(&options)
	}

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			request, ok := http.RequestFromServerContext(ctx)
			if !ok {
				helper.Error("failed to get request from ctx")
				return nil, iam.UnauthorizedError
			}

			if isRequestFromOpenAPI(request, options.H.GetSignatureKey()) {
				return signature.DoSignatureV2Auth(ctx, req, iamCli, uc, handler, helper, opts...)
			}

			return DoIAMAuth(ctx, req, iamCli, handler, helper)
		}
	}
}

// isRequestFromOpenAPI 判断是不是openapi的请求
func isRequestFromOpenAPI(request *http.Request, header string) bool {
	return request.Header.Get(header) != ""
}
