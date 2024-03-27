// Code generated by ent, DO NOT EDIT.

package timestamp

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/bom-squad/protobom/pkg/sbom/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Timestamp {
	return predicate.Timestamp(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Timestamp {
	return predicate.Timestamp(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Timestamp {
	return predicate.Timestamp(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Timestamp {
	return predicate.Timestamp(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Timestamp {
	return predicate.Timestamp(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Timestamp {
	return predicate.Timestamp(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Timestamp {
	return predicate.Timestamp(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Timestamp {
	return predicate.Timestamp(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Timestamp {
	return predicate.Timestamp(sql.FieldLTE(FieldID, id))
}

// HasDate applies the HasEdge predicate on the "date" edge.
func HasDate() predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, DateTable, DatePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDateWith applies the HasEdge predicate on the "date" edge with a given conditions (other predicates).
func HasDateWith(preds ...predicate.Timestamp) predicate.Timestamp {
	return predicate.Timestamp(func(s *sql.Selector) {
		step := newDateStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Timestamp) predicate.Timestamp {
	return predicate.Timestamp(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Timestamp) predicate.Timestamp {
	return predicate.Timestamp(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Timestamp) predicate.Timestamp {
	return predicate.Timestamp(sql.NotPredicates(p))
}
