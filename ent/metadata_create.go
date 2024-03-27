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

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bom-squad/protobom/ent/document"
	"github.com/bom-squad/protobom/ent/documenttype"
	"github.com/bom-squad/protobom/ent/metadata"
	"github.com/bom-squad/protobom/ent/person"
	"github.com/bom-squad/protobom/ent/timestamp"
	"github.com/bom-squad/protobom/ent/tool"
)

// MetadataCreate is the builder for creating a Metadata entity.
type MetadataCreate struct {
	config
	mutation *MetadataMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetVersion sets the "version" field.
func (mc *MetadataCreate) SetVersion(s string) *MetadataCreate {
	mc.mutation.SetVersion(s)
	return mc
}

// SetName sets the "name" field.
func (mc *MetadataCreate) SetName(s string) *MetadataCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetComment sets the "comment" field.
func (mc *MetadataCreate) SetComment(s string) *MetadataCreate {
	mc.mutation.SetComment(s)
	return mc
}

// SetID sets the "id" field.
func (mc *MetadataCreate) SetID(s string) *MetadataCreate {
	mc.mutation.SetID(s)
	return mc
}

// AddToolIDs adds the "tools" edge to the Tool entity by IDs.
func (mc *MetadataCreate) AddToolIDs(ids ...int) *MetadataCreate {
	mc.mutation.AddToolIDs(ids...)
	return mc
}

// AddTools adds the "tools" edges to the Tool entity.
func (mc *MetadataCreate) AddTools(t ...*Tool) *MetadataCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return mc.AddToolIDs(ids...)
}

// AddAuthorIDs adds the "authors" edge to the Person entity by IDs.
func (mc *MetadataCreate) AddAuthorIDs(ids ...int) *MetadataCreate {
	mc.mutation.AddAuthorIDs(ids...)
	return mc
}

// AddAuthors adds the "authors" edges to the Person entity.
func (mc *MetadataCreate) AddAuthors(p ...*Person) *MetadataCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mc.AddAuthorIDs(ids...)
}

// AddDocumentTypeIDs adds the "documentTypes" edge to the DocumentType entity by IDs.
func (mc *MetadataCreate) AddDocumentTypeIDs(ids ...int) *MetadataCreate {
	mc.mutation.AddDocumentTypeIDs(ids...)
	return mc
}

// AddDocumentTypes adds the "documentTypes" edges to the DocumentType entity.
func (mc *MetadataCreate) AddDocumentTypes(d ...*DocumentType) *MetadataCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mc.AddDocumentTypeIDs(ids...)
}

// AddDateIDs adds the "date" edge to the Timestamp entity by IDs.
func (mc *MetadataCreate) AddDateIDs(ids ...int) *MetadataCreate {
	mc.mutation.AddDateIDs(ids...)
	return mc
}

// AddDate adds the "date" edges to the Timestamp entity.
func (mc *MetadataCreate) AddDate(t ...*Timestamp) *MetadataCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return mc.AddDateIDs(ids...)
}

// SetDocumentID sets the "document" edge to the Document entity by ID.
func (mc *MetadataCreate) SetDocumentID(id int) *MetadataCreate {
	mc.mutation.SetDocumentID(id)
	return mc
}

// SetDocument sets the "document" edge to the Document entity.
func (mc *MetadataCreate) SetDocument(d *Document) *MetadataCreate {
	return mc.SetDocumentID(d.ID)
}

// Mutation returns the MetadataMutation object of the builder.
func (mc *MetadataCreate) Mutation() *MetadataMutation {
	return mc.mutation
}

// Save creates the Metadata in the database.
func (mc *MetadataCreate) Save(ctx context.Context) (*Metadata, error) {
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MetadataCreate) SaveX(ctx context.Context) *Metadata {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MetadataCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MetadataCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MetadataCreate) check() error {
	if _, ok := mc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "Metadata.version"`)}
	}
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Metadata.name"`)}
	}
	if _, ok := mc.mutation.Comment(); !ok {
		return &ValidationError{Name: "comment", err: errors.New(`ent: missing required field "Metadata.comment"`)}
	}
	if _, ok := mc.mutation.DocumentID(); !ok {
		return &ValidationError{Name: "document", err: errors.New(`ent: missing required edge "Metadata.document"`)}
	}
	return nil
}

