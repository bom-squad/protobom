// Code generated by ent, DO NOT EDIT.

package timestamp

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the timestamp type in the database.
	Label = "timestamp"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeDate holds the string denoting the date edge name in mutations.
	EdgeDate = "date"
	// Table holds the table name of the timestamp in the database.
	Table = "timestamps"
	// DateTable is the table that holds the date relation/edge. The primary key declared below.
	DateTable = "timestamp_date"
)

// Columns holds all SQL columns for timestamp fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "timestamps"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"metadata_date",
	"node_release_date",
	"node_build_date",
	"node_valid_until_date",
}

var (
	// DatePrimaryKey and DateColumn2 are the table columns denoting the
	// primary key for the date relation (M2M).
	DatePrimaryKey = []string{"timestamp_id", "date_id"}
)

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

// OrderOption defines the ordering options for the Timestamp queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
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
func newDateStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, DateTable, DatePrimaryKey...),
	)
}
