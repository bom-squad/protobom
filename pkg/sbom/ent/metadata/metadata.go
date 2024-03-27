// Code generated by ent, DO NOT EDIT.

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
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// EdgeTools holds the string denoting the tools edge name in mutations.
	EdgeTools = "tools"
	// EdgeAuthors holds the string denoting the authors edge name in mutations.
	EdgeAuthors = "authors"
	// EdgeDocumentTypes holds the string denoting the documenttypes edge name in mutations.
	EdgeDocumentTypes = "documentTypes"
	// EdgeDate holds the string denoting the date edge name in mutations.
	EdgeDate = "date"
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
	// DocumentTypesTable is the table that holds the documentTypes relation/edge.
	DocumentTypesTable = "document_types"
	// DocumentTypesInverseTable is the table name for the DocumentType entity.
	// It exists in this package in order to avoid circular dependency with the "documenttype" package.
	DocumentTypesInverseTable = "document_types"
	// DocumentTypesColumn is the table column denoting the documentTypes relation/edge.
	DocumentTypesColumn = "metadata_document_types"
	// DateTable is the table that holds the date relation/edge.
	DateTable = "timestamps"
	// DateInverseTable is the table name for the Timestamp entity.
	// It exists in this package in order to avoid circular dependency with the "timestamp" package.
	DateInverseTable = "timestamps"
	// DateColumn is the table column denoting the date relation/edge.
	DateColumn = "metadata_date"
)

// Columns holds all SQL columns for metadata fields.
var Columns = []string{
	FieldID,
	FieldVersion,
	FieldName,
	FieldComment,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "metadata"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"document_metadata",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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

// ByDocumentTypesCount orders the results by documentTypes count.
func ByDocumentTypesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDocumentTypesStep(), opts...)
	}
}

// ByDocumentTypes orders the results by documentTypes terms.
func ByDocumentTypes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDocumentTypesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDateCount orders the results by date count.
func ByDateCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDateStep(), opts...)
	}
}

// ByDate orders the results by date terms.
func ByDate(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDateStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newDateStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DateInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DateTable, DateColumn),
	)
}
