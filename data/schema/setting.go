package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/mixin"
)

// Setting holds the schema definition for the Setting entity.
type Setting struct {
	ent.Schema
}

// Fields of the Setting.
func (Setting) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("name").Comment("设备名称"),
		field.String("serial_number").Comment("设备序列号"),
		field.String("version").Comment("系统版本号"),
		field.String("model").Comment("设备型号"),
		field.String("workspace_id").Default("").Comment("纳管的工作区id"),
	}
}

// Annotations of the Setting
func (Setting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "setting"},
	}
}

// Mixin of the Setting
func (Setting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{}, // 添加创建时间、更新时间字段
		//mixin.TenantMixin{},       // 添加租户
		//mixin.OrganizationMixin{}, // 添加组织架构
	}
}

// Edges of the Setting.
func (Setting) Edges() []ent.Edge {
	return nil
}
