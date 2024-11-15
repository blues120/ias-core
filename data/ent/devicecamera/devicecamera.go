// Code generated by ent, DO NOT EDIT.

package devicecamera

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the devicecamera type in the database.
	Label = "device_camera"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeviceID holds the string denoting the device_id field in the database.
	FieldDeviceID = "device_id"
	// FieldCameraID holds the string denoting the camera_id field in the database.
	FieldCameraID = "camera_id"
	// EdgeCamera holds the string denoting the camera edge name in mutations.
	EdgeCamera = "camera"
	// EdgeDevice holds the string denoting the device edge name in mutations.
	EdgeDevice = "device"
	// Table holds the table name of the devicecamera in the database.
	Table = "device_camera"
	// CameraTable is the table that holds the camera relation/edge.
	CameraTable = "device_camera"
	// CameraInverseTable is the table name for the Camera entity.
	// It exists in this package in order to avoid circular dependency with the "camera" package.
	CameraInverseTable = "camera"
	// CameraColumn is the table column denoting the camera relation/edge.
	CameraColumn = "camera_id"
	// DeviceTable is the table that holds the device relation/edge.
	DeviceTable = "device_camera"
	// DeviceInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	DeviceInverseTable = "task"
	// DeviceColumn is the table column denoting the device relation/edge.
	DeviceColumn = "device_id"
)

// Columns holds all SQL columns for devicecamera fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeviceID,
	FieldCameraID,
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

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the DeviceCamera queries.
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

// ByDeviceID orders the results by the device_id field.
func ByDeviceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeviceID, opts...).ToFunc()
}

// ByCameraID orders the results by the camera_id field.
func ByCameraID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCameraID, opts...).ToFunc()
}

// ByCameraField orders the results by camera field.
func ByCameraField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCameraStep(), sql.OrderByField(field, opts...))
	}
}

// ByDeviceField orders the results by device field.
func ByDeviceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeviceStep(), sql.OrderByField(field, opts...))
	}
}
func newCameraStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CameraInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CameraTable, CameraColumn),
	)
}
func newDeviceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeviceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DeviceTable, DeviceColumn),
	)
}
