package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/mixin"
)

type WarnSmsErr struct {
	ent.Schema
}

func (WarnSmsErr) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("app_name").Comment("应用名称"),
		field.Uint64("record_id").Comment("告警记录ID"),
		field.String("error_msg").Comment("短信发送错误信息"),
	}
}
func (WarnSmsErr) Edges() []ent.Edge {
	return nil
}
func (WarnSmsErr) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "warn_sms_err"}}
}

func (WarnSmsErr) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},              // 添加创建时间、更新时间字段
		mixin.TenantMixin{},       // 添加租户
		mixin.OrganizationMixin{}, // 添加组织架构
	}
}
