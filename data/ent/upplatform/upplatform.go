// Code generated by ent, DO NOT EDIT.

package upplatform

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the upplatform type in the database.
	Label = "up_platform"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldAccessOrgList holds the string denoting the access_org_list field in the database.
	FieldAccessOrgList = "access_org_list"
	// FieldSipID holds the string denoting the sip_id field in the database.
	FieldSipID = "sip_id"
	// FieldSipDomain holds the string denoting the sip_domain field in the database.
	FieldSipDomain = "sip_domain"
	// FieldSipIP holds the string denoting the sip_ip field in the database.
	FieldSipIP = "sip_ip"
	// FieldSipPort holds the string denoting the sip_port field in the database.
	FieldSipPort = "sip_port"
	// FieldSipUser holds the string denoting the sip_user field in the database.
	FieldSipUser = "sip_user"
	// FieldSipPassword holds the string denoting the sip_password field in the database.
	FieldSipPassword = "sip_password"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldHeartbeatInterval holds the string denoting the heartbeat_interval field in the database.
	FieldHeartbeatInterval = "heartbeat_interval"
	// FieldRegisterInterval holds the string denoting the register_interval field in the database.
	FieldRegisterInterval = "register_interval"
	// FieldTransType holds the string denoting the trans_type field in the database.
	FieldTransType = "trans_type"
	// FieldGBID holds the string denoting the gb_id field in the database.
	FieldGBID = "gb_id"
	// FieldCascadestatus holds the string denoting the cascadestatus field in the database.
	FieldCascadestatus = "cascadestatus"
	// FieldRegistrationStatus holds the string denoting the registration_status field in the database.
	FieldRegistrationStatus = "registration_status"
	// Table holds the table name of the upplatform in the database.
	Table = "up_platforms"
)

// Columns holds all SQL columns for upplatform fields.
var Columns = []string{
	FieldID,
	FieldTenantID,
	FieldAccessOrgList,
	FieldSipID,
	FieldSipDomain,
	FieldSipIP,
	FieldSipPort,
	FieldSipUser,
	FieldSipPassword,
	FieldDescription,
	FieldHeartbeatInterval,
	FieldRegisterInterval,
	FieldTransType,
	FieldGBID,
	FieldCascadestatus,
	FieldRegistrationStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/blues120/ias-core/data/ent/runtime"
var (
	Hooks        [4]ent.Hook
	Interceptors [2]ent.Interceptor
)

// OrderOption defines the ordering options for the UpPlatform queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByAccessOrgList orders the results by the access_org_list field.
func ByAccessOrgList(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccessOrgList, opts...).ToFunc()
}

// BySipID orders the results by the sip_id field.
func BySipID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSipID, opts...).ToFunc()
}

// BySipDomain orders the results by the sip_domain field.
func BySipDomain(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSipDomain, opts...).ToFunc()
}

// BySipIP orders the results by the sip_ip field.
func BySipIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSipIP, opts...).ToFunc()
}

// BySipPort orders the results by the sip_port field.
func BySipPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSipPort, opts...).ToFunc()
}

// BySipUser orders the results by the sip_user field.
func BySipUser(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSipUser, opts...).ToFunc()
}

// BySipPassword orders the results by the sip_password field.
func BySipPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSipPassword, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByHeartbeatInterval orders the results by the heartbeat_interval field.
func ByHeartbeatInterval(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeartbeatInterval, opts...).ToFunc()
}

// ByRegisterInterval orders the results by the register_interval field.
func ByRegisterInterval(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegisterInterval, opts...).ToFunc()
}

// ByTransType orders the results by the trans_type field.
func ByTransType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTransType, opts...).ToFunc()
}

// ByGBID orders the results by the gb_id field.
func ByGBID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGBID, opts...).ToFunc()
}

// ByCascadestatus orders the results by the cascadestatus field.
func ByCascadestatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCascadestatus, opts...).ToFunc()
}

// ByRegistrationStatus orders the results by the registration_status field.
func ByRegistrationStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegistrationStatus, opts...).ToFunc()
}
