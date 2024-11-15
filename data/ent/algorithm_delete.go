// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/algorithm"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/predicate"
)

// AlgorithmDelete is the builder for deleting a Algorithm entity.
type AlgorithmDelete struct {
	config
	hooks    []Hook
	mutation *AlgorithmMutation
}

// Where appends a list predicates to the AlgorithmDelete builder.
func (ad *AlgorithmDelete) Where(ps ...predicate.Algorithm) *AlgorithmDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AlgorithmDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AlgorithmDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AlgorithmDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(algorithm.Table, sqlgraph.NewFieldSpec(algorithm.FieldID, field.TypeUint64))
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// AlgorithmDeleteOne is the builder for deleting a single Algorithm entity.
type AlgorithmDeleteOne struct {
	ad *AlgorithmDelete
}

// Where appends a list predicates to the AlgorithmDelete builder.
func (ado *AlgorithmDeleteOne) Where(ps ...predicate.Algorithm) *AlgorithmDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *AlgorithmDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{algorithm.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AlgorithmDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}