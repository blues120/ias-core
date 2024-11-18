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
	"github.com/blues120/ias-core/data/ent/algorithmalarmtype"
)

// AlgorithmAlarmTypeCreate is the builder for creating a AlgorithmAlarmType entity.
type AlgorithmAlarmTypeCreate struct {
	config
	mutation *AlgorithmAlarmTypeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (aatc *AlgorithmAlarmTypeCreate) SetCreatedAt(t time.Time) *AlgorithmAlarmTypeCreate {
	aatc.mutation.SetCreatedAt(t)
	return aatc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aatc *AlgorithmAlarmTypeCreate) SetNillableCreatedAt(t *time.Time) *AlgorithmAlarmTypeCreate {
	if t != nil {
		aatc.SetCreatedAt(*t)
	}
	return aatc
}

// SetUpdatedAt sets the "updated_at" field.
func (aatc *AlgorithmAlarmTypeCreate) SetUpdatedAt(t time.Time) *AlgorithmAlarmTypeCreate {
	aatc.mutation.SetUpdatedAt(t)
	return aatc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aatc *AlgorithmAlarmTypeCreate) SetNillableUpdatedAt(t *time.Time) *AlgorithmAlarmTypeCreate {
	if t != nil {
		aatc.SetUpdatedAt(*t)
	}
	return aatc
}

// SetDeletedAt sets the "deleted_at" field.
func (aatc *AlgorithmAlarmTypeCreate) SetDeletedAt(t time.Time) *AlgorithmAlarmTypeCreate {
	aatc.mutation.SetDeletedAt(t)
	return aatc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aatc *AlgorithmAlarmTypeCreate) SetNillableDeletedAt(t *time.Time) *AlgorithmAlarmTypeCreate {
	if t != nil {
		aatc.SetDeletedAt(*t)
	}
	return aatc
}

// SetTypeName sets the "type_name" field.
func (aatc *AlgorithmAlarmTypeCreate) SetTypeName(s string) *AlgorithmAlarmTypeCreate {
	aatc.mutation.SetTypeName(s)
	return aatc
}

// SetNillableTypeName sets the "type_name" field if the given value is not nil.
func (aatc *AlgorithmAlarmTypeCreate) SetNillableTypeName(s *string) *AlgorithmAlarmTypeCreate {
	if s != nil {
		aatc.SetTypeName(*s)
	}
	return aatc
}

// SetID sets the "id" field.
func (aatc *AlgorithmAlarmTypeCreate) SetID(u uint64) *AlgorithmAlarmTypeCreate {
	aatc.mutation.SetID(u)
	return aatc
}

// Mutation returns the AlgorithmAlarmTypeMutation object of the builder.
func (aatc *AlgorithmAlarmTypeCreate) Mutation() *AlgorithmAlarmTypeMutation {
	return aatc.mutation
}

// Save creates the AlgorithmAlarmType in the database.
func (aatc *AlgorithmAlarmTypeCreate) Save(ctx context.Context) (*AlgorithmAlarmType, error) {
	if err := aatc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, aatc.sqlSave, aatc.mutation, aatc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (aatc *AlgorithmAlarmTypeCreate) SaveX(ctx context.Context) *AlgorithmAlarmType {
	v, err := aatc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aatc *AlgorithmAlarmTypeCreate) Exec(ctx context.Context) error {
	_, err := aatc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aatc *AlgorithmAlarmTypeCreate) ExecX(ctx context.Context) {
	if err := aatc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aatc *AlgorithmAlarmTypeCreate) defaults() error {
	if _, ok := aatc.mutation.CreatedAt(); !ok {
		if algorithmalarmtype.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized algorithmalarmtype.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := algorithmalarmtype.DefaultCreatedAt()
		aatc.mutation.SetCreatedAt(v)
	}
	if _, ok := aatc.mutation.UpdatedAt(); !ok {
		if algorithmalarmtype.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized algorithmalarmtype.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := algorithmalarmtype.DefaultUpdatedAt()
		aatc.mutation.SetUpdatedAt(v)
	}
	if _, ok := aatc.mutation.TypeName(); !ok {
		v := algorithmalarmtype.DefaultTypeName
		aatc.mutation.SetTypeName(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (aatc *AlgorithmAlarmTypeCreate) check() error {
	if _, ok := aatc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AlgorithmAlarmType.created_at"`)}
	}
	if _, ok := aatc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AlgorithmAlarmType.updated_at"`)}
	}
	if _, ok := aatc.mutation.TypeName(); !ok {
		return &ValidationError{Name: "type_name", err: errors.New(`ent: missing required field "AlgorithmAlarmType.type_name"`)}
	}
	if v, ok := aatc.mutation.TypeName(); ok {
		if err := algorithmalarmtype.TypeNameValidator(v); err != nil {
			return &ValidationError{Name: "type_name", err: fmt.Errorf(`ent: validator failed for field "AlgorithmAlarmType.type_name": %w`, err)}
		}
	}
	return nil
}

func (aatc *AlgorithmAlarmTypeCreate) sqlSave(ctx context.Context) (*AlgorithmAlarmType, error) {
	if err := aatc.check(); err != nil {
		return nil, err
	}
	_node, _spec := aatc.createSpec()
	if err := sqlgraph.CreateNode(ctx, aatc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	aatc.mutation.id = &_node.ID
	aatc.mutation.done = true
	return _node, nil
}

func (aatc *AlgorithmAlarmTypeCreate) createSpec() (*AlgorithmAlarmType, *sqlgraph.CreateSpec) {
	var (
		_node = &AlgorithmAlarmType{config: aatc.config}
		_spec = sqlgraph.NewCreateSpec(algorithmalarmtype.Table, sqlgraph.NewFieldSpec(algorithmalarmtype.FieldID, field.TypeUint64))
	)
	_spec.OnConflict = aatc.conflict
	if id, ok := aatc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := aatc.mutation.CreatedAt(); ok {
		_spec.SetField(algorithmalarmtype.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := aatc.mutation.UpdatedAt(); ok {
		_spec.SetField(algorithmalarmtype.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := aatc.mutation.DeletedAt(); ok {
		_spec.SetField(algorithmalarmtype.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := aatc.mutation.TypeName(); ok {
		_spec.SetField(algorithmalarmtype.FieldTypeName, field.TypeString, value)
		_node.TypeName = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AlgorithmAlarmType.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AlgorithmAlarmTypeUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (aatc *AlgorithmAlarmTypeCreate) OnConflict(opts ...sql.ConflictOption) *AlgorithmAlarmTypeUpsertOne {
	aatc.conflict = opts
	return &AlgorithmAlarmTypeUpsertOne{
		create: aatc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AlgorithmAlarmType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (aatc *AlgorithmAlarmTypeCreate) OnConflictColumns(columns ...string) *AlgorithmAlarmTypeUpsertOne {
	aatc.conflict = append(aatc.conflict, sql.ConflictColumns(columns...))
	return &AlgorithmAlarmTypeUpsertOne{
		create: aatc,
	}
}

type (
	// AlgorithmAlarmTypeUpsertOne is the builder for "upsert"-ing
	//  one AlgorithmAlarmType node.
	AlgorithmAlarmTypeUpsertOne struct {
		create *AlgorithmAlarmTypeCreate
	}

	// AlgorithmAlarmTypeUpsert is the "OnConflict" setter.
	AlgorithmAlarmTypeUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *AlgorithmAlarmTypeUpsert) SetUpdatedAt(v time.Time) *AlgorithmAlarmTypeUpsert {
	u.Set(algorithmalarmtype.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AlgorithmAlarmTypeUpsert) UpdateUpdatedAt() *AlgorithmAlarmTypeUpsert {
	u.SetExcluded(algorithmalarmtype.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AlgorithmAlarmTypeUpsert) SetDeletedAt(v time.Time) *AlgorithmAlarmTypeUpsert {
	u.Set(algorithmalarmtype.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AlgorithmAlarmTypeUpsert) UpdateDeletedAt() *AlgorithmAlarmTypeUpsert {
	u.SetExcluded(algorithmalarmtype.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AlgorithmAlarmTypeUpsert) ClearDeletedAt() *AlgorithmAlarmTypeUpsert {
	u.SetNull(algorithmalarmtype.FieldDeletedAt)
	return u
}

// SetTypeName sets the "type_name" field.
func (u *AlgorithmAlarmTypeUpsert) SetTypeName(v string) *AlgorithmAlarmTypeUpsert {
	u.Set(algorithmalarmtype.FieldTypeName, v)
	return u
}

// UpdateTypeName sets the "type_name" field to the value that was provided on create.
func (u *AlgorithmAlarmTypeUpsert) UpdateTypeName() *AlgorithmAlarmTypeUpsert {
	u.SetExcluded(algorithmalarmtype.FieldTypeName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AlgorithmAlarmType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(algorithmalarmtype.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AlgorithmAlarmTypeUpsertOne) UpdateNewValues() *AlgorithmAlarmTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(algorithmalarmtype.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(algorithmalarmtype.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AlgorithmAlarmType.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AlgorithmAlarmTypeUpsertOne) Ignore() *AlgorithmAlarmTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AlgorithmAlarmTypeUpsertOne) DoNothing() *AlgorithmAlarmTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AlgorithmAlarmTypeCreate.OnConflict
// documentation for more info.
func (u *AlgorithmAlarmTypeUpsertOne) Update(set func(*AlgorithmAlarmTypeUpsert)) *AlgorithmAlarmTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AlgorithmAlarmTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AlgorithmAlarmTypeUpsertOne) SetUpdatedAt(v time.Time) *AlgorithmAlarmTypeUpsertOne {
	return u.Update(func(s *AlgorithmAlarmTypeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AlgorithmAlarmTypeUpsertOne) UpdateUpdatedAt() *AlgorithmAlarmTypeUpsertOne {
	return u.Update(func(s *AlgorithmAlarmTypeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AlgorithmAlarmTypeUpsertOne) SetDeletedAt(v time.Time) *AlgorithmAlarmTypeUpsertOne {
	return u.Update(func(s *AlgorithmAlarmTypeUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AlgorithmAlarmTypeUpsertOne) UpdateDeletedAt() *AlgorithmAlarmTypeUpsertOne {
	return u.Update(func(s *AlgorithmAlarmTypeUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AlgorithmAlarmTypeUpsertOne) ClearDeletedAt() *AlgorithmAlarmTypeUpsertOne {
	return u.Update(func(s *AlgorithmAlarmTypeUpsert) {
		s.ClearDeletedAt()
	})
}

// SetTypeName sets the "type_name" field.
func (u *AlgorithmAlarmTypeUpsertOne) SetTypeName(v string) *AlgorithmAlarmTypeUpsertOne {
	return u.Update(func(s *AlgorithmAlarmTypeUpsert) {
		s.SetTypeName(v)
	})
}

// UpdateTypeName sets the "type_name" field to the value that was provided on create.
func (u *AlgorithmAlarmTypeUpsertOne) UpdateTypeName() *AlgorithmAlarmTypeUpsertOne {
	return u.Update(func(s *AlgorithmAlarmTypeUpsert) {
		s.UpdateTypeName()
	})
}

// Exec executes the query.
func (u *AlgorithmAlarmTypeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AlgorithmAlarmTypeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AlgorithmAlarmTypeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AlgorithmAlarmTypeUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AlgorithmAlarmTypeUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AlgorithmAlarmTypeCreateBulk is the builder for creating many AlgorithmAlarmType entities in bulk.
type AlgorithmAlarmTypeCreateBulk struct {
	config
	err      error
	builders []*AlgorithmAlarmTypeCreate
	conflict []sql.ConflictOption
}

// Save creates the AlgorithmAlarmType entities in the database.
func (aatcb *AlgorithmAlarmTypeCreateBulk) Save(ctx context.Context) ([]*AlgorithmAlarmType, error) {
	if aatcb.err != nil {
		return nil, aatcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(aatcb.builders))
	nodes := make([]*AlgorithmAlarmType, len(aatcb.builders))
	mutators := make([]Mutator, len(aatcb.builders))
	for i := range aatcb.builders {
		func(i int, root context.Context) {
			builder := aatcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AlgorithmAlarmTypeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, aatcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = aatcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aatcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, aatcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aatcb *AlgorithmAlarmTypeCreateBulk) SaveX(ctx context.Context) []*AlgorithmAlarmType {
	v, err := aatcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aatcb *AlgorithmAlarmTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := aatcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aatcb *AlgorithmAlarmTypeCreateBulk) ExecX(ctx context.Context) {
	if err := aatcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AlgorithmAlarmType.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AlgorithmAlarmTypeUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (aatcb *AlgorithmAlarmTypeCreateBulk) OnConflict(opts ...sql.ConflictOption) *AlgorithmAlarmTypeUpsertBulk {
	aatcb.conflict = opts
	return &AlgorithmAlarmTypeUpsertBulk{
		create: aatcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AlgorithmAlarmType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (aatcb *AlgorithmAlarmTypeCreateBulk) OnConflictColumns(columns ...string) *AlgorithmAlarmTypeUpsertBulk {
	aatcb.conflict = append(aatcb.conflict, sql.ConflictColumns(columns...))
	return &AlgorithmAlarmTypeUpsertBulk{
		create: aatcb,
	}
}

// AlgorithmAlarmTypeUpsertBulk is the builder for "upsert"-ing
// a bulk of AlgorithmAlarmType nodes.
type AlgorithmAlarmTypeUpsertBulk struct {
	create *AlgorithmAlarmTypeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AlgorithmAlarmType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(algorithmalarmtype.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AlgorithmAlarmTypeUpsertBulk) UpdateNewValues() *AlgorithmAlarmTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(algorithmalarmtype.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(algorithmalarmtype.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AlgorithmAlarmType.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AlgorithmAlarmTypeUpsertBulk) Ignore() *AlgorithmAlarmTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AlgorithmAlarmTypeUpsertBulk) DoNothing() *AlgorithmAlarmTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AlgorithmAlarmTypeCreateBulk.OnConflict
// documentation for more info.
func (u *AlgorithmAlarmTypeUpsertBulk) Update(set func(*AlgorithmAlarmTypeUpsert)) *AlgorithmAlarmTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AlgorithmAlarmTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AlgorithmAlarmTypeUpsertBulk) SetUpdatedAt(v time.Time) *AlgorithmAlarmTypeUpsertBulk {
	return u.Update(func(s *AlgorithmAlarmTypeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AlgorithmAlarmTypeUpsertBulk) UpdateUpdatedAt() *AlgorithmAlarmTypeUpsertBulk {
	return u.Update(func(s *AlgorithmAlarmTypeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AlgorithmAlarmTypeUpsertBulk) SetDeletedAt(v time.Time) *AlgorithmAlarmTypeUpsertBulk {
	return u.Update(func(s *AlgorithmAlarmTypeUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AlgorithmAlarmTypeUpsertBulk) UpdateDeletedAt() *AlgorithmAlarmTypeUpsertBulk {
	return u.Update(func(s *AlgorithmAlarmTypeUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AlgorithmAlarmTypeUpsertBulk) ClearDeletedAt() *AlgorithmAlarmTypeUpsertBulk {
	return u.Update(func(s *AlgorithmAlarmTypeUpsert) {
		s.ClearDeletedAt()
	})
}

// SetTypeName sets the "type_name" field.
func (u *AlgorithmAlarmTypeUpsertBulk) SetTypeName(v string) *AlgorithmAlarmTypeUpsertBulk {
	return u.Update(func(s *AlgorithmAlarmTypeUpsert) {
		s.SetTypeName(v)
	})
}

// UpdateTypeName sets the "type_name" field to the value that was provided on create.
func (u *AlgorithmAlarmTypeUpsertBulk) UpdateTypeName() *AlgorithmAlarmTypeUpsertBulk {
	return u.Update(func(s *AlgorithmAlarmTypeUpsert) {
		s.UpdateTypeName()
	})
}

// Exec executes the query.
func (u *AlgorithmAlarmTypeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AlgorithmAlarmTypeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AlgorithmAlarmTypeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AlgorithmAlarmTypeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
