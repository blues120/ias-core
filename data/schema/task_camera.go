package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/blues120/ias-core/data/ent/mixin"
)

// TaskCamera holds the schema definition for the Task entity.
type TaskCamera struct {
	ent.Schema
}

// Fields of the Task.
func (TaskCamera) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.Uint64("task_id").Comment("任务id"),
		field.Uint64("camera_id").Comment("摄像头id"),
		field.Text("multi_img_box").Comment("划线区域"),
	}
}

// Annotations of the TaskCamera
func (TaskCamera) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "task_camera"},
	}
}

// Mixin of the TaskCamera
func (TaskCamera) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},              // 添加创建时间、更新时间字段
		mixin.TenantMixin{},       // 添加租户
		mixin.OrganizationMixin{}, // 添加组织架构
	}
}

// Edges of the TaskCamera.
func (TaskCamera) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("camera", Camera.Type).Unique().Required().Field("camera_id"),
		edge.To("task", Task.Type).Unique().Required().Field("task_id"),
	}
}