func (mc *MetadataCreate) sqlSave(ctx context.Context) (*Metadata, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Metadata.ID type: %T", _spec.ID.Value)
		}
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MetadataCreate) createSpec() (*Metadata, *sqlgraph.CreateSpec) {
	var (
		_node = &Metadata{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(metadata.Table, sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString))
	)
	_spec.OnConflict = mc.conflict
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.Version(); ok {
		_spec.SetField(metadata.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(metadata.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.Comment(); ok {
		_spec.SetField(metadata.FieldComment, field.TypeString, value)
		_node.Comment = value
	}
	if nodes := mc.mutation.ToolsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.ToolsTable,
			Columns: []string{metadata.ToolsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tool.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.AuthorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.AuthorsTable,
			Columns: []string{metadata.AuthorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.DocumentTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.DocumentTypesTable,
			Columns: []string{metadata.DocumentTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(documenttype.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.DateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   metadata.DateTable,
			Columns: []string{metadata.DateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timestamp.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.DocumentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   metadata.DocumentTable,
			Columns: []string{metadata.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(document.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.document_metadata = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Metadata.Create().
//		SetVersion(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MetadataUpsert) {
//			SetVersion(v+v).
//		}).
//		Exec(ctx)
func (mc *MetadataCreate) OnConflict(opts ...sql.ConflictOption) *MetadataUpsertOne {
	mc.conflict = opts
	return &MetadataUpsertOne{
		create: mc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Metadata.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mc *MetadataCreate) OnConflictColumns(columns ...string) *MetadataUpsertOne {
	mc.conflict = append(mc.conflict, sql.ConflictColumns(columns...))
	return &MetadataUpsertOne{
		create: mc,
	}
}

type (
	// MetadataUpsertOne is the builder for "upsert"-ing
	//  one Metadata node.
	MetadataUpsertOne struct {
		create *MetadataCreate
	}

	// MetadataUpsert is the "OnConflict" setter.
	MetadataUpsert struct {
		*sql.UpdateSet
	}
)

// SetVersion sets the "version" field.
func (u *MetadataUpsert) SetVersion(v string) *MetadataUpsert {
	u.Set(metadata.FieldVersion, v)
	return u
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *MetadataUpsert) UpdateVersion() *MetadataUpsert {
	u.SetExcluded(metadata.FieldVersion)
	return u
}

// SetName sets the "name" field.
func (u *MetadataUpsert) SetName(v string) *MetadataUpsert {
	u.Set(metadata.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MetadataUpsert) UpdateName() *MetadataUpsert {
	u.SetExcluded(metadata.FieldName)
	return u
}

// SetComment sets the "comment" field.
func (u *MetadataUpsert) SetComment(v string) *MetadataUpsert {
	u.Set(metadata.FieldComment, v)
	return u
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *MetadataUpsert) UpdateComment() *MetadataUpsert {
	u.SetExcluded(metadata.FieldComment)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Metadata.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(metadata.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MetadataUpsertOne) UpdateNewValues() *MetadataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(metadata.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Metadata.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MetadataUpsertOne) Ignore() *MetadataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MetadataUpsertOne) DoNothing() *MetadataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MetadataCreate.OnConflict
// documentation for more info.
func (u *MetadataUpsertOne) Update(set func(*MetadataUpsert)) *MetadataUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MetadataUpsert{UpdateSet: update})
	}))
	return u
}

// SetVersion sets the "version" field.
func (u *MetadataUpsertOne) SetVersion(v string) *MetadataUpsertOne {
	return u.Update(func(s *MetadataUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *MetadataUpsertOne) UpdateVersion() *MetadataUpsertOne {
	return u.Update(func(s *MetadataUpsert) {
		s.UpdateVersion()
	})
}

// SetName sets the "name" field.
func (u *MetadataUpsertOne) SetName(v string) *MetadataUpsertOne {
	return u.Update(func(s *MetadataUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MetadataUpsertOne) UpdateName() *MetadataUpsertOne {
	return u.Update(func(s *MetadataUpsert) {
		s.UpdateName()
	})
}

// SetComment sets the "comment" field.
func (u *MetadataUpsertOne) SetComment(v string) *MetadataUpsertOne {
	return u.Update(func(s *MetadataUpsert) {
		s.SetComment(v)
	})
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *MetadataUpsertOne) UpdateComment() *MetadataUpsertOne {
	return u.Update(func(s *MetadataUpsert) {
		s.UpdateComment()
	})
}

// Exec executes the query.
func (u *MetadataUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MetadataCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MetadataUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MetadataUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: MetadataUpsertOne.ID is not supported by MySQL driver. Use MetadataUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MetadataUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MetadataCreateBulk is the builder for creating many Metadata entities in bulk.
type MetadataCreateBulk struct {
	config
	err      error
	builders []*MetadataCreate
	conflict []sql.ConflictOption
}

// Save creates the Metadata entities in the database.
func (mcb *MetadataCreateBulk) Save(ctx context.Context) ([]*Metadata, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Metadata, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MetadataMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = mcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MetadataCreateBulk) SaveX(ctx context.Context) []*Metadata {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MetadataCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MetadataCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Metadata.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MetadataUpsert) {
//			SetVersion(v+v).
//		}).
//		Exec(ctx)
func (mcb *MetadataCreateBulk) OnConflict(opts ...sql.ConflictOption) *MetadataUpsertBulk {
	mcb.conflict = opts
	return &MetadataUpsertBulk{
		create: mcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Metadata.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mcb *MetadataCreateBulk) OnConflictColumns(columns ...string) *MetadataUpsertBulk {
	mcb.conflict = append(mcb.conflict, sql.ConflictColumns(columns...))
	return &MetadataUpsertBulk{
		create: mcb,
	}
}

// MetadataUpsertBulk is the builder for "upsert"-ing
// a bulk of Metadata nodes.
type MetadataUpsertBulk struct {
	create *MetadataCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Metadata.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(metadata.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *MetadataUpsertBulk) UpdateNewValues() *MetadataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(metadata.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Metadata.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MetadataUpsertBulk) Ignore() *MetadataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MetadataUpsertBulk) DoNothing() *MetadataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MetadataCreateBulk.OnConflict
// documentation for more info.
func (u *MetadataUpsertBulk) Update(set func(*MetadataUpsert)) *MetadataUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MetadataUpsert{UpdateSet: update})
	}))
	return u
}

// SetVersion sets the "version" field.
func (u *MetadataUpsertBulk) SetVersion(v string) *MetadataUpsertBulk {
	return u.Update(func(s *MetadataUpsert) {
		s.SetVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *MetadataUpsertBulk) UpdateVersion() *MetadataUpsertBulk {
	return u.Update(func(s *MetadataUpsert) {
		s.UpdateVersion()
	})
}

// SetName sets the "name" field.
func (u *MetadataUpsertBulk) SetName(v string) *MetadataUpsertBulk {
	return u.Update(func(s *MetadataUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *MetadataUpsertBulk) UpdateName() *MetadataUpsertBulk {
	return u.Update(func(s *MetadataUpsert) {
		s.UpdateName()
	})
}

// SetComment sets the "comment" field.
func (u *MetadataUpsertBulk) SetComment(v string) *MetadataUpsertBulk {
	return u.Update(func(s *MetadataUpsert) {
		s.SetComment(v)
	})
}

// UpdateComment sets the "comment" field to the value that was provided on create.
func (u *MetadataUpsertBulk) UpdateComment() *MetadataUpsertBulk {
	return u.Update(func(s *MetadataUpsert) {
		s.UpdateComment()
	})
}

// Exec executes the query.
func (u *MetadataUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MetadataCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MetadataCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MetadataUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
