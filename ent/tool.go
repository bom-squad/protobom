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
	"github.com/bom-squad/protobom/ent/tool"
)

// Tool is the model entity for the Tool schema.
type Tool struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// Vendor holds the value of the "vendor" field.
	Vendor string `json:"vendor,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ToolQuery when eager-loading is set.
	Edges                   ToolEdges `json:"edges"`
	metadata_metadata_tools *string
	selectValues            sql.SelectValues
}

// ToolEdges holds the relations/edges for other nodes in the graph.
type ToolEdges struct {
	// Metadata holds the value of the metadata edge.
	Metadata *Metadata `json:"metadata,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MetadataOrErr returns the Metadata value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ToolEdges) MetadataOrErr() (*Metadata, error) {
	if e.loadedTypes[0] {
		if e.Metadata == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: metadata.Label}
		}
		return e.Metadata, nil
	}
	return nil, &NotLoadedError{edge: "metadata"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tool) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tool.FieldID:
			values[i] = new(sql.NullInt64)
		case tool.FieldName, tool.FieldVersion, tool.FieldVendor:
			values[i] = new(sql.NullString)
		case tool.ForeignKeys[0]: // metadata_metadata_tools
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tool fields.
func (t *Tool) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tool.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case tool.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case tool.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				t.Version = value.String
			}
		case tool.FieldVendor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vendor", values[i])
			} else if value.Valid {
				t.Vendor = value.String
			}
		case tool.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field metadata_metadata_tools", values[i])
			} else if value.Valid {
				t.metadata_metadata_tools = new(string)
				*t.metadata_metadata_tools = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Tool.
// This includes values selected through modifiers, order, etc.
func (t *Tool) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryMetadata queries the "metadata" edge of the Tool entity.
func (t *Tool) QueryMetadata() *MetadataQuery {
	return NewToolClient(t.config).QueryMetadata(t)
}

// Update returns a builder for updating this Tool.
// Note that you need to call Tool.Unwrap() before calling this method if this Tool
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tool) Update() *ToolUpdateOne {
	return NewToolClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Tool entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Tool) Unwrap() *Tool {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tool is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tool) String() string {
	var builder strings.Builder
	builder.WriteString("Tool(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("name=")
	builder.WriteString(t.Name)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(t.Version)
	builder.WriteString(", ")
	builder.WriteString("vendor=")
	builder.WriteString(t.Vendor)
	builder.WriteByte(')')
	return builder.String()
}

// Tools is a parsable slice of Tool.
type Tools []*Tool
