package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/blues120/ias-core/data/ent/mixin"
)

// Signature holds the schema definition for the Signature entity.
type Signature struct {
	ent.Schema
}

// Fields of the Signature.
func (Signature) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("box_id").Comment("对应device表ext_id"),
		field.String("app_id").Comment("app_id"),
		field.String("app_secret").Comment("app秘钥"),
	}
}

// Annotations of the Signature
func (Signature) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "signature"},
	}
}

// Mixin of the Signature
func (Signature) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},       // 添加创建时间、更新时间字段
		mixin.SoftDelete{}, // 软删除
		//mixin.TenantMixin{},       // 添加租户
		//mixin.OrganizationMixin{}, // 添加组织架构
	}
}

// Edges of the Signature.
func (Signature) Edges() []ent.Edge {
	return nil
}
