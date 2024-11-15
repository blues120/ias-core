// Code generated by ent, DO NOT EDIT.

package area

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the area type in the database.
	Label = "area"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// FieldPid holds the string denoting the pid field in the database.
	FieldPid = "pid"
	// Table holds the table name of the area in the database.
	Table = "area"
)

// Columns holds all SQL columns for area fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldLevel,
	FieldPid,
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

// OrderOption defines the ordering options for the Area queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByLevel orders the results by the level field.
func ByLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLevel, opts...).ToFunc()
}

// ByPid orders the results by the pid field.
func ByPid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPid, opts...).ToFunc()
}