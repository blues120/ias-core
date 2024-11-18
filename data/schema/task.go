package schema

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/blues120/ias-core/biz"
	coreEnt "github.com/blues120/ias-core/data/ent"
	"github.com/blues120/ias-core/data/ent/mixin"
	"github.com/blues120/ias-core/data/ent/task"
	"github.com/blues120/ias-core/middleware"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("name").Comment("名称"),
		field.Enum("type").GoType(biz.TaskType("")).Comment("任务类型"),
		field.Uint64("algo_id").Comment("算法id"),
		field.Float("algo_interval").Comment("算法执行间隔"),
		field.Text("algo_extra").Comment("算法扩展数据"),
		field.Text("extend").Comment("扩展字段"),
		field.Uint64("device_id").Comment("任务执行的设备id"),
		// field.Time("last_start_time").Comment("最后启动时间").Nillable(),
		field.Time("last_start_time").GoType(&sql.NullTime{}).Comment("最后启动时间").Optional(),
		field.Enum("status").GoType(biz.TaskStatus("")).Comment("任务状态"),
		field.Uint("algo_group_id").Comment("算法组ID"),
		field.String("parent_id").Comment("父ID, 作为任务下发的ID"),
		field.Uint32("is_warn").Comment("首页是否告警"),
		field.Uint32("period").Comment("告警周期"),
		field.Text("algo_config").Comment("算法特有配置").Optional(),
		field.String("reason").Default("").Comment("任务状态变更原因"),
		field.String("allow_time_type").Comment("运行时段类型").Optional(),
	}
}

// Annotations of the Task
func (Task) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "task"},
	}
}

// Mixin of the Task
func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},              // 添加创建时间、更新时间字段
		mixin.SoftDelete{},        // 添加软删除功能
		mixin.TenantMixin{},       // 添加租户
		mixin.OrganizationMixin{}, // 添加组织架构
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("camera", Camera.Type).Through("task_camera", TaskCamera.Type),
		edge.From("algorithm", Algorithm.Type).Ref("tasks").Unique().Required().Field("algo_id"),
		edge.From("device", Device.Type).Ref("task_device").Unique().Required().Field("device_id"),
	}
}

// 查找任务时候的拦截器
func (t Task) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		ent.InterceptFunc(func(next ent.Querier) ent.Querier {
			return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
				if q, ok := query.(*coreEnt.TaskQuery); ok {
					deviceId, ok := ctx.Value(middleware.DeviceIdKey{}).(uint64)
					if ok && deviceId > 0 {
						q.Where(task.DeviceID(deviceId))
					}
				}
				return next.Query(ctx, query)
			})
		}),
	}
}
