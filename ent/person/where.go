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

package person

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/bom-squad/protobom/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldName, v))
}

// IsOrg applies equality check predicate on the "is_org" field. It's identical to IsOrgEQ.
func IsOrg(v bool) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldIsOrg, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldEmail, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldURL, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldPhone, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldName, v))
}

// IsOrgEQ applies the EQ predicate on the "is_org" field.
func IsOrgEQ(v bool) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldIsOrg, v))
}

// IsOrgNEQ applies the NEQ predicate on the "is_org" field.
func IsOrgNEQ(v bool) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldIsOrg, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldEmail, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldURL, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Person {
	return predicate.Person(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Person {
	return predicate.Person(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Person {
	return predicate.Person(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Person {
	return predicate.Person(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Person {
	return predicate.Person(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Person {
	return predicate.Person(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Person {
	return predicate.Person(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Person {
	return predicate.Person(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Person {
	return predicate.Person(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Person {
	return predicate.Person(sql.FieldContainsFold(FieldPhone, v))
}

// HasContacts applies the HasEdge predicate on the "contacts" edge.
func HasContacts() predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ContactsTable, ContactsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContactsWith applies the HasEdge predicate on the "contacts" edge with a given conditions (other predicates).
func HasContactsWith(preds ...predicate.Person) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := newContactsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMetadata applies the HasEdge predicate on the "metadata" edge.
func HasMetadata() predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MetadataTable, MetadataColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMetadataWith applies the HasEdge predicate on the "metadata" edge with a given conditions (other predicates).
func HasMetadataWith(preds ...predicate.Metadata) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := newMetadataStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNodeSupplier applies the HasEdge predicate on the "node_supplier" edge.
func HasNodeSupplier() predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NodeSupplierTable, NodeSupplierColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodeSupplierWith applies the HasEdge predicate on the "node_supplier" edge with a given conditions (other predicates).
func HasNodeSupplierWith(preds ...predicate.Node) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := newNodeSupplierStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNodeOriginator applies the HasEdge predicate on the "node_originator" edge.
func HasNodeOriginator() predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NodeOriginatorTable, NodeOriginatorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodeOriginatorWith applies the HasEdge predicate on the "node_originator" edge with a given conditions (other predicates).
func HasNodeOriginatorWith(preds ...predicate.Node) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := newNodeOriginatorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPersonContact applies the HasEdge predicate on the "person_contact" edge.
func HasPersonContact() predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PersonContactTable, PersonContactColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonContactWith applies the HasEdge predicate on the "person_contact" edge with a given conditions (other predicates).
func HasPersonContactWith(preds ...predicate.Person) predicate.Person {
	return predicate.Person(func(s *sql.Selector) {
		step := newPersonContactStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Person) predicate.Person {
	return predicate.Person(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Person) predicate.Person {
	return predicate.Person(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Person) predicate.Person {
	return predicate.Person(sql.NotPredicates(p))
}
