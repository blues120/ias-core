// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/biz/streaming"
	"github.com/blues120/ias-core/data/ent/camera"
)

// Camera is the model entity for the Camera schema.
type Camera struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint64 `json:"id,omitempty"`
	// 创建时间
	CreatedAt time.Time `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 删除时间
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID string `json:"tenant_id,omitempty"`
	// 授权的组织 id 列表，#分隔
	AccessOrgList string `json:"access_org_list,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 点位
	Position string `json:"position,omitempty"`
	// 区域,存id
	Region string `json:"region,omitempty"`
	// 区域,存字符串
	RegionStr string `json:"region_str,omitempty"`
	// 经度
	Longitude float64 `json:"longitude,omitempty"`
	// 纬度
	Latitude float64 `json:"latitude,omitempty"`
	// 自定义编号
	CustomNumber int `json:"custom_number,omitempty"`
	// 通道id
	ChannelID string `json:"channel_id,omitempty"`
	// 设备序列号
	SerialNumber string `json:"serial_number,omitempty"`
	// 杆号
	PoleNumber string `json:"pole_number,omitempty"`
	// 设备描述
	DeviceDescription string `json:"device_description,omitempty"`
	// 适用场景
	Scene string `json:"scene,omitempty"`
	// 所属场所
	Place string `json:"place,omitempty"`
	// 状态
	Status biz.CameraStatus `json:"status,omitempty"`
	// 流媒体协议类型
	SpType streaming.ProtocolType `json:"sp_type,omitempty"`
	// 流媒体协议地址
	SpSource string `json:"sp_source,omitempty"`
	// 流媒体协议编码名称
	SpCodecName string `json:"sp_codec_name,omitempty"`
	// 流媒体协议宽度
	SpWidth int32 `json:"sp_width,omitempty"`
	// 流媒体协议高度
	SpHeight int32 `json:"sp_height,omitempty"`
	// 国标传输协议 UDP/TCP
	TransType string `json:"trans_type,omitempty"`
	// IP地址
	DeviceIP string `json:"device_ip,omitempty"`
	// 端口号
	DevicePort int32 `json:"device_port,omitempty"`
	// 国标ID
	GBID string `json:"gb_id,omitempty"`
	// 国标信令SIP认证用户名
	SipUser string `json:"sip_user,omitempty"`
	// 国标信令SIP认证密码
	SipPassword string `json:"sip_password,omitempty"`
	// 国标通道编码
	GBChannelID string `json:"gb_channel_id,omitempty"`
	// 向上级联自定义国标通道编码
	UpGBChannelID string `json:"up_gb_channel_id,omitempty"`
	// 国标设备类型
	GBDeviceType string `json:"gb_device_type,omitempty"`
	// 多媒体设备类型
	Type biz.MediaType `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CameraQuery when eager-loading is set.
	Edges        CameraEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CameraEdges holds the relations/edges for other nodes in the graph.
type CameraEdges struct {
	// Task holds the value of the task edge.
	Task []*Task `json:"task,omitempty"`
	// Device holds the value of the device edge.
	Device []*Device `json:"device,omitempty"`
	// TaskCamera holds the value of the task_camera edge.
	TaskCamera []*TaskCamera `json:"task_camera,omitempty"`
	// DeviceCamera holds the value of the device_camera edge.
	DeviceCamera []*DeviceCamera `json:"device_camera,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// TaskOrErr returns the Task value or an error if the edge
// was not loaded in eager-loading.
func (e CameraEdges) TaskOrErr() ([]*Task, error) {
	if e.loadedTypes[0] {
		return e.Task, nil
	}
	return nil, &NotLoadedError{edge: "task"}
}

// DeviceOrErr returns the Device value or an error if the edge
// was not loaded in eager-loading.
func (e CameraEdges) DeviceOrErr() ([]*Device, error) {
	if e.loadedTypes[1] {
		return e.Device, nil
	}
	return nil, &NotLoadedError{edge: "device"}
}

// TaskCameraOrErr returns the TaskCamera value or an error if the edge
// was not loaded in eager-loading.
func (e CameraEdges) TaskCameraOrErr() ([]*TaskCamera, error) {
	if e.loadedTypes[2] {
		return e.TaskCamera, nil
	}
	return nil, &NotLoadedError{edge: "task_camera"}
}

