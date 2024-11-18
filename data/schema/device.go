package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/blues120/ias-core/biz"
	coreEnt "github.com/blues120/ias-core/data/ent"
	"github.com/blues120/ias-core/data/ent/device"
	"github.com/blues120/ias-core/data/ent/mixin"
	"github.com/blues120/ias-core/middleware"
)

type Device struct {
	ent.Schema
}

func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("name").Comment("设备名称"),
		field.String("display_name").Comment("设备展示名称"),
		field.Enum("type").GoType(biz.EdgeDeviceType("")).Comment("设备类型"),
		field.String("ext_id").Comment("设备id").Unique(),
		field.String("serial_no").Comment("序列号"),
		field.Enum("state").GoType(biz.DeviceState("")).Comment("状态"),
		field.String("mac").Comment("MAC地址"),
		field.String("zone_name").Comment("区域名称"),
		field.String("zone_id").Comment("区域ID"),
		field.String("workspace_id").Comment("工作空间ID"),
		field.String("equip_id").Comment("数生使用的设备ID"),
		field.String("equip_password").Comment("数生使用的设备密码"),
		field.String("device_info").Comment("设备信息"),
		field.String("model").Optional().Comment("设备型号"),
		field.Int64("auth_deadline").Optional().Comment("纳管授权时间"),
		field.Time("activated_at").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}).
			Comment("纳管激活时间"), // 设备的重新激活时间，因created_at不可重设，所以添加此字段
	}
}

// Annotations of the Device
func (Device) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "device"},
	}
}

// Mixin of the Device
func (Device) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},              // 添加创建时间、更新时间字段
		mixin.SoftDelete{},        // 添加软删除功能
		mixin.TenantMixin{},       // 添加租户
		mixin.OrganizationMixin{}, // 添加组织架构
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("camera", Camera.Type).Through("device_camera", DeviceCamera.Type),
		edge.To("task_device", Task.Type),
	}
}

func (Device) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("mac"),
	}
}

// 查找设备时候的拦截器
func (d Device) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		ent.InterceptFunc(func(next ent.Querier) ent.Querier {
			return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
				if q, ok := query.(*coreEnt.DeviceQuery); ok {
					deviceId, ok := ctx.Value(middleware.DeviceIdKey{}).(uint64)
					if ok && deviceId > 0 {
						q.Where(device.IDEQ(deviceId))
					}
				}
				return next.Query(ctx, query)
			})
		}),
	}
}
