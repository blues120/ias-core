package schema

import (
	"context"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz"
	"gitlab.ctyuncdn.cn/ias/ias-core/biz/streaming"
	coreEnt "gitlab.ctyuncdn.cn/ias/ias-core/data/ent"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/camera"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/devicecamera"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/mixin"
	"gitlab.ctyuncdn.cn/ias/ias-core/middleware"
)

// Camera holds the schema definition for the Camera entity.
type Camera struct {
	ent.Schema
}

// Fields of the Camera.
func (Camera) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Comment("id"),
		field.String("name").Comment("名称"),
		field.String("position").Comment("点位"),
		field.String("region").Comment("区域,存id").Optional(),
		field.String("region_str").Comment("区域,存字符串").Optional(),
		field.Float("longitude").Comment("经度"),
		field.Float("latitude").Comment("纬度"),
		field.Int("custom_number").Comment("自定义编号"),
		field.String("channel_id").Comment("通道id"),
		field.String("serial_number").Comment("设备序列号"),
		field.String("pole_number").Comment("杆号"),
		field.String("device_description").Comment("设备描述"),
		field.String("scene").Comment("适用场景"),
		field.String("place").Default("").Comment("所属场所"),
		field.Enum("status").GoType(biz.CameraStatus("")).Comment("状态"),
		field.Enum("sp_type").GoType(streaming.ProtocolType("")).Comment("流媒体协议类型"),
		field.String("sp_source").MaxLen(2000).Comment("流媒体协议地址"),
		field.String("sp_codec_name").Comment("流媒体协议编码名称"),
		field.Int32("sp_width").Comment("流媒体协议宽度"),
		field.Int32("sp_height").Comment("流媒体协议高度"),
		field.String("trans_type").Default("TCP").Comment("国标传输协议 UDP/TCP"),
		field.String("device_ip").Default("0.0.0.0").Comment("IP地址"),
		field.Int32("device_port").Default(5060).Comment("端口号"),
		field.String("gb_id").Default("34020000000000000000").Comment("国标ID"),
		field.String("sip_user").Default("").Comment("国标信令SIP认证用户名"),
		field.String("sip_password").Default("").Comment("国标信令SIP认证密码"),
		field.String("gb_channel_id").Default("").Comment("国标通道编码"),
		field.String("up_gb_channel_id").Default("").Comment("向上级联自定义国标通道编码"),
		field.String("gb_device_type").Default("").Comment("国标设备类型"),
		field.Enum("type").GoType(biz.MediaType("")).Comment("多媒体设备类型"),
	}
}

// Annotations of the Camera
func (Camera) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		// Override the default table name
		entsql.Annotation{Table: "camera"},
	}
}

// Mixin of the Camera
func (Camera) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},              // 添加创建时间、更新时间字段
		mixin.SoftDelete{},        // 添加软删除功能
		mixin.TenantMixin{},       // 添加租户
		mixin.OrganizationMixin{}, // 添加组织架构
	}
}

// Edges of the Camera.
func (Camera) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).Ref("camera").
			Through("task_camera", TaskCamera.Type), // 一个摄像机可以属于多个任务
		edge.From("device", Device.Type).Ref("camera").
			Through("device_camera", DeviceCamera.Type),
	}
}

// 查找camera表时候的拦截器
func (d Camera) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		ent.InterceptFunc(func(next ent.Querier) ent.Querier {
			return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
				if q, ok := query.(*coreEnt.CameraQuery); ok {

					deviceId, ok := ctx.Value(middleware.DeviceIdKey{}).(uint64)
					if !ok || deviceId == 0 {
						return next.Query(ctx, query)
					}

					// 需要清除分页参数。实测这里limit(0)和limit(-1)都不可行
					cameras, err := q.Clone().Limit(9999999).Offset(0).QueryDeviceCamera().Where(devicecamera.DeviceIDEQ(deviceId)).All(ctx)

					cameraIDs := []uint64{}
					if err != nil && !coreEnt.IsNotFound(err) {
						return next.Query(ctx, query)
					}
					for _, cameraInfo := range cameras {
						cameraIDs = append(cameraIDs, cameraInfo.CameraID)
					}

					// 在 camera 表中过滤结果
					q.Where(camera.IDIn(cameraIDs...))

				}
				return next.Query(ctx, query)
			})
		}),
	}
}