// DeviceCameraOrErr returns the DeviceCamera value or an error if the edge
// was not loaded in eager-loading.
func (e CameraEdges) DeviceCameraOrErr() ([]*DeviceCamera, error) {
	if e.loadedTypes[3] {
		return e.DeviceCamera, nil
	}
	return nil, &NotLoadedError{edge: "device_camera"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Camera) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case camera.FieldLongitude, camera.FieldLatitude:
			values[i] = new(sql.NullFloat64)
		case camera.FieldID, camera.FieldCustomNumber, camera.FieldSpWidth, camera.FieldSpHeight, camera.FieldDevicePort:
			values[i] = new(sql.NullInt64)
		case camera.FieldTenantID, camera.FieldAccessOrgList, camera.FieldName, camera.FieldPosition, camera.FieldRegion, camera.FieldRegionStr, camera.FieldChannelID, camera.FieldSerialNumber, camera.FieldPoleNumber, camera.FieldDeviceDescription, camera.FieldScene, camera.FieldPlace, camera.FieldStatus, camera.FieldSpType, camera.FieldSpSource, camera.FieldSpCodecName, camera.FieldTransType, camera.FieldDeviceIP, camera.FieldGBID, camera.FieldSipUser, camera.FieldSipPassword, camera.FieldGBChannelID, camera.FieldUpGBChannelID, camera.FieldGBDeviceType, camera.FieldType:
			values[i] = new(sql.NullString)
		case camera.FieldCreatedAt, camera.FieldUpdatedAt, camera.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Camera fields.
func (c *Camera) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case camera.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = uint64(value.Int64)
		case camera.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case camera.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case camera.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = value.Time
			}
		case camera.FieldTenantID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				c.TenantID = value.String
			}
		case camera.FieldAccessOrgList:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_org_list", values[i])
			} else if value.Valid {
				c.AccessOrgList = value.String
			}
		case camera.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case camera.FieldPosition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field position", values[i])
			} else if value.Valid {
				c.Position = value.String
			}
		case camera.FieldRegion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field region", values[i])
			} else if value.Valid {
				c.Region = value.String
			}
		case camera.FieldRegionStr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field region_str", values[i])
			} else if value.Valid {
				c.RegionStr = value.String
			}
		case camera.FieldLongitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				c.Longitude = value.Float64
			}
		case camera.FieldLatitude:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				c.Latitude = value.Float64
			}
		case camera.FieldCustomNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field custom_number", values[i])
			} else if value.Valid {
				c.CustomNumber = int(value.Int64)
			}
		case camera.FieldChannelID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field channel_id", values[i])
			} else if value.Valid {
				c.ChannelID = value.String
			}
		case camera.FieldSerialNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field serial_number", values[i])
			} else if value.Valid {
				c.SerialNumber = value.String
			}
		case camera.FieldPoleNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pole_number", values[i])
			} else if value.Valid {
				c.PoleNumber = value.String
			}
		case camera.FieldDeviceDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_description", values[i])
			} else if value.Valid {
				c.DeviceDescription = value.String
			}
		case camera.FieldScene:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field scene", values[i])
			} else if value.Valid {
				c.Scene = value.String
			}
		case camera.FieldPlace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field place", values[i])
			} else if value.Valid {
				c.Place = value.String
			}
		case camera.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				c.Status = biz.CameraStatus(value.String)
			}
		case camera.FieldSpType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sp_type", values[i])
			} else if value.Valid {
				c.SpType = streaming.ProtocolType(value.String)
			}
		case camera.FieldSpSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sp_source", values[i])
			} else if value.Valid {
				c.SpSource = value.String
			}
		case camera.FieldSpCodecName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sp_codec_name", values[i])
			} else if value.Valid {
				c.SpCodecName = value.String
			}
		case camera.FieldSpWidth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sp_width", values[i])
			} else if value.Valid {
				c.SpWidth = int32(value.Int64)
			}
		case camera.FieldSpHeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sp_height", values[i])
			} else if value.Valid {
				c.SpHeight = int32(value.Int64)
			}
		case camera.FieldTransType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field trans_type", values[i])
			} else if value.Valid {
				c.TransType = value.String
			}
		case camera.FieldDeviceIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_ip", values[i])
			} else if value.Valid {
				c.DeviceIP = value.String
			}
		case camera.FieldDevicePort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field device_port", values[i])
			} else if value.Valid {
				c.DevicePort = int32(value.Int64)
			}
		case camera.FieldGBID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gb_id", values[i])
			} else if value.Valid {
				c.GBID = value.String
			}
		case camera.FieldSipUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sip_user", values[i])
			} else if value.Valid {
				c.SipUser = value.String
			}
		case camera.FieldSipPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sip_password", values[i])
			} else if value.Valid {
				c.SipPassword = value.String
			}
		case camera.FieldGBChannelID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gb_channel_id", values[i])
			} else if value.Valid {
				c.GBChannelID = value.String
			}
		case camera.FieldUpGBChannelID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field up_gb_channel_id", values[i])
			} else if value.Valid {
				c.UpGBChannelID = value.String
			}
		case camera.FieldGBDeviceType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gb_device_type", values[i])
			} else if value.Valid {
				c.GBDeviceType = value.String
			}
		case camera.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				c.Type = biz.MediaType(value.String)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Camera.
