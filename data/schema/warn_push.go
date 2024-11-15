package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/mixin"
)

type WarnPush struct {
	ent.Schema
}

func (WarnPush) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("name").Comment("推送名称"),
		field.Enum("type").GoType(biz.WarnPushType("")).Comment("推送方式"),
		field.String("url").Comment("推送地址"),
		field.String("remark").Comment("备注"),
		field.Enum("mode").GoType(biz.WarnPushMode("")).Comment("推送模式"),
		field.Enum("status").GoType(biz.WarnPushStatus("")).Comment("推送状态"),
	}
}

// Annotations of the WarnPush
func (WarnPush) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "warn_push"},
	}
}

// Mixin of the WarnPush
func (WarnPush) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},              // 添加创建时间、更新时间字段
		mixin.SoftDelete{},        // 添加软删除功能
		mixin.TenantMixin{},       // 添加租户
		mixin.OrganizationMixin{}, // 添加组织架构
	}
}

// Edges of the WarnPush.
func (WarnPush) Edges() []ent.Edge {
	return nil
}
