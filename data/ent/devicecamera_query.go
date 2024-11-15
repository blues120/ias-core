// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/camera"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/devicecamera"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/predicate"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/task"
)

// DeviceCameraQuery is the builder for querying DeviceCamera entities.
type DeviceCameraQuery struct {
	config
	ctx        *QueryContext
	order      []devicecamera.OrderOption
	inters     []Interceptor
	predicates []predicate.DeviceCamera
	withCamera *CameraQuery
	withDevice *TaskQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeviceCameraQuery builder.
func (dcq *DeviceCameraQuery) Where(ps ...predicate.DeviceCamera) *DeviceCameraQuery {
	dcq.predicates = append(dcq.predicates, ps...)
	return dcq
}

// Limit the number of records to be returned by this query.
func (dcq *DeviceCameraQuery) Limit(limit int) *DeviceCameraQuery {
	dcq.ctx.Limit = &limit
	return dcq
}

// Offset to start from.
func (dcq *DeviceCameraQuery) Offset(offset int) *DeviceCameraQuery {
	dcq.ctx.Offset = &offset
	return dcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dcq *DeviceCameraQuery) Unique(unique bool) *DeviceCameraQuery {
	dcq.ctx.Unique = &unique
	return dcq
}

// Order specifies how the records should be ordered.
func (dcq *DeviceCameraQuery) Order(o ...devicecamera.OrderOption) *DeviceCameraQuery {
	dcq.order = append(dcq.order, o...)
	return dcq
}

// QueryCamera chains the current query on the "camera" edge.
func (dcq *DeviceCameraQuery) QueryCamera() *CameraQuery {
	query := (&CameraClient{config: dcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(devicecamera.Table, devicecamera.FieldID, selector),
			sqlgraph.To(camera.Table, camera.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, devicecamera.CameraTable, devicecamera.CameraColumn),
		)
		fromU = sqlgraph.SetNeighbors(dcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDevice chains the current query on the "device" edge.
func (dcq *DeviceCameraQuery) QueryDevice() *TaskQuery {
	query := (&TaskClient{config: dcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(devicecamera.Table, devicecamera.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, devicecamera.DeviceTable, devicecamera.DeviceColumn),
		)
		fromU = sqlgraph.SetNeighbors(dcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DeviceCamera entity from the query.
// Returns a *NotFoundError when no DeviceCamera was found.
func (dcq *DeviceCameraQuery) First(ctx context.Context) (*DeviceCamera, error) {
	nodes, err := dcq.Limit(1).All(setContextOp(ctx, dcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{devicecamera.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dcq *DeviceCameraQuery) FirstX(ctx context.Context) *DeviceCamera {
	node, err := dcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DeviceCamera ID from the query.
// Returns a *NotFoundError when no DeviceCamera ID was found.
func (dcq *DeviceCameraQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = dcq.Limit(1).IDs(setContextOp(ctx, dcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{devicecamera.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dcq *DeviceCameraQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := dcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DeviceCamera entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DeviceCamera entity is found.
// Returns a *NotFoundError when no DeviceCamera entities are found.
func (dcq *DeviceCameraQuery) Only(ctx context.Context) (*DeviceCamera, error) {
	nodes, err := dcq.Limit(2).All(setContextOp(ctx, dcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{devicecamera.Label}
	default:
		return nil, &NotSingularError{devicecamera.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dcq *DeviceCameraQuery) OnlyX(ctx context.Context) *DeviceCamera {
	node, err := dcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DeviceCamera ID in the query.
// Returns a *NotSingularError when more than one DeviceCamera ID is found.
// Returns a *NotFoundError when no entities are found.
func (dcq *DeviceCameraQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = dcq.Limit(2).IDs(setContextOp(ctx, dcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{devicecamera.Label}
	default:
		err = &NotSingularError{devicecamera.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dcq *DeviceCameraQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := dcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DeviceCameras.
func (dcq *DeviceCameraQuery) All(ctx context.Context) ([]*DeviceCamera, error) {
	ctx = setContextOp(ctx, dcq.ctx, "All")
	if err := dcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DeviceCamera, *DeviceCameraQuery]()
	return withInterceptors[[]*DeviceCamera](ctx, dcq, qr, dcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dcq *DeviceCameraQuery) AllX(ctx context.Context) []*DeviceCamera {
	nodes, err := dcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DeviceCamera IDs.
func (dcq *DeviceCameraQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if dcq.ctx.Unique == nil && dcq.path != nil {
		dcq.Unique(true)
	}
	ctx = setContextOp(ctx, dcq.ctx, "IDs")
	if err = dcq.Select(devicecamera.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dcq *DeviceCameraQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := dcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dcq *DeviceCameraQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dcq.ctx, "Count")
	if err := dcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dcq, querierCount[*DeviceCameraQuery](), dcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dcq *DeviceCameraQuery) CountX(ctx context.Context) int {
	count, err := dcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dcq *DeviceCameraQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dcq.ctx, "Exist")
	switch _, err := dcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dcq *DeviceCameraQuery) ExistX(ctx context.Context) bool {
	exist, err := dcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeviceCameraQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dcq *DeviceCameraQuery) Clone() *DeviceCameraQuery {
	if dcq == nil {
		return nil
	}
	return &DeviceCameraQuery{
		config:     dcq.config,
		ctx:        dcq.ctx.Clone(),
		order:      append([]devicecamera.OrderOption{}, dcq.order...),
		inters:     append([]Interceptor{}, dcq.inters...),
		predicates: append([]predicate.DeviceCamera{}, dcq.predicates...),
		withCamera: dcq.withCamera.Clone(),
		withDevice: dcq.withDevice.Clone(),
		// clone intermediate query.
		sql:  dcq.sql.Clone(),
		path: dcq.path,
	}
}

// WithCamera tells the query-builder to eager-load the nodes that are connected to
// the "camera" edge. The optional arguments are used to configure the query builder of the edge.
func (dcq *DeviceCameraQuery) WithCamera(opts ...func(*CameraQuery)) *DeviceCameraQuery {
	query := (&CameraClient{config: dcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dcq.withCamera = query
	return dcq
}

// WithDevice tells the query-builder to eager-load the nodes that are connected to
// the "device" edge. The optional arguments are used to configure the query builder of the edge.
func (dcq *DeviceCameraQuery) WithDevice(opts ...func(*TaskQuery)) *DeviceCameraQuery {
	query := (&TaskClient{config: dcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dcq.withDevice = query
	return dcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DeviceCamera.Query().
//		GroupBy(devicecamera.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dcq *DeviceCameraQuery) GroupBy(field string, fields ...string) *DeviceCameraGroupBy {
	dcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DeviceCameraGroupBy{build: dcq}
	grbuild.flds = &dcq.ctx.Fields
	grbuild.label = devicecamera.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.DeviceCamera.Query().
//		Select(devicecamera.FieldCreatedAt).
//		Scan(ctx, &v)
func (dcq *DeviceCameraQuery) Select(fields ...string) *DeviceCameraSelect {
	dcq.ctx.Fields = append(dcq.ctx.Fields, fields...)
	sbuild := &DeviceCameraSelect{DeviceCameraQuery: dcq}
	sbuild.label = devicecamera.Label
	sbuild.flds, sbuild.scan = &dcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DeviceCameraSelect configured with the given aggregations.
func (dcq *DeviceCameraQuery) Aggregate(fns ...AggregateFunc) *DeviceCameraSelect {
	return dcq.Select().Aggregate(fns...)
}

func (dcq *DeviceCameraQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dcq); err != nil {
				return err
			}
		}
	}
	for _, f := range dcq.ctx.Fields {
		if !devicecamera.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dcq.path != nil {
		prev, err := dcq.path(ctx)
		if err != nil {
			return err
		}
		dcq.sql = prev
	}
	return nil
}

func (dcq *DeviceCameraQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DeviceCamera, error) {
	var (
		nodes       = []*DeviceCamera{}
		_spec       = dcq.querySpec()
		loadedTypes = [2]bool{
			dcq.withCamera != nil,
			dcq.withDevice != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DeviceCamera).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DeviceCamera{config: dcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dcq.modifiers) > 0 {
		_spec.Modifiers = dcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dcq.withCamera; query != nil {
		if err := dcq.loadCamera(ctx, query, nodes, nil,
			func(n *DeviceCamera, e *Camera) { n.Edges.Camera = e }); err != nil {
			return nil, err
		}
	}
	if query := dcq.withDevice; query != nil {
		if err := dcq.loadDevice(ctx, query, nodes, nil,
			func(n *DeviceCamera, e *Task) { n.Edges.Device = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dcq *DeviceCameraQuery) loadCamera(ctx context.Context, query *CameraQuery, nodes []*DeviceCamera, init func(*DeviceCamera), assign func(*DeviceCamera, *Camera)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*DeviceCamera)
	for i := range nodes {
		fk := nodes[i].CameraID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(camera.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "camera_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dcq *DeviceCameraQuery) loadDevice(ctx context.Context, query *TaskQuery, nodes []*DeviceCamera, init func(*DeviceCamera), assign func(*DeviceCamera, *Task)) error {
	ids := make([]uint64, 0, len(nodes))
	nodeids := make(map[uint64][]*DeviceCamera)
	for i := range nodes {
		fk := nodes[i].DeviceID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(task.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "device_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dcq *DeviceCameraQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dcq.querySpec()
	if len(dcq.modifiers) > 0 {
		_spec.Modifiers = dcq.modifiers
	}
	_spec.Node.Columns = dcq.ctx.Fields
	if len(dcq.ctx.Fields) > 0 {
		_spec.Unique = dcq.ctx.Unique != nil && *dcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dcq.driver, _spec)
}

func (dcq *DeviceCameraQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(devicecamera.Table, devicecamera.Columns, sqlgraph.NewFieldSpec(devicecamera.FieldID, field.TypeUint64))
	_spec.From = dcq.sql
	if unique := dcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dcq.path != nil {
		_spec.Unique = true
	}
	if fields := dcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, devicecamera.FieldID)
		for i := range fields {
			if fields[i] != devicecamera.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if dcq.withCamera != nil {
			_spec.Node.AddColumnOnce(devicecamera.FieldCameraID)
		}
		if dcq.withDevice != nil {
			_spec.Node.AddColumnOnce(devicecamera.FieldDeviceID)
		}
	}
	if ps := dcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dcq *DeviceCameraQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dcq.driver.Dialect())
	t1 := builder.Table(devicecamera.Table)
	columns := dcq.ctx.Fields
	if len(columns) == 0 {
		columns = devicecamera.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dcq.sql != nil {
		selector = dcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dcq.ctx.Unique != nil && *dcq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range dcq.modifiers {
		m(selector)
	}
	for _, p := range dcq.predicates {
		p(selector)
	}
	for _, p := range dcq.order {
		p(selector)
	}
	if offset := dcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (dcq *DeviceCameraQuery) ForUpdate(opts ...sql.LockOption) *DeviceCameraQuery {
	if dcq.driver.Dialect() == dialect.Postgres {
		dcq.Unique(false)
	}
	dcq.modifiers = append(dcq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return dcq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (dcq *DeviceCameraQuery) ForShare(opts ...sql.LockOption) *DeviceCameraQuery {
	if dcq.driver.Dialect() == dialect.Postgres {
		dcq.Unique(false)
	}
	dcq.modifiers = append(dcq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return dcq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dcq *DeviceCameraQuery) Modify(modifiers ...func(s *sql.Selector)) *DeviceCameraSelect {
	dcq.modifiers = append(dcq.modifiers, modifiers...)
	return dcq.Select()
}

// DeviceCameraGroupBy is the group-by builder for DeviceCamera entities.
type DeviceCameraGroupBy struct {
	selector
	build *DeviceCameraQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dcgb *DeviceCameraGroupBy) Aggregate(fns ...AggregateFunc) *DeviceCameraGroupBy {
	dcgb.fns = append(dcgb.fns, fns...)
	return dcgb
}

// Scan applies the selector query and scans the result into the given value.
func (dcgb *DeviceCameraGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dcgb.build.ctx, "GroupBy")
	if err := dcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeviceCameraQuery, *DeviceCameraGroupBy](ctx, dcgb.build, dcgb, dcgb.build.inters, v)
}

func (dcgb *DeviceCameraGroupBy) sqlScan(ctx context.Context, root *DeviceCameraQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dcgb.fns))
	for _, fn := range dcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dcgb.flds)+len(dcgb.fns))
		for _, f := range *dcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DeviceCameraSelect is the builder for selecting fields of DeviceCamera entities.
type DeviceCameraSelect struct {
	*DeviceCameraQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (dcs *DeviceCameraSelect) Aggregate(fns ...AggregateFunc) *DeviceCameraSelect {
	dcs.fns = append(dcs.fns, fns...)
	return dcs
}

// Scan applies the selector query and scans the result into the given value.
func (dcs *DeviceCameraSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dcs.ctx, "Select")
	if err := dcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeviceCameraQuery, *DeviceCameraSelect](ctx, dcs.DeviceCameraQuery, dcs, dcs.inters, v)
}

func (dcs *DeviceCameraSelect) sqlScan(ctx context.Context, root *DeviceCameraQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(dcs.fns))
	for _, fn := range dcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*dcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dcs *DeviceCameraSelect) Modify(modifiers ...func(s *sql.Selector)) *DeviceCameraSelect {
	dcs.modifiers = append(dcs.modifiers, modifiers...)
	return dcs
}
