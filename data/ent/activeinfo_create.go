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
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/activeinfo"
)

// ActiveInfoCreate is the builder for creating a ActiveInfo entity.
type ActiveInfoCreate struct {
	config
	mutation *ActiveInfoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (aic *ActiveInfoCreate) SetCreatedAt(t time.Time) *ActiveInfoCreate {
	aic.mutation.SetCreatedAt(t)
	return aic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aic *ActiveInfoCreate) SetNillableCreatedAt(t *time.Time) *ActiveInfoCreate {
	if t != nil {
		aic.SetCreatedAt(*t)
	}
	return aic
}

// SetUpdatedAt sets the "updated_at" field.
func (aic *ActiveInfoCreate) SetUpdatedAt(t time.Time) *ActiveInfoCreate {
	aic.mutation.SetUpdatedAt(t)
	return aic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aic *ActiveInfoCreate) SetNillableUpdatedAt(t *time.Time) *ActiveInfoCreate {
	if t != nil {
		aic.SetUpdatedAt(*t)
	}
	return aic
}

// SetDeletedAt sets the "deleted_at" field.
func (aic *ActiveInfoCreate) SetDeletedAt(t time.Time) *ActiveInfoCreate {
	aic.mutation.SetDeletedAt(t)
	return aic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aic *ActiveInfoCreate) SetNillableDeletedAt(t *time.Time) *ActiveInfoCreate {
	if t != nil {
		aic.SetDeletedAt(*t)
	}
	return aic
}

// SetProcessID sets the "process_id" field.
func (aic *ActiveInfoCreate) SetProcessID(s string) *ActiveInfoCreate {
	aic.mutation.SetProcessID(s)
	return aic
}

// SetStartTime sets the "start_time" field.
func (aic *ActiveInfoCreate) SetStartTime(s string) *ActiveInfoCreate {
	aic.mutation.SetStartTime(s)
	return aic
}

// SetResult sets the "result" field.
func (aic *ActiveInfoCreate) SetResult(s string) *ActiveInfoCreate {
	aic.mutation.SetResult(s)
	return aic
}

// SetMsg sets the "msg" field.
func (aic *ActiveInfoCreate) SetMsg(s string) *ActiveInfoCreate {
	aic.mutation.SetMsg(s)
	return aic
}

// SetID sets the "id" field.
func (aic *ActiveInfoCreate) SetID(u uint64) *ActiveInfoCreate {
	aic.mutation.SetID(u)
	return aic
}

// Mutation returns the ActiveInfoMutation object of the builder.
func (aic *ActiveInfoCreate) Mutation() *ActiveInfoMutation {
	return aic.mutation
}

// Save creates the ActiveInfo in the database.
func (aic *ActiveInfoCreate) Save(ctx context.Context) (*ActiveInfo, error) {
	if err := aic.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, aic.sqlSave, aic.mutation, aic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (aic *ActiveInfoCreate) SaveX(ctx context.Context) *ActiveInfo {
	v, err := aic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aic *ActiveInfoCreate) Exec(ctx context.Context) error {
	_, err := aic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aic *ActiveInfoCreate) ExecX(ctx context.Context) {
	if err := aic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aic *ActiveInfoCreate) defaults() error {
	if _, ok := aic.mutation.CreatedAt(); !ok {
		if activeinfo.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized activeinfo.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := activeinfo.DefaultCreatedAt()
		aic.mutation.SetCreatedAt(v)
	}
	if _, ok := aic.mutation.UpdatedAt(); !ok {
		if activeinfo.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized activeinfo.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := activeinfo.DefaultUpdatedAt()
		aic.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (aic *ActiveInfoCreate) check() error {
	if _, ok := aic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ActiveInfo.created_at"`)}
	}
	if _, ok := aic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ActiveInfo.updated_at"`)}
	}
	if _, ok := aic.mutation.ProcessID(); !ok {
		return &ValidationError{Name: "process_id", err: errors.New(`ent: missing required field "ActiveInfo.process_id"`)}
	}
	if _, ok := aic.mutation.StartTime(); !ok {
		return &ValidationError{Name: "start_time", err: errors.New(`ent: missing required field "ActiveInfo.start_time"`)}
	}
	if _, ok := aic.mutation.Result(); !ok {
		return &ValidationError{Name: "result", err: errors.New(`ent: missing required field "ActiveInfo.result"`)}
	}
	if _, ok := aic.mutation.Msg(); !ok {
		return &ValidationError{Name: "msg", err: errors.New(`ent: missing required field "ActiveInfo.msg"`)}
	}
	return nil
}

func (aic *ActiveInfoCreate) sqlSave(ctx context.Context) (*ActiveInfo, error) {
	if err := aic.check(); err != nil {
		return nil, err
	}
	_node, _spec := aic.createSpec()
	if err := sqlgraph.CreateNode(ctx, aic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	aic.mutation.id = &_node.ID
	aic.mutation.done = true
	return _node, nil
}

func (aic *ActiveInfoCreate) createSpec() (*ActiveInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &ActiveInfo{config: aic.config}
		_spec = sqlgraph.NewCreateSpec(activeinfo.Table, sqlgraph.NewFieldSpec(activeinfo.FieldID, field.TypeUint64))
	)
	_spec.OnConflict = aic.conflict
	if id, ok := aic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := aic.mutation.CreatedAt(); ok {
		_spec.SetField(activeinfo.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := aic.mutation.UpdatedAt(); ok {
		_spec.SetField(activeinfo.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := aic.mutation.DeletedAt(); ok {
		_spec.SetField(activeinfo.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := aic.mutation.ProcessID(); ok {
		_spec.SetField(activeinfo.FieldProcessID, field.TypeString, value)
		_node.ProcessID = value
	}
	if value, ok := aic.mutation.StartTime(); ok {
		_spec.SetField(activeinfo.FieldStartTime, field.TypeString, value)
		_node.StartTime = value
	}
	if value, ok := aic.mutation.Result(); ok {
		_spec.SetField(activeinfo.FieldResult, field.TypeString, value)
		_node.Result = value
	}
	if value, ok := aic.mutation.Msg(); ok {
		_spec.SetField(activeinfo.FieldMsg, field.TypeString, value)
		_node.Msg = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ActiveInfo.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ActiveInfoUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (aic *ActiveInfoCreate) OnConflict(opts ...sql.ConflictOption) *ActiveInfoUpsertOne {
	aic.conflict = opts
	return &ActiveInfoUpsertOne{
		create: aic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ActiveInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (aic *ActiveInfoCreate) OnConflictColumns(columns ...string) *ActiveInfoUpsertOne {
	aic.conflict = append(aic.conflict, sql.ConflictColumns(columns...))
	return &ActiveInfoUpsertOne{
		create: aic,
	}
}

type (
	// ActiveInfoUpsertOne is the builder for "upsert"-ing
	//  one ActiveInfo node.
	ActiveInfoUpsertOne struct {
		create *ActiveInfoCreate
	}

	// ActiveInfoUpsert is the "OnConflict" setter.
	ActiveInfoUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *ActiveInfoUpsert) SetUpdatedAt(v time.Time) *ActiveInfoUpsert {
	u.Set(activeinfo.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ActiveInfoUpsert) UpdateUpdatedAt() *ActiveInfoUpsert {
	u.SetExcluded(activeinfo.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ActiveInfoUpsert) SetDeletedAt(v time.Time) *ActiveInfoUpsert {
	u.Set(activeinfo.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ActiveInfoUpsert) UpdateDeletedAt() *ActiveInfoUpsert {
	u.SetExcluded(activeinfo.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ActiveInfoUpsert) ClearDeletedAt() *ActiveInfoUpsert {
	u.SetNull(activeinfo.FieldDeletedAt)
	return u
}

// SetProcessID sets the "process_id" field.
func (u *ActiveInfoUpsert) SetProcessID(v string) *ActiveInfoUpsert {
	u.Set(activeinfo.FieldProcessID, v)
	return u
}

// UpdateProcessID sets the "process_id" field to the value that was provided on create.
func (u *ActiveInfoUpsert) UpdateProcessID() *ActiveInfoUpsert {
	u.SetExcluded(activeinfo.FieldProcessID)
	return u
}

// SetStartTime sets the "start_time" field.
func (u *ActiveInfoUpsert) SetStartTime(v string) *ActiveInfoUpsert {
	u.Set(activeinfo.FieldStartTime, v)
	return u
}

// UpdateStartTime sets the "start_time" field to the value that was provided on create.
func (u *ActiveInfoUpsert) UpdateStartTime() *ActiveInfoUpsert {
	u.SetExcluded(activeinfo.FieldStartTime)
	return u
}

// SetResult sets the "result" field.
func (u *ActiveInfoUpsert) SetResult(v string) *ActiveInfoUpsert {
	u.Set(activeinfo.FieldResult, v)
	return u
}

// UpdateResult sets the "result" field to the value that was provided on create.
func (u *ActiveInfoUpsert) UpdateResult() *ActiveInfoUpsert {
	u.SetExcluded(activeinfo.FieldResult)
	return u
}

// SetMsg sets the "msg" field.
func (u *ActiveInfoUpsert) SetMsg(v string) *ActiveInfoUpsert {
	u.Set(activeinfo.FieldMsg, v)
	return u
}

// UpdateMsg sets the "msg" field to the value that was provided on create.
func (u *ActiveInfoUpsert) UpdateMsg() *ActiveInfoUpsert {
	u.SetExcluded(activeinfo.FieldMsg)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ActiveInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(activeinfo.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ActiveInfoUpsertOne) UpdateNewValues() *ActiveInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(activeinfo.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(activeinfo.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ActiveInfo.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ActiveInfoUpsertOne) Ignore() *ActiveInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ActiveInfoUpsertOne) DoNothing() *ActiveInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ActiveInfoCreate.OnConflict
// documentation for more info.
func (u *ActiveInfoUpsertOne) Update(set func(*ActiveInfoUpsert)) *ActiveInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ActiveInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ActiveInfoUpsertOne) SetUpdatedAt(v time.Time) *ActiveInfoUpsertOne {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ActiveInfoUpsertOne) UpdateUpdatedAt() *ActiveInfoUpsertOne {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ActiveInfoUpsertOne) SetDeletedAt(v time.Time) *ActiveInfoUpsertOne {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ActiveInfoUpsertOne) UpdateDeletedAt() *ActiveInfoUpsertOne {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ActiveInfoUpsertOne) ClearDeletedAt() *ActiveInfoUpsertOne {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.ClearDeletedAt()
	})
}

// SetProcessID sets the "process_id" field.
func (u *ActiveInfoUpsertOne) SetProcessID(v string) *ActiveInfoUpsertOne {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.SetProcessID(v)
	})
}

// UpdateProcessID sets the "process_id" field to the value that was provided on create.
func (u *ActiveInfoUpsertOne) UpdateProcessID() *ActiveInfoUpsertOne {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.UpdateProcessID()
	})
}

// SetStartTime sets the "start_time" field.
func (u *ActiveInfoUpsertOne) SetStartTime(v string) *ActiveInfoUpsertOne {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.SetStartTime(v)
	})
}

// UpdateStartTime sets the "start_time" field to the value that was provided on create.
func (u *ActiveInfoUpsertOne) UpdateStartTime() *ActiveInfoUpsertOne {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.UpdateStartTime()
	})
}

// SetResult sets the "result" field.
func (u *ActiveInfoUpsertOne) SetResult(v string) *ActiveInfoUpsertOne {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.SetResult(v)
	})
}

// UpdateResult sets the "result" field to the value that was provided on create.
func (u *ActiveInfoUpsertOne) UpdateResult() *ActiveInfoUpsertOne {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.UpdateResult()
	})
}

// SetMsg sets the "msg" field.
func (u *ActiveInfoUpsertOne) SetMsg(v string) *ActiveInfoUpsertOne {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.SetMsg(v)
	})
}

