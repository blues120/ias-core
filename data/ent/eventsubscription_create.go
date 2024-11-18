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
	"github.com/blues120/ias-core/biz"
	"github.com/blues120/ias-core/data/ent/eventsubscription"
)

// EventSubscriptionCreate is the builder for creating a EventSubscription entity.
type EventSubscriptionCreate struct {
	config
	mutation *EventSubscriptionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (esc *EventSubscriptionCreate) SetCreatedAt(t time.Time) *EventSubscriptionCreate {
	esc.mutation.SetCreatedAt(t)
	return esc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (esc *EventSubscriptionCreate) SetNillableCreatedAt(t *time.Time) *EventSubscriptionCreate {
	if t != nil {
		esc.SetCreatedAt(*t)
	}
	return esc
}

// SetUpdatedAt sets the "updated_at" field.
func (esc *EventSubscriptionCreate) SetUpdatedAt(t time.Time) *EventSubscriptionCreate {
	esc.mutation.SetUpdatedAt(t)
	return esc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (esc *EventSubscriptionCreate) SetNillableUpdatedAt(t *time.Time) *EventSubscriptionCreate {
	if t != nil {
		esc.SetUpdatedAt(*t)
	}
	return esc
}

// SetDeletedAt sets the "deleted_at" field.
func (esc *EventSubscriptionCreate) SetDeletedAt(t time.Time) *EventSubscriptionCreate {
	esc.mutation.SetDeletedAt(t)
	return esc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (esc *EventSubscriptionCreate) SetNillableDeletedAt(t *time.Time) *EventSubscriptionCreate {
	if t != nil {
		esc.SetDeletedAt(*t)
	}
	return esc
}

// SetBoxID sets the "box_id" field.
func (esc *EventSubscriptionCreate) SetBoxID(s string) *EventSubscriptionCreate {
	esc.mutation.SetBoxID(s)
	return esc
}

// SetChannelID sets the "channel_id" field.
func (esc *EventSubscriptionCreate) SetChannelID(s string) *EventSubscriptionCreate {
	esc.mutation.SetChannelID(s)
	return esc
}

// SetCallback sets the "callback" field.
func (esc *EventSubscriptionCreate) SetCallback(s string) *EventSubscriptionCreate {
	esc.mutation.SetCallback(s)
	return esc
}

// SetTemplateID sets the "template_id" field.
func (esc *EventSubscriptionCreate) SetTemplateID(s string) *EventSubscriptionCreate {
	esc.mutation.SetTemplateID(s)
	return esc
}

// SetStatus sets the "status" field.
func (esc *EventSubscriptionCreate) SetStatus(bss biz.EventSubStatus) *EventSubscriptionCreate {
	esc.mutation.SetStatus(bss)
	return esc
}

// SetID sets the "id" field.
func (esc *EventSubscriptionCreate) SetID(u uint64) *EventSubscriptionCreate {
	esc.mutation.SetID(u)
	return esc
}

// Mutation returns the EventSubscriptionMutation object of the builder.
func (esc *EventSubscriptionCreate) Mutation() *EventSubscriptionMutation {
	return esc.mutation
}

// Save creates the EventSubscription in the database.
func (esc *EventSubscriptionCreate) Save(ctx context.Context) (*EventSubscription, error) {
	if err := esc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, esc.sqlSave, esc.mutation, esc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (esc *EventSubscriptionCreate) SaveX(ctx context.Context) *EventSubscription {
	v, err := esc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (esc *EventSubscriptionCreate) Exec(ctx context.Context) error {
	_, err := esc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (esc *EventSubscriptionCreate) ExecX(ctx context.Context) {
	if err := esc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (esc *EventSubscriptionCreate) defaults() error {
	if _, ok := esc.mutation.CreatedAt(); !ok {
		if eventsubscription.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized eventsubscription.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := eventsubscription.DefaultCreatedAt()
		esc.mutation.SetCreatedAt(v)
	}
	if _, ok := esc.mutation.UpdatedAt(); !ok {
		if eventsubscription.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized eventsubscription.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := eventsubscription.DefaultUpdatedAt()
		esc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (esc *EventSubscriptionCreate) check() error {
	if _, ok := esc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "EventSubscription.created_at"`)}
	}
	if _, ok := esc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "EventSubscription.updated_at"`)}
	}
	if _, ok := esc.mutation.BoxID(); !ok {
		return &ValidationError{Name: "box_id", err: errors.New(`ent: missing required field "EventSubscription.box_id"`)}
	}
	if _, ok := esc.mutation.ChannelID(); !ok {
		return &ValidationError{Name: "channel_id", err: errors.New(`ent: missing required field "EventSubscription.channel_id"`)}
	}
	if _, ok := esc.mutation.Callback(); !ok {
		return &ValidationError{Name: "callback", err: errors.New(`ent: missing required field "EventSubscription.callback"`)}
	}
	if _, ok := esc.mutation.TemplateID(); !ok {
		return &ValidationError{Name: "template_id", err: errors.New(`ent: missing required field "EventSubscription.template_id"`)}
	}
	if _, ok := esc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "EventSubscription.status"`)}
	}
	if v, ok := esc.mutation.Status(); ok {
		if err := eventsubscription.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "EventSubscription.status": %w`, err)}
		}
	}
	return nil
}

func (esc *EventSubscriptionCreate) sqlSave(ctx context.Context) (*EventSubscription, error) {
	if err := esc.check(); err != nil {
		return nil, err
	}
	_node, _spec := esc.createSpec()
	if err := sqlgraph.CreateNode(ctx, esc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	esc.mutation.id = &_node.ID
	esc.mutation.done = true
	return _node, nil
}

func (esc *EventSubscriptionCreate) createSpec() (*EventSubscription, *sqlgraph.CreateSpec) {
	var (
		_node = &EventSubscription{config: esc.config}
		_spec = sqlgraph.NewCreateSpec(eventsubscription.Table, sqlgraph.NewFieldSpec(eventsubscription.FieldID, field.TypeUint64))
	)
	_spec.OnConflict = esc.conflict
	if id, ok := esc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := esc.mutation.CreatedAt(); ok {
		_spec.SetField(eventsubscription.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := esc.mutation.UpdatedAt(); ok {
		_spec.SetField(eventsubscription.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := esc.mutation.DeletedAt(); ok {
		_spec.SetField(eventsubscription.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := esc.mutation.BoxID(); ok {
		_spec.SetField(eventsubscription.FieldBoxID, field.TypeString, value)
		_node.BoxID = value
	}
	if value, ok := esc.mutation.ChannelID(); ok {
		_spec.SetField(eventsubscription.FieldChannelID, field.TypeString, value)
		_node.ChannelID = value
	}
	if value, ok := esc.mutation.Callback(); ok {
		_spec.SetField(eventsubscription.FieldCallback, field.TypeString, value)
		_node.Callback = value
	}
	if value, ok := esc.mutation.TemplateID(); ok {
		_spec.SetField(eventsubscription.FieldTemplateID, field.TypeString, value)
		_node.TemplateID = value
	}
	if value, ok := esc.mutation.Status(); ok {
		_spec.SetField(eventsubscription.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.EventSubscription.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EventSubscriptionUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (esc *EventSubscriptionCreate) OnConflict(opts ...sql.ConflictOption) *EventSubscriptionUpsertOne {
	esc.conflict = opts
	return &EventSubscriptionUpsertOne{
		create: esc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.EventSubscription.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (esc *EventSubscriptionCreate) OnConflictColumns(columns ...string) *EventSubscriptionUpsertOne {
	esc.conflict = append(esc.conflict, sql.ConflictColumns(columns...))
	return &EventSubscriptionUpsertOne{
		create: esc,
	}
}

type (
	// EventSubscriptionUpsertOne is the builder for "upsert"-ing
	//  one EventSubscription node.
	EventSubscriptionUpsertOne struct {
		create *EventSubscriptionCreate
	}

	// EventSubscriptionUpsert is the "OnConflict" setter.
	EventSubscriptionUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *EventSubscriptionUpsert) SetUpdatedAt(v time.Time) *EventSubscriptionUpsert {
	u.Set(eventsubscription.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EventSubscriptionUpsert) UpdateUpdatedAt() *EventSubscriptionUpsert {
	u.SetExcluded(eventsubscription.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *EventSubscriptionUpsert) SetDeletedAt(v time.Time) *EventSubscriptionUpsert {
	u.Set(eventsubscription.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *EventSubscriptionUpsert) UpdateDeletedAt() *EventSubscriptionUpsert {
	u.SetExcluded(eventsubscription.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *EventSubscriptionUpsert) ClearDeletedAt() *EventSubscriptionUpsert {
	u.SetNull(eventsubscription.FieldDeletedAt)
	return u
}

// SetBoxID sets the "box_id" field.
func (u *EventSubscriptionUpsert) SetBoxID(v string) *EventSubscriptionUpsert {
	u.Set(eventsubscription.FieldBoxID, v)
	return u
}

// UpdateBoxID sets the "box_id" field to the value that was provided on create.
func (u *EventSubscriptionUpsert) UpdateBoxID() *EventSubscriptionUpsert {
	u.SetExcluded(eventsubscription.FieldBoxID)
	return u
}

// SetChannelID sets the "channel_id" field.
func (u *EventSubscriptionUpsert) SetChannelID(v string) *EventSubscriptionUpsert {
	u.Set(eventsubscription.FieldChannelID, v)
	return u
}

// UpdateChannelID sets the "channel_id" field to the value that was provided on create.
func (u *EventSubscriptionUpsert) UpdateChannelID() *EventSubscriptionUpsert {
	u.SetExcluded(eventsubscription.FieldChannelID)
	return u
}

// SetCallback sets the "callback" field.
func (u *EventSubscriptionUpsert) SetCallback(v string) *EventSubscriptionUpsert {
	u.Set(eventsubscription.FieldCallback, v)
	return u
}

// UpdateCallback sets the "callback" field to the value that was provided on create.
func (u *EventSubscriptionUpsert) UpdateCallback() *EventSubscriptionUpsert {
	u.SetExcluded(eventsubscription.FieldCallback)
	return u
}

// SetTemplateID sets the "template_id" field.
func (u *EventSubscriptionUpsert) SetTemplateID(v string) *EventSubscriptionUpsert {
	u.Set(eventsubscription.FieldTemplateID, v)
	return u
}

// UpdateTemplateID sets the "template_id" field to the value that was provided on create.
func (u *EventSubscriptionUpsert) UpdateTemplateID() *EventSubscriptionUpsert {
	u.SetExcluded(eventsubscription.FieldTemplateID)
	return u
}

// SetStatus sets the "status" field.
func (u *EventSubscriptionUpsert) SetStatus(v biz.EventSubStatus) *EventSubscriptionUpsert {
	u.Set(eventsubscription.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *EventSubscriptionUpsert) UpdateStatus() *EventSubscriptionUpsert {
	u.SetExcluded(eventsubscription.FieldStatus)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.EventSubscription.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(eventsubscription.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EventSubscriptionUpsertOne) UpdateNewValues() *EventSubscriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(eventsubscription.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(eventsubscription.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.EventSubscription.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *EventSubscriptionUpsertOne) Ignore() *EventSubscriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EventSubscriptionUpsertOne) DoNothing() *EventSubscriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EventSubscriptionCreate.OnConflict
// documentation for more info.
func (u *EventSubscriptionUpsertOne) Update(set func(*EventSubscriptionUpsert)) *EventSubscriptionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EventSubscriptionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EventSubscriptionUpsertOne) SetUpdatedAt(v time.Time) *EventSubscriptionUpsertOne {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EventSubscriptionUpsertOne) UpdateUpdatedAt() *EventSubscriptionUpsertOne {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *EventSubscriptionUpsertOne) SetDeletedAt(v time.Time) *EventSubscriptionUpsertOne {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *EventSubscriptionUpsertOne) UpdateDeletedAt() *EventSubscriptionUpsertOne {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *EventSubscriptionUpsertOne) ClearDeletedAt() *EventSubscriptionUpsertOne {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.ClearDeletedAt()
	})
}

// SetBoxID sets the "box_id" field.
func (u *EventSubscriptionUpsertOne) SetBoxID(v string) *EventSubscriptionUpsertOne {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.SetBoxID(v)
	})
}

// UpdateBoxID sets the "box_id" field to the value that was provided on create.
func (u *EventSubscriptionUpsertOne) UpdateBoxID() *EventSubscriptionUpsertOne {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.UpdateBoxID()
	})
}

// SetChannelID sets the "channel_id" field.
func (u *EventSubscriptionUpsertOne) SetChannelID(v string) *EventSubscriptionUpsertOne {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.SetChannelID(v)
	})
}

// UpdateChannelID sets the "channel_id" field to the value that was provided on create.
func (u *EventSubscriptionUpsertOne) UpdateChannelID() *EventSubscriptionUpsertOne {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.UpdateChannelID()
	})
}

// SetCallback sets the "callback" field.
func (u *EventSubscriptionUpsertOne) SetCallback(v string) *EventSubscriptionUpsertOne {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.SetCallback(v)
	})
}

// UpdateCallback sets the "callback" field to the value that was provided on create.
func (u *EventSubscriptionUpsertOne) UpdateCallback() *EventSubscriptionUpsertOne {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.UpdateCallback()
	})
}

// SetTemplateID sets the "template_id" field.
func (u *EventSubscriptionUpsertOne) SetTemplateID(v string) *EventSubscriptionUpsertOne {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.SetTemplateID(v)
	})
}

// UpdateTemplateID sets the "template_id" field to the value that was provided on create.
func (u *EventSubscriptionUpsertOne) UpdateTemplateID() *EventSubscriptionUpsertOne {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.UpdateTemplateID()
	})
}

// SetStatus sets the "status" field.
func (u *EventSubscriptionUpsertOne) SetStatus(v biz.EventSubStatus) *EventSubscriptionUpsertOne {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *EventSubscriptionUpsertOne) UpdateStatus() *EventSubscriptionUpsertOne {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *EventSubscriptionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EventSubscriptionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EventSubscriptionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EventSubscriptionUpsertOne) ID(ctx context.Context) (id uint64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EventSubscriptionUpsertOne) IDX(ctx context.Context) uint64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EventSubscriptionCreateBulk is the builder for creating many EventSubscription entities in bulk.
type EventSubscriptionCreateBulk struct {
	config
	err      error
	builders []*EventSubscriptionCreate
	conflict []sql.ConflictOption
}

// Save creates the EventSubscription entities in the database.
func (escb *EventSubscriptionCreateBulk) Save(ctx context.Context) ([]*EventSubscription, error) {
	if escb.err != nil {
		return nil, escb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(escb.builders))
	nodes := make([]*EventSubscription, len(escb.builders))
	mutators := make([]Mutator, len(escb.builders))
	for i := range escb.builders {
		func(i int, root context.Context) {
			builder := escb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventSubscriptionMutation)
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
					_, err = mutators[i+1].Mutate(root, escb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = escb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, escb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, escb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (escb *EventSubscriptionCreateBulk) SaveX(ctx context.Context) []*EventSubscription {
	v, err := escb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (escb *EventSubscriptionCreateBulk) Exec(ctx context.Context) error {
	_, err := escb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (escb *EventSubscriptionCreateBulk) ExecX(ctx context.Context) {
	if err := escb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.EventSubscription.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EventSubscriptionUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (escb *EventSubscriptionCreateBulk) OnConflict(opts ...sql.ConflictOption) *EventSubscriptionUpsertBulk {
	escb.conflict = opts
	return &EventSubscriptionUpsertBulk{
		create: escb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.EventSubscription.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (escb *EventSubscriptionCreateBulk) OnConflictColumns(columns ...string) *EventSubscriptionUpsertBulk {
	escb.conflict = append(escb.conflict, sql.ConflictColumns(columns...))
	return &EventSubscriptionUpsertBulk{
		create: escb,
	}
}

// EventSubscriptionUpsertBulk is the builder for "upsert"-ing
// a bulk of EventSubscription nodes.
type EventSubscriptionUpsertBulk struct {
	create *EventSubscriptionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.EventSubscription.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(eventsubscription.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *EventSubscriptionUpsertBulk) UpdateNewValues() *EventSubscriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(eventsubscription.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(eventsubscription.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.EventSubscription.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *EventSubscriptionUpsertBulk) Ignore() *EventSubscriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EventSubscriptionUpsertBulk) DoNothing() *EventSubscriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EventSubscriptionCreateBulk.OnConflict
// documentation for more info.
func (u *EventSubscriptionUpsertBulk) Update(set func(*EventSubscriptionUpsert)) *EventSubscriptionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EventSubscriptionUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *EventSubscriptionUpsertBulk) SetUpdatedAt(v time.Time) *EventSubscriptionUpsertBulk {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *EventSubscriptionUpsertBulk) UpdateUpdatedAt() *EventSubscriptionUpsertBulk {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *EventSubscriptionUpsertBulk) SetDeletedAt(v time.Time) *EventSubscriptionUpsertBulk {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *EventSubscriptionUpsertBulk) UpdateDeletedAt() *EventSubscriptionUpsertBulk {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *EventSubscriptionUpsertBulk) ClearDeletedAt() *EventSubscriptionUpsertBulk {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.ClearDeletedAt()
	})
}

// SetBoxID sets the "box_id" field.
func (u *EventSubscriptionUpsertBulk) SetBoxID(v string) *EventSubscriptionUpsertBulk {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.SetBoxID(v)
	})
}

// UpdateBoxID sets the "box_id" field to the value that was provided on create.
func (u *EventSubscriptionUpsertBulk) UpdateBoxID() *EventSubscriptionUpsertBulk {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.UpdateBoxID()
	})
}

// SetChannelID sets the "channel_id" field.
func (u *EventSubscriptionUpsertBulk) SetChannelID(v string) *EventSubscriptionUpsertBulk {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.SetChannelID(v)
	})
}

// UpdateChannelID sets the "channel_id" field to the value that was provided on create.
func (u *EventSubscriptionUpsertBulk) UpdateChannelID() *EventSubscriptionUpsertBulk {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.UpdateChannelID()
	})
}

