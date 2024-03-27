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
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bom-squad/protobom/ent/identifiersentry"
	"github.com/bom-squad/protobom/ent/node"
	"github.com/bom-squad/protobom/ent/predicate"
)

// IdentifiersEntryUpdate is the builder for updating IdentifiersEntry entities.
type IdentifiersEntryUpdate struct {
	config
	hooks    []Hook
	mutation *IdentifiersEntryMutation
}

// Where appends a list predicates to the IdentifiersEntryUpdate builder.
func (ieu *IdentifiersEntryUpdate) Where(ps ...predicate.IdentifiersEntry) *IdentifiersEntryUpdate {
	ieu.mutation.Where(ps...)
	return ieu
}

// SetSoftwareIdentifierType sets the "software_identifier_type" field.
func (ieu *IdentifiersEntryUpdate) SetSoftwareIdentifierType(iit identifiersentry.SoftwareIdentifierType) *IdentifiersEntryUpdate {
	ieu.mutation.SetSoftwareIdentifierType(iit)
	return ieu
}

// SetNillableSoftwareIdentifierType sets the "software_identifier_type" field if the given value is not nil.
func (ieu *IdentifiersEntryUpdate) SetNillableSoftwareIdentifierType(iit *identifiersentry.SoftwareIdentifierType) *IdentifiersEntryUpdate {
	if iit != nil {
		ieu.SetSoftwareIdentifierType(*iit)
	}
	return ieu
}

// SetSoftwareIdentifierValue sets the "software_identifier_value" field.
func (ieu *IdentifiersEntryUpdate) SetSoftwareIdentifierValue(s string) *IdentifiersEntryUpdate {
	ieu.mutation.SetSoftwareIdentifierValue(s)
	return ieu
}

// SetNillableSoftwareIdentifierValue sets the "software_identifier_value" field if the given value is not nil.
func (ieu *IdentifiersEntryUpdate) SetNillableSoftwareIdentifierValue(s *string) *IdentifiersEntryUpdate {
	if s != nil {
		ieu.SetSoftwareIdentifierValue(*s)
	}
	return ieu
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (ieu *IdentifiersEntryUpdate) AddNodeIDs(ids ...string) *IdentifiersEntryUpdate {
	ieu.mutation.AddNodeIDs(ids...)
	return ieu
}

// AddNodes adds the "nodes" edges to the Node entity.
func (ieu *IdentifiersEntryUpdate) AddNodes(n ...*Node) *IdentifiersEntryUpdate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ieu.AddNodeIDs(ids...)
}

// Mutation returns the IdentifiersEntryMutation object of the builder.
func (ieu *IdentifiersEntryUpdate) Mutation() *IdentifiersEntryMutation {
	return ieu.mutation
}

// ClearNodes clears all "nodes" edges to the Node entity.
func (ieu *IdentifiersEntryUpdate) ClearNodes() *IdentifiersEntryUpdate {
	ieu.mutation.ClearNodes()
	return ieu
}

// RemoveNodeIDs removes the "nodes" edge to Node entities by IDs.
func (ieu *IdentifiersEntryUpdate) RemoveNodeIDs(ids ...string) *IdentifiersEntryUpdate {
	ieu.mutation.RemoveNodeIDs(ids...)
	return ieu
}

// RemoveNodes removes "nodes" edges to Node entities.
func (ieu *IdentifiersEntryUpdate) RemoveNodes(n ...*Node) *IdentifiersEntryUpdate {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ieu.RemoveNodeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ieu *IdentifiersEntryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ieu.sqlSave, ieu.mutation, ieu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ieu *IdentifiersEntryUpdate) SaveX(ctx context.Context) int {
	affected, err := ieu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ieu *IdentifiersEntryUpdate) Exec(ctx context.Context) error {
	_, err := ieu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ieu *IdentifiersEntryUpdate) ExecX(ctx context.Context) {
	if err := ieu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ieu *IdentifiersEntryUpdate) check() error {
	if v, ok := ieu.mutation.SoftwareIdentifierType(); ok {
		if err := identifiersentry.SoftwareIdentifierTypeValidator(v); err != nil {
			return &ValidationError{Name: "software_identifier_type", err: fmt.Errorf(`ent: validator failed for field "IdentifiersEntry.software_identifier_type": %w`, err)}
		}
	}
	return nil
}

func (ieu *IdentifiersEntryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ieu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(identifiersentry.Table, identifiersentry.Columns, sqlgraph.NewFieldSpec(identifiersentry.FieldID, field.TypeInt))
	if ps := ieu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ieu.mutation.SoftwareIdentifierType(); ok {
		_spec.SetField(identifiersentry.FieldSoftwareIdentifierType, field.TypeEnum, value)
	}
	if value, ok := ieu.mutation.SoftwareIdentifierValue(); ok {
		_spec.SetField(identifiersentry.FieldSoftwareIdentifierValue, field.TypeString, value)
	}
	if ieu.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   identifiersentry.NodesTable,
			Columns: identifiersentry.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ieu.mutation.RemovedNodesIDs(); len(nodes) > 0 && !ieu.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   identifiersentry.NodesTable,
			Columns: identifiersentry.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ieu.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   identifiersentry.NodesTable,
			Columns: identifiersentry.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ieu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{identifiersentry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ieu.mutation.done = true
	return n, nil
}

// IdentifiersEntryUpdateOne is the builder for updating a single IdentifiersEntry entity.
type IdentifiersEntryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IdentifiersEntryMutation
}

// SetSoftwareIdentifierType sets the "software_identifier_type" field.
func (ieuo *IdentifiersEntryUpdateOne) SetSoftwareIdentifierType(iit identifiersentry.SoftwareIdentifierType) *IdentifiersEntryUpdateOne {
	ieuo.mutation.SetSoftwareIdentifierType(iit)
	return ieuo
}

