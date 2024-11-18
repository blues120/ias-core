// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/blues120/ias-core/data/ent/upplatform"
)

// UpPlatform is the model entity for the UpPlatform schema.
type UpPlatform struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID string `json:"tenant_id,omitempty"`
	// 授权的组织 id 列表，#分隔
	AccessOrgList string `json:"access_org_list,omitempty"`
	// SipId
	SipID string `json:"sip_id,omitempty"`
	// SipDomain
	SipDomain string `json:"sip_domain,omitempty"`
	// SipIp
	SipIP string `json:"sip_ip,omitempty"`
	// SipPort
	SipPort int32 `json:"sip_port,omitempty"`
	// SipUser
	SipUser string `json:"sip_user,omitempty"`
	// SipPassword
	SipPassword string `json:"sip_password,omitempty"`
	// Description
	Description string `json:"description,omitempty"`
	// HeartbeatInterval
	HeartbeatInterval int32 `json:"heartbeat_interval,omitempty"`
	// RegisterInterval
	RegisterInterval int32 `json:"register_interval,omitempty"`
	// TransType
	TransType string `json:"trans_type,omitempty"`
	// GbId
	GBID string `json:"gb_id,omitempty"`
	// Cascadestatus
	Cascadestatus string `json:"cascadestatus,omitempty"`
	// RegistrationStatus
	RegistrationStatus string `json:"registration_status,omitempty"`
	selectValues       sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UpPlatform) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case upplatform.FieldID, upplatform.FieldSipPort, upplatform.FieldHeartbeatInterval, upplatform.FieldRegisterInterval:
			values[i] = new(sql.NullInt64)
		case upplatform.FieldTenantID, upplatform.FieldAccessOrgList, upplatform.FieldSipID, upplatform.FieldSipDomain, upplatform.FieldSipIP, upplatform.FieldSipUser, upplatform.FieldSipPassword, upplatform.FieldDescription, upplatform.FieldTransType, upplatform.FieldGBID, upplatform.FieldCascadestatus, upplatform.FieldRegistrationStatus:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UpPlatform fields.
func (up *UpPlatform) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case upplatform.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			up.ID = int(value.Int64)
		case upplatform.FieldTenantID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				up.TenantID = value.String
			}
		case upplatform.FieldAccessOrgList:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field access_org_list", values[i])
			} else if value.Valid {
				up.AccessOrgList = value.String
			}
		case upplatform.FieldSipID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sip_id", values[i])
			} else if value.Valid {
				up.SipID = value.String
			}
		case upplatform.FieldSipDomain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sip_domain", values[i])
			} else if value.Valid {
				up.SipDomain = value.String
			}
		case upplatform.FieldSipIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sip_ip", values[i])
			} else if value.Valid {
				up.SipIP = value.String
			}
		case upplatform.FieldSipPort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sip_port", values[i])
			} else if value.Valid {
				up.SipPort = int32(value.Int64)
			}
		case upplatform.FieldSipUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sip_user", values[i])
			} else if value.Valid {
				up.SipUser = value.String
			}
		case upplatform.FieldSipPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sip_password", values[i])
			} else if value.Valid {
				up.SipPassword = value.String
			}
		case upplatform.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				up.Description = value.String
			}
		case upplatform.FieldHeartbeatInterval:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field heartbeat_interval", values[i])
			} else if value.Valid {
				up.HeartbeatInterval = int32(value.Int64)
			}
		case upplatform.FieldRegisterInterval:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field register_interval", values[i])
			} else if value.Valid {
				up.RegisterInterval = int32(value.Int64)
			}
		case upplatform.FieldTransType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field trans_type", values[i])
			} else if value.Valid {
				up.TransType = value.String
			}
		case upplatform.FieldGBID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gb_id", values[i])
			} else if value.Valid {
				up.GBID = value.String
			}
		case upplatform.FieldCascadestatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cascadestatus", values[i])
			} else if value.Valid {
				up.Cascadestatus = value.String
			}
		case upplatform.FieldRegistrationStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field registration_status", values[i])
			} else if value.Valid {
				up.RegistrationStatus = value.String
			}
		default:
			up.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UpPlatform.
// This includes values selected through modifiers, order, etc.
func (up *UpPlatform) Value(name string) (ent.Value, error) {
	return up.selectValues.Get(name)
}

// Update returns a builder for updating this UpPlatform.
// Note that you need to call UpPlatform.Unwrap() before calling this method if this UpPlatform
// was returned from a transaction, and the transaction was committed or rolled back.
func (up *UpPlatform) Update() *UpPlatformUpdateOne {
	return NewUpPlatformClient(up.config).UpdateOne(up)
}

// Unwrap unwraps the UpPlatform entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (up *UpPlatform) Unwrap() *UpPlatform {
	_tx, ok := up.config.driver.(*txDriver)
	if !ok {
		panic("ent: UpPlatform is not a transactional entity")
	}
	up.config.driver = _tx.drv
	return up
}

// String implements the fmt.Stringer.
func (up *UpPlatform) String() string {
	var builder strings.Builder
	builder.WriteString("UpPlatform(")
	builder.WriteString(fmt.Sprintf("id=%v, ", up.ID))
	builder.WriteString("tenant_id=")
	builder.WriteString(up.TenantID)
	builder.WriteString(", ")
	builder.WriteString("access_org_list=")
	builder.WriteString(up.AccessOrgList)
	builder.WriteString(", ")
	builder.WriteString("sip_id=")
	builder.WriteString(up.SipID)
	builder.WriteString(", ")
	builder.WriteString("sip_domain=")
	builder.WriteString(up.SipDomain)
	builder.WriteString(", ")
	builder.WriteString("sip_ip=")
	builder.WriteString(up.SipIP)
	builder.WriteString(", ")
	builder.WriteString("sip_port=")
	builder.WriteString(fmt.Sprintf("%v", up.SipPort))
	builder.WriteString(", ")
	builder.WriteString("sip_user=")
	builder.WriteString(up.SipUser)
	builder.WriteString(", ")
	builder.WriteString("sip_password=")
	builder.WriteString(up.SipPassword)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(up.Description)
	builder.WriteString(", ")
	builder.WriteString("heartbeat_interval=")
	builder.WriteString(fmt.Sprintf("%v", up.HeartbeatInterval))
	builder.WriteString(", ")
	builder.WriteString("register_interval=")
	builder.WriteString(fmt.Sprintf("%v", up.RegisterInterval))
	builder.WriteString(", ")
	builder.WriteString("trans_type=")
	builder.WriteString(up.TransType)
	builder.WriteString(", ")
	builder.WriteString("gb_id=")
	builder.WriteString(up.GBID)
	builder.WriteString(", ")
	builder.WriteString("cascadestatus=")
	builder.WriteString(up.Cascadestatus)
	builder.WriteString(", ")
	builder.WriteString("registration_status=")
	builder.WriteString(up.RegistrationStatus)
	builder.WriteByte(')')
	return builder.String()
}

// UpPlatforms is a parsable slice of UpPlatform.
type UpPlatforms []*UpPlatform
