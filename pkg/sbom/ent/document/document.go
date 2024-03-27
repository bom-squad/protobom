// Code generated by ent, DO NOT EDIT.

package document

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the document type in the database.
	Label = "document"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeMetadata holds the string denoting the metadata edge name in mutations.
	EdgeMetadata = "metadata"
	// EdgeNodeList holds the string denoting the node_list edge name in mutations.
	EdgeNodeList = "node_list"
	// Table holds the table name of the document in the database.
	Table = "documents"
	// MetadataTable is the table that holds the metadata relation/edge.
	MetadataTable = "metadata"
	// MetadataInverseTable is the table name for the Metadata entity.
	// It exists in this package in order to avoid circular dependency with the "metadata" package.
	MetadataInverseTable = "metadata"
	// MetadataColumn is the table column denoting the metadata relation/edge.
	MetadataColumn = "document_metadata"
	// NodeListTable is the table that holds the node_list relation/edge.
	NodeListTable = "node_lists"
	// NodeListInverseTable is the table name for the NodeList entity.
	// It exists in this package in order to avoid circular dependency with the "nodelist" package.
	NodeListInverseTable = "node_lists"
	// NodeListColumn is the table column denoting the node_list relation/edge.
	NodeListColumn = "document_node_list"
)

// Columns holds all SQL columns for document fields.
var Columns = []string{
	FieldID,
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

// OrderOption defines the ordering options for the Document queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByMetadataCount orders the results by metadata count.
func ByMetadataCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMetadataStep(), opts...)
	}
}

// ByMetadata orders the results by metadata terms.
func ByMetadata(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMetadataStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByNodeListCount orders the results by node_list count.
func ByNodeListCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNodeListStep(), opts...)
	}
}

// ByNodeList orders the results by node_list terms.
func ByNodeList(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNodeListStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMetadataStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MetadataInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MetadataTable, MetadataColumn),
	)
}
func newNodeListStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NodeListInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, NodeListTable, NodeListColumn),
	)
}
