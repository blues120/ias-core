// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/activeinfo"
)

// ActiveInfo is the model entity for the ActiveInfo schema.
type ActiveInfo struct {
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
	// 流程id
	ProcessID string `json:"process_id,omitempty"`
	// 激活开始时间
	StartTime string `json:"start_time,omitempty"`
	// 激活结果
	Result string `json:"result,omitempty"`
	// 消息
	Msg          string `json:"msg,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ActiveInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case activeinfo.FieldID:
			values[i] = new(sql.NullInt64)
		case activeinfo.FieldProcessID, activeinfo.FieldStartTime, activeinfo.FieldResult, activeinfo.FieldMsg:
			values[i] = new(sql.NullString)
		case activeinfo.FieldCreatedAt, activeinfo.FieldUpdatedAt, activeinfo.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ActiveInfo fields.
func (ai *ActiveInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case activeinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ai.ID = uint64(value.Int64)
		case activeinfo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ai.CreatedAt = value.Time
			}
		case activeinfo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ai.UpdatedAt = value.Time
			}
		case activeinfo.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ai.DeletedAt = value.Time
			}
		case activeinfo.FieldProcessID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field process_id", values[i])
			} else if value.Valid {
				ai.ProcessID = value.String
			}
		case activeinfo.FieldStartTime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				ai.StartTime = value.String
			}
		case activeinfo.FieldResult:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field result", values[i])
			} else if value.Valid {
				ai.Result = value.String
			}
		case activeinfo.FieldMsg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field msg", values[i])
			} else if value.Valid {
				ai.Msg = value.String
			}
		default:
			ai.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ActiveInfo.
// This includes values selected through modifiers, order, etc.
func (ai *ActiveInfo) Value(name string) (ent.Value, error) {
	return ai.selectValues.Get(name)
}

// Update returns a builder for updating this ActiveInfo.
// Note that you need to call ActiveInfo.Unwrap() before calling this method if this ActiveInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (ai *ActiveInfo) Update() *ActiveInfoUpdateOne {
	return NewActiveInfoClient(ai.config).UpdateOne(ai)
}

// Unwrap unwraps the ActiveInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ai *ActiveInfo) Unwrap() *ActiveInfo {
	_tx, ok := ai.config.driver.(*txDriver)
	if !ok {
		panic("ent: ActiveInfo is not a transactional entity")
	}
	ai.config.driver = _tx.drv
	return ai
}

// String implements the fmt.Stringer.
func (ai *ActiveInfo) String() string {
	var builder strings.Builder
	builder.WriteString("ActiveInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ai.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ai.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ai.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(ai.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("process_id=")
	builder.WriteString(ai.ProcessID)
	builder.WriteString(", ")
	builder.WriteString("start_time=")
	builder.WriteString(ai.StartTime)
	builder.WriteString(", ")
	builder.WriteString("result=")
	builder.WriteString(ai.Result)
	builder.WriteString(", ")
	builder.WriteString("msg=")
	builder.WriteString(ai.Msg)
	builder.WriteByte(')')
	return builder.String()
}

// ActiveInfos is a parsable slice of ActiveInfo.
type ActiveInfos []*ActiveInfo
