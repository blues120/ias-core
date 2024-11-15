// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/devicecamera"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/predicate"
)

// DeviceCameraDelete is the builder for deleting a DeviceCamera entity.
type DeviceCameraDelete struct {
	config
	hooks    []Hook
	mutation *DeviceCameraMutation
}

// Where appends a list predicates to the DeviceCameraDelete builder.
func (dcd *DeviceCameraDelete) Where(ps ...predicate.DeviceCamera) *DeviceCameraDelete {
	dcd.mutation.Where(ps...)
	return dcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dcd *DeviceCameraDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dcd.sqlExec, dcd.mutation, dcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dcd *DeviceCameraDelete) ExecX(ctx context.Context) int {
	n, err := dcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dcd *DeviceCameraDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(devicecamera.Table, sqlgraph.NewFieldSpec(devicecamera.FieldID, field.TypeUint64))
	if ps := dcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dcd.mutation.done = true
	return affected, err
}

// DeviceCameraDeleteOne is the builder for deleting a single DeviceCamera entity.
type DeviceCameraDeleteOne struct {
	dcd *DeviceCameraDelete
}

// Where appends a list predicates to the DeviceCameraDelete builder.
func (dcdo *DeviceCameraDeleteOne) Where(ps ...predicate.DeviceCamera) *DeviceCameraDeleteOne {
	dcdo.dcd.mutation.Where(ps...)
	return dcdo
}

// Exec executes the deletion query.
func (dcdo *DeviceCameraDeleteOne) Exec(ctx context.Context) error {
	n, err := dcdo.dcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{devicecamera.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dcdo *DeviceCameraDeleteOne) ExecX(ctx context.Context) {
	if err := dcdo.Exec(ctx); err != nil {
		panic(err)
	}
}