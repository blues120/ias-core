// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/warningtype"
)

// WarningType is the model entity for the WarningType schema.
type WarningType struct {
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
	// 告警I级类型
	AlarmType string `json:"alarm_type,omitempty"`
	// 告警II级类型
	AlarmName    string `json:"alarm_name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WarningType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case warningtype.FieldID:
			values[i] = new(sql.NullInt64)
		case warningtype.FieldAlarmType, warningtype.FieldAlarmName:
			values[i] = new(sql.NullString)
		case warningtype.FieldCreatedAt, warningtype.FieldUpdatedAt, warningtype.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WarningType fields.
func (wt *WarningType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case warningtype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			wt.ID = uint64(value.Int64)
		case warningtype.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				wt.CreatedAt = value.Time
			}
		case warningtype.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				wt.UpdatedAt = value.Time
			}
		case warningtype.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				wt.DeletedAt = value.Time
			}
		case warningtype.FieldAlarmType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field alarm_type", values[i])
			} else if value.Valid {
				wt.AlarmType = value.String
			}
		case warningtype.FieldAlarmName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field alarm_name", values[i])
			} else if value.Valid {
				wt.AlarmName = value.String
			}
		default:
			wt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WarningType.
// This includes values selected through modifiers, order, etc.
func (wt *WarningType) Value(name string) (ent.Value, error) {
	return wt.selectValues.Get(name)
}

// Update returns a builder for updating this WarningType.
// Note that you need to call WarningType.Unwrap() before calling this method if this WarningType
// was returned from a transaction, and the transaction was committed or rolled back.
func (wt *WarningType) Update() *WarningTypeUpdateOne {
	return NewWarningTypeClient(wt.config).UpdateOne(wt)
}

// Unwrap unwraps the WarningType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wt *WarningType) Unwrap() *WarningType {
	_tx, ok := wt.config.driver.(*txDriver)
	if !ok {
		panic("ent: WarningType is not a transactional entity")
	}
	wt.config.driver = _tx.drv
	return wt
}

// String implements the fmt.Stringer.
func (wt *WarningType) String() string {
	var builder strings.Builder
	builder.WriteString("WarningType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wt.ID))
	builder.WriteString("created_at=")
	builder.WriteString(wt.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(wt.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(wt.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("alarm_type=")
	builder.WriteString(wt.AlarmType)
	builder.WriteString(", ")
	builder.WriteString("alarm_name=")
	builder.WriteString(wt.AlarmName)
	builder.WriteByte(')')
	return builder.String()
}

// WarningTypes is a parsable slice of WarningType.
type WarningTypes []*WarningType