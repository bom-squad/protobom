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
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/bom-squad/protobom/ent/metadata"
	"github.com/bom-squad/protobom/ent/node"
	"github.com/bom-squad/protobom/ent/person"
)

// Person is the model entity for the Person schema.
type Person struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// IsOrg holds the value of the "is_org" field.
	IsOrg bool `json:"is_org,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonQuery when eager-loading is set.
	Edges            PersonEdges `json:"edges"`
	metadata_authors *string
	node_suppliers   *string
	node_originators *string
	selectValues     sql.SelectValues
}

// PersonEdges holds the relations/edges for other nodes in the graph.
type PersonEdges struct {
	// Contacts holds the value of the contacts edge.
	Contacts []*Person `json:"contacts,omitempty"`
	// Metadata holds the value of the metadata edge.
	Metadata *Metadata `json:"metadata,omitempty"`
	// NodeSupplier holds the value of the node_supplier edge.
	NodeSupplier *Node `json:"node_supplier,omitempty"`
	// NodeOriginator holds the value of the node_originator edge.
	NodeOriginator *Node `json:"node_originator,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ContactsOrErr returns the Contacts value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) ContactsOrErr() ([]*Person, error) {
	if e.loadedTypes[0] {
		return e.Contacts, nil
	}
	return nil, &NotLoadedError{edge: "contacts"}
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonEdges) MetadataOrErr() (*Metadata, error) {
	if e.loadedTypes[1] {
		if e.Metadata == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: metadata.Label}
		}
		return e.Metadata, nil
	}
	return nil, &NotLoadedError{edge: "metadata"}
}

// NodeSupplierOrErr returns the NodeSupplier value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonEdges) NodeSupplierOrErr() (*Node, error) {
	if e.loadedTypes[2] {
		if e.NodeSupplier == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: node.Label}
		}
		return e.NodeSupplier, nil
	}
	return nil, &NotLoadedError{edge: "node_supplier"}
}

// NodeOriginatorOrErr returns the NodeOriginator value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PersonEdges) NodeOriginatorOrErr() (*Node, error) {
	if e.loadedTypes[3] {
		if e.NodeOriginator == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: node.Label}
		}
		return e.NodeOriginator, nil
	}
	return nil, &NotLoadedError{edge: "node_originator"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Person) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case person.FieldIsOrg:
			values[i] = new(sql.NullBool)
		case person.FieldID:
			values[i] = new(sql.NullInt64)
		case person.FieldName, person.FieldEmail, person.FieldURL, person.FieldPhone:
			values[i] = new(sql.NullString)
		case person.ForeignKeys[0]: // metadata_authors
			values[i] = new(sql.NullString)
		case person.ForeignKeys[1]: // node_suppliers
			values[i] = new(sql.NullString)
		case person.ForeignKeys[2]: // node_originators
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Person fields.
func (pe *Person) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case person.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case person.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pe.Name = value.String
			}
		case person.FieldIsOrg:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_org", values[i])
			} else if value.Valid {
				pe.IsOrg = value.Bool
			}
		case person.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				pe.Email = value.String
			}
		case person.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				pe.URL = value.String
			}
		case person.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				pe.Phone = value.String
			}
		case person.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field metadata_authors", values[i])
			} else if value.Valid {
				pe.metadata_authors = new(string)
				*pe.metadata_authors = value.String
			}
		case person.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field node_suppliers", values[i])
			} else if value.Valid {
				pe.node_suppliers = new(string)
				*pe.node_suppliers = value.String
			}
		case person.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field node_originators", values[i])
			} else if value.Valid {
				pe.node_originators = new(string)
				*pe.node_originators = value.String
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Person.
// This includes values selected through modifiers, order, etc.
func (pe *Person) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// QueryContacts queries the "contacts" edge of the Person entity.
func (pe *Person) QueryContacts() *PersonQuery {
	return NewPersonClient(pe.config).QueryContacts(pe)
}

// QueryMetadata queries the "metadata" edge of the Person entity.
func (pe *Person) QueryMetadata() *MetadataQuery {
	return NewPersonClient(pe.config).QueryMetadata(pe)
}

// QueryNodeSupplier queries the "node_supplier" edge of the Person entity.
func (pe *Person) QueryNodeSupplier() *NodeQuery {
	return NewPersonClient(pe.config).QueryNodeSupplier(pe)
}

// QueryNodeOriginator queries the "node_originator" edge of the Person entity.
func (pe *Person) QueryNodeOriginator() *NodeQuery {
	return NewPersonClient(pe.config).QueryNodeOriginator(pe)
}

// Update returns a builder for updating this Person.
// Note that you need to call Person.Unwrap() before calling this method if this Person
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Person) Update() *PersonUpdateOne {
	return NewPersonClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Person entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Person) Unwrap() *Person {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Person is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Person) String() string {
	var builder strings.Builder
	builder.WriteString("Person(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", ")
	builder.WriteString("is_org=")
	builder.WriteString(fmt.Sprintf("%v", pe.IsOrg))
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(pe.Email)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(pe.URL)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(pe.Phone)
	builder.WriteByte(')')
	return builder.String()
}

// Persons is a parsable slice of Person.
type Persons []*Person
