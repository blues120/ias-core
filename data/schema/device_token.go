package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/mixin"
)

// DeviceToken holds the schema definition for the Task entity.
type DeviceToken struct {
	ent.Schema
}

// Fields of the Task.
func (DeviceToken) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("token").Unique().Comment("设备接入码/激活码"),
		field.String("device_ext_id").Comment("设备ext id"),
	}
}

// Annotations of the DeviceToken
func (DeviceToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "device_token"},
	}
}

// Mixin of the DeviceToken
func (DeviceToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},       // 添加创建时间、更新时间字段
		mixin.SoftDelete{}, // 添加软删除功能
		//mixin.TenantMixin{},       // 添加租户
		//mixin.OrganizationMixin{}, // 添加组织架构
	}
}
