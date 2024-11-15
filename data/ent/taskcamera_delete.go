// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/predicate"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/taskcamera"
)

// TaskCameraDelete is the builder for deleting a TaskCamera entity.
type TaskCameraDelete struct {
	config
	hooks    []Hook
	mutation *TaskCameraMutation
}

// Where appends a list predicates to the TaskCameraDelete builder.
func (tcd *TaskCameraDelete) Where(ps ...predicate.TaskCamera) *TaskCameraDelete {
	tcd.mutation.Where(ps...)
	return tcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tcd *TaskCameraDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tcd.sqlExec, tcd.mutation, tcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tcd *TaskCameraDelete) ExecX(ctx context.Context) int {
	n, err := tcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tcd *TaskCameraDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(taskcamera.Table, sqlgraph.NewFieldSpec(taskcamera.FieldID, field.TypeUint64))
	if ps := tcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tcd.mutation.done = true
	return affected, err
}

// TaskCameraDeleteOne is the builder for deleting a single TaskCamera entity.
type TaskCameraDeleteOne struct {
	tcd *TaskCameraDelete
}

// Where appends a list predicates to the TaskCameraDelete builder.
func (tcdo *TaskCameraDeleteOne) Where(ps ...predicate.TaskCamera) *TaskCameraDeleteOne {
	tcdo.tcd.mutation.Where(ps...)
	return tcdo
}

// Exec executes the deletion query.
func (tcdo *TaskCameraDeleteOne) Exec(ctx context.Context) error {
	n, err := tcdo.tcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{taskcamera.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tcdo *TaskCameraDeleteOne) ExecX(ctx context.Context) {
	if err := tcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
