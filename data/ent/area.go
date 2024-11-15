// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/area"
)

// Area is the model entity for the Area schema.
type Area struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint64 `json:"id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 层级
	Level uint64 `json:"level,omitempty"`
	// 父id
	Pid          int64 `json:"pid,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Area) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case area.FieldID, area.FieldLevel, area.FieldPid:
			values[i] = new(sql.NullInt64)
		case area.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Area fields.
func (a *Area) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case area.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint64(value.Int64)
		case area.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case area.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				a.Level = uint64(value.Int64)
			}
		case area.FieldPid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pid", values[i])
			} else if value.Valid {
				a.Pid = value.Int64
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Area.
// This includes values selected through modifiers, order, etc.
func (a *Area) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// Update returns a builder for updating this Area.
// Note that you need to call Area.Unwrap() before calling this method if this Area
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Area) Update() *AreaUpdateOne {
	return NewAreaClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Area entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Area) Unwrap() *Area {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Area is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Area) String() string {
	var builder strings.Builder
	builder.WriteString("Area(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("level=")
	builder.WriteString(fmt.Sprintf("%v", a.Level))
	builder.WriteString(", ")
	builder.WriteString("pid=")
	builder.WriteString(fmt.Sprintf("%v", a.Pid))
	builder.WriteByte(')')
	return builder.String()
}

// Areas is a parsable slice of Area.
type Areas []*Area