package iam

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// orgInfoKey 组织架构信息key
type orgInfoKey struct{}

// OrgInfo 组织架构信息
type OrgInfo struct {
	OrgList       string // 用户所在的层级列表，多个用逗号分隔，如 "2,5"
	CurrentOrg    string // 用户当前所在层级，如 "#2#"
	AccessOrgList string // 能访问用户资源的所有层级，从当前所在层级回溯到父层级获得，使用 # 分隔，如 "#1#2#"
}

// SetOrgInfo 设置组织架构信息
func SetOrgInfo(ctx context.Context, info *OrgInfo) context.Context {
	return context.WithValue(ctx, orgInfoKey{}, info)
}

// GetOrgInfo 获取组织架构信息
func GetOrgInfo(ctx context.Context) (*OrgInfo, bool) {
	u, ok := ctx.Value(orgInfoKey{}).(*OrgInfo)
	return u, ok
}

// SkipOrgInfo 跳过组织架构
func SkipOrgInfo(ctx context.Context) context.Context {
	return context.WithValue(ctx, orgInfoKey{}, true)
}

// IsSkipOrgInfo 是否跳过组织架构
func IsSkipOrgInfo(ctx context.Context) bool {
	_, ok := ctx.Value(orgInfoKey{}).(bool)
	return ok
}

// OrgID returns a OrgID valuer for logging.
func OrgID() log.Valuer {
	return func(ctx context.Context) interface{} {
		if u, ok := GetOrgInfo(ctx); ok {
			return u.CurrentOrg
		}
		return ""
	}
}