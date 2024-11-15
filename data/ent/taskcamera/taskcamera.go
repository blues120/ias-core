// Code generated by ent, DO NOT EDIT.

package taskcamera

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the taskcamera type in the database.
	Label = "task_camera"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldAccessOrgList holds the string denoting the access_org_list field in the database.
	FieldAccessOrgList = "access_org_list"
	// FieldTaskID holds the string denoting the task_id field in the database.
	FieldTaskID = "task_id"
	// FieldCameraID holds the string denoting the camera_id field in the database.
	FieldCameraID = "camera_id"
	// FieldMultiImgBox holds the string denoting the multi_img_box field in the database.
	FieldMultiImgBox = "multi_img_box"
	// EdgeCamera holds the string denoting the camera edge name in mutations.
	EdgeCamera = "camera"
	// EdgeTask holds the string denoting the task edge name in mutations.
	EdgeTask = "task"
	// Table holds the table name of the taskcamera in the database.
	Table = "task_camera"
	// CameraTable is the table that holds the camera relation/edge.
	CameraTable = "task_camera"
	// CameraInverseTable is the table name for the Camera entity.
	// It exists in this package in order to avoid circular dependency with the "camera" package.
	CameraInverseTable = "camera"
	// CameraColumn is the table column denoting the camera relation/edge.
	CameraColumn = "camera_id"
	// TaskTable is the table that holds the task relation/edge.
	TaskTable = "task_camera"
	// TaskInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskInverseTable = "task"
	// TaskColumn is the table column denoting the task relation/edge.
	TaskColumn = "task_id"
)

// Columns holds all SQL columns for taskcamera fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTenantID,
	FieldAccessOrgList,
	FieldTaskID,
	FieldCameraID,
	FieldMultiImgBox,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the TaskCamera queries.
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

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByAccessOrgList orders the results by the access_org_list field.
func ByAccessOrgList(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccessOrgList, opts...).ToFunc()
}

// ByTaskID orders the results by the task_id field.
func ByTaskID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTaskID, opts...).ToFunc()
}

// ByCameraID orders the results by the camera_id field.
func ByCameraID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCameraID, opts...).ToFunc()
}

// ByMultiImgBox orders the results by the multi_img_box field.
func ByMultiImgBox(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMultiImgBox, opts...).ToFunc()
}

// ByCameraField orders the results by camera field.
func ByCameraField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCameraStep(), sql.OrderByField(field, opts...))
	}
}

// ByTaskField orders the results by task field.
func ByTaskField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTaskStep(), sql.OrderByField(field, opts...))
	}
}
func newCameraStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CameraInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CameraTable, CameraColumn),
	)
}
func newTaskStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TaskInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, TaskTable, TaskColumn),
	)
}