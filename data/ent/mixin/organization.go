package mixin

import (
	"context"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/blues120/ias-core/data/ent/hook"
	"github.com/blues120/ias-core/data/ent/intercept"
	"github.com/blues120/ias-core/data/iam"
)

// OrganizationMixin for embedding the tenant info in different schemas.
type OrganizationMixin struct {
	mixin.Schema
}

// Fields for all schemas that embed OrganizationMixin.
func (OrganizationMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("access_org_list").Optional().Comment("授权的组织 id 列表，#分隔"),
	}
}

// Hooks of the OrganizationMixin.
func (t OrganizationMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					if iam.IsSkipOrgInfo(ctx) {
						return next.Mutate(ctx, m)
					}
					info, ok := iam.GetOrgInfo(ctx)
					// 未设置组织架构时跳过
					if !ok {
						return next.Mutate(ctx, m)
					}
					if s, ok := m.(interface{ SetAccessOrgList(string) }); ok {
						s.SetAccessOrgList(info.AccessOrgList)
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
					if iam.IsSkipOrgInfo(ctx) {
						return next.Mutate(ctx, m)
					}
					info, ok := iam.GetOrgInfo(ctx)
					// 未设置组织架构时跳过
					if !ok {
						return next.Mutate(ctx, m)
					}
					mx, ok := m.(interface {
						WhereP(...func(*sql.Selector))
					})
					if !ok {
						return nil, fmt.Errorf("unexpected mutation type %T", m)
					}
					t.P(mx, info.CurrentOrg)
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
	}
}

// Interceptors of the OrganizationMixin.
func (t OrganizationMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
			if iam.IsSkipOrgInfo(ctx) {
				return nil
			}
			info, ok := iam.GetOrgInfo(ctx)
			if ok {
				t.P(q, info.CurrentOrg)
			}
			return nil
		}),
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (t OrganizationMixin) P(w interface{ WhereP(...func(*sql.Selector)) }, currentOrg string) {
	w.WhereP(
		sql.FieldContains(t.Fields()[0].Descriptor().Name, currentOrg),
	)
}

// Null adds a storage-level predicate to the queries and mutations.
func (t OrganizationMixin) Null(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull(t.Fields()[0].Descriptor().Name),
	)
}

// Empty adds a storage-level predicate to the queries and mutations.
func (t OrganizationMixin) Empty(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldEQ(t.Fields()[0].Descriptor().Name, ""),
	)
}
