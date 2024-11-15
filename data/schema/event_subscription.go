package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/mixin"
)

// Event_subscription holds the schema definition for the Event_subscription entity.
type EventSubscription struct {
	ent.Schema
}

// Fields of the EventSubscription.
func (EventSubscription) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("box_id").Comment("盒子id"),
		field.String("channel_id").Comment("通道id"),
		field.String("callback").Comment("回调地址"),
		field.String("template_id").Comment("模版id"),
		field.Enum("status").GoType(biz.EventSubStatus("")).Comment("订阅状态"),
	}
}

// Annotations of the EventSubscription
func (EventSubscription) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "event_subs"},
	}
}

// Mixin of the EventSubscription
func (EventSubscription) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},       // 添加创建时间、更新时间字段
		mixin.SoftDelete{}, // 添加软删除功能
		//mixin.TenantMixin{},       // 添加租户
		//mixin.OrganizationMixin{}, // 添加组织架构
	}
}

// Edges of the EventSubscription.
func (EventSubscription) Edges() []ent.Edge {
	return nil
}
