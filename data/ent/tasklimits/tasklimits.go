// Code generated by ent, DO NOT EDIT.

package tasklimits

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the tasklimits type in the database.
	Label = "task_limits"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldModel holds the string denoting the model field in the database.
	FieldModel = "model"
	// FieldMaxCameraNum holds the string denoting the maxcameranum field in the database.
	FieldMaxCameraNum = "max_camera_num"
	// FieldAlgoNum holds the string denoting the algonum field in the database.
	FieldAlgoNum = "algo_num"
	// FieldMaxSubTaskNum holds the string denoting the maxsubtasknum field in the database.
	FieldMaxSubTaskNum = "max_sub_task_num"
	// Table holds the table name of the tasklimits in the database.
	Table = "task_limits"
)

// Columns holds all SQL columns for tasklimits fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldModel,
	FieldMaxCameraNum,
	FieldAlgoNum,
	FieldMaxSubTaskNum,
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

// OrderOption defines the ordering options for the TaskLimits queries.
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

// ByModel orders the results by the model field.
func ByModel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModel, opts...).ToFunc()
}

// ByMaxCameraNum orders the results by the maxCameraNum field.
func ByMaxCameraNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMaxCameraNum, opts...).ToFunc()
}

// ByAlgoNum orders the results by the algoNum field.
func ByAlgoNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlgoNum, opts...).ToFunc()
}

// ByMaxSubTaskNum orders the results by the maxSubTaskNum field.
func ByMaxSubTaskNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMaxSubTaskNum, opts...).ToFunc()
}
