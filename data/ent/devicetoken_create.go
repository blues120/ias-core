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
	"github.com/blues120/ias-core/data/ent/devicetoken"
)

// DeviceTokenCreate is the builder for creating a DeviceToken entity.
type DeviceTokenCreate struct {
	config
	mutation *DeviceTokenMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (dtc *DeviceTokenCreate) SetCreatedAt(t time.Time) *DeviceTokenCreate {
	dtc.mutation.SetCreatedAt(t)
	return dtc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dtc *DeviceTokenCreate) SetNillableCreatedAt(t *time.Time) *DeviceTokenCreate {
	if t != nil {
		dtc.SetCreatedAt(*t)
	}
	return dtc
}

// SetUpdatedAt sets the "updated_at" field.
func (dtc *DeviceTokenCreate) SetUpdatedAt(t time.Time) *DeviceTokenCreate {
	dtc.mutation.SetUpdatedAt(t)
	return dtc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dtc *DeviceTokenCreate) SetNillableUpdatedAt(t *time.Time) *DeviceTokenCreate {
	if t != nil {
		dtc.SetUpdatedAt(*t)
	}
	return dtc
}

// SetDeletedAt sets the "deleted_at" field.
func (dtc *DeviceTokenCreate) SetDeletedAt(t time.Time) *DeviceTokenCreate {
	dtc.mutation.SetDeletedAt(t)
	return dtc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dtc *DeviceTokenCreate) SetNillableDeletedAt(t *time.Time) *DeviceTokenCreate {
	if t != nil {
		dtc.SetDeletedAt(*t)
	}
	return dtc
}

// SetToken sets the "token" field.
func (dtc *DeviceTokenCreate) SetToken(s string) *DeviceTokenCreate {
	dtc.mutation.SetToken(s)
	return dtc
}

// SetDeviceExtID sets the "device_ext_id" field.
func (dtc *DeviceTokenCreate) SetDeviceExtID(s string) *DeviceTokenCreate {
	dtc.mutation.SetDeviceExtID(s)
	return dtc
}

// SetID sets the "id" field.
func (dtc *DeviceTokenCreate) SetID(u uint64) *DeviceTokenCreate {
	dtc.mutation.SetID(u)
	return dtc
}

// Mutation returns the DeviceTokenMutation object of the builder.
func (dtc *DeviceTokenCreate) Mutation() *DeviceTokenMutation {
	return dtc.mutation
}

// Save creates the DeviceToken in the database.
func (dtc *DeviceTokenCreate) Save(ctx context.Context) (*DeviceToken, error) {
	if err := dtc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, dtc.sqlSave, dtc.mutation, dtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dtc *DeviceTokenCreate) SaveX(ctx context.Context) *DeviceToken {
	v, err := dtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dtc *DeviceTokenCreate) Exec(ctx context.Context) error {
	_, err := dtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtc *DeviceTokenCreate) ExecX(ctx context.Context) {
	if err := dtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dtc *DeviceTokenCreate) defaults() error {
	if _, ok := dtc.mutation.CreatedAt(); !ok {
		if devicetoken.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized devicetoken.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := devicetoken.DefaultCreatedAt()
		dtc.mutation.SetCreatedAt(v)
	}
	if _, ok := dtc.mutation.UpdatedAt(); !ok {
		if devicetoken.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized devicetoken.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := devicetoken.DefaultUpdatedAt()
		dtc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (dtc *DeviceTokenCreate) check() error {
	if _, ok := dtc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DeviceToken.created_at"`)}
	}
	if _, ok := dtc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "DeviceToken.updated_at"`)}
	}
	if _, ok := dtc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`ent: missing required field "DeviceToken.token"`)}
	}
	if _, ok := dtc.mutation.DeviceExtID(); !ok {
		return &ValidationError{Name: "device_ext_id", err: errors.New(`ent: missing required field "DeviceToken.device_ext_id"`)}
	}
	return nil
}

func (dtc *DeviceTokenCreate) sqlSave(ctx context.Context) (*DeviceToken, error) {
	if err := dtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	dtc.mutation.id = &_node.ID
	dtc.mutation.done = true
	return _node, nil
}

func (dtc *DeviceTokenCreate) createSpec() (*DeviceToken, *sqlgraph.CreateSpec) {
	var (
		_node = &DeviceToken{config: dtc.config}
		_spec = sqlgraph.NewCreateSpec(devicetoken.Table, sqlgraph.NewFieldSpec(devicetoken.FieldID, field.TypeUint64))
	)
	_spec.OnConflict = dtc.conflict
	if id, ok := dtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dtc.mutation.CreatedAt(); ok {
		_spec.SetField(devicetoken.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dtc.mutation.UpdatedAt(); ok {
		_spec.SetField(devicetoken.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dtc.mutation.DeletedAt(); ok {
		_spec.SetField(devicetoken.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := dtc.mutation.Token(); ok {
		_spec.SetField(devicetoken.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := dtc.mutation.DeviceExtID(); ok {
		_spec.SetField(devicetoken.FieldDeviceExtID, field.TypeString, value)
		_node.DeviceExtID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeviceToken.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceTokenUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dtc *DeviceTokenCreate) OnConflict(opts ...sql.ConflictOption) *DeviceTokenUpsertOne {
	dtc.conflict = opts
	return &DeviceTokenUpsertOne{
		create: dtc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeviceToken.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dtc *DeviceTokenCreate) OnConflictColumns(columns ...string) *DeviceTokenUpsertOne {
	dtc.conflict = append(dtc.conflict, sql.ConflictColumns(columns...))
	return &DeviceTokenUpsertOne{
		create: dtc,
	}
}

type (
	// DeviceTokenUpsertOne is the builder for "upsert"-ing
	//  one DeviceToken node.
	DeviceTokenUpsertOne struct {
		create *DeviceTokenCreate
	}

	// DeviceTokenUpsert is the "OnConflict" setter.
	DeviceTokenUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceTokenUpsert) SetUpdatedAt(v time.Time) *DeviceTokenUpsert {
	u.Set(devicetoken.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceTokenUpsert) UpdateUpdatedAt() *DeviceTokenUpsert {
	u.SetExcluded(devicetoken.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceTokenUpsert) SetDeletedAt(v time.Time) *DeviceTokenUpsert {
	u.Set(devicetoken.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceTokenUpsert) UpdateDeletedAt() *DeviceTokenUpsert {
	u.SetExcluded(devicetoken.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *DeviceTokenUpsert) ClearDeletedAt() *DeviceTokenUpsert {
	u.SetNull(devicetoken.FieldDeletedAt)
	return u
}

// SetToken sets the "token" field.
func (u *DeviceTokenUpsert) SetToken(v string) *DeviceTokenUpsert {
	u.Set(devicetoken.FieldToken, v)
	return u
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *DeviceTokenUpsert) UpdateToken() *DeviceTokenUpsert {
	u.SetExcluded(devicetoken.FieldToken)
	return u
}

// SetDeviceExtID sets the "device_ext_id" field.
func (u *DeviceTokenUpsert) SetDeviceExtID(v string) *DeviceTokenUpsert {
	u.Set(devicetoken.FieldDeviceExtID, v)
	return u
}

// UpdateDeviceExtID sets the "device_ext_id" field to the value that was provided on create.
func (u *DeviceTokenUpsert) UpdateDeviceExtID() *DeviceTokenUpsert {
	u.SetExcluded(devicetoken.FieldDeviceExtID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DeviceToken.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(devicetoken.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceTokenUpsertOne) UpdateNewValues() *DeviceTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(devicetoken.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(devicetoken.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeviceToken.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DeviceTokenUpsertOne) Ignore() *DeviceTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceTokenUpsertOne) DoNothing() *DeviceTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceTokenCreate.OnConflict
// documentation for more info.
func (u *DeviceTokenUpsertOne) Update(set func(*DeviceTokenUpsert)) *DeviceTokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceTokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceTokenUpsertOne) SetUpdatedAt(v time.Time) *DeviceTokenUpsertOne {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceTokenUpsertOne) UpdateUpdatedAt() *DeviceTokenUpsertOne {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceTokenUpsertOne) SetDeletedAt(v time.Time) *DeviceTokenUpsertOne {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceTokenUpsertOne) UpdateDeletedAt() *DeviceTokenUpsertOne {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *DeviceTokenUpsertOne) ClearDeletedAt() *DeviceTokenUpsertOne {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.ClearDeletedAt()
	})
}

// SetToken sets the "token" field.
func (u *DeviceTokenUpsertOne) SetToken(v string) *DeviceTokenUpsertOne {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.SetToken(v)
	})
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *DeviceTokenUpsertOne) UpdateToken() *DeviceTokenUpsertOne {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.UpdateToken()
	})
}

// SetDeviceExtID sets the "device_ext_id" field.
func (u *DeviceTokenUpsertOne) SetDeviceExtID(v string) *DeviceTokenUpsertOne {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.SetDeviceExtID(v)
	})
}

// UpdateDeviceExtID sets the "device_ext_id" field to the value that was provided on create.
func (u *DeviceTokenUpsertOne) UpdateDeviceExtID() *DeviceTokenUpsertOne {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.UpdateDeviceExtID()
	})
}

// Exec executes the query.
func (u *DeviceTokenUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeviceTokenCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceTokenUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeviceTokenUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeviceTokenUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeviceTokenCreateBulk is the builder for creating many DeviceToken entities in bulk.
type DeviceTokenCreateBulk struct {
	config
	err      error
	builders []*DeviceTokenCreate
	conflict []sql.ConflictOption
}

// Save creates the DeviceToken entities in the database.
func (dtcb *DeviceTokenCreateBulk) Save(ctx context.Context) ([]*DeviceToken, error) {
	if dtcb.err != nil {
		return nil, dtcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dtcb.builders))
	nodes := make([]*DeviceToken, len(dtcb.builders))
	mutators := make([]Mutator, len(dtcb.builders))
	for i := range dtcb.builders {
		func(i int, root context.Context) {
			builder := dtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceTokenMutation)
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
					_, err = mutators[i+1].Mutate(root, dtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dtcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dtcb *DeviceTokenCreateBulk) SaveX(ctx context.Context) []*DeviceToken {
	v, err := dtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dtcb *DeviceTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := dtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtcb *DeviceTokenCreateBulk) ExecX(ctx context.Context) {
	if err := dtcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeviceToken.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceTokenUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dtcb *DeviceTokenCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeviceTokenUpsertBulk {
	dtcb.conflict = opts
	return &DeviceTokenUpsertBulk{
		create: dtcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeviceToken.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dtcb *DeviceTokenCreateBulk) OnConflictColumns(columns ...string) *DeviceTokenUpsertBulk {
	dtcb.conflict = append(dtcb.conflict, sql.ConflictColumns(columns...))
	return &DeviceTokenUpsertBulk{
		create: dtcb,
	}
}

// DeviceTokenUpsertBulk is the builder for "upsert"-ing
// a bulk of DeviceToken nodes.
type DeviceTokenUpsertBulk struct {
	create *DeviceTokenCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DeviceToken.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(devicetoken.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceTokenUpsertBulk) UpdateNewValues() *DeviceTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(devicetoken.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(devicetoken.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeviceToken.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DeviceTokenUpsertBulk) Ignore() *DeviceTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceTokenUpsertBulk) DoNothing() *DeviceTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceTokenCreateBulk.OnConflict
// documentation for more info.
func (u *DeviceTokenUpsertBulk) Update(set func(*DeviceTokenUpsert)) *DeviceTokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceTokenUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceTokenUpsertBulk) SetUpdatedAt(v time.Time) *DeviceTokenUpsertBulk {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceTokenUpsertBulk) UpdateUpdatedAt() *DeviceTokenUpsertBulk {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceTokenUpsertBulk) SetDeletedAt(v time.Time) *DeviceTokenUpsertBulk {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceTokenUpsertBulk) UpdateDeletedAt() *DeviceTokenUpsertBulk {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *DeviceTokenUpsertBulk) ClearDeletedAt() *DeviceTokenUpsertBulk {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.ClearDeletedAt()
	})
}

// SetToken sets the "token" field.
func (u *DeviceTokenUpsertBulk) SetToken(v string) *DeviceTokenUpsertBulk {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.SetToken(v)
	})
}

// UpdateToken sets the "token" field to the value that was provided on create.
func (u *DeviceTokenUpsertBulk) UpdateToken() *DeviceTokenUpsertBulk {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.UpdateToken()
	})
}

// SetDeviceExtID sets the "device_ext_id" field.
func (u *DeviceTokenUpsertBulk) SetDeviceExtID(v string) *DeviceTokenUpsertBulk {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.SetDeviceExtID(v)
	})
}

// UpdateDeviceExtID sets the "device_ext_id" field to the value that was provided on create.
func (u *DeviceTokenUpsertBulk) UpdateDeviceExtID() *DeviceTokenUpsertBulk {
	return u.Update(func(s *DeviceTokenUpsert) {
		s.UpdateDeviceExtID()
	})
}

// Exec executes the query.
func (u *DeviceTokenUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DeviceTokenCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeviceTokenCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceTokenUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
