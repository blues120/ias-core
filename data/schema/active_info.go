package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/mixin"
)

type ActiveInfo struct {
	ent.Schema
}

func (ActiveInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("process_id").Comment("流程id"),
		field.String("start_time").Comment("激活开始时间"),
		field.String("result").Comment("激活结果"),
		field.String("msg").Comment("消息"),
	}
}

// Annotations of the ActiveInfo
func (ActiveInfo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "active_info"},
	}
}

// Mixin of the ActiveInfo
func (ActiveInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},       // 添加创建时间、更新时间字段
		mixin.SoftDelete{}, // 添加软删除功能
		//mixin.TenantMixin{},       // 添加租户
		//mixin.OrganizationMixin{}, // 添加组织架构
	}
}

// Edges of the ActiveInfo.
func (ActiveInfo) Edges() []ent.Edge {
	return nil
}
