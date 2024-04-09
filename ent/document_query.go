// Code generated by ent, DO NOT EDIT.
// ------------------------------------------------------------------------
// SPDX-FileCopyrightText: Copyright © 2024 The Protobom Authors
// SPDX-FileType: SOURCE
// SPDX-License-Identifier: Apache-2.0
// ------------------------------------------------------------------------
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ------------------------------------------------------------------------
package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bom-squad/protobom/ent/document"
	"github.com/bom-squad/protobom/ent/metadata"
	"github.com/bom-squad/protobom/ent/nodelist"
	"github.com/bom-squad/protobom/ent/predicate"
)

// DocumentQuery is the builder for querying Document entities.
type DocumentQuery struct {
	config
	ctx                  *QueryContext
	order                []document.OrderOption
	inters               []Interceptor
	predicates           []predicate.Document
	withDocumentMetadata *MetadataQuery
	withDocumentNodeList *NodeListQuery
	withFKs              bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DocumentQuery builder.
func (dq *DocumentQuery) Where(ps ...predicate.Document) *DocumentQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DocumentQuery) Limit(limit int) *DocumentQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DocumentQuery) Offset(offset int) *DocumentQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DocumentQuery) Unique(unique bool) *DocumentQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DocumentQuery) Order(o ...document.OrderOption) *DocumentQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryDocumentMetadata chains the current query on the "document_metadata" edge.
func (dq *DocumentQuery) QueryDocumentMetadata() *MetadataQuery {
	query := (&MetadataClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(document.Table, document.FieldID, selector),
			sqlgraph.To(metadata.Table, metadata.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, document.DocumentMetadataTable, document.DocumentMetadataColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDocumentNodeList chains the current query on the "document_node_list" edge.
func (dq *DocumentQuery) QueryDocumentNodeList() *NodeListQuery {
	query := (&NodeListClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(document.Table, document.FieldID, selector),
			sqlgraph.To(nodelist.Table, nodelist.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, document.DocumentNodeListTable, document.DocumentNodeListColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Document entity from the query.
// Returns a *NotFoundError when no Document was found.
func (dq *DocumentQuery) First(ctx context.Context) (*Document, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{document.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DocumentQuery) FirstX(ctx context.Context) *Document {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Document ID from the query.
// Returns a *NotFoundError when no Document ID was found.
func (dq *DocumentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{document.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DocumentQuery) FirstIDX(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Document entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Document entity is found.
// Returns a *NotFoundError when no Document entities are found.
func (dq *DocumentQuery) Only(ctx context.Context) (*Document, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{document.Label}
	default:
		return nil, &NotSingularError{document.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DocumentQuery) OnlyX(ctx context.Context) *Document {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Document ID in the query.
// Returns a *NotSingularError when more than one Document ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DocumentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{document.Label}
	default:
		err = &NotSingularError{document.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DocumentQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Documents.
func (dq *DocumentQuery) All(ctx context.Context) ([]*Document, error) {
	ctx = setContextOp(ctx, dq.ctx, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Document, *DocumentQuery]()
	return withInterceptors[[]*Document](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DocumentQuery) AllX(ctx context.Context) []*Document {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Document IDs.
func (dq *DocumentQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, "IDs")
	if err = dq.Select(document.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DocumentQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DocumentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DocumentQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DocumentQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DocumentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DocumentQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DocumentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DocumentQuery) Clone() *DocumentQuery {
	if dq == nil {
		return nil
	}
	return &DocumentQuery{
		config:               dq.config,
		ctx:                  dq.ctx.Clone(),
		order:                append([]document.OrderOption{}, dq.order...),
		inters:               append([]Interceptor{}, dq.inters...),
		predicates:           append([]predicate.Document{}, dq.predicates...),
		withDocumentMetadata: dq.withDocumentMetadata.Clone(),
		withDocumentNodeList: dq.withDocumentNodeList.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithDocumentMetadata tells the query-builder to eager-load the nodes that are connected to
// the "document_metadata" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DocumentQuery) WithDocumentMetadata(opts ...func(*MetadataQuery)) *DocumentQuery {
	query := (&MetadataClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withDocumentMetadata = query
	return dq
}

// WithDocumentNodeList tells the query-builder to eager-load the nodes that are connected to
// the "document_node_list" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DocumentQuery) WithDocumentNodeList(opts ...func(*NodeListQuery)) *DocumentQuery {
	query := (&NodeListClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withDocumentNodeList = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Metadata *sbom.Metadata `json:"metadata,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Document.Query().
//		GroupBy(document.FieldMetadata).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DocumentQuery) GroupBy(field string, fields ...string) *DocumentGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DocumentGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = document.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Metadata *sbom.Metadata `json:"metadata,omitempty"`
//	}
//
//	client.Document.Query().
//		Select(document.FieldMetadata).
//		Scan(ctx, &v)
func (dq *DocumentQuery) Select(fields ...string) *DocumentSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DocumentSelect{DocumentQuery: dq}
	sbuild.label = document.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DocumentSelect configured with the given aggregations.
func (dq *DocumentQuery) Aggregate(fns ...AggregateFunc) *DocumentSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DocumentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !document.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DocumentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Document, error) {
	var (
		nodes       = []*Document{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [2]bool{
			dq.withDocumentMetadata != nil,
			dq.withDocumentNodeList != nil,
		}
	)
	if dq.withDocumentMetadata != nil || dq.withDocumentNodeList != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, document.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Document).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Document{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withDocumentMetadata; query != nil {
		if err := dq.loadDocumentMetadata(ctx, query, nodes, nil,
			func(n *Document, e *Metadata) { n.Edges.DocumentMetadata = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withDocumentNodeList; query != nil {
		if err := dq.loadDocumentNodeList(ctx, query, nodes, nil,
			func(n *Document, e *NodeList) { n.Edges.DocumentNodeList = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DocumentQuery) loadDocumentMetadata(ctx context.Context, query *MetadataQuery, nodes []*Document, init func(*Document), assign func(*Document, *Metadata)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*Document)
	for i := range nodes {
		if nodes[i].metadata_document == nil {
			continue
		}
		fk := *nodes[i].metadata_document
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(metadata.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "metadata_document" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dq *DocumentQuery) loadDocumentNodeList(ctx context.Context, query *NodeListQuery, nodes []*Document, init func(*Document), assign func(*Document, *NodeList)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Document)
	for i := range nodes {
		if nodes[i].node_list_document == nil {
			continue
		}
		fk := *nodes[i].node_list_document
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(nodelist.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "node_list_document" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dq *DocumentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DocumentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(document.Table, document.Columns, sqlgraph.NewFieldSpec(document.FieldID, field.TypeInt))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, document.FieldID)
		for i := range fields {
			if fields[i] != document.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DocumentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(document.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = document.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DocumentGroupBy is the group-by builder for Document entities.
type DocumentGroupBy struct {
	selector
	build *DocumentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DocumentGroupBy) Aggregate(fns ...AggregateFunc) *DocumentGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DocumentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DocumentQuery, *DocumentGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DocumentGroupBy) sqlScan(ctx context.Context, root *DocumentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DocumentSelect is the builder for selecting fields of Document entities.
type DocumentSelect struct {
	*DocumentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DocumentSelect) Aggregate(fns ...AggregateFunc) *DocumentSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DocumentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DocumentQuery, *DocumentSelect](ctx, ds.DocumentQuery, ds, ds.inters, v)
}

func (ds *DocumentSelect) sqlScan(ctx context.Context, root *DocumentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
