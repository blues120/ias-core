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
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/signature"
)

// SignatureCreate is the builder for creating a Signature entity.
type SignatureCreate struct {
	config
	mutation *SignatureMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (sc *SignatureCreate) SetCreatedAt(t time.Time) *SignatureCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SignatureCreate) SetNillableCreatedAt(t *time.Time) *SignatureCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SignatureCreate) SetUpdatedAt(t time.Time) *SignatureCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SignatureCreate) SetNillableUpdatedAt(t *time.Time) *SignatureCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetDeletedAt sets the "deleted_at" field.
func (sc *SignatureCreate) SetDeletedAt(t time.Time) *SignatureCreate {
	sc.mutation.SetDeletedAt(t)
	return sc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sc *SignatureCreate) SetNillableDeletedAt(t *time.Time) *SignatureCreate {
	if t != nil {
		sc.SetDeletedAt(*t)
	}
	return sc
}

// SetBoxID sets the "box_id" field.
func (sc *SignatureCreate) SetBoxID(s string) *SignatureCreate {
	sc.mutation.SetBoxID(s)
	return sc
}

// SetAppID sets the "app_id" field.
func (sc *SignatureCreate) SetAppID(s string) *SignatureCreate {
	sc.mutation.SetAppID(s)
	return sc
}

// SetAppSecret sets the "app_secret" field.
func (sc *SignatureCreate) SetAppSecret(s string) *SignatureCreate {
	sc.mutation.SetAppSecret(s)
	return sc
}

// SetID sets the "id" field.
func (sc *SignatureCreate) SetID(u uint64) *SignatureCreate {
	sc.mutation.SetID(u)
	return sc
}

// Mutation returns the SignatureMutation object of the builder.
func (sc *SignatureCreate) Mutation() *SignatureMutation {
	return sc.mutation
}

