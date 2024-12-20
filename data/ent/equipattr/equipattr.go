// Code generated by ent, DO NOT EDIT.

package equipattr

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the equipattr type in the database.
	Label = "equip_attr"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAttrKey holds the string denoting the attr_key field in the database.
	FieldAttrKey = "attr_key"
	// FieldAttrValue holds the string denoting the attr_value field in the database.
	FieldAttrValue = "attr_value"
	// FieldExtend holds the string denoting the extend field in the database.
	FieldExtend = "extend"
	// Table holds the table name of the equipattr in the database.
	Table = "equip_attr"
)

// Columns holds all SQL columns for equipattr fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAttrKey,
	FieldAttrValue,
	FieldExtend,
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
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the EquipAttr queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
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

// ByAttrKey orders the results by the attr_key field.
func ByAttrKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAttrKey, opts...).ToFunc()
}

// ByAttrValue orders the results by the attr_value field.
func ByAttrValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAttrValue, opts...).ToFunc()
}

// ByExtend orders the results by the extend field.
func ByExtend(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExtend, opts...).ToFunc()
}
