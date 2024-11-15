package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/mixin"
)

// TaskLimits holds the schema definition for the TaskLimits entity.
type TaskLimits struct {
	ent.Schema
}

// Fields of the TaskLimits.
func (TaskLimits) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("model").Optional().Comment("设备型号"),
		field.Uint64("maxCameraNum").Comment("最大摄像头数"),
		field.Uint64("algoNum").Comment("算法数"),
		field.Uint64("maxSubTaskNum").Comment("最大子任务路数"),
	}
}

// Annotations of the TaskLimits.
func (TaskLimits) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "task_limits"},
	}
}

// Mixin of the TaskLimits.
func (TaskLimits) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},       // 添加创建时间、更新时间字段
		mixin.SoftDelete{}, // 添加软删除功能
		//mixin.TenantMixin{},       // 添加租户
		//mixin.OrganizationMixin{}, // 添加组织架构
	}
}

// Edges of the TaskLimits.
func (TaskLimits) Edges() []ent.Edge {
	return nil
}
