package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/mixin"
)

// DeviceCamera holds the schema definition for the Task entity.
type DeviceCamera struct {
	ent.Schema
}

// Fields of the Task.
func (DeviceCamera) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.Uint64("device_id").Comment("任务id"),
		field.Uint64("camera_id").Comment("摄像头id"),
	}
}

// Annotations of the DeviceCamera
func (DeviceCamera) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.WithComments(true),
		entsql.Annotation{Table: "device_camera"},
	}
}

// Mixin of the DeviceCamera
func (DeviceCamera) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{}, // 添加创建时间、更新时间字段
		//mixin.TenantMixin{},       // 添加租户
		//mixin.OrganizationMixin{}, // 添加组织架构
	}
}

// Edges of the DeviceCamera.
func (DeviceCamera) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("camera", Camera.Type).Unique().Required().Field("camera_id"),
		edge.To("device", Task.Type).Unique().Required().Field("device_id"),
	}
}
