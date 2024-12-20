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
	"github.com/blues120/ias-core/data/ent/inform"
	"github.com/blues120/ias-core/data/ent/predicate"
)

// InformUpdate is the builder for updating Inform entities.
type InformUpdate struct {
	config
	hooks     []Hook
	mutation  *InformMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the InformUpdate builder.
func (iu *InformUpdate) Where(ps ...predicate.Inform) *InformUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetTenantID sets the "tenant_id" field.
func (iu *InformUpdate) SetTenantID(s string) *InformUpdate {
	iu.mutation.SetTenantID(s)
	return iu
}

// SetNillableTenantID sets the "tenant_id" field if the given value is not nil.
func (iu *InformUpdate) SetNillableTenantID(s *string) *InformUpdate {
	if s != nil {
		iu.SetTenantID(*s)
	}
	return iu
}

// ClearTenantID clears the value of the "tenant_id" field.
func (iu *InformUpdate) ClearTenantID() *InformUpdate {
	iu.mutation.ClearTenantID()
	return iu
}

// SetAccessOrgList sets the "access_org_list" field.
func (iu *InformUpdate) SetAccessOrgList(s string) *InformUpdate {
	iu.mutation.SetAccessOrgList(s)
	return iu
}

// SetNillableAccessOrgList sets the "access_org_list" field if the given value is not nil.
func (iu *InformUpdate) SetNillableAccessOrgList(s *string) *InformUpdate {
	if s != nil {
		iu.SetAccessOrgList(*s)
	}
	return iu
}

// ClearAccessOrgList clears the value of the "access_org_list" field.
func (iu *InformUpdate) ClearAccessOrgList() *InformUpdate {
	iu.mutation.ClearAccessOrgList()
	return iu
}

// SetCreatedAt sets the "created_at" field.
func (iu *InformUpdate) SetCreatedAt(t time.Time) *InformUpdate {
	iu.mutation.SetCreatedAt(t)
	return iu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (iu *InformUpdate) SetNillableCreatedAt(t *time.Time) *InformUpdate {
	if t != nil {
		iu.SetCreatedAt(*t)
	}
	return iu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (iu *InformUpdate) ClearCreatedAt() *InformUpdate {
	iu.mutation.ClearCreatedAt()
	return iu
}

// SetUpdatedAt sets the "updated_at" field.
func (iu *InformUpdate) SetUpdatedAt(t time.Time) *InformUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (iu *InformUpdate) SetNillableUpdatedAt(t *time.Time) *InformUpdate {
	if t != nil {
		iu.SetUpdatedAt(*t)
	}
	return iu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (iu *InformUpdate) ClearUpdatedAt() *InformUpdate {
	iu.mutation.ClearUpdatedAt()
	return iu
}

// SetDeletedAt sets the "deleted_at" field.
func (iu *InformUpdate) SetDeletedAt(t time.Time) *InformUpdate {
	iu.mutation.SetDeletedAt(t)
	return iu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (iu *InformUpdate) SetNillableDeletedAt(t *time.Time) *InformUpdate {
	if t != nil {
		iu.SetDeletedAt(*t)
	}
	return iu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (iu *InformUpdate) ClearDeletedAt() *InformUpdate {
	iu.mutation.ClearDeletedAt()
	return iu
}

// SetAlarmName sets the "alarm_name" field.
func (iu *InformUpdate) SetAlarmName(s string) *InformUpdate {
	iu.mutation.SetAlarmName(s)
	return iu
}

// SetNillableAlarmName sets the "alarm_name" field if the given value is not nil.
func (iu *InformUpdate) SetNillableAlarmName(s *string) *InformUpdate {
	if s != nil {
		iu.SetAlarmName(*s)
	}
	return iu
}

// SetAlarmType sets the "alarm_type" field.
func (iu *InformUpdate) SetAlarmType(s string) *InformUpdate {
	iu.mutation.SetAlarmType(s)
	return iu
}

// SetNillableAlarmType sets the "alarm_type" field if the given value is not nil.
func (iu *InformUpdate) SetNillableAlarmType(s *string) *InformUpdate {
	if s != nil {
		iu.SetAlarmType(*s)
	}
	return iu
}

// SetSignName sets the "sign_name" field.
func (iu *InformUpdate) SetSignName(s string) *InformUpdate {
	iu.mutation.SetSignName(s)
	return iu
}

// SetNillableSignName sets the "sign_name" field if the given value is not nil.
func (iu *InformUpdate) SetNillableSignName(s *string) *InformUpdate {
	if s != nil {
		iu.SetSignName(*s)
	}
	return iu
}

// SetNotifyTemplate sets the "notify_template" field.
func (iu *InformUpdate) SetNotifyTemplate(s string) *InformUpdate {
	iu.mutation.SetNotifyTemplate(s)
	return iu
}

// SetNillableNotifyTemplate sets the "notify_template" field if the given value is not nil.
func (iu *InformUpdate) SetNillableNotifyTemplate(s *string) *InformUpdate {
	if s != nil {
		iu.SetNotifyTemplate(*s)
	}
	return iu
}

// SetTemplateCode sets the "template_code" field.
func (iu *InformUpdate) SetTemplateCode(s string) *InformUpdate {
	iu.mutation.SetTemplateCode(s)
	return iu
}

// SetNillableTemplateCode sets the "template_code" field if the given value is not nil.
func (iu *InformUpdate) SetNillableTemplateCode(s *string) *InformUpdate {
	if s != nil {
		iu.SetTemplateCode(*s)
	}
	return iu
}

// SetPhoneNumbers sets the "phone_numbers" field.
func (iu *InformUpdate) SetPhoneNumbers(s string) *InformUpdate {
	iu.mutation.SetPhoneNumbers(s)
	return iu
}

// SetNillablePhoneNumbers sets the "phone_numbers" field if the given value is not nil.
func (iu *InformUpdate) SetNillablePhoneNumbers(s *string) *InformUpdate {
	if s != nil {
		iu.SetPhoneNumbers(*s)
	}
	return iu
}

// SetNotifySwitch sets the "notify_switch" field.
func (iu *InformUpdate) SetNotifySwitch(s string) *InformUpdate {
	iu.mutation.SetNotifySwitch(s)
	return iu
}

// SetNillableNotifySwitch sets the "notify_switch" field if the given value is not nil.
func (iu *InformUpdate) SetNillableNotifySwitch(s *string) *InformUpdate {
	if s != nil {
		iu.SetNotifySwitch(*s)
	}
	return iu
}

// SetTaskName sets the "task_name" field.
func (iu *InformUpdate) SetTaskName(s string) *InformUpdate {
	iu.mutation.SetTaskName(s)
	return iu
}

// SetNillableTaskName sets the "task_name" field if the given value is not nil.
func (iu *InformUpdate) SetNillableTaskName(s *string) *InformUpdate {
	if s != nil {
		iu.SetTaskName(*s)
	}
	return iu
}

// SetTaskID sets the "task_id" field.
func (iu *InformUpdate) SetTaskID(u uint64) *InformUpdate {
	iu.mutation.ResetTaskID()
	iu.mutation.SetTaskID(u)
	return iu
}

// SetNillableTaskID sets the "task_id" field if the given value is not nil.
func (iu *InformUpdate) SetNillableTaskID(u *uint64) *InformUpdate {
	if u != nil {
		iu.SetTaskID(*u)
	}
	return iu
}

// AddTaskID adds u to the "task_id" field.
func (iu *InformUpdate) AddTaskID(u int64) *InformUpdate {
	iu.mutation.AddTaskID(u)
	return iu
}

// Mutation returns the InformMutation object of the builder.
func (iu *InformUpdate) Mutation() *InformMutation {
	return iu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *InformUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InformUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InformUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InformUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (iu *InformUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *InformUpdate {
	iu.modifiers = append(iu.modifiers, modifiers...)
	return iu
}

func (iu *InformUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(inform.Table, inform.Columns, sqlgraph.NewFieldSpec(inform.FieldID, field.TypeUint64))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.TenantID(); ok {
		_spec.SetField(inform.FieldTenantID, field.TypeString, value)
	}
	if iu.mutation.TenantIDCleared() {
		_spec.ClearField(inform.FieldTenantID, field.TypeString)
	}
	if value, ok := iu.mutation.AccessOrgList(); ok {
		_spec.SetField(inform.FieldAccessOrgList, field.TypeString, value)
	}
	if iu.mutation.AccessOrgListCleared() {
		_spec.ClearField(inform.FieldAccessOrgList, field.TypeString)
	}
	if value, ok := iu.mutation.CreatedAt(); ok {
		_spec.SetField(inform.FieldCreatedAt, field.TypeTime, value)
	}
	if iu.mutation.CreatedAtCleared() {
		_spec.ClearField(inform.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.SetField(inform.FieldUpdatedAt, field.TypeTime, value)
	}
	if iu.mutation.UpdatedAtCleared() {
		_spec.ClearField(inform.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := iu.mutation.DeletedAt(); ok {
		_spec.SetField(inform.FieldDeletedAt, field.TypeTime, value)
	}
	if iu.mutation.DeletedAtCleared() {
		_spec.ClearField(inform.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := iu.mutation.AlarmName(); ok {
		_spec.SetField(inform.FieldAlarmName, field.TypeString, value)
	}
	if value, ok := iu.mutation.AlarmType(); ok {
		_spec.SetField(inform.FieldAlarmType, field.TypeString, value)
	}
	if value, ok := iu.mutation.SignName(); ok {
		_spec.SetField(inform.FieldSignName, field.TypeString, value)
	}
	if value, ok := iu.mutation.NotifyTemplate(); ok {
		_spec.SetField(inform.FieldNotifyTemplate, field.TypeString, value)
	}
	if value, ok := iu.mutation.TemplateCode(); ok {
		_spec.SetField(inform.FieldTemplateCode, field.TypeString, value)
	}
	if value, ok := iu.mutation.PhoneNumbers(); ok {
		_spec.SetField(inform.FieldPhoneNumbers, field.TypeString, value)
	}
	if value, ok := iu.mutation.NotifySwitch(); ok {
		_spec.SetField(inform.FieldNotifySwitch, field.TypeString, value)
	}
	if value, ok := iu.mutation.TaskName(); ok {
		_spec.SetField(inform.FieldTaskName, field.TypeString, value)
	}
	if value, ok := iu.mutation.TaskID(); ok {
		_spec.SetField(inform.FieldTaskID, field.TypeUint64, value)
	}
	if value, ok := iu.mutation.AddedTaskID(); ok {
		_spec.AddField(inform.FieldTaskID, field.TypeUint64, value)
	}
	_spec.AddModifiers(iu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{inform.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// InformUpdateOne is the builder for updating a single Inform entity.
type InformUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *InformMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetTenantID sets the "tenant_id" field.
func (iuo *InformUpdateOne) SetTenantID(s string) *InformUpdateOne {
	iuo.mutation.SetTenantID(s)
	return iuo
}

// SetNillableTenantID sets the "tenant_id" field if the given value is not nil.
func (iuo *InformUpdateOne) SetNillableTenantID(s *string) *InformUpdateOne {
	if s != nil {
		iuo.SetTenantID(*s)
	}
	return iuo
}

// ClearTenantID clears the value of the "tenant_id" field.
func (iuo *InformUpdateOne) ClearTenantID() *InformUpdateOne {
	iuo.mutation.ClearTenantID()
	return iuo
}

// SetAccessOrgList sets the "access_org_list" field.
func (iuo *InformUpdateOne) SetAccessOrgList(s string) *InformUpdateOne {
	iuo.mutation.SetAccessOrgList(s)
	return iuo
}

// SetNillableAccessOrgList sets the "access_org_list" field if the given value is not nil.
func (iuo *InformUpdateOne) SetNillableAccessOrgList(s *string) *InformUpdateOne {
	if s != nil {
		iuo.SetAccessOrgList(*s)
	}
	return iuo
}

// ClearAccessOrgList clears the value of the "access_org_list" field.
func (iuo *InformUpdateOne) ClearAccessOrgList() *InformUpdateOne {
	iuo.mutation.ClearAccessOrgList()
	return iuo
}

// SetCreatedAt sets the "created_at" field.
func (iuo *InformUpdateOne) SetCreatedAt(t time.Time) *InformUpdateOne {
	iuo.mutation.SetCreatedAt(t)
	return iuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (iuo *InformUpdateOne) SetNillableCreatedAt(t *time.Time) *InformUpdateOne {
	if t != nil {
		iuo.SetCreatedAt(*t)
	}
	return iuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (iuo *InformUpdateOne) ClearCreatedAt() *InformUpdateOne {
	iuo.mutation.ClearCreatedAt()
	return iuo
}

// SetUpdatedAt sets the "updated_at" field.
func (iuo *InformUpdateOne) SetUpdatedAt(t time.Time) *InformUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (iuo *InformUpdateOne) SetNillableUpdatedAt(t *time.Time) *InformUpdateOne {
	if t != nil {
		iuo.SetUpdatedAt(*t)
	}
	return iuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (iuo *InformUpdateOne) ClearUpdatedAt() *InformUpdateOne {
	iuo.mutation.ClearUpdatedAt()
	return iuo
}

// SetDeletedAt sets the "deleted_at" field.
func (iuo *InformUpdateOne) SetDeletedAt(t time.Time) *InformUpdateOne {
	iuo.mutation.SetDeletedAt(t)
	return iuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (iuo *InformUpdateOne) SetNillableDeletedAt(t *time.Time) *InformUpdateOne {
	if t != nil {
		iuo.SetDeletedAt(*t)
	}
	return iuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (iuo *InformUpdateOne) ClearDeletedAt() *InformUpdateOne {
	iuo.mutation.ClearDeletedAt()
	return iuo
}

// SetAlarmName sets the "alarm_name" field.
func (iuo *InformUpdateOne) SetAlarmName(s string) *InformUpdateOne {
	iuo.mutation.SetAlarmName(s)
	return iuo
}

// SetNillableAlarmName sets the "alarm_name" field if the given value is not nil.
func (iuo *InformUpdateOne) SetNillableAlarmName(s *string) *InformUpdateOne {
	if s != nil {
		iuo.SetAlarmName(*s)
	}
	return iuo
}

// SetAlarmType sets the "alarm_type" field.
func (iuo *InformUpdateOne) SetAlarmType(s string) *InformUpdateOne {
	iuo.mutation.SetAlarmType(s)
	return iuo
}

// SetNillableAlarmType sets the "alarm_type" field if the given value is not nil.
func (iuo *InformUpdateOne) SetNillableAlarmType(s *string) *InformUpdateOne {
	if s != nil {
		iuo.SetAlarmType(*s)
	}
	return iuo
}

// SetSignName sets the "sign_name" field.
func (iuo *InformUpdateOne) SetSignName(s string) *InformUpdateOne {
	iuo.mutation.SetSignName(s)
	return iuo
}

// SetNillableSignName sets the "sign_name" field if the given value is not nil.
func (iuo *InformUpdateOne) SetNillableSignName(s *string) *InformUpdateOne {
	if s != nil {
		iuo.SetSignName(*s)
	}
	return iuo
}

// SetNotifyTemplate sets the "notify_template" field.
func (iuo *InformUpdateOne) SetNotifyTemplate(s string) *InformUpdateOne {
	iuo.mutation.SetNotifyTemplate(s)
	return iuo
}

// SetNillableNotifyTemplate sets the "notify_template" field if the given value is not nil.
func (iuo *InformUpdateOne) SetNillableNotifyTemplate(s *string) *InformUpdateOne {
	if s != nil {
		iuo.SetNotifyTemplate(*s)
	}
	return iuo
}

// SetTemplateCode sets the "template_code" field.
func (iuo *InformUpdateOne) SetTemplateCode(s string) *InformUpdateOne {
	iuo.mutation.SetTemplateCode(s)
	return iuo
}

// SetNillableTemplateCode sets the "template_code" field if the given value is not nil.
func (iuo *InformUpdateOne) SetNillableTemplateCode(s *string) *InformUpdateOne {
	if s != nil {
		iuo.SetTemplateCode(*s)
	}
	return iuo
}

// SetPhoneNumbers sets the "phone_numbers" field.
func (iuo *InformUpdateOne) SetPhoneNumbers(s string) *InformUpdateOne {
	iuo.mutation.SetPhoneNumbers(s)
	return iuo
}

// SetNillablePhoneNumbers sets the "phone_numbers" field if the given value is not nil.
func (iuo *InformUpdateOne) SetNillablePhoneNumbers(s *string) *InformUpdateOne {
	if s != nil {
		iuo.SetPhoneNumbers(*s)
	}
	return iuo
}

// SetNotifySwitch sets the "notify_switch" field.
func (iuo *InformUpdateOne) SetNotifySwitch(s string) *InformUpdateOne {
	iuo.mutation.SetNotifySwitch(s)
	return iuo
}

// SetNillableNotifySwitch sets the "notify_switch" field if the given value is not nil.
func (iuo *InformUpdateOne) SetNillableNotifySwitch(s *string) *InformUpdateOne {
	if s != nil {
		iuo.SetNotifySwitch(*s)
	}
	return iuo
}

// SetTaskName sets the "task_name" field.
func (iuo *InformUpdateOne) SetTaskName(s string) *InformUpdateOne {
	iuo.mutation.SetTaskName(s)
	return iuo
}

// SetNillableTaskName sets the "task_name" field if the given value is not nil.
func (iuo *InformUpdateOne) SetNillableTaskName(s *string) *InformUpdateOne {
	if s != nil {
		iuo.SetTaskName(*s)
	}
	return iuo
}

// SetTaskID sets the "task_id" field.
func (iuo *InformUpdateOne) SetTaskID(u uint64) *InformUpdateOne {
	iuo.mutation.ResetTaskID()
	iuo.mutation.SetTaskID(u)
	return iuo
}

// SetNillableTaskID sets the "task_id" field if the given value is not nil.
func (iuo *InformUpdateOne) SetNillableTaskID(u *uint64) *InformUpdateOne {
	if u != nil {
		iuo.SetTaskID(*u)
	}
	return iuo
}

// AddTaskID adds u to the "task_id" field.
func (iuo *InformUpdateOne) AddTaskID(u int64) *InformUpdateOne {
	iuo.mutation.AddTaskID(u)
	return iuo
}

// Mutation returns the InformMutation object of the builder.
func (iuo *InformUpdateOne) Mutation() *InformMutation {
	return iuo.mutation
}

// Where appends a list predicates to the InformUpdate builder.
func (iuo *InformUpdateOne) Where(ps ...predicate.Inform) *InformUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *InformUpdateOne) Select(field string, fields ...string) *InformUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Inform entity.
func (iuo *InformUpdateOne) Save(ctx context.Context) (*Inform, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InformUpdateOne) SaveX(ctx context.Context) *Inform {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *InformUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InformUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (iuo *InformUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *InformUpdateOne {
	iuo.modifiers = append(iuo.modifiers, modifiers...)
	return iuo
}

func (iuo *InformUpdateOne) sqlSave(ctx context.Context) (_node *Inform, err error) {
	_spec := sqlgraph.NewUpdateSpec(inform.Table, inform.Columns, sqlgraph.NewFieldSpec(inform.FieldID, field.TypeUint64))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Inform.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, inform.FieldID)
		for _, f := range fields {
			if !inform.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != inform.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.TenantID(); ok {
		_spec.SetField(inform.FieldTenantID, field.TypeString, value)
	}
	if iuo.mutation.TenantIDCleared() {
		_spec.ClearField(inform.FieldTenantID, field.TypeString)
	}
	if value, ok := iuo.mutation.AccessOrgList(); ok {
		_spec.SetField(inform.FieldAccessOrgList, field.TypeString, value)
	}
	if iuo.mutation.AccessOrgListCleared() {
		_spec.ClearField(inform.FieldAccessOrgList, field.TypeString)
	}
	if value, ok := iuo.mutation.CreatedAt(); ok {
		_spec.SetField(inform.FieldCreatedAt, field.TypeTime, value)
	}
	if iuo.mutation.CreatedAtCleared() {
		_spec.ClearField(inform.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.SetField(inform.FieldUpdatedAt, field.TypeTime, value)
	}
	if iuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(inform.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := iuo.mutation.DeletedAt(); ok {
		_spec.SetField(inform.FieldDeletedAt, field.TypeTime, value)
	}
	if iuo.mutation.DeletedAtCleared() {
		_spec.ClearField(inform.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := iuo.mutation.AlarmName(); ok {
		_spec.SetField(inform.FieldAlarmName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.AlarmType(); ok {
		_spec.SetField(inform.FieldAlarmType, field.TypeString, value)
	}
	if value, ok := iuo.mutation.SignName(); ok {
		_spec.SetField(inform.FieldSignName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.NotifyTemplate(); ok {
		_spec.SetField(inform.FieldNotifyTemplate, field.TypeString, value)
	}
	if value, ok := iuo.mutation.TemplateCode(); ok {
		_spec.SetField(inform.FieldTemplateCode, field.TypeString, value)
	}
	if value, ok := iuo.mutation.PhoneNumbers(); ok {
		_spec.SetField(inform.FieldPhoneNumbers, field.TypeString, value)
	}
	if value, ok := iuo.mutation.NotifySwitch(); ok {
		_spec.SetField(inform.FieldNotifySwitch, field.TypeString, value)
	}
	if value, ok := iuo.mutation.TaskName(); ok {
		_spec.SetField(inform.FieldTaskName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.TaskID(); ok {
		_spec.SetField(inform.FieldTaskID, field.TypeUint64, value)
	}
	if value, ok := iuo.mutation.AddedTaskID(); ok {
		_spec.AddField(inform.FieldTaskID, field.TypeUint64, value)
	}
	_spec.AddModifiers(iuo.modifiers...)
	_node = &Inform{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{inform.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
