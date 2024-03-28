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

package metadata

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the metadata type in the database.
	Label = "metadata"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// EdgeTools holds the string denoting the tools edge name in mutations.
	EdgeTools = "tools"
	// EdgeAuthors holds the string denoting the authors edge name in mutations.
	EdgeAuthors = "authors"
	// EdgeDocumentTypes holds the string denoting the document_types edge name in mutations.
	EdgeDocumentTypes = "document_types"
	// EdgeDocument holds the string denoting the document edge name in mutations.
	EdgeDocument = "document"
	// Table holds the table name of the metadata in the database.
	Table = "metadata"
	// ToolsTable is the table that holds the tools relation/edge.
	ToolsTable = "tools"
	// ToolsInverseTable is the table name for the Tool entity.
	// It exists in this package in order to avoid circular dependency with the "tool" package.
	ToolsInverseTable = "tools"
	// ToolsColumn is the table column denoting the tools relation/edge.
	ToolsColumn = "metadata_tools"
	// AuthorsTable is the table that holds the authors relation/edge.
	AuthorsTable = "persons"
	// AuthorsInverseTable is the table name for the Person entity.
	// It exists in this package in order to avoid circular dependency with the "person" package.
	AuthorsInverseTable = "persons"
	// AuthorsColumn is the table column denoting the authors relation/edge.
	AuthorsColumn = "metadata_authors"
	// DocumentTypesTable is the table that holds the document_types relation/edge.
	DocumentTypesTable = "document_types"
	// DocumentTypesInverseTable is the table name for the DocumentType entity.
	// It exists in this package in order to avoid circular dependency with the "documenttype" package.
	DocumentTypesInverseTable = "document_types"
	// DocumentTypesColumn is the table column denoting the document_types relation/edge.
	DocumentTypesColumn = "metadata_document_types"
	// DocumentTable is the table that holds the document relation/edge.
	DocumentTable = "documents"
	// DocumentInverseTable is the table name for the Document entity.
	// It exists in this package in order to avoid circular dependency with the "document" package.
	DocumentInverseTable = "documents"
	// DocumentColumn is the table column denoting the document relation/edge.
	DocumentColumn = "metadata_document"
)

// Columns holds all SQL columns for metadata fields.
var Columns = []string{
	FieldID,
	FieldVersion,
	FieldName,
	FieldDate,
	FieldComment,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Metadata queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByComment orders the results by the comment field.
func ByComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComment, opts...).ToFunc()
}

// ByToolsCount orders the results by tools count.
func ByToolsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newToolsStep(), opts...)
	}
}

// ByTools orders the results by tools terms.
func ByTools(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newToolsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAuthorsCount orders the results by authors count.
func ByAuthorsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAuthorsStep(), opts...)
	}
}

// ByAuthors orders the results by authors terms.
func ByAuthors(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuthorsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDocumentTypesCount orders the results by document_types count.
func ByDocumentTypesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDocumentTypesStep(), opts...)
	}
}

// ByDocumentTypes orders the results by document_types terms.
func ByDocumentTypes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDocumentTypesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDocumentField orders the results by document field.
func ByDocumentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDocumentStep(), sql.OrderByField(field, opts...))
	}
}
func newToolsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ToolsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ToolsTable, ToolsColumn),
	)
}
func newAuthorsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuthorsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AuthorsTable, AuthorsColumn),
	)
}
func newDocumentTypesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DocumentTypesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DocumentTypesTable, DocumentTypesColumn),
	)
}
func newDocumentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DocumentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, DocumentTable, DocumentColumn),
	)
}
