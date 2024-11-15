// Code generated by ent, DO NOT EDIT.

package algorithm

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the algorithm type in the database.
	Label = "algorithm"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldAppName holds the string denoting the app_name field in the database.
	FieldAppName = "app_name"
	// FieldAlarmType holds the string denoting the alarm_type field in the database.
	FieldAlarmType = "alarm_type"
	// FieldAlarmName holds the string denoting the alarm_name field in the database.
	FieldAlarmName = "alarm_name"
	// FieldNotify holds the string denoting the notify field in the database.
	FieldNotify = "notify"
	// FieldExtend holds the string denoting the extend field in the database.
	FieldExtend = "extend"
	// FieldDrawType holds the string denoting the draw_type field in the database.
	FieldDrawType = "draw_type"
	// FieldBaseType holds the string denoting the base_type field in the database.
	FieldBaseType = "base_type"
	// FieldAvailable holds the string denoting the available field in the database.
	FieldAvailable = "available"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldLabelMap holds the string denoting the label_map field in the database.
	FieldLabelMap = "label_map"
	// FieldTarget holds the string denoting the target field in the database.
	FieldTarget = "target"
	// FieldAlgoNameEn holds the string denoting the algo_name_en field in the database.
	FieldAlgoNameEn = "algo_name_en"
	// FieldAlgoGroupID holds the string denoting the algo_group_id field in the database.
	FieldAlgoGroupID = "algo_group_id"
	// FieldAlgoGroupName holds the string denoting the algo_group_name field in the database.
	FieldAlgoGroupName = "algo_group_name"
	// FieldAlgoGroupVersion holds the string denoting the algo_group_version field in the database.
	FieldAlgoGroupVersion = "algo_group_version"
	// FieldConfig holds the string denoting the config field in the database.
	FieldConfig = "config"
	// FieldProvider holds the string denoting the provider field in the database.
	FieldProvider = "provider"
	// FieldAlgoID holds the string denoting the algo_id field in the database.
	FieldAlgoID = "algo_id"
	// FieldPlatform holds the string denoting the platform field in the database.
	FieldPlatform = "platform"
	// FieldDeviceModel holds the string denoting the device_model field in the database.
	FieldDeviceModel = "device_model"
	// FieldIsGroupType holds the string denoting the is_group_type field in the database.
	FieldIsGroupType = "is_group_type"
	// FieldPrefix holds the string denoting the prefix field in the database.
	FieldPrefix = "prefix"
	// EdgeTasks holds the string denoting the tasks edge name in mutations.
	EdgeTasks = "tasks"
	// Table holds the table name of the algorithm in the database.
	Table = "algorithm"
	// TasksTable is the table that holds the tasks relation/edge.
	TasksTable = "task"
	// TasksInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TasksInverseTable = "task"
	// TasksColumn is the table column denoting the tasks relation/edge.
	TasksColumn = "algo_id"
)

// Columns holds all SQL columns for algorithm fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldType,
	FieldDescription,
	FieldVersion,
	FieldAppName,
	FieldAlarmType,
	FieldAlarmName,
	FieldNotify,
	FieldExtend,
	FieldDrawType,
	FieldBaseType,
	FieldAvailable,
	FieldImage,
	FieldLabelMap,
	FieldTarget,
	FieldAlgoNameEn,
	FieldAlgoGroupID,
	FieldAlgoGroupName,
	FieldAlgoGroupVersion,
	FieldConfig,
	FieldProvider,
	FieldAlgoID,
	FieldPlatform,
	FieldDeviceModel,
	FieldIsGroupType,
	FieldPrefix,
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
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultAvailable holds the default value on creation for the "available" field.
	DefaultAvailable uint
	// ImageValidator is a validator for the "image" field. It is called by the builders before save.
	ImageValidator func(string) error
	// LabelMapValidator is a validator for the "label_map" field. It is called by the builders before save.
	LabelMapValidator func(string) error
	// DefaultIsGroupType holds the default value on creation for the "is_group_type" field.
	DefaultIsGroupType uint
)

// OrderOption defines the ordering options for the Algorithm queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByAppName orders the results by the app_name field.
func ByAppName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAppName, opts...).ToFunc()
}

// ByAlarmType orders the results by the alarm_type field.
func ByAlarmType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlarmType, opts...).ToFunc()
}

// ByAlarmName orders the results by the alarm_name field.
func ByAlarmName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlarmName, opts...).ToFunc()
}

// ByNotify orders the results by the notify field.
func ByNotify(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotify, opts...).ToFunc()
}

// ByDrawType orders the results by the draw_type field.
func ByDrawType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDrawType, opts...).ToFunc()
}

// ByBaseType orders the results by the base_type field.
func ByBaseType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBaseType, opts...).ToFunc()
}

// ByAvailable orders the results by the available field.
func ByAvailable(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvailable, opts...).ToFunc()
}

// ByImage orders the results by the image field.
func ByImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImage, opts...).ToFunc()
}

// ByLabelMap orders the results by the label_map field.
func ByLabelMap(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLabelMap, opts...).ToFunc()
}

// ByTarget orders the results by the target field.
func ByTarget(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTarget, opts...).ToFunc()
}

// ByAlgoNameEn orders the results by the algo_name_en field.
func ByAlgoNameEn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlgoNameEn, opts...).ToFunc()
}

// ByAlgoGroupID orders the results by the algo_group_id field.
func ByAlgoGroupID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlgoGroupID, opts...).ToFunc()
}

// ByAlgoGroupName orders the results by the algo_group_name field.
func ByAlgoGroupName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlgoGroupName, opts...).ToFunc()
}

// ByAlgoGroupVersion orders the results by the algo_group_version field.
func ByAlgoGroupVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlgoGroupVersion, opts...).ToFunc()
}

// ByConfig orders the results by the config field.
func ByConfig(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConfig, opts...).ToFunc()
}

// ByProvider orders the results by the provider field.
func ByProvider(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProvider, opts...).ToFunc()
}

// ByAlgoID orders the results by the algo_id field.
func ByAlgoID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlgoID, opts...).ToFunc()
}

// ByPlatform orders the results by the platform field.
func ByPlatform(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlatform, opts...).ToFunc()
}

// ByDeviceModel orders the results by the device_model field.
func ByDeviceModel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeviceModel, opts...).ToFunc()
}

// ByIsGroupType orders the results by the is_group_type field.
func ByIsGroupType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsGroupType, opts...).ToFunc()
}

// ByPrefix orders the results by the prefix field.
func ByPrefix(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrefix, opts...).ToFunc()
}

// ByTasksCount orders the results by tasks count.
func ByTasksCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTasksStep(), opts...)
	}
}

// ByTasks orders the results by tasks terms.
func ByTasks(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTasksStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTasksStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TasksInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TasksTable, TasksColumn),
	)
}