// UpdateMsg sets the "msg" field to the value that was provided on create.
func (u *ActiveInfoUpsertOne) UpdateMsg() *ActiveInfoUpsertOne {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.UpdateMsg()
	})
}

// Exec executes the query.
func (u *ActiveInfoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ActiveInfoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ActiveInfoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ActiveInfoUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ActiveInfoUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ActiveInfoCreateBulk is the builder for creating many ActiveInfo entities in bulk.
type ActiveInfoCreateBulk struct {
	config
	err      error
	builders []*ActiveInfoCreate
	conflict []sql.ConflictOption
}

// Save creates the ActiveInfo entities in the database.
func (aicb *ActiveInfoCreateBulk) Save(ctx context.Context) ([]*ActiveInfo, error) {
	if aicb.err != nil {
		return nil, aicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(aicb.builders))
	nodes := make([]*ActiveInfo, len(aicb.builders))
	mutators := make([]Mutator, len(aicb.builders))
	for i := range aicb.builders {
		func(i int, root context.Context) {
			builder := aicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ActiveInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, aicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = aicb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, aicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aicb *ActiveInfoCreateBulk) SaveX(ctx context.Context) []*ActiveInfo {
	v, err := aicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aicb *ActiveInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := aicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aicb *ActiveInfoCreateBulk) ExecX(ctx context.Context) {
	if err := aicb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ActiveInfo.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ActiveInfoUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (aicb *ActiveInfoCreateBulk) OnConflict(opts ...sql.ConflictOption) *ActiveInfoUpsertBulk {
	aicb.conflict = opts
	return &ActiveInfoUpsertBulk{
		create: aicb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ActiveInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (aicb *ActiveInfoCreateBulk) OnConflictColumns(columns ...string) *ActiveInfoUpsertBulk {
	aicb.conflict = append(aicb.conflict, sql.ConflictColumns(columns...))
	return &ActiveInfoUpsertBulk{
		create: aicb,
	}
}

// ActiveInfoUpsertBulk is the builder for "upsert"-ing
// a bulk of ActiveInfo nodes.
type ActiveInfoUpsertBulk struct {
	create *ActiveInfoCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ActiveInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(activeinfo.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ActiveInfoUpsertBulk) UpdateNewValues() *ActiveInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(activeinfo.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(activeinfo.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ActiveInfo.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ActiveInfoUpsertBulk) Ignore() *ActiveInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ActiveInfoUpsertBulk) DoNothing() *ActiveInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ActiveInfoCreateBulk.OnConflict
// documentation for more info.
func (u *ActiveInfoUpsertBulk) Update(set func(*ActiveInfoUpsert)) *ActiveInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ActiveInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ActiveInfoUpsertBulk) SetUpdatedAt(v time.Time) *ActiveInfoUpsertBulk {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ActiveInfoUpsertBulk) UpdateUpdatedAt() *ActiveInfoUpsertBulk {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ActiveInfoUpsertBulk) SetDeletedAt(v time.Time) *ActiveInfoUpsertBulk {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ActiveInfoUpsertBulk) UpdateDeletedAt() *ActiveInfoUpsertBulk {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *ActiveInfoUpsertBulk) ClearDeletedAt() *ActiveInfoUpsertBulk {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.ClearDeletedAt()
	})
}

// SetProcessID sets the "process_id" field.
func (u *ActiveInfoUpsertBulk) SetProcessID(v string) *ActiveInfoUpsertBulk {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.SetProcessID(v)
	})
}

// UpdateProcessID sets the "process_id" field to the value that was provided on create.
func (u *ActiveInfoUpsertBulk) UpdateProcessID() *ActiveInfoUpsertBulk {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.UpdateProcessID()
	})
}

