// Code generated by ent, DO NOT EDIT.

package inform

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the inform type in the database.
	Label = "inform"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldAccessOrgList holds the string denoting the access_org_list field in the database.
	FieldAccessOrgList = "access_org_list"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAlarmName holds the string denoting the alarm_name field in the database.
	FieldAlarmName = "alarm_name"
	// FieldAlarmType holds the string denoting the alarm_type field in the database.
	FieldAlarmType = "alarm_type"
	// FieldSignName holds the string denoting the sign_name field in the database.
	FieldSignName = "sign_name"
	// FieldNotifyTemplate holds the string denoting the notify_template field in the database.
	FieldNotifyTemplate = "notify_template"
	// FieldTemplateCode holds the string denoting the template_code field in the database.
	FieldTemplateCode = "template_code"
	// FieldPhoneNumbers holds the string denoting the phone_numbers field in the database.
	FieldPhoneNumbers = "phone_numbers"
	// FieldNotifySwitch holds the string denoting the notify_switch field in the database.
	FieldNotifySwitch = "notify_switch"
	// FieldTaskName holds the string denoting the task_name field in the database.
	FieldTaskName = "task_name"
	// FieldTaskID holds the string denoting the task_id field in the database.
	FieldTaskID = "task_id"
	// Table holds the table name of the inform in the database.
	Table = "inform"
)

// Columns holds all SQL columns for inform fields.
var Columns = []string{
	FieldID,
	FieldTenantID,
	FieldAccessOrgList,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAlarmName,
	FieldAlarmType,
	FieldSignName,
	FieldNotifyTemplate,
	FieldTemplateCode,
	FieldPhoneNumbers,
	FieldNotifySwitch,
	FieldTaskName,
	FieldTaskID,
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
//	import _ "gitlab.ctyuncdn.cn/ias/ias-core/data/ent/runtime"
var (
	Hooks        [4]ent.Hook
	Interceptors [2]ent.Interceptor
)

// OrderOption defines the ordering options for the Inform queries.
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

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByAlarmName orders the results by the alarm_name field.
func ByAlarmName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlarmName, opts...).ToFunc()
}

// ByAlarmType orders the results by the alarm_type field.
func ByAlarmType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlarmType, opts...).ToFunc()
}

// BySignName orders the results by the sign_name field.
func BySignName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSignName, opts...).ToFunc()
}

// ByNotifyTemplate orders the results by the notify_template field.
func ByNotifyTemplate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotifyTemplate, opts...).ToFunc()
}

// ByTemplateCode orders the results by the template_code field.
func ByTemplateCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTemplateCode, opts...).ToFunc()
}

// ByPhoneNumbers orders the results by the phone_numbers field.
func ByPhoneNumbers(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneNumbers, opts...).ToFunc()
}

// ByNotifySwitch orders the results by the notify_switch field.
func ByNotifySwitch(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotifySwitch, opts...).ToFunc()
}

// ByTaskName orders the results by the task_name field.
func ByTaskName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTaskName, opts...).ToFunc()
}

// ByTaskID orders the results by the task_id field.
func ByTaskID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTaskID, opts...).ToFunc()
}