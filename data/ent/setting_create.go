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
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/setting"
)

// SettingCreate is the builder for creating a Setting entity.
type SettingCreate struct {
	config
	mutation *SettingMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (sc *SettingCreate) SetCreatedAt(t time.Time) *SettingCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SettingCreate) SetNillableCreatedAt(t *time.Time) *SettingCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SettingCreate) SetUpdatedAt(t time.Time) *SettingCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SettingCreate) SetNillableUpdatedAt(t *time.Time) *SettingCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *SettingCreate) SetName(s string) *SettingCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetSerialNumber sets the "serial_number" field.
func (sc *SettingCreate) SetSerialNumber(s string) *SettingCreate {
	sc.mutation.SetSerialNumber(s)
	return sc
}

// SetVersion sets the "version" field.
func (sc *SettingCreate) SetVersion(s string) *SettingCreate {
	sc.mutation.SetVersion(s)
	return sc
}

// SetModel sets the "model" field.
func (sc *SettingCreate) SetModel(s string) *SettingCreate {
	sc.mutation.SetModel(s)
	return sc
}

// SetWorkspaceID sets the "workspace_id" field.
func (sc *SettingCreate) SetWorkspaceID(s string) *SettingCreate {
	sc.mutation.SetWorkspaceID(s)
	return sc
}

// SetNillableWorkspaceID sets the "workspace_id" field if the given value is not nil.
func (sc *SettingCreate) SetNillableWorkspaceID(s *string) *SettingCreate {
	if s != nil {
		sc.SetWorkspaceID(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SettingCreate) SetID(u uint64) *SettingCreate {
	sc.mutation.SetID(u)
	return sc
}

// Mutation returns the SettingMutation object of the builder.
func (sc *SettingCreate) Mutation() *SettingMutation {
	return sc.mutation
}

// Save creates the Setting in the database.
func (sc *SettingCreate) Save(ctx context.Context) (*Setting, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SettingCreate) SaveX(ctx context.Context) *Setting {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SettingCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SettingCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SettingCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := setting.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := setting.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.WorkspaceID(); !ok {
		v := setting.DefaultWorkspaceID
		sc.mutation.SetWorkspaceID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SettingCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Setting.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Setting.updated_at"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Setting.name"`)}
	}
	if _, ok := sc.mutation.SerialNumber(); !ok {
		return &ValidationError{Name: "serial_number", err: errors.New(`ent: missing required field "Setting.serial_number"`)}
	}
	if _, ok := sc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "Setting.version"`)}
	}
	if _, ok := sc.mutation.Model(); !ok {
		return &ValidationError{Name: "model", err: errors.New(`ent: missing required field "Setting.model"`)}
	}
	if _, ok := sc.mutation.WorkspaceID(); !ok {
		return &ValidationError{Name: "workspace_id", err: errors.New(`ent: missing required field "Setting.workspace_id"`)}
	}
	return nil
}

func (sc *SettingCreate) sqlSave(ctx context.Context) (*Setting, error) {
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

func (sc *SettingCreate) createSpec() (*Setting, *sqlgraph.CreateSpec) {
	var (
		_node = &Setting{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(setting.Table, sqlgraph.NewFieldSpec(setting.FieldID, field.TypeUint64))
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(setting.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(setting.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(setting.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.SerialNumber(); ok {
		_spec.SetField(setting.FieldSerialNumber, field.TypeString, value)
		_node.SerialNumber = value
	}
	if value, ok := sc.mutation.Version(); ok {
		_spec.SetField(setting.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := sc.mutation.Model(); ok {
		_spec.SetField(setting.FieldModel, field.TypeString, value)
		_node.Model = value
	}
	if value, ok := sc.mutation.WorkspaceID(); ok {
		_spec.SetField(setting.FieldWorkspaceID, field.TypeString, value)
		_node.WorkspaceID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Setting.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SettingUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (sc *SettingCreate) OnConflict(opts ...sql.ConflictOption) *SettingUpsertOne {
	sc.conflict = opts
	return &SettingUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SettingCreate) OnConflictColumns(columns ...string) *SettingUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SettingUpsertOne{
		create: sc,
	}
}

type (
	// SettingUpsertOne is the builder for "upsert"-ing
	//  one Setting node.
	SettingUpsertOne struct {
		create *SettingCreate
	}

	// SettingUpsert is the "OnConflict" setter.
	SettingUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *SettingUpsert) SetUpdatedAt(v time.Time) *SettingUpsert {
	u.Set(setting.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SettingUpsert) UpdateUpdatedAt() *SettingUpsert {
	u.SetExcluded(setting.FieldUpdatedAt)
	return u
}

// SetName sets the "name" field.
func (u *SettingUpsert) SetName(v string) *SettingUpsert {
	u.Set(setting.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SettingUpsert) UpdateName() *SettingUpsert {
	u.SetExcluded(setting.FieldName)
	return u
}

// SetSerialNumber sets the "serial_number" field.
func (u *SettingUpsert) SetSerialNumber(v string) *SettingUpsert {
	u.Set(setting.FieldSerialNumber, v)
	return u
}

// UpdateSerialNumber sets the "serial_number" field to the value that was provided on create.
func (u *SettingUpsert) UpdateSerialNumber() *SettingUpsert {
	u.SetExcluded(setting.FieldSerialNumber)
	return u
}

// SetVersion sets the "version" field.
func (u *SettingUpsert) SetVersion(v string) *SettingUpsert {
	u.Set(setting.FieldVersion, v)
	return u
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *SettingUpsert) UpdateVersion() *SettingUpsert {
	u.SetExcluded(setting.FieldVersion)
	return u
}

// SetModel sets the "model" field.
func (u *SettingUpsert) SetModel(v string) *SettingUpsert {
	u.Set(setting.FieldModel, v)
	return u
}

// UpdateModel sets the "model" field to the value that was provided on create.
func (u *SettingUpsert) UpdateModel() *SettingUpsert {
	u.SetExcluded(setting.FieldModel)
	return u
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *SettingUpsert) SetWorkspaceID(v string) *SettingUpsert {
	u.Set(setting.FieldWorkspaceID, v)
	return u
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *SettingUpsert) UpdateWorkspaceID() *SettingUpsert {
	u.SetExcluded(setting.FieldWorkspaceID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(setting.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SettingUpsertOne) UpdateNewValues() *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(setting.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(setting.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Setting.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SettingUpsertOne) Ignore() *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SettingUpsertOne) DoNothing() *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SettingCreate.OnConflict
// documentation for more info.
func (u *SettingUpsertOne) Update(set func(*SettingUpsert)) *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SettingUpsertOne) SetUpdatedAt(v time.Time) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateUpdatedAt() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *SettingUpsertOne) SetName(v string) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateName() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateName()
	})
}

// SetSerialNumber sets the "serial_number" field.
func (u *SettingUpsertOne) SetSerialNumber(v string) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetSerialNumber(v)
	})
}

// UpdateSerialNumber sets the "serial_number" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateSerialNumber() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateSerialNumber()
	})
}

// SetVersion sets the "version" field.
func (u *SettingUpsertOne) SetVersion(v string) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateVersion() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateVersion()
	})
}

