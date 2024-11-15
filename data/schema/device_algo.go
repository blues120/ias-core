package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type DeviceAlgo struct {
	ent.Schema
}

func (DeviceAlgo) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.Uint64("device_id").Comment("设备id"),
		field.Uint("algo_group_id").Optional().Comment("算法组ID"),
		field.String("algo_group_name").Optional().Comment("算法组name(如十二合一算法)，前端显示用"),
		field.String("algo_group_version").Optional().Comment("算法组版本号，用于算法组整体更新场景"),
		field.String("name").Comment("算法名称"),
		field.String("version").Comment("算法版本"),
		field.Time("install_time").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Comment("安装时间"),
	}
}

// Annotations of the DeviceAlgo
func (DeviceAlgo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "device_algo"},
	}
}

// Edges of the DeviceAlgo.
func (DeviceAlgo) Edges() []ent.Edge {
	return nil
}

// Mixin of the DeviceAlgo
//func (DeviceAlgo) Mixin() []ent.Mixin {
//	return []ent.Mixin{
//		mixin.TenantMixin{},       // 添加租户
//		mixin.OrganizationMixin{}, // 添加组织架构
//	}
//}

// index of the DeviceAlgo
func (DeviceAlgo) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("device_id", "algo_group_id", "name").
			Unique(),
	}
}
