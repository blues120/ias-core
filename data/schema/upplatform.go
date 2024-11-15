package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/mixin"
)

// UpPlatform holds the schema definition for the UpPlatform entity.
type UpPlatform struct {
	ent.Schema
}

// Fields of the UpPlatform.
func (UpPlatform) Fields() []ent.Field {
	return []ent.Field{
		field.String("sip_id").Comment("SipId"),
		field.String("sip_domain").Comment("SipDomain"),
		field.String("sip_ip").Comment("SipIp"),
		field.Int32("sip_port").Comment("SipPort"),
		field.String("sip_user").Comment("SipUser"),
		field.String("sip_password").Comment("SipPassword"),
		field.String("description").Comment("Description"),
		field.Int32("heartbeat_interval").Comment("HeartbeatInterval"),
		field.Int32("register_interval").Comment("RegisterInterval"),
		field.String("trans_type").Comment("TransType"),
		field.String("gb_id").Comment("GbId"),
		field.String("cascadestatus").Comment("Cascadestatus"),
		field.String("registration_status").Comment("RegistrationStatus"),
	}
}

// Edges of the UpPlatform.
func (UpPlatform) Edges() []ent.Edge {
	return nil
}

// Mixin of the UpPlatform
func (UpPlatform) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TenantMixin{},       // 添加租户
		mixin.OrganizationMixin{}, // 添加组织架构
	}
}
