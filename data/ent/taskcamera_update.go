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
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/camera"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/predicate"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/task"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/taskcamera"
)

// TaskCameraUpdate is the builder for updating TaskCamera entities.
type TaskCameraUpdate struct {
	config
	hooks     []Hook
	mutation  *TaskCameraMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TaskCameraUpdate builder.
func (tcu *TaskCameraUpdate) Where(ps ...predicate.TaskCamera) *TaskCameraUpdate {
	tcu.mutation.Where(ps...)
	return tcu
}

// SetUpdatedAt sets the "updated_at" field.
func (tcu *TaskCameraUpdate) SetUpdatedAt(t time.Time) *TaskCameraUpdate {
	tcu.mutation.SetUpdatedAt(t)
	return tcu
}

// SetTenantID sets the "tenant_id" field.
func (tcu *TaskCameraUpdate) SetTenantID(s string) *TaskCameraUpdate {
	tcu.mutation.SetTenantID(s)
	return tcu
}

// SetNillableTenantID sets the "tenant_id" field if the given value is not nil.
func (tcu *TaskCameraUpdate) SetNillableTenantID(s *string) *TaskCameraUpdate {
	if s != nil {
		tcu.SetTenantID(*s)
	}
	return tcu
}

// ClearTenantID clears the value of the "tenant_id" field.
func (tcu *TaskCameraUpdate) ClearTenantID() *TaskCameraUpdate {
	tcu.mutation.ClearTenantID()
	return tcu
}

// SetAccessOrgList sets the "access_org_list" field.
func (tcu *TaskCameraUpdate) SetAccessOrgList(s string) *TaskCameraUpdate {
	tcu.mutation.SetAccessOrgList(s)
	return tcu
}

// SetNillableAccessOrgList sets the "access_org_list" field if the given value is not nil.
func (tcu *TaskCameraUpdate) SetNillableAccessOrgList(s *string) *TaskCameraUpdate {
	if s != nil {
		tcu.SetAccessOrgList(*s)
	}
	return tcu
}

// ClearAccessOrgList clears the value of the "access_org_list" field.
func (tcu *TaskCameraUpdate) ClearAccessOrgList() *TaskCameraUpdate {
	tcu.mutation.ClearAccessOrgList()
	return tcu
}

// SetTaskID sets the "task_id" field.
func (tcu *TaskCameraUpdate) SetTaskID(u uint64) *TaskCameraUpdate {
	tcu.mutation.SetTaskID(u)
	return tcu
}

// SetNillableTaskID sets the "task_id" field if the given value is not nil.
func (tcu *TaskCameraUpdate) SetNillableTaskID(u *uint64) *TaskCameraUpdate {
	if u != nil {
		tcu.SetTaskID(*u)
	}
	return tcu
}

// SetCameraID sets the "camera_id" field.
func (tcu *TaskCameraUpdate) SetCameraID(u uint64) *TaskCameraUpdate {
	tcu.mutation.SetCameraID(u)
	return tcu
}

// SetNillableCameraID sets the "camera_id" field if the given value is not nil.
func (tcu *TaskCameraUpdate) SetNillableCameraID(u *uint64) *TaskCameraUpdate {
	if u != nil {
		tcu.SetCameraID(*u)
	}
	return tcu
}

// SetMultiImgBox sets the "multi_img_box" field.
func (tcu *TaskCameraUpdate) SetMultiImgBox(s string) *TaskCameraUpdate {
	tcu.mutation.SetMultiImgBox(s)
	return tcu
}

// SetNillableMultiImgBox sets the "multi_img_box" field if the given value is not nil.
func (tcu *TaskCameraUpdate) SetNillableMultiImgBox(s *string) *TaskCameraUpdate {
	if s != nil {
		tcu.SetMultiImgBox(*s)
	}
	return tcu
}

// SetCamera sets the "camera" edge to the Camera entity.
func (tcu *TaskCameraUpdate) SetCamera(c *Camera) *TaskCameraUpdate {
	return tcu.SetCameraID(c.ID)
}

// SetTask sets the "task" edge to the Task entity.
func (tcu *TaskCameraUpdate) SetTask(t *Task) *TaskCameraUpdate {
	return tcu.SetTaskID(t.ID)
}

// Mutation returns the TaskCameraMutation object of the builder.
func (tcu *TaskCameraUpdate) Mutation() *TaskCameraMutation {
	return tcu.mutation
}

// ClearCamera clears the "camera" edge to the Camera entity.
func (tcu *TaskCameraUpdate) ClearCamera() *TaskCameraUpdate {
	tcu.mutation.ClearCamera()
	return tcu
}

// ClearTask clears the "task" edge to the Task entity.
func (tcu *TaskCameraUpdate) ClearTask() *TaskCameraUpdate {
	tcu.mutation.ClearTask()
	return tcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tcu *TaskCameraUpdate) Save(ctx context.Context) (int, error) {
	if err := tcu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, tcu.sqlSave, tcu.mutation, tcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tcu *TaskCameraUpdate) SaveX(ctx context.Context) int {
	affected, err := tcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tcu *TaskCameraUpdate) Exec(ctx context.Context) error {
	_, err := tcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcu *TaskCameraUpdate) ExecX(ctx context.Context) {
	if err := tcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcu *TaskCameraUpdate) defaults() error {
	if _, ok := tcu.mutation.UpdatedAt(); !ok {
		if taskcamera.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized taskcamera.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := taskcamera.UpdateDefaultUpdatedAt()
		tcu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tcu *TaskCameraUpdate) check() error {
	if _, ok := tcu.mutation.CameraID(); tcu.mutation.CameraCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TaskCamera.camera"`)
	}
	if _, ok := tcu.mutation.TaskID(); tcu.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TaskCamera.task"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tcu *TaskCameraUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TaskCameraUpdate {
	tcu.modifiers = append(tcu.modifiers, modifiers...)
	return tcu
}

func (tcu *TaskCameraUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(taskcamera.Table, taskcamera.Columns, sqlgraph.NewFieldSpec(taskcamera.FieldID, field.TypeUint64))
	if ps := tcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcu.mutation.UpdatedAt(); ok {
		_spec.SetField(taskcamera.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tcu.mutation.TenantID(); ok {
		_spec.SetField(taskcamera.FieldTenantID, field.TypeString, value)
	}
	if tcu.mutation.TenantIDCleared() {
		_spec.ClearField(taskcamera.FieldTenantID, field.TypeString)
	}
	if value, ok := tcu.mutation.AccessOrgList(); ok {
		_spec.SetField(taskcamera.FieldAccessOrgList, field.TypeString, value)
	}
	if tcu.mutation.AccessOrgListCleared() {
		_spec.ClearField(taskcamera.FieldAccessOrgList, field.TypeString)
	}
	if value, ok := tcu.mutation.MultiImgBox(); ok {
		_spec.SetField(taskcamera.FieldMultiImgBox, field.TypeString, value)
	}
	if tcu.mutation.CameraCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   taskcamera.CameraTable,
			Columns: []string{taskcamera.CameraColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(camera.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcu.mutation.CameraIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   taskcamera.CameraTable,
			Columns: []string{taskcamera.CameraColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(camera.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tcu.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   taskcamera.TaskTable,
			Columns: []string{taskcamera.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcu.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   taskcamera.TaskTable,
			Columns: []string{taskcamera.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(tcu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taskcamera.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tcu.mutation.done = true
	return n, nil
}

// TaskCameraUpdateOne is the builder for updating a single TaskCamera entity.
type TaskCameraUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TaskCameraMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (tcuo *TaskCameraUpdateOne) SetUpdatedAt(t time.Time) *TaskCameraUpdateOne {
	tcuo.mutation.SetUpdatedAt(t)
	return tcuo
}

// SetTenantID sets the "tenant_id" field.
func (tcuo *TaskCameraUpdateOne) SetTenantID(s string) *TaskCameraUpdateOne {
	tcuo.mutation.SetTenantID(s)
	return tcuo
}

// SetNillableTenantID sets the "tenant_id" field if the given value is not nil.
func (tcuo *TaskCameraUpdateOne) SetNillableTenantID(s *string) *TaskCameraUpdateOne {
	if s != nil {
		tcuo.SetTenantID(*s)
	}
	return tcuo
}

// ClearTenantID clears the value of the "tenant_id" field.
func (tcuo *TaskCameraUpdateOne) ClearTenantID() *TaskCameraUpdateOne {
	tcuo.mutation.ClearTenantID()
	return tcuo
}

// SetAccessOrgList sets the "access_org_list" field.
func (tcuo *TaskCameraUpdateOne) SetAccessOrgList(s string) *TaskCameraUpdateOne {
	tcuo.mutation.SetAccessOrgList(s)
	return tcuo
}

// SetNillableAccessOrgList sets the "access_org_list" field if the given value is not nil.
func (tcuo *TaskCameraUpdateOne) SetNillableAccessOrgList(s *string) *TaskCameraUpdateOne {
	if s != nil {
		tcuo.SetAccessOrgList(*s)
	}
	return tcuo
}

// ClearAccessOrgList clears the value of the "access_org_list" field.
func (tcuo *TaskCameraUpdateOne) ClearAccessOrgList() *TaskCameraUpdateOne {
	tcuo.mutation.ClearAccessOrgList()
	return tcuo
}

// SetTaskID sets the "task_id" field.
func (tcuo *TaskCameraUpdateOne) SetTaskID(u uint64) *TaskCameraUpdateOne {
	tcuo.mutation.SetTaskID(u)
	return tcuo
}

// SetNillableTaskID sets the "task_id" field if the given value is not nil.
func (tcuo *TaskCameraUpdateOne) SetNillableTaskID(u *uint64) *TaskCameraUpdateOne {
	if u != nil {
		tcuo.SetTaskID(*u)
	}
	return tcuo
}

// SetCameraID sets the "camera_id" field.
func (tcuo *TaskCameraUpdateOne) SetCameraID(u uint64) *TaskCameraUpdateOne {
	tcuo.mutation.SetCameraID(u)
	return tcuo
}

// SetNillableCameraID sets the "camera_id" field if the given value is not nil.
func (tcuo *TaskCameraUpdateOne) SetNillableCameraID(u *uint64) *TaskCameraUpdateOne {
	if u != nil {
		tcuo.SetCameraID(*u)
	}
	return tcuo
}

// SetMultiImgBox sets the "multi_img_box" field.
func (tcuo *TaskCameraUpdateOne) SetMultiImgBox(s string) *TaskCameraUpdateOne {
	tcuo.mutation.SetMultiImgBox(s)
	return tcuo
}

// SetNillableMultiImgBox sets the "multi_img_box" field if the given value is not nil.
func (tcuo *TaskCameraUpdateOne) SetNillableMultiImgBox(s *string) *TaskCameraUpdateOne {
	if s != nil {
		tcuo.SetMultiImgBox(*s)
	}
	return tcuo
}

// SetCamera sets the "camera" edge to the Camera entity.
func (tcuo *TaskCameraUpdateOne) SetCamera(c *Camera) *TaskCameraUpdateOne {
	return tcuo.SetCameraID(c.ID)
}

// SetTask sets the "task" edge to the Task entity.
func (tcuo *TaskCameraUpdateOne) SetTask(t *Task) *TaskCameraUpdateOne {
	return tcuo.SetTaskID(t.ID)
}

// Mutation returns the TaskCameraMutation object of the builder.
func (tcuo *TaskCameraUpdateOne) Mutation() *TaskCameraMutation {
	return tcuo.mutation
}

// ClearCamera clears the "camera" edge to the Camera entity.
func (tcuo *TaskCameraUpdateOne) ClearCamera() *TaskCameraUpdateOne {
	tcuo.mutation.ClearCamera()
	return tcuo
}

// ClearTask clears the "task" edge to the Task entity.
func (tcuo *TaskCameraUpdateOne) ClearTask() *TaskCameraUpdateOne {
	tcuo.mutation.ClearTask()
	return tcuo
}

// Where appends a list predicates to the TaskCameraUpdate builder.
func (tcuo *TaskCameraUpdateOne) Where(ps ...predicate.TaskCamera) *TaskCameraUpdateOne {
	tcuo.mutation.Where(ps...)
	return tcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tcuo *TaskCameraUpdateOne) Select(field string, fields ...string) *TaskCameraUpdateOne {
	tcuo.fields = append([]string{field}, fields...)
	return tcuo
}

// Save executes the query and returns the updated TaskCamera entity.
func (tcuo *TaskCameraUpdateOne) Save(ctx context.Context) (*TaskCamera, error) {
	if err := tcuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tcuo.sqlSave, tcuo.mutation, tcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tcuo *TaskCameraUpdateOne) SaveX(ctx context.Context) *TaskCamera {
	node, err := tcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tcuo *TaskCameraUpdateOne) Exec(ctx context.Context) error {
	_, err := tcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcuo *TaskCameraUpdateOne) ExecX(ctx context.Context) {
	if err := tcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcuo *TaskCameraUpdateOne) defaults() error {
	if _, ok := tcuo.mutation.UpdatedAt(); !ok {
		if taskcamera.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized taskcamera.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := taskcamera.UpdateDefaultUpdatedAt()
		tcuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tcuo *TaskCameraUpdateOne) check() error {
	if _, ok := tcuo.mutation.CameraID(); tcuo.mutation.CameraCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TaskCamera.camera"`)
	}
	if _, ok := tcuo.mutation.TaskID(); tcuo.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TaskCamera.task"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tcuo *TaskCameraUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TaskCameraUpdateOne {
	tcuo.modifiers = append(tcuo.modifiers, modifiers...)
	return tcuo
}

func (tcuo *TaskCameraUpdateOne) sqlSave(ctx context.Context) (_node *TaskCamera, err error) {
	if err := tcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(taskcamera.Table, taskcamera.Columns, sqlgraph.NewFieldSpec(taskcamera.FieldID, field.TypeUint64))
	id, ok := tcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TaskCamera.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taskcamera.FieldID)
		for _, f := range fields {
			if !taskcamera.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != taskcamera.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(taskcamera.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tcuo.mutation.TenantID(); ok {
		_spec.SetField(taskcamera.FieldTenantID, field.TypeString, value)
	}
	if tcuo.mutation.TenantIDCleared() {
		_spec.ClearField(taskcamera.FieldTenantID, field.TypeString)
	}
	if value, ok := tcuo.mutation.AccessOrgList(); ok {
		_spec.SetField(taskcamera.FieldAccessOrgList, field.TypeString, value)
	}
	if tcuo.mutation.AccessOrgListCleared() {
		_spec.ClearField(taskcamera.FieldAccessOrgList, field.TypeString)
	}
	if value, ok := tcuo.mutation.MultiImgBox(); ok {
		_spec.SetField(taskcamera.FieldMultiImgBox, field.TypeString, value)
	}
	if tcuo.mutation.CameraCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   taskcamera.CameraTable,
			Columns: []string{taskcamera.CameraColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(camera.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcuo.mutation.CameraIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   taskcamera.CameraTable,
			Columns: []string{taskcamera.CameraColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(camera.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tcuo.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   taskcamera.TaskTable,
			Columns: []string{taskcamera.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tcuo.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   taskcamera.TaskTable,
			Columns: []string{taskcamera.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(tcuo.modifiers...)
	_node = &TaskCamera{config: tcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taskcamera.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tcuo.mutation.done = true
	return _node, nil
}