// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/data/ent/algorithm"
	"github.com/blues120/ias-core/data/ent/device"
	"github.com/blues120/ias-core/data/ent/task"
)

// Task is the model entity for the Task schema.
type Task struct {
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
	// 任务类型
	Type biz.TaskType `json:"type,omitempty"`
	// 算法id
	AlgoID uint64 `json:"algo_id,omitempty"`
	// 算法执行间隔
	AlgoInterval float64 `json:"algo_interval,omitempty"`
	// 算法扩展数据
	AlgoExtra string `json:"algo_extra,omitempty"`
	// 扩展字段
	Extend string `json:"extend,omitempty"`
	// 任务执行的设备id
	DeviceID uint64 `json:"device_id,omitempty"`
	// 最后启动时间
	LastStartTime *sql.NullTime `json:"last_start_time,omitempty"`
	// 任务状态
	Status biz.TaskStatus `json:"status,omitempty"`
	// 算法组ID
	AlgoGroupID uint `json:"algo_group_id,omitempty"`
	// 父ID, 作为任务下发的ID
	ParentID string `json:"parent_id,omitempty"`
	// 首页是否告警
	IsWarn uint32 `json:"is_warn,omitempty"`
	// 告警周期
	Period uint32 `json:"period,omitempty"`
	// 算法特有配置
	AlgoConfig string `json:"algo_config,omitempty"`
	// 任务状态变更原因
	Reason string `json:"reason,omitempty"`
	// 运行时段类型
	AllowTimeType string `json:"allow_time_type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TaskQuery when eager-loading is set.
	Edges        TaskEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TaskEdges holds the relations/edges for other nodes in the graph.
type TaskEdges struct {
	// Camera holds the value of the camera edge.
	Camera []*Camera `json:"camera,omitempty"`
	// Algorithm holds the value of the algorithm edge.
	Algorithm *Algorithm `json:"algorithm,omitempty"`
	// Device holds the value of the device edge.
	Device *Device `json:"device,omitempty"`
	// TaskCamera holds the value of the task_camera edge.
	TaskCamera []*TaskCamera `json:"task_camera,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// CameraOrErr returns the Camera value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) CameraOrErr() ([]*Camera, error) {
	if e.loadedTypes[0] {
		return e.Camera, nil
	}
	return nil, &NotLoadedError{edge: "camera"}
}

// AlgorithmOrErr returns the Algorithm value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskEdges) AlgorithmOrErr() (*Algorithm, error) {
	if e.loadedTypes[1] {
		if e.Algorithm == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: algorithm.Label}
		}
		return e.Algorithm, nil
	}
	return nil, &NotLoadedError{edge: "algorithm"}
}

// DeviceOrErr returns the Device value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TaskEdges) DeviceOrErr() (*Device, error) {
	if e.loadedTypes[2] {
		if e.Device == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: device.Label}
		}
		return e.Device, nil
	}
	return nil, &NotLoadedError{edge: "device"}
}

