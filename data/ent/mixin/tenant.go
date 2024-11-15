package mixin

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/hook"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/intercept"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/iam"
)

// TenantMixin for embedding the tenant info in different schemas.
type TenantMixin struct {
	mixin.Schema
}

// Fields for all schemas that embed TenantMixin.
func (TenantMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("tenant_id").Optional(),
	}
}

// Hooks of the TenantMixin.
func (t TenantMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					if iam.IsSkipUserInfo(ctx) {
						return next.Mutate(ctx, m)
					}
					info, ok := iam.GetUserInfo(ctx)
					// 未设置租户id时跳过
					if !ok || info.TenantID == "" {
						return next.Mutate(ctx, m)
					}
					if s, ok := m.(interface{ SetTenantID(string) }); ok {
						s.SetTenantID(info.TenantID)
					}
					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for these operations.
			ent.OpCreate,
		),
		hook.Unless(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					if iam.IsSkipUserInfo(ctx) {
						return next.Mutate(ctx, m)
					}
					info, ok := iam.GetUserInfo(ctx)
					// 未设置租户id时跳过
					if !ok || info.TenantID == "" {
						return next.Mutate(ctx, m)
					}
					mx, ok := m.(interface {
						WhereP(...func(*sql.Selector))
					})
					if !ok {
						return nil, fmt.Errorf("unexpected mutation type %T", m)
					}
					t.P(mx, info.TenantID)
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
	}
}

// Interceptors of the TenantMixin.
func (t TenantMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
			if iam.IsSkipUserInfo(ctx) {
				return nil
			}
			info, ok := iam.GetUserInfo(ctx)
			if ok && info.TenantID != "" {
				t.P(q, info.TenantID)
			}
			return nil
		}),
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (t TenantMixin) P(w interface{ WhereP(...func(*sql.Selector)) }, tid string) {
	w.WhereP(
		sql.FieldEQ(t.Fields()[0].Descriptor().Name, tid),
	)
}