// SetCallback sets the "callback" field.
func (u *EventSubscriptionUpsertBulk) SetCallback(v string) *EventSubscriptionUpsertBulk {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.SetCallback(v)
	})
}

// UpdateCallback sets the "callback" field to the value that was provided on create.
func (u *EventSubscriptionUpsertBulk) UpdateCallback() *EventSubscriptionUpsertBulk {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.UpdateCallback()
	})
}

// SetTemplateID sets the "template_id" field.
func (u *EventSubscriptionUpsertBulk) SetTemplateID(v string) *EventSubscriptionUpsertBulk {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.SetTemplateID(v)
	})
}

// UpdateTemplateID sets the "template_id" field to the value that was provided on create.
func (u *EventSubscriptionUpsertBulk) UpdateTemplateID() *EventSubscriptionUpsertBulk {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.UpdateTemplateID()
	})
}

// SetStatus sets the "status" field.
func (u *EventSubscriptionUpsertBulk) SetStatus(v biz.EventSubStatus) *EventSubscriptionUpsertBulk {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *EventSubscriptionUpsertBulk) UpdateStatus() *EventSubscriptionUpsertBulk {
	return u.Update(func(s *EventSubscriptionUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *EventSubscriptionUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the EventSubscriptionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EventSubscriptionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EventSubscriptionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
