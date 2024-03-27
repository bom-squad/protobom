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

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/bom-squad/protobom/ent/document"
	"github.com/bom-squad/protobom/ent/predicate"
)

// DocumentDelete is the builder for deleting a Document entity.
type DocumentDelete struct {
	config
	hooks    []Hook
	mutation *DocumentMutation
}

// Where appends a list predicates to the DocumentDelete builder.
func (dd *DocumentDelete) Where(ps ...predicate.Document) *DocumentDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DocumentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dd.sqlExec, dd.mutation, dd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DocumentDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DocumentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(document.Table, sqlgraph.NewFieldSpec(document.FieldID, field.TypeInt))
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dd.mutation.done = true
	return affected, err
}

// DocumentDeleteOne is the builder for deleting a single Document entity.
type DocumentDeleteOne struct {
	dd *DocumentDelete
}

// Where appends a list predicates to the DocumentDelete builder.
func (ddo *DocumentDeleteOne) Where(ps ...predicate.Document) *DocumentDeleteOne {
	ddo.dd.mutation.Where(ps...)
	return ddo
}

// Exec executes the deletion query.
func (ddo *DocumentDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{document.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DocumentDeleteOne) ExecX(ctx context.Context) {
	if err := ddo.Exec(ctx); err != nil {
		panic(err)
	}
}
