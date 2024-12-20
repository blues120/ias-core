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
	"github.com/blues120/ias-core/data/ent/activeinfo"
	"github.com/blues120/ias-core/data/ent/predicate"
)

// ActiveInfoQuery is the builder for querying ActiveInfo entities.
type ActiveInfoQuery struct {
	config
	ctx        *QueryContext
	order      []activeinfo.OrderOption
	inters     []Interceptor
	predicates []predicate.ActiveInfo
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ActiveInfoQuery builder.
func (aiq *ActiveInfoQuery) Where(ps ...predicate.ActiveInfo) *ActiveInfoQuery {
	aiq.predicates = append(aiq.predicates, ps...)
	return aiq
}

// Limit the number of records to be returned by this query.
func (aiq *ActiveInfoQuery) Limit(limit int) *ActiveInfoQuery {
	aiq.ctx.Limit = &limit
	return aiq
}

// Offset to start from.
func (aiq *ActiveInfoQuery) Offset(offset int) *ActiveInfoQuery {
	aiq.ctx.Offset = &offset
	return aiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aiq *ActiveInfoQuery) Unique(unique bool) *ActiveInfoQuery {
	aiq.ctx.Unique = &unique
	return aiq
}

// Order specifies how the records should be ordered.
func (aiq *ActiveInfoQuery) Order(o ...activeinfo.OrderOption) *ActiveInfoQuery {
	aiq.order = append(aiq.order, o...)
	return aiq
}

// First returns the first ActiveInfo entity from the query.
// Returns a *NotFoundError when no ActiveInfo was found.
func (aiq *ActiveInfoQuery) First(ctx context.Context) (*ActiveInfo, error) {
	nodes, err := aiq.Limit(1).All(setContextOp(ctx, aiq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{activeinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aiq *ActiveInfoQuery) FirstX(ctx context.Context) *ActiveInfo {
	node, err := aiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ActiveInfo ID from the query.
// Returns a *NotFoundError when no ActiveInfo ID was found.
func (aiq *ActiveInfoQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = aiq.Limit(1).IDs(setContextOp(ctx, aiq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{activeinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aiq *ActiveInfoQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := aiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ActiveInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ActiveInfo entity is found.
// Returns a *NotFoundError when no ActiveInfo entities are found.
func (aiq *ActiveInfoQuery) Only(ctx context.Context) (*ActiveInfo, error) {
	nodes, err := aiq.Limit(2).All(setContextOp(ctx, aiq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{activeinfo.Label}
	default:
		return nil, &NotSingularError{activeinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aiq *ActiveInfoQuery) OnlyX(ctx context.Context) *ActiveInfo {
	node, err := aiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ActiveInfo ID in the query.
// Returns a *NotSingularError when more than one ActiveInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (aiq *ActiveInfoQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = aiq.Limit(2).IDs(setContextOp(ctx, aiq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{activeinfo.Label}
	default:
		err = &NotSingularError{activeinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aiq *ActiveInfoQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := aiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ActiveInfos.
func (aiq *ActiveInfoQuery) All(ctx context.Context) ([]*ActiveInfo, error) {
	ctx = setContextOp(ctx, aiq.ctx, "All")
	if err := aiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ActiveInfo, *ActiveInfoQuery]()
	return withInterceptors[[]*ActiveInfo](ctx, aiq, qr, aiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aiq *ActiveInfoQuery) AllX(ctx context.Context) []*ActiveInfo {
	nodes, err := aiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ActiveInfo IDs.
func (aiq *ActiveInfoQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if aiq.ctx.Unique == nil && aiq.path != nil {
		aiq.Unique(true)
	}
	ctx = setContextOp(ctx, aiq.ctx, "IDs")
	if err = aiq.Select(activeinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aiq *ActiveInfoQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := aiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aiq *ActiveInfoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aiq.ctx, "Count")
	if err := aiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aiq, querierCount[*ActiveInfoQuery](), aiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aiq *ActiveInfoQuery) CountX(ctx context.Context) int {
	count, err := aiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aiq *ActiveInfoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aiq.ctx, "Exist")
	switch _, err := aiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aiq *ActiveInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := aiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ActiveInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aiq *ActiveInfoQuery) Clone() *ActiveInfoQuery {
	if aiq == nil {
		return nil
	}
	return &ActiveInfoQuery{
		config:     aiq.config,
		ctx:        aiq.ctx.Clone(),
		order:      append([]activeinfo.OrderOption{}, aiq.order...),
		inters:     append([]Interceptor{}, aiq.inters...),
		predicates: append([]predicate.ActiveInfo{}, aiq.predicates...),
		// clone intermediate query.
		sql:  aiq.sql.Clone(),
		path: aiq.path,
	}
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
//	client.ActiveInfo.Query().
//		GroupBy(activeinfo.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aiq *ActiveInfoQuery) GroupBy(field string, fields ...string) *ActiveInfoGroupBy {
	aiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ActiveInfoGroupBy{build: aiq}
	grbuild.flds = &aiq.ctx.Fields
	grbuild.label = activeinfo.Label
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
//	client.ActiveInfo.Query().
//		Select(activeinfo.FieldCreatedAt).
//		Scan(ctx, &v)
func (aiq *ActiveInfoQuery) Select(fields ...string) *ActiveInfoSelect {
	aiq.ctx.Fields = append(aiq.ctx.Fields, fields...)
	sbuild := &ActiveInfoSelect{ActiveInfoQuery: aiq}
	sbuild.label = activeinfo.Label
	sbuild.flds, sbuild.scan = &aiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ActiveInfoSelect configured with the given aggregations.
func (aiq *ActiveInfoQuery) Aggregate(fns ...AggregateFunc) *ActiveInfoSelect {
	return aiq.Select().Aggregate(fns...)
}

func (aiq *ActiveInfoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aiq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aiq); err != nil {
				return err
			}
		}
	}
	for _, f := range aiq.ctx.Fields {
		if !activeinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aiq.path != nil {
		prev, err := aiq.path(ctx)
		if err != nil {
			return err
		}
		aiq.sql = prev
	}
	return nil
}

func (aiq *ActiveInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ActiveInfo, error) {
	var (
		nodes = []*ActiveInfo{}
		_spec = aiq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ActiveInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ActiveInfo{config: aiq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(aiq.modifiers) > 0 {
		_spec.Modifiers = aiq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aiq *ActiveInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aiq.querySpec()
	if len(aiq.modifiers) > 0 {
		_spec.Modifiers = aiq.modifiers
	}
	_spec.Node.Columns = aiq.ctx.Fields
	if len(aiq.ctx.Fields) > 0 {
		_spec.Unique = aiq.ctx.Unique != nil && *aiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aiq.driver, _spec)
}

func (aiq *ActiveInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(activeinfo.Table, activeinfo.Columns, sqlgraph.NewFieldSpec(activeinfo.FieldID, field.TypeUint64))
	_spec.From = aiq.sql
	if unique := aiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aiq.path != nil {
		_spec.Unique = true
	}
	if fields := aiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, activeinfo.FieldID)
		for i := range fields {
			if fields[i] != activeinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aiq *ActiveInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aiq.driver.Dialect())
	t1 := builder.Table(activeinfo.Table)
	columns := aiq.ctx.Fields
	if len(columns) == 0 {
		columns = activeinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aiq.sql != nil {
		selector = aiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aiq.ctx.Unique != nil && *aiq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range aiq.modifiers {
		m(selector)
	}
	for _, p := range aiq.predicates {
		p(selector)
	}
	for _, p := range aiq.order {
		p(selector)
	}
	if offset := aiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (aiq *ActiveInfoQuery) ForUpdate(opts ...sql.LockOption) *ActiveInfoQuery {
	if aiq.driver.Dialect() == dialect.Postgres {
		aiq.Unique(false)
	}
	aiq.modifiers = append(aiq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return aiq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (aiq *ActiveInfoQuery) ForShare(opts ...sql.LockOption) *ActiveInfoQuery {
	if aiq.driver.Dialect() == dialect.Postgres {
		aiq.Unique(false)
	}
	aiq.modifiers = append(aiq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return aiq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aiq *ActiveInfoQuery) Modify(modifiers ...func(s *sql.Selector)) *ActiveInfoSelect {
	aiq.modifiers = append(aiq.modifiers, modifiers...)
	return aiq.Select()
}

// ActiveInfoGroupBy is the group-by builder for ActiveInfo entities.
type ActiveInfoGroupBy struct {
	selector
	build *ActiveInfoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aigb *ActiveInfoGroupBy) Aggregate(fns ...AggregateFunc) *ActiveInfoGroupBy {
	aigb.fns = append(aigb.fns, fns...)
	return aigb
}

// Scan applies the selector query and scans the result into the given value.
func (aigb *ActiveInfoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aigb.build.ctx, "GroupBy")
	if err := aigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ActiveInfoQuery, *ActiveInfoGroupBy](ctx, aigb.build, aigb, aigb.build.inters, v)
}

func (aigb *ActiveInfoGroupBy) sqlScan(ctx context.Context, root *ActiveInfoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(aigb.fns))
	for _, fn := range aigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*aigb.flds)+len(aigb.fns))
		for _, f := range *aigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*aigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ActiveInfoSelect is the builder for selecting fields of ActiveInfo entities.
type ActiveInfoSelect struct {
	*ActiveInfoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ais *ActiveInfoSelect) Aggregate(fns ...AggregateFunc) *ActiveInfoSelect {
	ais.fns = append(ais.fns, fns...)
	return ais
}

// Scan applies the selector query and scans the result into the given value.
func (ais *ActiveInfoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ais.ctx, "Select")
	if err := ais.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ActiveInfoQuery, *ActiveInfoSelect](ctx, ais.ActiveInfoQuery, ais, ais.inters, v)
}

func (ais *ActiveInfoSelect) sqlScan(ctx context.Context, root *ActiveInfoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ais.fns))
	for _, fn := range ais.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ais.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ais.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ais *ActiveInfoSelect) Modify(modifiers ...func(s *sql.Selector)) *ActiveInfoSelect {
	ais.modifiers = append(ais.modifiers, modifiers...)
	return ais
}
