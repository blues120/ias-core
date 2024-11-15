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
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/equipattr"
)

// EquipAttrCreate is the builder for creating a EquipAttr entity.
type EquipAttrCreate struct {
	config
	mutation *EquipAttrMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (eac *EquipAttrCreate) SetCreatedAt(t time.Time) *EquipAttrCreate {
	eac.mutation.SetCreatedAt(t)
	return eac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eac *EquipAttrCreate) SetNillableCreatedAt(t *time.Time) *EquipAttrCreate {
	if t != nil {
		eac.SetCreatedAt(*t)
	}
	return eac
}

// SetUpdatedAt sets the "updated_at" field.
func (eac *EquipAttrCreate) SetUpdatedAt(t time.Time) *EquipAttrCreate {
	eac.mutation.SetUpdatedAt(t)
	return eac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (eac *EquipAttrCreate) SetNillableUpdatedAt(t *time.Time) *EquipAttrCreate {
	if t != nil {
		eac.SetUpdatedAt(*t)
	}
	return eac
}

// SetDeletedAt sets the "deleted_at" field.
func (eac *EquipAttrCreate) SetDeletedAt(t time.Time) *EquipAttrCreate {
	eac.mutation.SetDeletedAt(t)
	return eac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (eac *EquipAttrCreate) SetNillableDeletedAt(t *time.Time) *EquipAttrCreate {
	if t != nil {
		eac.SetDeletedAt(*t)
	}
	return eac
}

// SetAttrKey sets the "attr_key" field.
func (eac *EquipAttrCreate) SetAttrKey(s string) *EquipAttrCreate {
	eac.mutation.SetAttrKey(s)
	return eac
}

// SetAttrValue sets the "attr_value" field.
func (eac *EquipAttrCreate) SetAttrValue(s string) *EquipAttrCreate {
	eac.mutation.SetAttrValue(s)
	return eac
}

// SetExtend sets the "extend" field.
func (eac *EquipAttrCreate) SetExtend(s string) *EquipAttrCreate {
	eac.mutation.SetExtend(s)
	return eac
}

// SetID sets the "id" field.
func (eac *EquipAttrCreate) SetID(u uint64) *EquipAttrCreate {
	eac.mutation.SetID(u)
	return eac
}

// Mutation returns the EquipAttrMutation object of the builder.
func (eac *EquipAttrCreate) Mutation() *EquipAttrMutation {
	return eac.mutation
}

// Save creates the EquipAttr in the database.
func (eac *EquipAttrCreate) Save(ctx context.Context) (*EquipAttr, error) {
	if err := eac.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, eac.sqlSave, eac.mutation, eac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (eac *EquipAttrCreate) SaveX(ctx context.Context) *EquipAttr {
	v, err := eac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eac *EquipAttrCreate) Exec(ctx context.Context) error {
	_, err := eac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eac *EquipAttrCreate) ExecX(ctx context.Context) {
	if err := eac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eac *EquipAttrCreate) defaults() error {
	if _, ok := eac.mutation.CreatedAt(); !ok {
		if equipattr.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized equipattr.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := equipattr.DefaultCreatedAt()
		eac.mutation.SetCreatedAt(v)
	}
	if _, ok := eac.mutation.UpdatedAt(); !ok {
		if equipattr.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized equipattr.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := equipattr.DefaultUpdatedAt()
		eac.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (eac *EquipAttrCreate) check() error {
	if _, ok := eac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "EquipAttr.created_at"`)}
	}
	if _, ok := eac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "EquipAttr.updated_at"`)}
	}
	if _, ok := eac.mutation.AttrKey(); !ok {
		return &ValidationError{Name: "attr_key", err: errors.New(`ent: missing required field "EquipAttr.attr_key"`)}
	}
	if _, ok := eac.mutation.AttrValue(); !ok {
		return &ValidationError{Name: "attr_value", err: errors.New(`ent: missing required field "EquipAttr.attr_value"`)}
	}
	if _, ok := eac.mutation.Extend(); !ok {
		return &ValidationError{Name: "extend", err: errors.New(`ent: missing required field "EquipAttr.extend"`)}
	}
	return nil
}

func (eac *EquipAttrCreate) sqlSave(ctx context.Context) (*EquipAttr, error) {
	if err := eac.check(); err != nil {
		return nil, err
	}
	_node, _spec := eac.createSpec()
	if err := sqlgraph.CreateNode(ctx, eac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	eac.mutation.id = &_node.ID
	eac.mutation.done = true
	return _node, nil
}

func (eac *EquipAttrCreate) createSpec() (*EquipAttr, *sqlgraph.CreateSpec) {
	var (
		_node = &EquipAttr{config: eac.config}
		_spec = sqlgraph.NewCreateSpec(equipattr.Table, sqlgraph.NewFieldSpec(equipattr.FieldID, field.TypeUint64))
	)
	_spec.OnConflict = eac.conflict
	if id, ok := eac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := eac.mutation.CreatedAt(); ok {
		_spec.SetField(equipattr.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := eac.mutation.UpdatedAt(); ok {
		_spec.SetField(equipattr.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := eac.mutation.DeletedAt(); ok {
		_spec.SetField(equipattr.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := eac.mutation.AttrKey(); ok {
		_spec.SetField(equipattr.FieldAttrKey, field.TypeString, value)
		_node.AttrKey = value
	}
	if value, ok := eac.mutation.AttrValue(); ok {
		_spec.SetField(equipattr.FieldAttrValue, field.TypeString, value)
		_node.AttrValue = value
	}
	if value, ok := eac.mutation.Extend(); ok {
		_spec.SetField(equipattr.FieldExtend, field.TypeString, value)
		_node.Extend = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.EquipAttr.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EquipAttrUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (eac *EquipAttrCreate) OnConflict(opts ...sql.ConflictOption) *EquipAttrUpsertOne {
	eac.conflict = opts
	return &EquipAttrUpsertOne{
		create: eac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.EquipAttr.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (eac *EquipAttrCreate) OnConflictColumns(columns ...string) *EquipAttrUpsertOne {
	eac.conflict = append(eac.conflict, sql.ConflictColumns(columns...))
	return &EquipAttrUpsertOne{
		create: eac,
	}
}

type (
	// EquipAttrUpsertOne is the builder for "upsert"-ing
	//  one EquipAttr node.
	EquipAttrUpsertOne struct {
		create *EquipAttrCreate
	}

	// EquipAttrUpsert is the "OnConflict" setter.
	EquipAttrUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *EquipAttrUpsert) SetUpdatedAt(v time.Time) *EquipAttrUpsert {
	u.Set(equipattr.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EquipAttrUpsert) UpdateUpdatedAt() *EquipAttrUpsert {
	u.SetExcluded(equipattr.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *EquipAttrUpsert) SetDeletedAt(v time.Time) *EquipAttrUpsert {
	u.Set(equipattr.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *EquipAttrUpsert) UpdateDeletedAt() *EquipAttrUpsert {
	u.SetExcluded(equipattr.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *EquipAttrUpsert) ClearDeletedAt() *EquipAttrUpsert {
	u.SetNull(equipattr.FieldDeletedAt)
	return u
}

// SetAttrKey sets the "attr_key" field.
func (u *EquipAttrUpsert) SetAttrKey(v string) *EquipAttrUpsert {
	u.Set(equipattr.FieldAttrKey, v)
	return u
}

// UpdateAttrKey sets the "attr_key" field to the value that was provided on create.
func (u *EquipAttrUpsert) UpdateAttrKey() *EquipAttrUpsert {
	u.SetExcluded(equipattr.FieldAttrKey)
	return u
}

// SetAttrValue sets the "attr_value" field.
func (u *EquipAttrUpsert) SetAttrValue(v string) *EquipAttrUpsert {
	u.Set(equipattr.FieldAttrValue, v)
	return u
}

// UpdateAttrValue sets the "attr_value" field to the value that was provided on create.
func (u *EquipAttrUpsert) UpdateAttrValue() *EquipAttrUpsert {
	u.SetExcluded(equipattr.FieldAttrValue)
	return u
}

// SetExtend sets the "extend" field.
func (u *EquipAttrUpsert) SetExtend(v string) *EquipAttrUpsert {
	u.Set(equipattr.FieldExtend, v)
	return u
}

// UpdateExtend sets the "extend" field to the value that was provided on create.
func (u *EquipAttrUpsert) UpdateExtend() *EquipAttrUpsert {
	u.SetExcluded(equipattr.FieldExtend)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.EquipAttr.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(equipattr.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EquipAttrUpsertOne) UpdateNewValues() *EquipAttrUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(equipattr.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(equipattr.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.EquipAttr.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *EquipAttrUpsertOne) Ignore() *EquipAttrUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EquipAttrUpsertOne) DoNothing() *EquipAttrUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EquipAttrCreate.OnConflict
// documentation for more info.
func (u *EquipAttrUpsertOne) Update(set func(*EquipAttrUpsert)) *EquipAttrUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EquipAttrUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EquipAttrUpsertOne) SetUpdatedAt(v time.Time) *EquipAttrUpsertOne {
	return u.Update(func(s *EquipAttrUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EquipAttrUpsertOne) UpdateUpdatedAt() *EquipAttrUpsertOne {
	return u.Update(func(s *EquipAttrUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *EquipAttrUpsertOne) SetDeletedAt(v time.Time) *EquipAttrUpsertOne {
	return u.Update(func(s *EquipAttrUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *EquipAttrUpsertOne) UpdateDeletedAt() *EquipAttrUpsertOne {
	return u.Update(func(s *EquipAttrUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *EquipAttrUpsertOne) ClearDeletedAt() *EquipAttrUpsertOne {
	return u.Update(func(s *EquipAttrUpsert) {
		s.ClearDeletedAt()
	})
}

// SetAttrKey sets the "attr_key" field.
func (u *EquipAttrUpsertOne) SetAttrKey(v string) *EquipAttrUpsertOne {
	return u.Update(func(s *EquipAttrUpsert) {
		s.SetAttrKey(v)
	})
}

// UpdateAttrKey sets the "attr_key" field to the value that was provided on create.
func (u *EquipAttrUpsertOne) UpdateAttrKey() *EquipAttrUpsertOne {
	return u.Update(func(s *EquipAttrUpsert) {
		s.UpdateAttrKey()
	})
}

// SetAttrValue sets the "attr_value" field.
func (u *EquipAttrUpsertOne) SetAttrValue(v string) *EquipAttrUpsertOne {
	return u.Update(func(s *EquipAttrUpsert) {
		s.SetAttrValue(v)
	})
}

// UpdateAttrValue sets the "attr_value" field to the value that was provided on create.
func (u *EquipAttrUpsertOne) UpdateAttrValue() *EquipAttrUpsertOne {
	return u.Update(func(s *EquipAttrUpsert) {
		s.UpdateAttrValue()
	})
}

// SetExtend sets the "extend" field.
func (u *EquipAttrUpsertOne) SetExtend(v string) *EquipAttrUpsertOne {
	return u.Update(func(s *EquipAttrUpsert) {
		s.SetExtend(v)
	})
}

// UpdateExtend sets the "extend" field to the value that was provided on create.
func (u *EquipAttrUpsertOne) UpdateExtend() *EquipAttrUpsertOne {
	return u.Update(func(s *EquipAttrUpsert) {
		s.UpdateExtend()
	})
}

// Exec executes the query.
func (u *EquipAttrUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EquipAttrCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EquipAttrUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EquipAttrUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EquipAttrUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EquipAttrCreateBulk is the builder for creating many EquipAttr entities in bulk.
type EquipAttrCreateBulk struct {
	config
	err      error
	builders []*EquipAttrCreate
	conflict []sql.ConflictOption
}

// Save creates the EquipAttr entities in the database.
func (eacb *EquipAttrCreateBulk) Save(ctx context.Context) ([]*EquipAttr, error) {
	if eacb.err != nil {
		return nil, eacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(eacb.builders))
	nodes := make([]*EquipAttr, len(eacb.builders))
	mutators := make([]Mutator, len(eacb.builders))
	for i := range eacb.builders {
		func(i int, root context.Context) {
			builder := eacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EquipAttrMutation)
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
					_, err = mutators[i+1].Mutate(root, eacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = eacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, eacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eacb *EquipAttrCreateBulk) SaveX(ctx context.Context) []*EquipAttr {
	v, err := eacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eacb *EquipAttrCreateBulk) Exec(ctx context.Context) error {
	_, err := eacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eacb *EquipAttrCreateBulk) ExecX(ctx context.Context) {
	if err := eacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.EquipAttr.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EquipAttrUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (eacb *EquipAttrCreateBulk) OnConflict(opts ...sql.ConflictOption) *EquipAttrUpsertBulk {
	eacb.conflict = opts
	return &EquipAttrUpsertBulk{
		create: eacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.EquipAttr.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (eacb *EquipAttrCreateBulk) OnConflictColumns(columns ...string) *EquipAttrUpsertBulk {
	eacb.conflict = append(eacb.conflict, sql.ConflictColumns(columns...))
	return &EquipAttrUpsertBulk{
		create: eacb,
	}
}

// EquipAttrUpsertBulk is the builder for "upsert"-ing
// a bulk of EquipAttr nodes.
type EquipAttrUpsertBulk struct {
	create *EquipAttrCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.EquipAttr.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(equipattr.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EquipAttrUpsertBulk) UpdateNewValues() *EquipAttrUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(equipattr.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(equipattr.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.EquipAttr.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *EquipAttrUpsertBulk) Ignore() *EquipAttrUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EquipAttrUpsertBulk) DoNothing() *EquipAttrUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EquipAttrCreateBulk.OnConflict
// documentation for more info.
func (u *EquipAttrUpsertBulk) Update(set func(*EquipAttrUpsert)) *EquipAttrUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EquipAttrUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EquipAttrUpsertBulk) SetUpdatedAt(v time.Time) *EquipAttrUpsertBulk {
	return u.Update(func(s *EquipAttrUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EquipAttrUpsertBulk) UpdateUpdatedAt() *EquipAttrUpsertBulk {
	return u.Update(func(s *EquipAttrUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *EquipAttrUpsertBulk) SetDeletedAt(v time.Time) *EquipAttrUpsertBulk {
	return u.Update(func(s *EquipAttrUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *EquipAttrUpsertBulk) UpdateDeletedAt() *EquipAttrUpsertBulk {
	return u.Update(func(s *EquipAttrUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *EquipAttrUpsertBulk) ClearDeletedAt() *EquipAttrUpsertBulk {
	return u.Update(func(s *EquipAttrUpsert) {
		s.ClearDeletedAt()
	})
}

// SetAttrKey sets the "attr_key" field.
func (u *EquipAttrUpsertBulk) SetAttrKey(v string) *EquipAttrUpsertBulk {
	return u.Update(func(s *EquipAttrUpsert) {
		s.SetAttrKey(v)
	})
}

// UpdateAttrKey sets the "attr_key" field to the value that was provided on create.
func (u *EquipAttrUpsertBulk) UpdateAttrKey() *EquipAttrUpsertBulk {
	return u.Update(func(s *EquipAttrUpsert) {
		s.UpdateAttrKey()
	})
}

// SetAttrValue sets the "attr_value" field.
func (u *EquipAttrUpsertBulk) SetAttrValue(v string) *EquipAttrUpsertBulk {
	return u.Update(func(s *EquipAttrUpsert) {
		s.SetAttrValue(v)
	})
}

// UpdateAttrValue sets the "attr_value" field to the value that was provided on create.
func (u *EquipAttrUpsertBulk) UpdateAttrValue() *EquipAttrUpsertBulk {
	return u.Update(func(s *EquipAttrUpsert) {
		s.UpdateAttrValue()
	})
}

// SetExtend sets the "extend" field.
func (u *EquipAttrUpsertBulk) SetExtend(v string) *EquipAttrUpsertBulk {
	return u.Update(func(s *EquipAttrUpsert) {
		s.SetExtend(v)
	})
}

// UpdateExtend sets the "extend" field to the value that was provided on create.
func (u *EquipAttrUpsertBulk) UpdateExtend() *EquipAttrUpsertBulk {
	return u.Update(func(s *EquipAttrUpsert) {
		s.UpdateExtend()
	})
}

// Exec executes the query.
func (u *EquipAttrUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the EquipAttrCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EquipAttrCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EquipAttrUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