// This includes values selected through modifiers, order, etc.
func (c *Camera) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryTask queries the "task" edge of the Camera entity.
func (c *Camera) QueryTask() *TaskQuery {
	return NewCameraClient(c.config).QueryTask(c)
}

// QueryDevice queries the "device" edge of the Camera entity.
func (c *Camera) QueryDevice() *DeviceQuery {
	return NewCameraClient(c.config).QueryDevice(c)
}

// QueryTaskCamera queries the "task_camera" edge of the Camera entity.
func (c *Camera) QueryTaskCamera() *TaskCameraQuery {
	return NewCameraClient(c.config).QueryTaskCamera(c)
}

// QueryDeviceCamera queries the "device_camera" edge of the Camera entity.
func (c *Camera) QueryDeviceCamera() *DeviceCameraQuery {
	return NewCameraClient(c.config).QueryDeviceCamera(c)
}

// Update returns a builder for updating this Camera.
// Note that you need to call Camera.Unwrap() before calling this method if this Camera
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Camera) Update() *CameraUpdateOne {
	return NewCameraClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Camera entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Camera) Unwrap() *Camera {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Camera is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Camera) String() string {
	var builder strings.Builder
	builder.WriteString("Camera(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(c.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("tenant_id=")
	builder.WriteString(c.TenantID)
	builder.WriteString(", ")
	builder.WriteString("access_org_list=")
	builder.WriteString(c.AccessOrgList)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("position=")
	builder.WriteString(c.Position)
	builder.WriteString(", ")
	builder.WriteString("region=")
	builder.WriteString(c.Region)
	builder.WriteString(", ")
	builder.WriteString("region_str=")
	builder.WriteString(c.RegionStr)
	builder.WriteString(", ")
	builder.WriteString("longitude=")
	builder.WriteString(fmt.Sprintf("%v", c.Longitude))
	builder.WriteString(", ")
	builder.WriteString("latitude=")
	builder.WriteString(fmt.Sprintf("%v", c.Latitude))
	builder.WriteString(", ")
	builder.WriteString("custom_number=")
	builder.WriteString(fmt.Sprintf("%v", c.CustomNumber))
	builder.WriteString(", ")
	builder.WriteString("channel_id=")
	builder.WriteString(c.ChannelID)
	builder.WriteString(", ")
	builder.WriteString("serial_number=")
	builder.WriteString(c.SerialNumber)
	builder.WriteString(", ")
	builder.WriteString("pole_number=")
	builder.WriteString(c.PoleNumber)
	builder.WriteString(", ")
	builder.WriteString("device_description=")
	builder.WriteString(c.DeviceDescription)
	builder.WriteString(", ")
	builder.WriteString("scene=")
	builder.WriteString(c.Scene)
	builder.WriteString(", ")
	builder.WriteString("place=")
	builder.WriteString(c.Place)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", c.Status))
	builder.WriteString(", ")
	builder.WriteString("sp_type=")
	builder.WriteString(fmt.Sprintf("%v", c.SpType))
	builder.WriteString(", ")
	builder.WriteString("sp_source=")
	builder.WriteString(c.SpSource)
	builder.WriteString(", ")
	builder.WriteString("sp_codec_name=")
	builder.WriteString(c.SpCodecName)
	builder.WriteString(", ")
	builder.WriteString("sp_width=")
	builder.WriteString(fmt.Sprintf("%v", c.SpWidth))
	builder.WriteString(", ")
	builder.WriteString("sp_height=")
	builder.WriteString(fmt.Sprintf("%v", c.SpHeight))
	builder.WriteString(", ")
	builder.WriteString("trans_type=")
	builder.WriteString(c.TransType)
	builder.WriteString(", ")
	builder.WriteString("device_ip=")
	builder.WriteString(c.DeviceIP)
	builder.WriteString(", ")
	builder.WriteString("device_port=")
	builder.WriteString(fmt.Sprintf("%v", c.DevicePort))
	builder.WriteString(", ")
	builder.WriteString("gb_id=")
	builder.WriteString(c.GBID)
	builder.WriteString(", ")
	builder.WriteString("sip_user=")
	builder.WriteString(c.SipUser)
	builder.WriteString(", ")
	builder.WriteString("sip_password=")
	builder.WriteString(c.SipPassword)
	builder.WriteString(", ")
	builder.WriteString("gb_channel_id=")
	builder.WriteString(c.GBChannelID)
	builder.WriteString(", ")
	builder.WriteString("up_gb_channel_id=")
	builder.WriteString(c.UpGBChannelID)
	builder.WriteString(", ")
	builder.WriteString("gb_device_type=")
	builder.WriteString(c.GBDeviceType)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", c.Type))
	builder.WriteByte(')')
	return builder.String()
}

// Cameras is a parsable slice of Camera.
type Cameras []*Camera
