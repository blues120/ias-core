package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/blues120/ias-core/data/ent/mixin"
)

type EquipAttr struct {
	ent.Schema
}

func (EquipAttr) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("attr_key").Comment("属性key"),
		field.String("attr_value").Comment("属性值"),
		field.String("extend").Comment("扩展字段"),
	}
}

// Annotations of the EquipAttr
func (EquipAttr) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "equip_attr"},
	}
}

// Mixin of the EquipAttr
func (EquipAttr) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},       // 添加创建时间、更新时间字段
		mixin.SoftDelete{}, // 添加软删除功能
		//mixin.TenantMixin{},       // 添加租户
		//mixin.OrganizationMixin{}, // 添加组织架构
	}
}

// Edges of the EquipAttr.
func (EquipAttr) Edges() []ent.Edge {
	return nil
}
