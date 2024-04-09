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
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/bom-squad/protobom/ent/document"
	"github.com/bom-squad/protobom/ent/metadata"
	"github.com/bom-squad/protobom/pkg/sbom"
)

// Metadata is the model entity for the Metadata schema.
type Metadata struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment string `json:"comment,omitempty"`
	// Tools holds the value of the "tools" field.
	Tools []*sbom.Tool `json:"tools,omitempty"`
	// Authors holds the value of the "authors" field.
	Authors []*sbom.Person `json:"authors,omitempty"`
	// DocumentTypes holds the value of the "document_types" field.
	DocumentTypes []*sbom.DocumentType `json:"document_types,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MetadataQuery when eager-loading is set.
	Edges        MetadataEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MetadataEdges holds the relations/edges for other nodes in the graph.
type MetadataEdges struct {
	// MetadataTools holds the value of the metadata_tools edge.
	MetadataTools []*Tool `json:"metadata_tools,omitempty"`
	// MetadataAuthors holds the value of the metadata_authors edge.
	MetadataAuthors []*Person `json:"metadata_authors,omitempty"`
	// MetadataDocumentTypes holds the value of the metadata_document_types edge.
	MetadataDocumentTypes []*DocumentType `json:"metadata_document_types,omitempty"`
	// Document holds the value of the document edge.
	Document *Document `json:"document,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// MetadataToolsOrErr returns the MetadataTools value or an error if the edge
// was not loaded in eager-loading.
func (e MetadataEdges) MetadataToolsOrErr() ([]*Tool, error) {
	if e.loadedTypes[0] {
		return e.MetadataTools, nil
	}
	return nil, &NotLoadedError{edge: "metadata_tools"}
}

// MetadataAuthorsOrErr returns the MetadataAuthors value or an error if the edge
// was not loaded in eager-loading.
func (e MetadataEdges) MetadataAuthorsOrErr() ([]*Person, error) {
	if e.loadedTypes[1] {
		return e.MetadataAuthors, nil
	}
	return nil, &NotLoadedError{edge: "metadata_authors"}
}

// MetadataDocumentTypesOrErr returns the MetadataDocumentTypes value or an error if the edge
// was not loaded in eager-loading.
func (e MetadataEdges) MetadataDocumentTypesOrErr() ([]*DocumentType, error) {
	if e.loadedTypes[2] {
		return e.MetadataDocumentTypes, nil
	}
	return nil, &NotLoadedError{edge: "metadata_document_types"}
}

// DocumentOrErr returns the Document value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MetadataEdges) DocumentOrErr() (*Document, error) {
	if e.loadedTypes[3] {
		if e.Document == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: document.Label}
		}
		return e.Document, nil
	}
	return nil, &NotLoadedError{edge: "document"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Metadata) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case metadata.FieldTools, metadata.FieldAuthors, metadata.FieldDocumentTypes:
			values[i] = new([]byte)
		case metadata.FieldID, metadata.FieldVersion, metadata.FieldName, metadata.FieldComment:
			values[i] = new(sql.NullString)
		case metadata.FieldDate:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Metadata fields.
func (m *Metadata) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case metadata.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				m.ID = value.String
			}
		case metadata.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				m.Version = value.String
			}
		case metadata.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		case metadata.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				m.Date = value.Time
			}
		case metadata.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				m.Comment = value.String
			}
		case metadata.FieldTools:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tools", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &m.Tools); err != nil {
					return fmt.Errorf("unmarshal field tools: %w", err)
				}
			}
		case metadata.FieldAuthors:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field authors", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &m.Authors); err != nil {
					return fmt.Errorf("unmarshal field authors: %w", err)
				}
			}
		case metadata.FieldDocumentTypes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field document_types", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &m.DocumentTypes); err != nil {
					return fmt.Errorf("unmarshal field document_types: %w", err)
				}
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Metadata.
// This includes values selected through modifiers, order, etc.
func (m *Metadata) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryMetadataTools queries the "metadata_tools" edge of the Metadata entity.
func (m *Metadata) QueryMetadataTools() *ToolQuery {
	return NewMetadataClient(m.config).QueryMetadataTools(m)
}

// QueryMetadataAuthors queries the "metadata_authors" edge of the Metadata entity.
func (m *Metadata) QueryMetadataAuthors() *PersonQuery {
	return NewMetadataClient(m.config).QueryMetadataAuthors(m)
}

// QueryMetadataDocumentTypes queries the "metadata_document_types" edge of the Metadata entity.
func (m *Metadata) QueryMetadataDocumentTypes() *DocumentTypeQuery {
	return NewMetadataClient(m.config).QueryMetadataDocumentTypes(m)
}

// QueryDocument queries the "document" edge of the Metadata entity.
func (m *Metadata) QueryDocument() *DocumentQuery {
	return NewMetadataClient(m.config).QueryDocument(m)
}

// Update returns a builder for updating this Metadata.
// Note that you need to call Metadata.Unwrap() before calling this method if this Metadata
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Metadata) Update() *MetadataUpdateOne {
	return NewMetadataClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Metadata entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Metadata) Unwrap() *Metadata {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Metadata is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Metadata) String() string {
	var builder strings.Builder
	builder.WriteString("Metadata(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("version=")
	builder.WriteString(m.Version)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteString(", ")
	builder.WriteString("date=")
	builder.WriteString(m.Date.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("comment=")
	builder.WriteString(m.Comment)
	builder.WriteString(", ")
	builder.WriteString("tools=")
	builder.WriteString(fmt.Sprintf("%v", m.Tools))
	builder.WriteString(", ")
	builder.WriteString("authors=")
	builder.WriteString(fmt.Sprintf("%v", m.Authors))
	builder.WriteString(", ")
	builder.WriteString("document_types=")
	builder.WriteString(fmt.Sprintf("%v", m.DocumentTypes))
	builder.WriteByte(')')
	return builder.String()
}

// MetadataSlice is a parsable slice of Metadata.
type MetadataSlice []*Metadata
