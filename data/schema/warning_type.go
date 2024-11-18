package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/blues120/ias-core/data/ent/mixin"
)

// WarningType holds the schema definition for the Task entity.
type WarningType struct {
	ent.Schema
}

// Fields of the WarningType.
func (WarningType) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("alarm_type").Comment("告警I级类型"),
		field.String("alarm_name").Comment("告警II级类型"),
	}
}

// Indexes of the WarningType.
func (WarningType) Indexes() []ent.Index {
	return []ent.Index{
		// 唯一约束索引
		index.Fields("alarm_type", "alarm_name").
			Unique(),
	}
}

// Annotations of the WarningType
func (WarningType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "warning_type"},
	}
}

// Mixin of the WarningType
func (WarningType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},       // 添加创建时间、更新时间字段
		mixin.SoftDelete{}, // 添加软删除功能
		//mixin.TenantMixin{},       // 添加租户
		//mixin.OrganizationMixin{}, // 添加组织架构
	}
}
