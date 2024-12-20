// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/blues120/ias-core/data/ent/area"
	"github.com/blues120/ias-core/data/ent/predicate"
)

// AreaUpdate is the builder for updating Area entities.
type AreaUpdate struct {
	config
	hooks     []Hook
	mutation  *AreaMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AreaUpdate builder.
func (au *AreaUpdate) Where(ps ...predicate.Area) *AreaUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AreaUpdate) SetName(s string) *AreaUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *AreaUpdate) SetNillableName(s *string) *AreaUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetLevel sets the "level" field.
func (au *AreaUpdate) SetLevel(u uint64) *AreaUpdate {
	au.mutation.ResetLevel()
	au.mutation.SetLevel(u)
	return au
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (au *AreaUpdate) SetNillableLevel(u *uint64) *AreaUpdate {
	if u != nil {
		au.SetLevel(*u)
	}
	return au
}

// AddLevel adds u to the "level" field.
func (au *AreaUpdate) AddLevel(u int64) *AreaUpdate {
	au.mutation.AddLevel(u)
	return au
}

// SetPid sets the "pid" field.
func (au *AreaUpdate) SetPid(i int64) *AreaUpdate {
	au.mutation.ResetPid()
	au.mutation.SetPid(i)
	return au
}

// SetNillablePid sets the "pid" field if the given value is not nil.
func (au *AreaUpdate) SetNillablePid(i *int64) *AreaUpdate {
	if i != nil {
		au.SetPid(*i)
	}
	return au
}

// AddPid adds i to the "pid" field.
func (au *AreaUpdate) AddPid(i int64) *AreaUpdate {
	au.mutation.AddPid(i)
	return au
}

// Mutation returns the AreaMutation object of the builder.
func (au *AreaUpdate) Mutation() *AreaMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AreaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AreaUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AreaUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AreaUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (au *AreaUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AreaUpdate {
	au.modifiers = append(au.modifiers, modifiers...)
	return au
}

func (au *AreaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(area.Table, area.Columns, sqlgraph.NewFieldSpec(area.FieldID, field.TypeUint64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(area.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.Level(); ok {
		_spec.SetField(area.FieldLevel, field.TypeUint64, value)
	}
	if value, ok := au.mutation.AddedLevel(); ok {
		_spec.AddField(area.FieldLevel, field.TypeUint64, value)
	}
	if value, ok := au.mutation.Pid(); ok {
		_spec.SetField(area.FieldPid, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedPid(); ok {
		_spec.AddField(area.FieldPid, field.TypeInt64, value)
	}
	_spec.AddModifiers(au.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{area.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AreaUpdateOne is the builder for updating a single Area entity.
type AreaUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AreaMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (auo *AreaUpdateOne) SetName(s string) *AreaUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *AreaUpdateOne) SetNillableName(s *string) *AreaUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetLevel sets the "level" field.
func (auo *AreaUpdateOne) SetLevel(u uint64) *AreaUpdateOne {
	auo.mutation.ResetLevel()
	auo.mutation.SetLevel(u)
	return auo
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (auo *AreaUpdateOne) SetNillableLevel(u *uint64) *AreaUpdateOne {
	if u != nil {
		auo.SetLevel(*u)
	}
	return auo
}

// AddLevel adds u to the "level" field.
func (auo *AreaUpdateOne) AddLevel(u int64) *AreaUpdateOne {
	auo.mutation.AddLevel(u)
	return auo
}

// SetPid sets the "pid" field.
func (auo *AreaUpdateOne) SetPid(i int64) *AreaUpdateOne {
	auo.mutation.ResetPid()
	auo.mutation.SetPid(i)
	return auo
}

// SetNillablePid sets the "pid" field if the given value is not nil.
func (auo *AreaUpdateOne) SetNillablePid(i *int64) *AreaUpdateOne {
	if i != nil {
		auo.SetPid(*i)
	}
	return auo
}

// AddPid adds i to the "pid" field.
func (auo *AreaUpdateOne) AddPid(i int64) *AreaUpdateOne {
	auo.mutation.AddPid(i)
	return auo
}

// Mutation returns the AreaMutation object of the builder.
func (auo *AreaUpdateOne) Mutation() *AreaMutation {
	return auo.mutation
}

// Where appends a list predicates to the AreaUpdate builder.
func (auo *AreaUpdateOne) Where(ps ...predicate.Area) *AreaUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AreaUpdateOne) Select(field string, fields ...string) *AreaUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Area entity.
func (auo *AreaUpdateOne) Save(ctx context.Context) (*Area, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AreaUpdateOne) SaveX(ctx context.Context) *Area {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AreaUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AreaUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auo *AreaUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AreaUpdateOne {
	auo.modifiers = append(auo.modifiers, modifiers...)
	return auo
}

func (auo *AreaUpdateOne) sqlSave(ctx context.Context) (_node *Area, err error) {
	_spec := sqlgraph.NewUpdateSpec(area.Table, area.Columns, sqlgraph.NewFieldSpec(area.FieldID, field.TypeUint64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Area.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, area.FieldID)
		for _, f := range fields {
			if !area.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != area.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(area.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.Level(); ok {
		_spec.SetField(area.FieldLevel, field.TypeUint64, value)
	}
	if value, ok := auo.mutation.AddedLevel(); ok {
		_spec.AddField(area.FieldLevel, field.TypeUint64, value)
	}
	if value, ok := auo.mutation.Pid(); ok {
		_spec.SetField(area.FieldPid, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedPid(); ok {
		_spec.AddField(area.FieldPid, field.TypeInt64, value)
	}
	_spec.AddModifiers(auo.modifiers...)
	_node = &Area{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{area.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