// SetStartTime sets the "start_time" field.
func (u *ActiveInfoUpsertBulk) SetStartTime(v string) *ActiveInfoUpsertBulk {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.SetStartTime(v)
	})
}

// UpdateStartTime sets the "start_time" field to the value that was provided on create.
func (u *ActiveInfoUpsertBulk) UpdateStartTime() *ActiveInfoUpsertBulk {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.UpdateStartTime()
	})
}

// SetResult sets the "result" field.
func (u *ActiveInfoUpsertBulk) SetResult(v string) *ActiveInfoUpsertBulk {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.SetResult(v)
	})
}

// UpdateResult sets the "result" field to the value that was provided on create.
func (u *ActiveInfoUpsertBulk) UpdateResult() *ActiveInfoUpsertBulk {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.UpdateResult()
	})
}

// SetMsg sets the "msg" field.
func (u *ActiveInfoUpsertBulk) SetMsg(v string) *ActiveInfoUpsertBulk {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.SetMsg(v)
	})
}

// UpdateMsg sets the "msg" field to the value that was provided on create.
func (u *ActiveInfoUpsertBulk) UpdateMsg() *ActiveInfoUpsertBulk {
	return u.Update(func(s *ActiveInfoUpsert) {
		s.UpdateMsg()
	})
}

// Exec executes the query.
func (u *ActiveInfoUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ActiveInfoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ActiveInfoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ActiveInfoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
