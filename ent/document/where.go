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

package document

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/bom-squad/protobom/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Document {
	return predicate.Document(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Document {
	return predicate.Document(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Document {
	return predicate.Document(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Document {
	return predicate.Document(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Document {
	return predicate.Document(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Document {
	return predicate.Document(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Document {
	return predicate.Document(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Document {
	return predicate.Document(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Document {
	return predicate.Document(sql.FieldLTE(FieldID, id))
}

// HasDocumentMetadata applies the HasEdge predicate on the "document_metadata" edge.
func HasDocumentMetadata() predicate.Document {
	return predicate.Document(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, DocumentMetadataTable, DocumentMetadataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDocumentMetadataWith applies the HasEdge predicate on the "document_metadata" edge with a given conditions (other predicates).
func HasDocumentMetadataWith(preds ...predicate.Metadata) predicate.Document {
	return predicate.Document(func(s *sql.Selector) {
		step := newDocumentMetadataStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDocumentNodeList applies the HasEdge predicate on the "document_node_list" edge.
func HasDocumentNodeList() predicate.Document {
	return predicate.Document(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, DocumentNodeListTable, DocumentNodeListColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDocumentNodeListWith applies the HasEdge predicate on the "document_node_list" edge with a given conditions (other predicates).
func HasDocumentNodeListWith(preds ...predicate.NodeList) predicate.Document {
	return predicate.Document(func(s *sql.Selector) {
		step := newDocumentNodeListStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Document) predicate.Document {
	return predicate.Document(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Document) predicate.Document {
	return predicate.Document(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Document) predicate.Document {
	return predicate.Document(sql.NotPredicates(p))
}