// SetNillableSoftwareIdentifierType sets the "software_identifier_type" field if the given value is not nil.
func (ieuo *IdentifiersEntryUpdateOne) SetNillableSoftwareIdentifierType(iit *identifiersentry.SoftwareIdentifierType) *IdentifiersEntryUpdateOne {
	if iit != nil {
		ieuo.SetSoftwareIdentifierType(*iit)
	}
	return ieuo
}

// SetSoftwareIdentifierValue sets the "software_identifier_value" field.
func (ieuo *IdentifiersEntryUpdateOne) SetSoftwareIdentifierValue(s string) *IdentifiersEntryUpdateOne {
	ieuo.mutation.SetSoftwareIdentifierValue(s)
	return ieuo
}

// SetNillableSoftwareIdentifierValue sets the "software_identifier_value" field if the given value is not nil.
func (ieuo *IdentifiersEntryUpdateOne) SetNillableSoftwareIdentifierValue(s *string) *IdentifiersEntryUpdateOne {
	if s != nil {
		ieuo.SetSoftwareIdentifierValue(*s)
	}
	return ieuo
}

// AddNodeIDs adds the "nodes" edge to the Node entity by IDs.
func (ieuo *IdentifiersEntryUpdateOne) AddNodeIDs(ids ...string) *IdentifiersEntryUpdateOne {
	ieuo.mutation.AddNodeIDs(ids...)
	return ieuo
}

// AddNodes adds the "nodes" edges to the Node entity.
func (ieuo *IdentifiersEntryUpdateOne) AddNodes(n ...*Node) *IdentifiersEntryUpdateOne {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ieuo.AddNodeIDs(ids...)
}

// Mutation returns the IdentifiersEntryMutation object of the builder.
func (ieuo *IdentifiersEntryUpdateOne) Mutation() *IdentifiersEntryMutation {
	return ieuo.mutation
}

// ClearNodes clears all "nodes" edges to the Node entity.
func (ieuo *IdentifiersEntryUpdateOne) ClearNodes() *IdentifiersEntryUpdateOne {
	ieuo.mutation.ClearNodes()
	return ieuo
}

// RemoveNodeIDs removes the "nodes" edge to Node entities by IDs.
func (ieuo *IdentifiersEntryUpdateOne) RemoveNodeIDs(ids ...string) *IdentifiersEntryUpdateOne {
	ieuo.mutation.RemoveNodeIDs(ids...)
	return ieuo
}

// RemoveNodes removes "nodes" edges to Node entities.
func (ieuo *IdentifiersEntryUpdateOne) RemoveNodes(n ...*Node) *IdentifiersEntryUpdateOne {
	ids := make([]string, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return ieuo.RemoveNodeIDs(ids...)
}

// Where appends a list predicates to the IdentifiersEntryUpdate builder.
func (ieuo *IdentifiersEntryUpdateOne) Where(ps ...predicate.IdentifiersEntry) *IdentifiersEntryUpdateOne {
	ieuo.mutation.Where(ps...)
	return ieuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ieuo *IdentifiersEntryUpdateOne) Select(field string, fields ...string) *IdentifiersEntryUpdateOne {
	ieuo.fields = append([]string{field}, fields...)
	return ieuo
}

// Save executes the query and returns the updated IdentifiersEntry entity.
func (ieuo *IdentifiersEntryUpdateOne) Save(ctx context.Context) (*IdentifiersEntry, error) {
	return withHooks(ctx, ieuo.sqlSave, ieuo.mutation, ieuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ieuo *IdentifiersEntryUpdateOne) SaveX(ctx context.Context) *IdentifiersEntry {
	node, err := ieuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ieuo *IdentifiersEntryUpdateOne) Exec(ctx context.Context) error {
	_, err := ieuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ieuo *IdentifiersEntryUpdateOne) ExecX(ctx context.Context) {
	if err := ieuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ieuo *IdentifiersEntryUpdateOne) check() error {
	if v, ok := ieuo.mutation.SoftwareIdentifierType(); ok {
		if err := identifiersentry.SoftwareIdentifierTypeValidator(v); err != nil {
			return &ValidationError{Name: "software_identifier_type", err: fmt.Errorf(`ent: validator failed for field "IdentifiersEntry.software_identifier_type": %w`, err)}
		}
	}
	return nil
}

func (ieuo *IdentifiersEntryUpdateOne) sqlSave(ctx context.Context) (_node *IdentifiersEntry, err error) {
	if err := ieuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(identifiersentry.Table, identifiersentry.Columns, sqlgraph.NewFieldSpec(identifiersentry.FieldID, field.TypeInt))
	id, ok := ieuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "IdentifiersEntry.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ieuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, identifiersentry.FieldID)
		for _, f := range fields {
			if !identifiersentry.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != identifiersentry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ieuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ieuo.mutation.SoftwareIdentifierType(); ok {
		_spec.SetField(identifiersentry.FieldSoftwareIdentifierType, field.TypeEnum, value)
	}
	if value, ok := ieuo.mutation.SoftwareIdentifierValue(); ok {
		_spec.SetField(identifiersentry.FieldSoftwareIdentifierValue, field.TypeString, value)
	}
	if ieuo.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   identifiersentry.NodesTable,
			Columns: identifiersentry.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ieuo.mutation.RemovedNodesIDs(); len(nodes) > 0 && !ieuo.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   identifiersentry.NodesTable,
			Columns: identifiersentry.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ieuo.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   identifiersentry.NodesTable,
			Columns: identifiersentry.NodesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(node.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &IdentifiersEntry{config: ieuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ieuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{identifiersentry.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ieuo.mutation.done = true
	return _node, nil
}
