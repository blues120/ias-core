package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/blues120/ias-core/data/ent/mixin"
)

type Algorithm struct {
	ent.Schema
}

func (Algorithm) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("name").Comment("算法名称"),
		field.String("type").Comment("算法类型:图片帧image/视频流video"),
		field.String("description").Comment("算法描述"),
		field.String("version").Comment("算法版本"),
		field.String("app_name").Comment("应用名称"),
		field.String("alarm_type").Comment("告警类型"),
		field.String("alarm_name").Comment("告警名称"),
		field.Uint("notify").Comment("是否告警"),
		field.JSON("extend", map[string]interface{}{}).Comment("非通用属性"),
		field.Uint("draw_type").Comment("绘制区域类型 多边形区域/流量方向及界线 1/2"),
		field.Uint("base_type").Comment("底库类型 无/人员/车辆 0/1/2"),
		field.Uint("available").Default(1).Comment("是否可用 默认为1,不可用为0"),
		field.String("image").MaxLen(1000).Comment("算法镜像"),
		field.String("label_map").MaxLen(2000).Optional().Comment("英中文label映射"),
		field.String("target").Optional().Comment("检测目标"),
		field.String("algo_name_en").Optional().Comment("算法英文名，下发给agent配置需要"),
		field.Uint("algo_group_id").Optional().Comment("算法组ID"),
		field.String("algo_group_name").Optional().Comment("算法组name(如十二合一算法)，前端显示用"),
		field.String("algo_group_version").Optional().Comment("算法组版本号，用于算法组整体更新场景"),
		field.Text("config").Optional().Comment("算法特有配置，如算能算法的单独配置"),
		field.String("provider").Optional().Comment("算法供应商 ctyun_ias/sophgo_park/sophgo_city"),
		field.String("algo_id").Optional().Comment("算法 id，非自增 id，用于填充园区算法 alg_flag 字段（使用时需转换为 Uint）"),
		field.String("platform").Optional().Comment("平台类型,服务器/边缘设备"),
		field.String("device_model").Optional().Comment("设备型号"),
		field.Uint("is_group_type").Optional().Default(0).Comment("是否是多合一算法组类型算法, 否/是 0/1"),
		field.String("prefix").Optional().Comment("算法包服务启动api前缀 ip:port"),
		// field.Uint("control_period").Default(0).Comment("是否打开布控时段"),
	}
}

// Annotations of the Algorithm
func (Algorithm) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "algorithm"},
	}
}

// Mixin of the Algorithm
func (Algorithm) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},       // 添加创建时间、更新时间字段
		mixin.SoftDelete{}, // 添加软删除功能
		// 算法表不需要租户和组织架构信息
		// mixin.TenantMixin{},       // 添加租户
		// mixin.OrganizationMixin{}, // 添加组织架构
	}
}

type AlgorithmAlarmType struct {
	ent.Schema
}

func (AlgorithmAlarmType) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("type_name").NotEmpty().Default("").Comment("告警类型"),
	}
}

// Annotations of the AlgorithmAlarmType
func (AlgorithmAlarmType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "algorithm_alarm_type"},
	}
}

// Mixin of the AlgorithmAlarmType
func (AlgorithmAlarmType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},       // 添加创建时间、更新时间字段
		mixin.SoftDelete{}, // 添加软删除功能
	}
}

// Edges of the Algorithm.
func (Algorithm) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tasks", Task.Type),
	}
}
