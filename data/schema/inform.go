// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/mixin"
)

type Inform struct {
	ent.Schema
}

func (Inform) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Time("created_at").Optional(),
		field.Time("updated_at").Optional(),
		field.Time("deleted_at").Optional(),
		field.String("alarm_name").Comment("告警名称"),
		field.String("alarm_type").Comment("告警名称"),
		field.String("sign_name").Comment("签名"),
		field.String("notify_template").Comment("通知模板内容"),
		field.String("template_code").Comment("通知模板"),
		field.String("phone_numbers").Comment("通知号码"),
		field.String("notify_switch").Comment("通知开关"),
		field.String("task_name").Comment("任务名称"),
		field.Uint64("task_id").Comment("任务id"),
	}
}

func (Inform) Edges() []ent.Edge {
	return nil
}

func (Inform) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "inform"}}
}

// Mixin of the Inform
func (Inform) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TenantMixin{},       // 添加租户
		mixin.OrganizationMixin{}, // 添加组织架构
	}
}
