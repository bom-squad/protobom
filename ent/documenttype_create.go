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
	"github.com/bom-squad/protobom/ent/documenttype"
	"github.com/bom-squad/protobom/ent/metadata"
)

// DocumentTypeCreate is the builder for creating a DocumentType entity.
type DocumentTypeCreate struct {
	config
	mutation *DocumentTypeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetType sets the "type" field.
func (dtc *DocumentTypeCreate) SetType(d documenttype.Type) *DocumentTypeCreate {
	dtc.mutation.SetType(d)
	return dtc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (dtc *DocumentTypeCreate) SetNillableType(d *documenttype.Type) *DocumentTypeCreate {
	if d != nil {
		dtc.SetType(*d)
	}
	return dtc
}

// SetName sets the "name" field.
func (dtc *DocumentTypeCreate) SetName(s string) *DocumentTypeCreate {
	dtc.mutation.SetName(s)
	return dtc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (dtc *DocumentTypeCreate) SetNillableName(s *string) *DocumentTypeCreate {
	if s != nil {
		dtc.SetName(*s)
	}
	return dtc
}

// SetDescription sets the "description" field.
func (dtc *DocumentTypeCreate) SetDescription(s string) *DocumentTypeCreate {
	dtc.mutation.SetDescription(s)
	return dtc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (dtc *DocumentTypeCreate) SetNillableDescription(s *string) *DocumentTypeCreate {
	if s != nil {
		dtc.SetDescription(*s)
	}
	return dtc
}

// SetMetadataID sets the "metadata" edge to the Metadata entity by ID.
func (dtc *DocumentTypeCreate) SetMetadataID(id string) *DocumentTypeCreate {
	dtc.mutation.SetMetadataID(id)
	return dtc
}

// SetNillableMetadataID sets the "metadata" edge to the Metadata entity by ID if the given value is not nil.
func (dtc *DocumentTypeCreate) SetNillableMetadataID(id *string) *DocumentTypeCreate {
	if id != nil {
		dtc = dtc.SetMetadataID(*id)
	}
	return dtc
}

// SetMetadata sets the "metadata" edge to the Metadata entity.
func (dtc *DocumentTypeCreate) SetMetadata(m *Metadata) *DocumentTypeCreate {
	return dtc.SetMetadataID(m.ID)
}

// Mutation returns the DocumentTypeMutation object of the builder.
func (dtc *DocumentTypeCreate) Mutation() *DocumentTypeMutation {
	return dtc.mutation
}

// Save creates the DocumentType in the database.
func (dtc *DocumentTypeCreate) Save(ctx context.Context) (*DocumentType, error) {
	return withHooks(ctx, dtc.sqlSave, dtc.mutation, dtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dtc *DocumentTypeCreate) SaveX(ctx context.Context) *DocumentType {
	v, err := dtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dtc *DocumentTypeCreate) Exec(ctx context.Context) error {
	_, err := dtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtc *DocumentTypeCreate) ExecX(ctx context.Context) {
	if err := dtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dtc *DocumentTypeCreate) check() error {
	if v, ok := dtc.mutation.GetType(); ok {
		if err := documenttype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "DocumentType.type": %w`, err)}
		}
	}
	return nil
}

func (dtc *DocumentTypeCreate) sqlSave(ctx context.Context) (*DocumentType, error) {
	if err := dtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dtc.mutation.id = &_node.ID
	dtc.mutation.done = true
	return _node, nil
}

func (dtc *DocumentTypeCreate) createSpec() (*DocumentType, *sqlgraph.CreateSpec) {
	var (
		_node = &DocumentType{config: dtc.config}
		_spec = sqlgraph.NewCreateSpec(documenttype.Table, sqlgraph.NewFieldSpec(documenttype.FieldID, field.TypeInt))
	)
	_spec.OnConflict = dtc.conflict
	if value, ok := dtc.mutation.GetType(); ok {
		_spec.SetField(documenttype.FieldType, field.TypeEnum, value)
		_node.Type = &value
	}
	if value, ok := dtc.mutation.Name(); ok {
		_spec.SetField(documenttype.FieldName, field.TypeString, value)
		_node.Name = &value
	}
	if value, ok := dtc.mutation.Description(); ok {
		_spec.SetField(documenttype.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	if nodes := dtc.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   documenttype.MetadataTable,
			Columns: []string{documenttype.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.metadata_metadata_document_types = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DocumentType.Create().
//		SetType(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DocumentTypeUpsert) {
//			SetType(v+v).
//		}).
//		Exec(ctx)
func (dtc *DocumentTypeCreate) OnConflict(opts ...sql.ConflictOption) *DocumentTypeUpsertOne {
	dtc.conflict = opts
	return &DocumentTypeUpsertOne{
		create: dtc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DocumentType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dtc *DocumentTypeCreate) OnConflictColumns(columns ...string) *DocumentTypeUpsertOne {
	dtc.conflict = append(dtc.conflict, sql.ConflictColumns(columns...))
	return &DocumentTypeUpsertOne{
		create: dtc,
	}
}

type (
	// DocumentTypeUpsertOne is the builder for "upsert"-ing
	//  one DocumentType node.
	DocumentTypeUpsertOne struct {
		create *DocumentTypeCreate
	}

	// DocumentTypeUpsert is the "OnConflict" setter.
	DocumentTypeUpsert struct {
		*sql.UpdateSet
	}
)

// SetType sets the "type" field.
func (u *DocumentTypeUpsert) SetType(v documenttype.Type) *DocumentTypeUpsert {
	u.Set(documenttype.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DocumentTypeUpsert) UpdateType() *DocumentTypeUpsert {
	u.SetExcluded(documenttype.FieldType)
	return u
}

// ClearType clears the value of the "type" field.
func (u *DocumentTypeUpsert) ClearType() *DocumentTypeUpsert {
	u.SetNull(documenttype.FieldType)
	return u
}

// SetName sets the "name" field.
func (u *DocumentTypeUpsert) SetName(v string) *DocumentTypeUpsert {
	u.Set(documenttype.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DocumentTypeUpsert) UpdateName() *DocumentTypeUpsert {
	u.SetExcluded(documenttype.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *DocumentTypeUpsert) ClearName() *DocumentTypeUpsert {
	u.SetNull(documenttype.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *DocumentTypeUpsert) SetDescription(v string) *DocumentTypeUpsert {
	u.Set(documenttype.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *DocumentTypeUpsert) UpdateDescription() *DocumentTypeUpsert {
	u.SetExcluded(documenttype.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *DocumentTypeUpsert) ClearDescription() *DocumentTypeUpsert {
	u.SetNull(documenttype.FieldDescription)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.DocumentType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DocumentTypeUpsertOne) UpdateNewValues() *DocumentTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DocumentType.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DocumentTypeUpsertOne) Ignore() *DocumentTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DocumentTypeUpsertOne) DoNothing() *DocumentTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DocumentTypeCreate.OnConflict
// documentation for more info.
func (u *DocumentTypeUpsertOne) Update(set func(*DocumentTypeUpsert)) *DocumentTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DocumentTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetType sets the "type" field.
func (u *DocumentTypeUpsertOne) SetType(v documenttype.Type) *DocumentTypeUpsertOne {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DocumentTypeUpsertOne) UpdateType() *DocumentTypeUpsertOne {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.UpdateType()
	})
}

// ClearType clears the value of the "type" field.
func (u *DocumentTypeUpsertOne) ClearType() *DocumentTypeUpsertOne {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.ClearType()
	})
}

// SetName sets the "name" field.
func (u *DocumentTypeUpsertOne) SetName(v string) *DocumentTypeUpsertOne {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DocumentTypeUpsertOne) UpdateName() *DocumentTypeUpsertOne {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *DocumentTypeUpsertOne) ClearName() *DocumentTypeUpsertOne {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.ClearName()
	})
}

// SetDescription sets the "description" field.
func (u *DocumentTypeUpsertOne) SetDescription(v string) *DocumentTypeUpsertOne {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *DocumentTypeUpsertOne) UpdateDescription() *DocumentTypeUpsertOne {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *DocumentTypeUpsertOne) ClearDescription() *DocumentTypeUpsertOne {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.ClearDescription()
	})
}

// Exec executes the query.
func (u *DocumentTypeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DocumentTypeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DocumentTypeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DocumentTypeUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DocumentTypeUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DocumentTypeCreateBulk is the builder for creating many DocumentType entities in bulk.
type DocumentTypeCreateBulk struct {
	config
	err      error
	builders []*DocumentTypeCreate
	conflict []sql.ConflictOption
}

// Save creates the DocumentType entities in the database.
func (dtcb *DocumentTypeCreateBulk) Save(ctx context.Context) ([]*DocumentType, error) {
	if dtcb.err != nil {
		return nil, dtcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dtcb.builders))
	nodes := make([]*DocumentType, len(dtcb.builders))
	mutators := make([]Mutator, len(dtcb.builders))
	for i := range dtcb.builders {
		func(i int, root context.Context) {
			builder := dtcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DocumentTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, dtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dtcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dtcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, dtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dtcb *DocumentTypeCreateBulk) SaveX(ctx context.Context) []*DocumentType {
	v, err := dtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dtcb *DocumentTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := dtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dtcb *DocumentTypeCreateBulk) ExecX(ctx context.Context) {
	if err := dtcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DocumentType.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DocumentTypeUpsert) {
//			SetType(v+v).
//		}).
//		Exec(ctx)
func (dtcb *DocumentTypeCreateBulk) OnConflict(opts ...sql.ConflictOption) *DocumentTypeUpsertBulk {
	dtcb.conflict = opts
	return &DocumentTypeUpsertBulk{
		create: dtcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DocumentType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dtcb *DocumentTypeCreateBulk) OnConflictColumns(columns ...string) *DocumentTypeUpsertBulk {
	dtcb.conflict = append(dtcb.conflict, sql.ConflictColumns(columns...))
	return &DocumentTypeUpsertBulk{
		create: dtcb,
	}
}

// DocumentTypeUpsertBulk is the builder for "upsert"-ing
// a bulk of DocumentType nodes.
type DocumentTypeUpsertBulk struct {
	create *DocumentTypeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DocumentType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DocumentTypeUpsertBulk) UpdateNewValues() *DocumentTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DocumentType.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DocumentTypeUpsertBulk) Ignore() *DocumentTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DocumentTypeUpsertBulk) DoNothing() *DocumentTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DocumentTypeCreateBulk.OnConflict
// documentation for more info.
func (u *DocumentTypeUpsertBulk) Update(set func(*DocumentTypeUpsert)) *DocumentTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DocumentTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetType sets the "type" field.
func (u *DocumentTypeUpsertBulk) SetType(v documenttype.Type) *DocumentTypeUpsertBulk {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *DocumentTypeUpsertBulk) UpdateType() *DocumentTypeUpsertBulk {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.UpdateType()
	})
}

// ClearType clears the value of the "type" field.
func (u *DocumentTypeUpsertBulk) ClearType() *DocumentTypeUpsertBulk {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.ClearType()
	})
}

// SetName sets the "name" field.
func (u *DocumentTypeUpsertBulk) SetName(v string) *DocumentTypeUpsertBulk {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DocumentTypeUpsertBulk) UpdateName() *DocumentTypeUpsertBulk {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *DocumentTypeUpsertBulk) ClearName() *DocumentTypeUpsertBulk {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.ClearName()
	})
}

// SetDescription sets the "description" field.
func (u *DocumentTypeUpsertBulk) SetDescription(v string) *DocumentTypeUpsertBulk {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *DocumentTypeUpsertBulk) UpdateDescription() *DocumentTypeUpsertBulk {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *DocumentTypeUpsertBulk) ClearDescription() *DocumentTypeUpsertBulk {
	return u.Update(func(s *DocumentTypeUpsert) {
		s.ClearDescription()
	})
}

// Exec executes the query.
func (u *DocumentTypeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DocumentTypeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DocumentTypeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DocumentTypeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
