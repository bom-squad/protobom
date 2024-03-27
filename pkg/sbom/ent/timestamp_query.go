// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bom-squad/protobom/pkg/sbom/ent/predicate"
	"github.com/bom-squad/protobom/pkg/sbom/ent/timestamp"
)

// TimestampQuery is the builder for querying Timestamp entities.
type TimestampQuery struct {
	config
	ctx        *QueryContext
	order      []timestamp.OrderOption
	inters     []Interceptor
	predicates []predicate.Timestamp
	withDate   *TimestampQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TimestampQuery builder.
func (tq *TimestampQuery) Where(ps ...predicate.Timestamp) *TimestampQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TimestampQuery) Limit(limit int) *TimestampQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TimestampQuery) Offset(offset int) *TimestampQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TimestampQuery) Unique(unique bool) *TimestampQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TimestampQuery) Order(o ...timestamp.OrderOption) *TimestampQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryDate chains the current query on the "date" edge.
func (tq *TimestampQuery) QueryDate() *TimestampQuery {
	query := (&TimestampClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(timestamp.Table, timestamp.FieldID, selector),
			sqlgraph.To(timestamp.Table, timestamp.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, timestamp.DateTable, timestamp.DatePrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Timestamp entity from the query.
// Returns a *NotFoundError when no Timestamp was found.
func (tq *TimestampQuery) First(ctx context.Context) (*Timestamp, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{timestamp.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TimestampQuery) FirstX(ctx context.Context) *Timestamp {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Timestamp ID from the query.
// Returns a *NotFoundError when no Timestamp ID was found.
func (tq *TimestampQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{timestamp.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TimestampQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Timestamp entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Timestamp entity is found.
// Returns a *NotFoundError when no Timestamp entities are found.
func (tq *TimestampQuery) Only(ctx context.Context) (*Timestamp, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{timestamp.Label}
	default:
		return nil, &NotSingularError{timestamp.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TimestampQuery) OnlyX(ctx context.Context) *Timestamp {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Timestamp ID in the query.
// Returns a *NotSingularError when more than one Timestamp ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TimestampQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{timestamp.Label}
	default:
		err = &NotSingularError{timestamp.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TimestampQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Timestamps.
func (tq *TimestampQuery) All(ctx context.Context) ([]*Timestamp, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Timestamp, *TimestampQuery]()
	return withInterceptors[[]*Timestamp](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TimestampQuery) AllX(ctx context.Context) []*Timestamp {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Timestamp IDs.
func (tq *TimestampQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(timestamp.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TimestampQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TimestampQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TimestampQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TimestampQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TimestampQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TimestampQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TimestampQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TimestampQuery) Clone() *TimestampQuery {
	if tq == nil {
		return nil
	}
	return &TimestampQuery{
		config:     tq.config,
		ctx:        tq.ctx.Clone(),
		order:      append([]timestamp.OrderOption{}, tq.order...),
		inters:     append([]Interceptor{}, tq.inters...),
		predicates: append([]predicate.Timestamp{}, tq.predicates...),
		withDate:   tq.withDate.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithDate tells the query-builder to eager-load the nodes that are connected to
// the "date" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TimestampQuery) WithDate(opts ...func(*TimestampQuery)) *TimestampQuery {
	query := (&TimestampClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withDate = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (tq *TimestampQuery) GroupBy(field string, fields ...string) *TimestampGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TimestampGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = timestamp.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (tq *TimestampQuery) Select(fields ...string) *TimestampSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TimestampSelect{TimestampQuery: tq}
	sbuild.label = timestamp.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TimestampSelect configured with the given aggregations.
func (tq *TimestampQuery) Aggregate(fns ...AggregateFunc) *TimestampSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TimestampQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !timestamp.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TimestampQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Timestamp, error) {
	var (
		nodes       = []*Timestamp{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withDate != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, timestamp.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Timestamp).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Timestamp{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withDate; query != nil {
		if err := tq.loadDate(ctx, query, nodes,
			func(n *Timestamp) { n.Edges.Date = []*Timestamp{} },
			func(n *Timestamp, e *Timestamp) { n.Edges.Date = append(n.Edges.Date, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TimestampQuery) loadDate(ctx context.Context, query *TimestampQuery, nodes []*Timestamp, init func(*Timestamp), assign func(*Timestamp, *Timestamp)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Timestamp)
	nids := make(map[int]map[*Timestamp]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(timestamp.DateTable)
		s.Join(joinT).On(s.C(timestamp.FieldID), joinT.C(timestamp.DatePrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(timestamp.DatePrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(timestamp.DatePrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Timestamp]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Timestamp](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "date" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (tq *TimestampQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TimestampQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(timestamp.Table, timestamp.Columns, sqlgraph.NewFieldSpec(timestamp.FieldID, field.TypeInt))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, timestamp.FieldID)
		for i := range fields {
			if fields[i] != timestamp.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TimestampQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(timestamp.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = timestamp.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TimestampGroupBy is the group-by builder for Timestamp entities.
type TimestampGroupBy struct {
	selector
	build *TimestampQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TimestampGroupBy) Aggregate(fns ...AggregateFunc) *TimestampGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TimestampGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TimestampQuery, *TimestampGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TimestampGroupBy) sqlScan(ctx context.Context, root *TimestampQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TimestampSelect is the builder for selecting fields of Timestamp entities.
type TimestampSelect struct {
	*TimestampQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TimestampSelect) Aggregate(fns ...AggregateFunc) *TimestampSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TimestampSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TimestampQuery, *TimestampSelect](ctx, ts.TimestampQuery, ts, ts.inters, v)
}

func (ts *TimestampSelect) sqlScan(ctx context.Context, root *TimestampQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
