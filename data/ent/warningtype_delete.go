// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/blues120/ias-core/data/ent/predicate"
	"github.com/blues120/ias-core/data/ent/warningtype"
)

// WarningTypeDelete is the builder for deleting a WarningType entity.
type WarningTypeDelete struct {
	config
	hooks    []Hook
	mutation *WarningTypeMutation
}

// Where appends a list predicates to the WarningTypeDelete builder.
func (wtd *WarningTypeDelete) Where(ps ...predicate.WarningType) *WarningTypeDelete {
	wtd.mutation.Where(ps...)
	return wtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wtd *WarningTypeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wtd.sqlExec, wtd.mutation, wtd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wtd *WarningTypeDelete) ExecX(ctx context.Context) int {
	n, err := wtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wtd *WarningTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(warningtype.Table, sqlgraph.NewFieldSpec(warningtype.FieldID, field.TypeUint64))
	if ps := wtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wtd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wtd.mutation.done = true
	return affected, err
}

// WarningTypeDeleteOne is the builder for deleting a single WarningType entity.
type WarningTypeDeleteOne struct {
	wtd *WarningTypeDelete
}

// Where appends a list predicates to the WarningTypeDelete builder.
func (wtdo *WarningTypeDeleteOne) Where(ps ...predicate.WarningType) *WarningTypeDeleteOne {
	wtdo.wtd.mutation.Where(ps...)
	return wtdo
}

// Exec executes the deletion query.
func (wtdo *WarningTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := wtdo.wtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{warningtype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wtdo *WarningTypeDeleteOne) ExecX(ctx context.Context) {
	if err := wtdo.Exec(ctx); err != nil {
		panic(err)
	}
}
