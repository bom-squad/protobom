// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/bom-squad/protobom/pkg/sbom/ent/edge"
)

// Edge is the model entity for the Edge schema.
type Edge struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type edge.Type `json:"type,omitempty"`
	// From holds the value of the "from" field.
	From string `json:"from,omitempty"`
	// To holds the value of the "to" field.
	To              string `json:"to,omitempty"`
	node_list_edges *int
	selectValues    sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Edge) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case edge.FieldID:
			values[i] = new(sql.NullInt64)
		case edge.FieldType, edge.FieldFrom, edge.FieldTo:
			values[i] = new(sql.NullString)
		case edge.ForeignKeys[0]: // node_list_edges
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Edge fields.
func (e *Edge) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case edge.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case edge.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				e.Type = edge.Type(value.String)
			}
		case edge.FieldFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				e.From = value.String
			}
		case edge.FieldTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field to", values[i])
			} else if value.Valid {
				e.To = value.String
			}
		case edge.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field node_list_edges", value)
			} else if value.Valid {
				e.node_list_edges = new(int)
				*e.node_list_edges = int(value.Int64)
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Edge.
// This includes values selected through modifiers, order, etc.
func (e *Edge) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// Update returns a builder for updating this Edge.
// Note that you need to call Edge.Unwrap() before calling this method if this Edge
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Edge) Update() *EdgeUpdateOne {
	return NewEdgeClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Edge entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Edge) Unwrap() *Edge {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Edge is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Edge) String() string {
	var builder strings.Builder
	builder.WriteString("Edge(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", e.Type))
	builder.WriteString(", ")
	builder.WriteString("from=")
	builder.WriteString(e.From)
	builder.WriteString(", ")
	builder.WriteString("to=")
	builder.WriteString(e.To)
	builder.WriteByte(')')
	return builder.String()
}

// Edges is a parsable slice of Edge.
type Edges []*Edge
