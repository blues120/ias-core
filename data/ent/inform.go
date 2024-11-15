// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/inform"
)

// Inform is the model entity for the Inform schema.
type Inform struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID string `json:"tenant_id,omitempty"`
	// 授权的组织 id 列表，#分隔
	AccessOrgList string `json:"access_org_list,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// 告警名称
	AlarmName string `json:"alarm_name,omitempty"`
	// 告警名称
	AlarmType string `json:"alarm_type,omitempty"`
	// 签名
	SignName string `json:"sign_name,omitempty"`
	// 通知模板内容
	NotifyTemplate string `json:"notify_template,omitempty"`
	// 通知模板
	TemplateCode string `json:"template_code,omitempty"`
	// 通知号码
	PhoneNumbers string `json:"phone_numbers,omitempty"`
	// 通知开关
	NotifySwitch string `json:"notify_switch,omitempty"`
	// 任务名称
	TaskName string `json:"task_name,omitempty"`
	// 任务id
	TaskID       uint64 `json:"task_id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Inform) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case inform.FieldID, inform.FieldTaskID:
			values[i] = new(sql.NullInt64)
		case inform.FieldTenantID, inform.FieldAccessOrgList, inform.FieldAlarmName, inform.FieldAlarmType, inform.FieldSignName, inform.FieldNotifyTemplate, inform.FieldTemplateCode, inform.FieldPhoneNumbers, inform.FieldNotifySwitch, inform.FieldTaskName:
			values[i] = new(sql.NullString)
		case inform.FieldCreatedAt, inform.FieldUpdatedAt, inform.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Inform fields.
func (i *Inform) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case inform.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = uint64(value.Int64)
		case inform.FieldTenantID:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[j])
			} else if value.Valid {
				i.TenantID = value.String
			}
		case inform.FieldAccessOrgList:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_org_list", values[j])
			} else if value.Valid {
				i.AccessOrgList = value.String
			}
		case inform.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case inform.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case inform.FieldDeletedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[j])
			} else if value.Valid {
				i.DeletedAt = value.Time
			}
		case inform.FieldAlarmName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field alarm_name", values[j])
			} else if value.Valid {
				i.AlarmName = value.String
			}
		case inform.FieldAlarmType:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field alarm_type", values[j])
			} else if value.Valid {
				i.AlarmType = value.String
			}
		case inform.FieldSignName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sign_name", values[j])
			} else if value.Valid {
				i.SignName = value.String
			}
		case inform.FieldNotifyTemplate:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notify_template", values[j])
			} else if value.Valid {
				i.NotifyTemplate = value.String
			}
		case inform.FieldTemplateCode:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field template_code", values[j])
			} else if value.Valid {
				i.TemplateCode = value.String
			}
		case inform.FieldPhoneNumbers:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_numbers", values[j])
			} else if value.Valid {
				i.PhoneNumbers = value.String
			}
		case inform.FieldNotifySwitch:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notify_switch", values[j])
			} else if value.Valid {
				i.NotifySwitch = value.String
			}
		case inform.FieldTaskName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field task_name", values[j])
			} else if value.Valid {
				i.TaskName = value.String
			}
		case inform.FieldTaskID:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field task_id", values[j])
			} else if value.Valid {
				i.TaskID = uint64(value.Int64)
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Inform.
// This includes values selected through modifiers, order, etc.
func (i *Inform) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// Update returns a builder for updating this Inform.
// Note that you need to call Inform.Unwrap() before calling this method if this Inform
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Inform) Update() *InformUpdateOne {
	return NewInformClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Inform entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Inform) Unwrap() *Inform {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Inform is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Inform) String() string {
	var builder strings.Builder
	builder.WriteString("Inform(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("tenant_id=")
	builder.WriteString(i.TenantID)
	builder.WriteString(", ")
	builder.WriteString("access_org_list=")
	builder.WriteString(i.AccessOrgList)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(i.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("alarm_name=")
	builder.WriteString(i.AlarmName)
	builder.WriteString(", ")
	builder.WriteString("alarm_type=")
	builder.WriteString(i.AlarmType)
	builder.WriteString(", ")
	builder.WriteString("sign_name=")
	builder.WriteString(i.SignName)
	builder.WriteString(", ")
	builder.WriteString("notify_template=")
	builder.WriteString(i.NotifyTemplate)
	builder.WriteString(", ")
	builder.WriteString("template_code=")
	builder.WriteString(i.TemplateCode)
	builder.WriteString(", ")
	builder.WriteString("phone_numbers=")
	builder.WriteString(i.PhoneNumbers)
	builder.WriteString(", ")
	builder.WriteString("notify_switch=")
	builder.WriteString(i.NotifySwitch)
	builder.WriteString(", ")
	builder.WriteString("task_name=")
	builder.WriteString(i.TaskName)
	builder.WriteString(", ")
	builder.WriteString("task_id=")
	builder.WriteString(fmt.Sprintf("%v", i.TaskID))
	builder.WriteByte(')')
	return builder.String()
}

// Informs is a parsable slice of Inform.
type Informs []*Inform