// Save creates the Signature in the database.
func (sc *SignatureCreate) Save(ctx context.Context) (*Signature, error) {
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SignatureCreate) SaveX(ctx context.Context) *Signature {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SignatureCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SignatureCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SignatureCreate) defaults() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		if signature.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized signature.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := signature.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		if signature.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized signature.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := signature.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *SignatureCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Signature.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Signature.updated_at"`)}
	}
	if _, ok := sc.mutation.BoxID(); !ok {
		return &ValidationError{Name: "box_id", err: errors.New(`ent: missing required field "Signature.box_id"`)}
	}
	if _, ok := sc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "Signature.app_id"`)}
	}
	if _, ok := sc.mutation.AppSecret(); !ok {
		return &ValidationError{Name: "app_secret", err: errors.New(`ent: missing required field "Signature.app_secret"`)}
	}
	return nil
}

func (sc *SignatureCreate) sqlSave(ctx context.Context) (*Signature, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SignatureCreate) createSpec() (*Signature, *sqlgraph.CreateSpec) {
	var (
		_node = &Signature{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(signature.Table, sqlgraph.NewFieldSpec(signature.FieldID, field.TypeUint64))
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(signature.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(signature.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.DeletedAt(); ok {
		_spec.SetField(signature.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := sc.mutation.BoxID(); ok {
		_spec.SetField(signature.FieldBoxID, field.TypeString, value)
		_node.BoxID = value
	}
	if value, ok := sc.mutation.AppID(); ok {
		_spec.SetField(signature.FieldAppID, field.TypeString, value)
		_node.AppID = value
	}
	if value, ok := sc.mutation.AppSecret(); ok {
		_spec.SetField(signature.FieldAppSecret, field.TypeString, value)
		_node.AppSecret = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Signature.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SignatureUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (sc *SignatureCreate) OnConflict(opts ...sql.ConflictOption) *SignatureUpsertOne {
	sc.conflict = opts
	return &SignatureUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Signature.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SignatureCreate) OnConflictColumns(columns ...string) *SignatureUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SignatureUpsertOne{
		create: sc,
	}
}

type (
	// SignatureUpsertOne is the builder for "upsert"-ing
	//  one Signature node.
	SignatureUpsertOne struct {
		create *SignatureCreate
	}

	// SignatureUpsert is the "OnConflict" setter.
	SignatureUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *SignatureUpsert) SetUpdatedAt(v time.Time) *SignatureUpsert {
	u.Set(signature.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SignatureUpsert) UpdateUpdatedAt() *SignatureUpsert {
	u.SetExcluded(signature.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SignatureUpsert) SetDeletedAt(v time.Time) *SignatureUpsert {
	u.Set(signature.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SignatureUpsert) UpdateDeletedAt() *SignatureUpsert {
	u.SetExcluded(signature.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SignatureUpsert) ClearDeletedAt() *SignatureUpsert {
	u.SetNull(signature.FieldDeletedAt)
	return u
}

// SetBoxID sets the "box_id" field.
func (u *SignatureUpsert) SetBoxID(v string) *SignatureUpsert {
	u.Set(signature.FieldBoxID, v)
	return u
}

// UpdateBoxID sets the "box_id" field to the value that was provided on create.
func (u *SignatureUpsert) UpdateBoxID() *SignatureUpsert {
	u.SetExcluded(signature.FieldBoxID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *SignatureUpsert) SetAppID(v string) *SignatureUpsert {
	u.Set(signature.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *SignatureUpsert) UpdateAppID() *SignatureUpsert {
	u.SetExcluded(signature.FieldAppID)
	return u
}

// SetAppSecret sets the "app_secret" field.
func (u *SignatureUpsert) SetAppSecret(v string) *SignatureUpsert {
	u.Set(signature.FieldAppSecret, v)
	return u
}

// UpdateAppSecret sets the "app_secret" field to the value that was provided on create.
func (u *SignatureUpsert) UpdateAppSecret() *SignatureUpsert {
	u.SetExcluded(signature.FieldAppSecret)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Signature.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(signature.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SignatureUpsertOne) UpdateNewValues() *SignatureUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(signature.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(signature.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Signature.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SignatureUpsertOne) Ignore() *SignatureUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SignatureUpsertOne) DoNothing() *SignatureUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SignatureCreate.OnConflict
// documentation for more info.
func (u *SignatureUpsertOne) Update(set func(*SignatureUpsert)) *SignatureUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SignatureUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SignatureUpsertOne) SetUpdatedAt(v time.Time) *SignatureUpsertOne {
	return u.Update(func(s *SignatureUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SignatureUpsertOne) UpdateUpdatedAt() *SignatureUpsertOne {
	return u.Update(func(s *SignatureUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SignatureUpsertOne) SetDeletedAt(v time.Time) *SignatureUpsertOne {
	return u.Update(func(s *SignatureUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SignatureUpsertOne) UpdateDeletedAt() *SignatureUpsertOne {
	return u.Update(func(s *SignatureUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SignatureUpsertOne) ClearDeletedAt() *SignatureUpsertOne {
	return u.Update(func(s *SignatureUpsert) {
		s.ClearDeletedAt()
	})
}

// SetBoxID sets the "box_id" field.
func (u *SignatureUpsertOne) SetBoxID(v string) *SignatureUpsertOne {
	return u.Update(func(s *SignatureUpsert) {
		s.SetBoxID(v)
	})
}

// UpdateBoxID sets the "box_id" field to the value that was provided on create.
func (u *SignatureUpsertOne) UpdateBoxID() *SignatureUpsertOne {
	return u.Update(func(s *SignatureUpsert) {
		s.UpdateBoxID()
	})
}

// SetAppID sets the "app_id" field.
func (u *SignatureUpsertOne) SetAppID(v string) *SignatureUpsertOne {
	return u.Update(func(s *SignatureUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *SignatureUpsertOne) UpdateAppID() *SignatureUpsertOne {
	return u.Update(func(s *SignatureUpsert) {
		s.UpdateAppID()
	})
}

// SetAppSecret sets the "app_secret" field.
func (u *SignatureUpsertOne) SetAppSecret(v string) *SignatureUpsertOne {
	return u.Update(func(s *SignatureUpsert) {
		s.SetAppSecret(v)
	})
}

// UpdateAppSecret sets the "app_secret" field to the value that was provided on create.
func (u *SignatureUpsertOne) UpdateAppSecret() *SignatureUpsertOne {
	return u.Update(func(s *SignatureUpsert) {
		s.UpdateAppSecret()
	})
}

// Exec executes the query.
func (u *SignatureUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SignatureCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SignatureUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SignatureUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SignatureUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SignatureCreateBulk is the builder for creating many Signature entities in bulk.
type SignatureCreateBulk struct {
	config
	err      error
	builders []*SignatureCreate
	conflict []sql.ConflictOption
}

// Save creates the Signature entities in the database.
func (scb *SignatureCreateBulk) Save(ctx context.Context) ([]*Signature, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Signature, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SignatureMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SignatureCreateBulk) SaveX(ctx context.Context) []*Signature {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SignatureCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SignatureCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Signature.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SignatureUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (scb *SignatureCreateBulk) OnConflict(opts ...sql.ConflictOption) *SignatureUpsertBulk {
	scb.conflict = opts
	return &SignatureUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Signature.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SignatureCreateBulk) OnConflictColumns(columns ...string) *SignatureUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SignatureUpsertBulk{
		create: scb,
	}
}

// SignatureUpsertBulk is the builder for "upsert"-ing
// a bulk of Signature nodes.
type SignatureUpsertBulk struct {
	create *SignatureCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Signature.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(signature.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SignatureUpsertBulk) UpdateNewValues() *SignatureUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(signature.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(signature.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Signature.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SignatureUpsertBulk) Ignore() *SignatureUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SignatureUpsertBulk) DoNothing() *SignatureUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SignatureCreateBulk.OnConflict
// documentation for more info.
func (u *SignatureUpsertBulk) Update(set func(*SignatureUpsert)) *SignatureUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SignatureUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SignatureUpsertBulk) SetUpdatedAt(v time.Time) *SignatureUpsertBulk {
	return u.Update(func(s *SignatureUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SignatureUpsertBulk) UpdateUpdatedAt() *SignatureUpsertBulk {
	return u.Update(func(s *SignatureUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SignatureUpsertBulk) SetDeletedAt(v time.Time) *SignatureUpsertBulk {
	return u.Update(func(s *SignatureUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SignatureUpsertBulk) UpdateDeletedAt() *SignatureUpsertBulk {
	return u.Update(func(s *SignatureUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SignatureUpsertBulk) ClearDeletedAt() *SignatureUpsertBulk {
	return u.Update(func(s *SignatureUpsert) {
		s.ClearDeletedAt()
	})
}

// SetBoxID sets the "box_id" field.
func (u *SignatureUpsertBulk) SetBoxID(v string) *SignatureUpsertBulk {
	return u.Update(func(s *SignatureUpsert) {
		s.SetBoxID(v)
	})
}

// UpdateBoxID sets the "box_id" field to the value that was provided on create.
func (u *SignatureUpsertBulk) UpdateBoxID() *SignatureUpsertBulk {
	return u.Update(func(s *SignatureUpsert) {
		s.UpdateBoxID()
	})
}

// SetAppID sets the "app_id" field.
func (u *SignatureUpsertBulk) SetAppID(v string) *SignatureUpsertBulk {
	return u.Update(func(s *SignatureUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *SignatureUpsertBulk) UpdateAppID() *SignatureUpsertBulk {
	return u.Update(func(s *SignatureUpsert) {
		s.UpdateAppID()
	})
}

// SetAppSecret sets the "app_secret" field.
func (u *SignatureUpsertBulk) SetAppSecret(v string) *SignatureUpsertBulk {
	return u.Update(func(s *SignatureUpsert) {
		s.SetAppSecret(v)
	})
}

// UpdateAppSecret sets the "app_secret" field to the value that was provided on create.
func (u *SignatureUpsertBulk) UpdateAppSecret() *SignatureUpsertBulk {
	return u.Update(func(s *SignatureUpsert) {
		s.UpdateAppSecret()
	})
}

// Exec executes the query.
func (u *SignatureUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SignatureCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SignatureCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SignatureUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
