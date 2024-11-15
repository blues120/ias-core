package iam

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// userInfoKey 用户信息key
type userInfoKey struct{}

// UserInfo 用户信息
type UserInfo struct {
	TenantID string // 租户ID
	UserID   string // 用户ID
}

// SetUserInfo 设置用户信息
func SetUserInfo(ctx context.Context, info *UserInfo) context.Context {
	return context.WithValue(ctx, userInfoKey{}, info)
}

// GetUserInfo 获取用户信息
func GetUserInfo(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userInfoKey{}).(*UserInfo)
	return u, ok
}

// SkipUserInfo 跳过用户信息
func SkipUserInfo(ctx context.Context) context.Context {
	return context.WithValue(ctx, userInfoKey{}, true)
}

// IsSkipUserInfo 是否跳过用户信息
func IsSkipUserInfo(ctx context.Context) bool {
	_, ok := ctx.Value(userInfoKey{}).(bool)
	return ok
}

// UserId returns a UserId valuer for logging.
func UserID() log.Valuer {
	return func(ctx context.Context) interface{} {
		if u, ok := GetUserInfo(ctx); ok {
			return u.UserID
		}
		return ""
	}
}

// TenantId returns a TenantId valuer for logging.
func TenantID() log.Valuer {
	return func(ctx context.Context) interface{} {
		if u, ok := GetUserInfo(ctx); ok {
			return u.TenantID
		}
		return ""
	}
}