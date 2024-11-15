package middleware

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/go-kratos/kratos/v2/middleware"
)

type DeviceIdKey struct{}

const DeviceIdFromLiteKey = "deviceIdFromLite"

func LiteMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

			httpReq, ok := http.RequestFromServerContext(ctx)
			if ok {
				deviceIdStr := httpReq.URL.Query().Get(DeviceIdFromLiteKey)
				deviceId, err := strconv.Atoi(deviceIdStr)
				if err == nil && deviceId != 0 {
					ctx = context.WithValue(ctx, DeviceIdKey{}, uint64(deviceId))
				}
			}

			return handler(ctx, req)
		}
	}
}
