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
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/predicate"
	"gitlab.ctyuncdn.cn/ias/ias-core/data/ent/upplatform"
)

// UpPlatformQuery is the builder for querying UpPlatform entities.
type UpPlatformQuery struct {
	config
	ctx        *QueryContext
	order      []upplatform.OrderOption
	inters     []Interceptor
	predicates []predicate.UpPlatform
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UpPlatformQuery builder.
func (upq *UpPlatformQuery) Where(ps ...predicate.UpPlatform) *UpPlatformQuery {
	upq.predicates = append(upq.predicates, ps...)
	return upq
}

// Limit the number of records to be returned by this query.
func (upq *UpPlatformQuery) Limit(limit int) *UpPlatformQuery {
	upq.ctx.Limit = &limit
	return upq
}

// Offset to start from.
func (upq *UpPlatformQuery) Offset(offset int) *UpPlatformQuery {
	upq.ctx.Offset = &offset
	return upq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (upq *UpPlatformQuery) Unique(unique bool) *UpPlatformQuery {
	upq.ctx.Unique = &unique
	return upq
}

// Order specifies how the records should be ordered.
func (upq *UpPlatformQuery) Order(o ...upplatform.OrderOption) *UpPlatformQuery {
	upq.order = append(upq.order, o...)
	return upq
}

// First returns the first UpPlatform entity from the query.
// Returns a *NotFoundError when no UpPlatform was found.
func (upq *UpPlatformQuery) First(ctx context.Context) (*UpPlatform, error) {
	nodes, err := upq.Limit(1).All(setContextOp(ctx, upq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{upplatform.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (upq *UpPlatformQuery) FirstX(ctx context.Context) *UpPlatform {
	node, err := upq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UpPlatform ID from the query.
// Returns a *NotFoundError when no UpPlatform ID was found.
func (upq *UpPlatformQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = upq.Limit(1).IDs(setContextOp(ctx, upq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{upplatform.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (upq *UpPlatformQuery) FirstIDX(ctx context.Context) int {
	id, err := upq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UpPlatform entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UpPlatform entity is found.
// Returns a *NotFoundError when no UpPlatform entities are found.
func (upq *UpPlatformQuery) Only(ctx context.Context) (*UpPlatform, error) {
	nodes, err := upq.Limit(2).All(setContextOp(ctx, upq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{upplatform.Label}
	default:
		return nil, &NotSingularError{upplatform.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (upq *UpPlatformQuery) OnlyX(ctx context.Context) *UpPlatform {
	node, err := upq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UpPlatform ID in the query.
// Returns a *NotSingularError when more than one UpPlatform ID is found.
// Returns a *NotFoundError when no entities are found.
func (upq *UpPlatformQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = upq.Limit(2).IDs(setContextOp(ctx, upq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{upplatform.Label}
	default:
		err = &NotSingularError{upplatform.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (upq *UpPlatformQuery) OnlyIDX(ctx context.Context) int {
	id, err := upq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UpPlatforms.
func (upq *UpPlatformQuery) All(ctx context.Context) ([]*UpPlatform, error) {
	ctx = setContextOp(ctx, upq.ctx, "All")
	if err := upq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UpPlatform, *UpPlatformQuery]()
	return withInterceptors[[]*UpPlatform](ctx, upq, qr, upq.inters)
}

// AllX is like All, but panics if an error occurs.
func (upq *UpPlatformQuery) AllX(ctx context.Context) []*UpPlatform {
	nodes, err := upq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UpPlatform IDs.
func (upq *UpPlatformQuery) IDs(ctx context.Context) (ids []int, err error) {
	if upq.ctx.Unique == nil && upq.path != nil {
		upq.Unique(true)
	}
	ctx = setContextOp(ctx, upq.ctx, "IDs")
	if err = upq.Select(upplatform.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (upq *UpPlatformQuery) IDsX(ctx context.Context) []int {
	ids, err := upq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (upq *UpPlatformQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, upq.ctx, "Count")
	if err := upq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, upq, querierCount[*UpPlatformQuery](), upq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (upq *UpPlatformQuery) CountX(ctx context.Context) int {
	count, err := upq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (upq *UpPlatformQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, upq.ctx, "Exist")
	switch _, err := upq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (upq *UpPlatformQuery) ExistX(ctx context.Context) bool {
	exist, err := upq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UpPlatformQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (upq *UpPlatformQuery) Clone() *UpPlatformQuery {
	if upq == nil {
		return nil
	}
	return &UpPlatformQuery{
		config:     upq.config,
		ctx:        upq.ctx.Clone(),
		order:      append([]upplatform.OrderOption{}, upq.order...),
		inters:     append([]Interceptor{}, upq.inters...),
		predicates: append([]predicate.UpPlatform{}, upq.predicates...),
		// clone intermediate query.
		sql:  upq.sql.Clone(),
		path: upq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TenantID string `json:"tenant_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UpPlatform.Query().
//		GroupBy(upplatform.FieldTenantID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (upq *UpPlatformQuery) GroupBy(field string, fields ...string) *UpPlatformGroupBy {
	upq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UpPlatformGroupBy{build: upq}
	grbuild.flds = &upq.ctx.Fields
	grbuild.label = upplatform.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TenantID string `json:"tenant_id,omitempty"`
//	}
//
//	client.UpPlatform.Query().
//		Select(upplatform.FieldTenantID).
//		Scan(ctx, &v)
func (upq *UpPlatformQuery) Select(fields ...string) *UpPlatformSelect {
	upq.ctx.Fields = append(upq.ctx.Fields, fields...)
	sbuild := &UpPlatformSelect{UpPlatformQuery: upq}
	sbuild.label = upplatform.Label
	sbuild.flds, sbuild.scan = &upq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UpPlatformSelect configured with the given aggregations.
func (upq *UpPlatformQuery) Aggregate(fns ...AggregateFunc) *UpPlatformSelect {
	return upq.Select().Aggregate(fns...)
}

func (upq *UpPlatformQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range upq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, upq); err != nil {
				return err
			}
		}
	}
	for _, f := range upq.ctx.Fields {
		if !upplatform.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if upq.path != nil {
		prev, err := upq.path(ctx)
		if err != nil {
			return err
		}
		upq.sql = prev
	}
	return nil
}

func (upq *UpPlatformQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UpPlatform, error) {
	var (
		nodes = []*UpPlatform{}
		_spec = upq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UpPlatform).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UpPlatform{config: upq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(upq.modifiers) > 0 {
		_spec.Modifiers = upq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, upq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (upq *UpPlatformQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := upq.querySpec()
	if len(upq.modifiers) > 0 {
		_spec.Modifiers = upq.modifiers
	}
	_spec.Node.Columns = upq.ctx.Fields
	if len(upq.ctx.Fields) > 0 {
		_spec.Unique = upq.ctx.Unique != nil && *upq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, upq.driver, _spec)
}

func (upq *UpPlatformQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(upplatform.Table, upplatform.Columns, sqlgraph.NewFieldSpec(upplatform.FieldID, field.TypeInt))
	_spec.From = upq.sql
	if unique := upq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if upq.path != nil {
		_spec.Unique = true
	}
	if fields := upq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, upplatform.FieldID)
		for i := range fields {
			if fields[i] != upplatform.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := upq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := upq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := upq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := upq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (upq *UpPlatformQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(upq.driver.Dialect())
	t1 := builder.Table(upplatform.Table)
	columns := upq.ctx.Fields
	if len(columns) == 0 {
		columns = upplatform.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if upq.sql != nil {
		selector = upq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if upq.ctx.Unique != nil && *upq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range upq.modifiers {
		m(selector)
	}
	for _, p := range upq.predicates {
		p(selector)
	}
	for _, p := range upq.order {
		p(selector)
	}
	if offset := upq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := upq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (upq *UpPlatformQuery) ForUpdate(opts ...sql.LockOption) *UpPlatformQuery {
	if upq.driver.Dialect() == dialect.Postgres {
		upq.Unique(false)
	}
	upq.modifiers = append(upq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return upq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (upq *UpPlatformQuery) ForShare(opts ...sql.LockOption) *UpPlatformQuery {
	if upq.driver.Dialect() == dialect.Postgres {
		upq.Unique(false)
	}
	upq.modifiers = append(upq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return upq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (upq *UpPlatformQuery) Modify(modifiers ...func(s *sql.Selector)) *UpPlatformSelect {
	upq.modifiers = append(upq.modifiers, modifiers...)
	return upq.Select()
}

// UpPlatformGroupBy is the group-by builder for UpPlatform entities.
type UpPlatformGroupBy struct {
	selector
	build *UpPlatformQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (upgb *UpPlatformGroupBy) Aggregate(fns ...AggregateFunc) *UpPlatformGroupBy {
	upgb.fns = append(upgb.fns, fns...)
	return upgb
}

// Scan applies the selector query and scans the result into the given value.
func (upgb *UpPlatformGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, upgb.build.ctx, "GroupBy")
	if err := upgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UpPlatformQuery, *UpPlatformGroupBy](ctx, upgb.build, upgb, upgb.build.inters, v)
}

func (upgb *UpPlatformGroupBy) sqlScan(ctx context.Context, root *UpPlatformQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(upgb.fns))
	for _, fn := range upgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*upgb.flds)+len(upgb.fns))
		for _, f := range *upgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*upgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := upgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UpPlatformSelect is the builder for selecting fields of UpPlatform entities.
type UpPlatformSelect struct {
	*UpPlatformQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ups *UpPlatformSelect) Aggregate(fns ...AggregateFunc) *UpPlatformSelect {
	ups.fns = append(ups.fns, fns...)
	return ups
}

// Scan applies the selector query and scans the result into the given value.
func (ups *UpPlatformSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ups.ctx, "Select")
	if err := ups.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UpPlatformQuery, *UpPlatformSelect](ctx, ups.UpPlatformQuery, ups, ups.inters, v)
}

func (ups *UpPlatformSelect) sqlScan(ctx context.Context, root *UpPlatformQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ups.fns))
	for _, fn := range ups.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ups.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ups.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ups *UpPlatformSelect) Modify(modifiers ...func(s *sql.Selector)) *UpPlatformSelect {
	ups.modifiers = append(ups.modifiers, modifiers...)
	return ups
}