// TaskCameraOrErr returns the TaskCamera value or an error if the edge
// was not loaded in eager-loading.
func (e TaskEdges) TaskCameraOrErr() ([]*TaskCamera, error) {
	if e.loadedTypes[3] {
		return e.TaskCamera, nil
	}
	return nil, &NotLoadedError{edge: "task_camera"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Task) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case task.FieldAlgoInterval:
			values[i] = new(sql.NullFloat64)
		case task.FieldID, task.FieldAlgoID, task.FieldDeviceID, task.FieldAlgoGroupID, task.FieldIsWarn, task.FieldPeriod:
			values[i] = new(sql.NullInt64)
		case task.FieldTenantID, task.FieldAccessOrgList, task.FieldName, task.FieldType, task.FieldAlgoExtra, task.FieldExtend, task.FieldStatus, task.FieldParentID, task.FieldAlgoConfig, task.FieldReason, task.FieldAllowTimeType:
			values[i] = new(sql.NullString)
		case task.FieldCreatedAt, task.FieldUpdatedAt, task.FieldDeletedAt, task.FieldLastStartTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Task fields.
func (t *Task) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case task.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = uint64(value.Int64)
		case task.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Time
			}
		case task.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Time
			}
		case task.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				t.DeletedAt = value.Time
			}
		case task.FieldTenantID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				t.TenantID = value.String
			}
		case task.FieldAccessOrgList:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_org_list", values[i])
			} else if value.Valid {
				t.AccessOrgList = value.String
			}
		case task.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case task.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = biz.TaskType(value.String)
			}
		case task.FieldAlgoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field algo_id", values[i])
			} else if value.Valid {
				t.AlgoID = uint64(value.Int64)
			}
		case task.FieldAlgoInterval:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field algo_interval", values[i])
			} else if value.Valid {
				t.AlgoInterval = value.Float64
			}
		case task.FieldAlgoExtra:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field algo_extra", values[i])
			} else if value.Valid {
				t.AlgoExtra = value.String
			}
		case task.FieldExtend:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field extend", values[i])
			} else if value.Valid {
				t.Extend = value.String
			}
		case task.FieldDeviceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field device_id", values[i])
			} else if value.Valid {
				t.DeviceID = uint64(value.Int64)
			}
		case task.FieldLastStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_start_time", values[i])
			} else if value.Valid {
				t.LastStartTime = value
			}
		case task.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = biz.TaskStatus(value.String)
			}
		case task.FieldAlgoGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field algo_group_id", values[i])
			} else if value.Valid {
				t.AlgoGroupID = uint(value.Int64)
			}
		case task.FieldParentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				t.ParentID = value.String
			}
		case task.FieldIsWarn:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_warn", values[i])
			} else if value.Valid {
				t.IsWarn = uint32(value.Int64)
			}
		case task.FieldPeriod:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field period", values[i])
			} else if value.Valid {
				t.Period = uint32(value.Int64)
			}
		case task.FieldAlgoConfig:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field algo_config", values[i])
			} else if value.Valid {
				t.AlgoConfig = value.String
			}
		case task.FieldReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				t.Reason = value.String
			}
		case task.FieldAllowTimeType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field allow_time_type", values[i])
			} else if value.Valid {
				t.AllowTimeType = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Task.
// This includes values selected through modifiers, order, etc.
func (t *Task) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryCamera queries the "camera" edge of the Task entity.
func (t *Task) QueryCamera() *CameraQuery {
	return NewTaskClient(t.config).QueryCamera(t)
}

// QueryAlgorithm queries the "algorithm" edge of the Task entity.
func (t *Task) QueryAlgorithm() *AlgorithmQuery {
	return NewTaskClient(t.config).QueryAlgorithm(t)
}

// QueryDevice queries the "device" edge of the Task entity.
func (t *Task) QueryDevice() *DeviceQuery {
	return NewTaskClient(t.config).QueryDevice(t)
}

// QueryTaskCamera queries the "task_camera" edge of the Task entity.
func (t *Task) QueryTaskCamera() *TaskCameraQuery {
	return NewTaskClient(t.config).QueryTaskCamera(t)
}

// Update returns a builder for updating this Task.
// Note that you need to call Task.Unwrap() before calling this method if this Task
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Task) Update() *TaskUpdateOne {
	return NewTaskClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Task entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Task) Unwrap() *Task {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Task is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Task) String() string {
	var builder strings.Builder
	builder.WriteString("Task(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(t.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(t.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(t.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("tenant_id=")
	builder.WriteString(t.TenantID)
	builder.WriteString(", ")
	builder.WriteString("access_org_list=")
	builder.WriteString(t.AccessOrgList)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", t.Type))
	builder.WriteString(", ")
	builder.WriteString("algo_id=")
	builder.WriteString(fmt.Sprintf("%v", t.AlgoID))
	builder.WriteString(", ")
	builder.WriteString("algo_interval=")
	builder.WriteString(fmt.Sprintf("%v", t.AlgoInterval))
	builder.WriteString(", ")
	builder.WriteString("algo_extra=")
	builder.WriteString(t.AlgoExtra)
	builder.WriteString(", ")
	builder.WriteString("extend=")
	builder.WriteString(t.Extend)
	builder.WriteString(", ")
	builder.WriteString("device_id=")
	builder.WriteString(fmt.Sprintf("%v", t.DeviceID))
	builder.WriteString(", ")
	builder.WriteString("last_start_time=")
	builder.WriteString(fmt.Sprintf("%v", t.LastStartTime))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteString(", ")
	builder.WriteString("algo_group_id=")
	builder.WriteString(fmt.Sprintf("%v", t.AlgoGroupID))
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(t.ParentID)
	builder.WriteString(", ")
	builder.WriteString("is_warn=")
	builder.WriteString(fmt.Sprintf("%v", t.IsWarn))
	builder.WriteString(", ")
	builder.WriteString("period=")
	builder.WriteString(fmt.Sprintf("%v", t.Period))
	builder.WriteString(", ")
	builder.WriteString("algo_config=")
	builder.WriteString(t.AlgoConfig)
	builder.WriteString(", ")
	builder.WriteString("reason=")
	builder.WriteString(t.Reason)
	builder.WriteString(", ")
	builder.WriteString("allow_time_type=")
	builder.WriteString(t.AllowTimeType)
	builder.WriteByte(')')
	return builder.String()
}

// Tasks is a parsable slice of Task.
type Tasks []*Task
