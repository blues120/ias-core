// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/devicealgo"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/predicate"
)

// DeviceAlgoUpdate is the builder for updating DeviceAlgo entities.
type DeviceAlgoUpdate struct {
	config
	hooks     []Hook
	mutation  *DeviceAlgoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the DeviceAlgoUpdate builder.
func (dau *DeviceAlgoUpdate) Where(ps ...predicate.DeviceAlgo) *DeviceAlgoUpdate {
	dau.mutation.Where(ps...)
	return dau
}

// SetDeviceID sets the "device_id" field.
func (dau *DeviceAlgoUpdate) SetDeviceID(u uint64) *DeviceAlgoUpdate {
	dau.mutation.ResetDeviceID()
	dau.mutation.SetDeviceID(u)
	return dau
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (dau *DeviceAlgoUpdate) SetNillableDeviceID(u *uint64) *DeviceAlgoUpdate {
	if u != nil {
		dau.SetDeviceID(*u)
	}
	return dau
}

// AddDeviceID adds u to the "device_id" field.
func (dau *DeviceAlgoUpdate) AddDeviceID(u int64) *DeviceAlgoUpdate {
	dau.mutation.AddDeviceID(u)
	return dau
}

// SetAlgoGroupID sets the "algo_group_id" field.
func (dau *DeviceAlgoUpdate) SetAlgoGroupID(u uint) *DeviceAlgoUpdate {
	dau.mutation.ResetAlgoGroupID()
	dau.mutation.SetAlgoGroupID(u)
	return dau
}

// SetNillableAlgoGroupID sets the "algo_group_id" field if the given value is not nil.
func (dau *DeviceAlgoUpdate) SetNillableAlgoGroupID(u *uint) *DeviceAlgoUpdate {
	if u != nil {
		dau.SetAlgoGroupID(*u)
	}
	return dau
}

// AddAlgoGroupID adds u to the "algo_group_id" field.
func (dau *DeviceAlgoUpdate) AddAlgoGroupID(u int) *DeviceAlgoUpdate {
	dau.mutation.AddAlgoGroupID(u)
	return dau
}

// ClearAlgoGroupID clears the value of the "algo_group_id" field.
func (dau *DeviceAlgoUpdate) ClearAlgoGroupID() *DeviceAlgoUpdate {
	dau.mutation.ClearAlgoGroupID()
	return dau
}

// SetAlgoGroupName sets the "algo_group_name" field.
func (dau *DeviceAlgoUpdate) SetAlgoGroupName(s string) *DeviceAlgoUpdate {
	dau.mutation.SetAlgoGroupName(s)
	return dau
}

// SetNillableAlgoGroupName sets the "algo_group_name" field if the given value is not nil.
func (dau *DeviceAlgoUpdate) SetNillableAlgoGroupName(s *string) *DeviceAlgoUpdate {
	if s != nil {
		dau.SetAlgoGroupName(*s)
	}
	return dau
}

// ClearAlgoGroupName clears the value of the "algo_group_name" field.
func (dau *DeviceAlgoUpdate) ClearAlgoGroupName() *DeviceAlgoUpdate {
	dau.mutation.ClearAlgoGroupName()
	return dau
}

// SetAlgoGroupVersion sets the "algo_group_version" field.
func (dau *DeviceAlgoUpdate) SetAlgoGroupVersion(s string) *DeviceAlgoUpdate {
	dau.mutation.SetAlgoGroupVersion(s)
	return dau
}

// SetNillableAlgoGroupVersion sets the "algo_group_version" field if the given value is not nil.
func (dau *DeviceAlgoUpdate) SetNillableAlgoGroupVersion(s *string) *DeviceAlgoUpdate {
	if s != nil {
		dau.SetAlgoGroupVersion(*s)
	}
	return dau
}

// ClearAlgoGroupVersion clears the value of the "algo_group_version" field.
func (dau *DeviceAlgoUpdate) ClearAlgoGroupVersion() *DeviceAlgoUpdate {
	dau.mutation.ClearAlgoGroupVersion()
	return dau
}

// SetName sets the "name" field.
func (dau *DeviceAlgoUpdate) SetName(s string) *DeviceAlgoUpdate {
	dau.mutation.SetName(s)
	return dau
}

// SetNillableName sets the "name" field if the given value is not nil.
func (dau *DeviceAlgoUpdate) SetNillableName(s *string) *DeviceAlgoUpdate {
	if s != nil {
		dau.SetName(*s)
	}
	return dau
}

// SetVersion sets the "version" field.
func (dau *DeviceAlgoUpdate) SetVersion(s string) *DeviceAlgoUpdate {
	dau.mutation.SetVersion(s)
	return dau
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (dau *DeviceAlgoUpdate) SetNillableVersion(s *string) *DeviceAlgoUpdate {
	if s != nil {
		dau.SetVersion(*s)
	}
	return dau
}

// SetInstallTime sets the "install_time" field.
func (dau *DeviceAlgoUpdate) SetInstallTime(t time.Time) *DeviceAlgoUpdate {
	dau.mutation.SetInstallTime(t)
	return dau
}

// SetNillableInstallTime sets the "install_time" field if the given value is not nil.
func (dau *DeviceAlgoUpdate) SetNillableInstallTime(t *time.Time) *DeviceAlgoUpdate {
	if t != nil {
		dau.SetInstallTime(*t)
	}
	return dau
}

// Mutation returns the DeviceAlgoMutation object of the builder.
func (dau *DeviceAlgoUpdate) Mutation() *DeviceAlgoMutation {
	return dau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dau *DeviceAlgoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, dau.sqlSave, dau.mutation, dau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dau *DeviceAlgoUpdate) SaveX(ctx context.Context) int {
	affected, err := dau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dau *DeviceAlgoUpdate) Exec(ctx context.Context) error {
	_, err := dau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dau *DeviceAlgoUpdate) ExecX(ctx context.Context) {
	if err := dau.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (dau *DeviceAlgoUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DeviceAlgoUpdate {
	dau.modifiers = append(dau.modifiers, modifiers...)
	return dau
}

func (dau *DeviceAlgoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(devicealgo.Table, devicealgo.Columns, sqlgraph.NewFieldSpec(devicealgo.FieldID, field.TypeUint64))
	if ps := dau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dau.mutation.DeviceID(); ok {
		_spec.SetField(devicealgo.FieldDeviceID, field.TypeUint64, value)
	}
	if value, ok := dau.mutation.AddedDeviceID(); ok {
		_spec.AddField(devicealgo.FieldDeviceID, field.TypeUint64, value)
	}
	if value, ok := dau.mutation.AlgoGroupID(); ok {
		_spec.SetField(devicealgo.FieldAlgoGroupID, field.TypeUint, value)
	}
	if value, ok := dau.mutation.AddedAlgoGroupID(); ok {
		_spec.AddField(devicealgo.FieldAlgoGroupID, field.TypeUint, value)
	}
	if dau.mutation.AlgoGroupIDCleared() {
		_spec.ClearField(devicealgo.FieldAlgoGroupID, field.TypeUint)
	}
	if value, ok := dau.mutation.AlgoGroupName(); ok {
		_spec.SetField(devicealgo.FieldAlgoGroupName, field.TypeString, value)
	}
	if dau.mutation.AlgoGroupNameCleared() {
		_spec.ClearField(devicealgo.FieldAlgoGroupName, field.TypeString)
	}
	if value, ok := dau.mutation.AlgoGroupVersion(); ok {
		_spec.SetField(devicealgo.FieldAlgoGroupVersion, field.TypeString, value)
	}
	if dau.mutation.AlgoGroupVersionCleared() {
		_spec.ClearField(devicealgo.FieldAlgoGroupVersion, field.TypeString)
	}
	if value, ok := dau.mutation.Name(); ok {
		_spec.SetField(devicealgo.FieldName, field.TypeString, value)
	}
	if value, ok := dau.mutation.Version(); ok {
		_spec.SetField(devicealgo.FieldVersion, field.TypeString, value)
	}
	if value, ok := dau.mutation.InstallTime(); ok {
		_spec.SetField(devicealgo.FieldInstallTime, field.TypeTime, value)
	}
	_spec.AddModifiers(dau.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, dau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{devicealgo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dau.mutation.done = true
	return n, nil
}

// DeviceAlgoUpdateOne is the builder for updating a single DeviceAlgo entity.
type DeviceAlgoUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *DeviceAlgoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetDeviceID sets the "device_id" field.
func (dauo *DeviceAlgoUpdateOne) SetDeviceID(u uint64) *DeviceAlgoUpdateOne {
	dauo.mutation.ResetDeviceID()
	dauo.mutation.SetDeviceID(u)
	return dauo
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (dauo *DeviceAlgoUpdateOne) SetNillableDeviceID(u *uint64) *DeviceAlgoUpdateOne {
	if u != nil {
		dauo.SetDeviceID(*u)
	}
	return dauo
}

// AddDeviceID adds u to the "device_id" field.
func (dauo *DeviceAlgoUpdateOne) AddDeviceID(u int64) *DeviceAlgoUpdateOne {
	dauo.mutation.AddDeviceID(u)
	return dauo
}

// SetAlgoGroupID sets the "algo_group_id" field.
func (dauo *DeviceAlgoUpdateOne) SetAlgoGroupID(u uint) *DeviceAlgoUpdateOne {
	dauo.mutation.ResetAlgoGroupID()
	dauo.mutation.SetAlgoGroupID(u)
	return dauo
}

// SetNillableAlgoGroupID sets the "algo_group_id" field if the given value is not nil.
func (dauo *DeviceAlgoUpdateOne) SetNillableAlgoGroupID(u *uint) *DeviceAlgoUpdateOne {
	if u != nil {
		dauo.SetAlgoGroupID(*u)
	}
	return dauo
}

// AddAlgoGroupID adds u to the "algo_group_id" field.
func (dauo *DeviceAlgoUpdateOne) AddAlgoGroupID(u int) *DeviceAlgoUpdateOne {
	dauo.mutation.AddAlgoGroupID(u)
	return dauo
}

// ClearAlgoGroupID clears the value of the "algo_group_id" field.
func (dauo *DeviceAlgoUpdateOne) ClearAlgoGroupID() *DeviceAlgoUpdateOne {
	dauo.mutation.ClearAlgoGroupID()
	return dauo
}

// SetAlgoGroupName sets the "algo_group_name" field.
func (dauo *DeviceAlgoUpdateOne) SetAlgoGroupName(s string) *DeviceAlgoUpdateOne {
	dauo.mutation.SetAlgoGroupName(s)
	return dauo
}

// SetNillableAlgoGroupName sets the "algo_group_name" field if the given value is not nil.
func (dauo *DeviceAlgoUpdateOne) SetNillableAlgoGroupName(s *string) *DeviceAlgoUpdateOne {
	if s != nil {
		dauo.SetAlgoGroupName(*s)
	}
	return dauo
}

// ClearAlgoGroupName clears the value of the "algo_group_name" field.
func (dauo *DeviceAlgoUpdateOne) ClearAlgoGroupName() *DeviceAlgoUpdateOne {
	dauo.mutation.ClearAlgoGroupName()
	return dauo
}

// SetAlgoGroupVersion sets the "algo_group_version" field.
func (dauo *DeviceAlgoUpdateOne) SetAlgoGroupVersion(s string) *DeviceAlgoUpdateOne {
	dauo.mutation.SetAlgoGroupVersion(s)
	return dauo
}

// SetNillableAlgoGroupVersion sets the "algo_group_version" field if the given value is not nil.
func (dauo *DeviceAlgoUpdateOne) SetNillableAlgoGroupVersion(s *string) *DeviceAlgoUpdateOne {
	if s != nil {
		dauo.SetAlgoGroupVersion(*s)
	}
	return dauo
}

// ClearAlgoGroupVersion clears the value of the "algo_group_version" field.
func (dauo *DeviceAlgoUpdateOne) ClearAlgoGroupVersion() *DeviceAlgoUpdateOne {
	dauo.mutation.ClearAlgoGroupVersion()
	return dauo
}

// SetName sets the "name" field.
func (dauo *DeviceAlgoUpdateOne) SetName(s string) *DeviceAlgoUpdateOne {
	dauo.mutation.SetName(s)
	return dauo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (dauo *DeviceAlgoUpdateOne) SetNillableName(s *string) *DeviceAlgoUpdateOne {
	if s != nil {
		dauo.SetName(*s)
	}
	return dauo
}

// SetVersion sets the "version" field.
func (dauo *DeviceAlgoUpdateOne) SetVersion(s string) *DeviceAlgoUpdateOne {
	dauo.mutation.SetVersion(s)
	return dauo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (dauo *DeviceAlgoUpdateOne) SetNillableVersion(s *string) *DeviceAlgoUpdateOne {
	if s != nil {
		dauo.SetVersion(*s)
	}
	return dauo
}

// SetInstallTime sets the "install_time" field.
func (dauo *DeviceAlgoUpdateOne) SetInstallTime(t time.Time) *DeviceAlgoUpdateOne {
	dauo.mutation.SetInstallTime(t)
	return dauo
}

// SetNillableInstallTime sets the "install_time" field if the given value is not nil.
func (dauo *DeviceAlgoUpdateOne) SetNillableInstallTime(t *time.Time) *DeviceAlgoUpdateOne {
	if t != nil {
		dauo.SetInstallTime(*t)
	}
	return dauo
}

// Mutation returns the DeviceAlgoMutation object of the builder.
func (dauo *DeviceAlgoUpdateOne) Mutation() *DeviceAlgoMutation {
	return dauo.mutation
}

// Where appends a list predicates to the DeviceAlgoUpdate builder.
func (dauo *DeviceAlgoUpdateOne) Where(ps ...predicate.DeviceAlgo) *DeviceAlgoUpdateOne {
	dauo.mutation.Where(ps...)
	return dauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dauo *DeviceAlgoUpdateOne) Select(field string, fields ...string) *DeviceAlgoUpdateOne {
	dauo.fields = append([]string{field}, fields...)
	return dauo
}

// Save executes the query and returns the updated DeviceAlgo entity.
func (dauo *DeviceAlgoUpdateOne) Save(ctx context.Context) (*DeviceAlgo, error) {
	return withHooks(ctx, dauo.sqlSave, dauo.mutation, dauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dauo *DeviceAlgoUpdateOne) SaveX(ctx context.Context) *DeviceAlgo {
	node, err := dauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dauo *DeviceAlgoUpdateOne) Exec(ctx context.Context) error {
	_, err := dauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dauo *DeviceAlgoUpdateOne) ExecX(ctx context.Context) {
	if err := dauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (dauo *DeviceAlgoUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DeviceAlgoUpdateOne {
	dauo.modifiers = append(dauo.modifiers, modifiers...)
	return dauo
}

func (dauo *DeviceAlgoUpdateOne) sqlSave(ctx context.Context) (_node *DeviceAlgo, err error) {
	_spec := sqlgraph.NewUpdateSpec(devicealgo.Table, devicealgo.Columns, sqlgraph.NewFieldSpec(devicealgo.FieldID, field.TypeUint64))
	id, ok := dauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "DeviceAlgo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, devicealgo.FieldID)
		for _, f := range fields {
			if !devicealgo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != devicealgo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dauo.mutation.DeviceID(); ok {
		_spec.SetField(devicealgo.FieldDeviceID, field.TypeUint64, value)
	}
	if value, ok := dauo.mutation.AddedDeviceID(); ok {
		_spec.AddField(devicealgo.FieldDeviceID, field.TypeUint64, value)
	}
	if value, ok := dauo.mutation.AlgoGroupID(); ok {
		_spec.SetField(devicealgo.FieldAlgoGroupID, field.TypeUint, value)
	}
	if value, ok := dauo.mutation.AddedAlgoGroupID(); ok {
		_spec.AddField(devicealgo.FieldAlgoGroupID, field.TypeUint, value)
	}
	if dauo.mutation.AlgoGroupIDCleared() {
		_spec.ClearField(devicealgo.FieldAlgoGroupID, field.TypeUint)
	}
	if value, ok := dauo.mutation.AlgoGroupName(); ok {
		_spec.SetField(devicealgo.FieldAlgoGroupName, field.TypeString, value)
	}
	if dauo.mutation.AlgoGroupNameCleared() {
		_spec.ClearField(devicealgo.FieldAlgoGroupName, field.TypeString)
	}
	if value, ok := dauo.mutation.AlgoGroupVersion(); ok {
		_spec.SetField(devicealgo.FieldAlgoGroupVersion, field.TypeString, value)
	}
	if dauo.mutation.AlgoGroupVersionCleared() {
		_spec.ClearField(devicealgo.FieldAlgoGroupVersion, field.TypeString)
	}
	if value, ok := dauo.mutation.Name(); ok {
		_spec.SetField(devicealgo.FieldName, field.TypeString, value)
	}
	if value, ok := dauo.mutation.Version(); ok {
		_spec.SetField(devicealgo.FieldVersion, field.TypeString, value)
	}
	if value, ok := dauo.mutation.InstallTime(); ok {
		_spec.SetField(devicealgo.FieldInstallTime, field.TypeTime, value)
	}
	_spec.AddModifiers(dauo.modifiers...)
	_node = &DeviceAlgo{config: dauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{devicealgo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dauo.mutation.done = true
	return _node, nil
}