// SetModel sets the "model" field.
func (u *SettingUpsertOne) SetModel(v string) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetModel(v)
	})
}

// UpdateModel sets the "model" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateModel() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateModel()
	})
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *SettingUpsertOne) SetWorkspaceID(v string) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetWorkspaceID(v)
	})
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateWorkspaceID() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateWorkspaceID()
	})
}

// Exec executes the query.
func (u *SettingUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SettingCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SettingUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SettingUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SettingUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SettingCreateBulk is the builder for creating many Setting entities in bulk.
type SettingCreateBulk struct {
	config
	err      error
	builders []*SettingCreate
	conflict []sql.ConflictOption
}

// Save creates the Setting entities in the database.
func (scb *SettingCreateBulk) Save(ctx context.Context) ([]*Setting, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Setting, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SettingMutation)
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
func (scb *SettingCreateBulk) SaveX(ctx context.Context) []*Setting {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SettingCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SettingCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Setting.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SettingUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (scb *SettingCreateBulk) OnConflict(opts ...sql.ConflictOption) *SettingUpsertBulk {
	scb.conflict = opts
	return &SettingUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SettingCreateBulk) OnConflictColumns(columns ...string) *SettingUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SettingUpsertBulk{
		create: scb,
	}
}

// SettingUpsertBulk is the builder for "upsert"-ing
// a bulk of Setting nodes.
type SettingUpsertBulk struct {
	create *SettingCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(setting.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SettingUpsertBulk) UpdateNewValues() *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(setting.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(setting.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SettingUpsertBulk) Ignore() *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SettingUpsertBulk) DoNothing() *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SettingCreateBulk.OnConflict
// documentation for more info.
func (u *SettingUpsertBulk) Update(set func(*SettingUpsert)) *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SettingUpsertBulk) SetUpdatedAt(v time.Time) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateUpdatedAt() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetName sets the "name" field.
func (u *SettingUpsertBulk) SetName(v string) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateName() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateName()
	})
}

// SetSerialNumber sets the "serial_number" field.
func (u *SettingUpsertBulk) SetSerialNumber(v string) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetSerialNumber(v)
	})
}

// UpdateSerialNumber sets the "serial_number" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateSerialNumber() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateSerialNumber()
	})
}

// SetVersion sets the "version" field.
func (u *SettingUpsertBulk) SetVersion(v string) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateVersion() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateVersion()
	})
}

// SetModel sets the "model" field.
func (u *SettingUpsertBulk) SetModel(v string) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetModel(v)
	})
}

// UpdateModel sets the "model" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateModel() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateModel()
	})
}

// SetWorkspaceID sets the "workspace_id" field.
func (u *SettingUpsertBulk) SetWorkspaceID(v string) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetWorkspaceID(v)
	})
}

// UpdateWorkspaceID sets the "workspace_id" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateWorkspaceID() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateWorkspaceID()
	})
}

// Exec executes the query.
func (u *SettingUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SettingCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SettingCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SettingUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
