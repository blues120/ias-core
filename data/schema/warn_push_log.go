package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/mixin"
)

type WarnPushLog struct {
	ent.Schema
}

func (WarnPushLog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.Uint64("push_id").Comment("推送id"),
		field.Text("param").Comment("推送参数"),
		field.String("remark").Comment("备注"),
		field.Enum("status").GoType(biz.WarnPushLogStatus("")).Comment("推送状态"),
	}
}

// Annotations of the WarnPush
func (WarnPushLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "warn_push_log"},
	}
}

// Mixin of the WarnPush
func (WarnPushLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},              // 添加创建时间、更新时间字段
		mixin.SoftDelete{},        // 添加软删除功能
		mixin.TenantMixin{},       // 添加租户
		mixin.OrganizationMixin{}, // 添加组织架构
	}
}

// Edges of the WarnPushLog.
func (WarnPushLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("push", WarnPush.Type).Unique().Required().Field("push_id"),
	}
}
