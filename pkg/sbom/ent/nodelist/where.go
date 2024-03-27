// Code generated by ent, DO NOT EDIT.

package nodelist

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/bom-squad/protobom/pkg/sbom/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.NodeList {
	return predicate.NodeList(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.NodeList {
	return predicate.NodeList(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.NodeList {
	return predicate.NodeList(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.NodeList {
	return predicate.NodeList(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.NodeList {
	return predicate.NodeList(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.NodeList {
	return predicate.NodeList(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.NodeList {
	return predicate.NodeList(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.NodeList {
	return predicate.NodeList(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.NodeList {
	return predicate.NodeList(sql.FieldLTE(FieldID, id))
}

// RootElements applies equality check predicate on the "root_elements" field. It's identical to RootElementsEQ.
func RootElements(v string) predicate.NodeList {
	return predicate.NodeList(sql.FieldEQ(FieldRootElements, v))
}

// RootElementsEQ applies the EQ predicate on the "root_elements" field.
func RootElementsEQ(v string) predicate.NodeList {
	return predicate.NodeList(sql.FieldEQ(FieldRootElements, v))
}

// RootElementsNEQ applies the NEQ predicate on the "root_elements" field.
func RootElementsNEQ(v string) predicate.NodeList {
	return predicate.NodeList(sql.FieldNEQ(FieldRootElements, v))
}

// RootElementsIn applies the In predicate on the "root_elements" field.
func RootElementsIn(vs ...string) predicate.NodeList {
	return predicate.NodeList(sql.FieldIn(FieldRootElements, vs...))
}

// RootElementsNotIn applies the NotIn predicate on the "root_elements" field.
func RootElementsNotIn(vs ...string) predicate.NodeList {
	return predicate.NodeList(sql.FieldNotIn(FieldRootElements, vs...))
}

// RootElementsGT applies the GT predicate on the "root_elements" field.
func RootElementsGT(v string) predicate.NodeList {
	return predicate.NodeList(sql.FieldGT(FieldRootElements, v))
}

// RootElementsGTE applies the GTE predicate on the "root_elements" field.
func RootElementsGTE(v string) predicate.NodeList {
	return predicate.NodeList(sql.FieldGTE(FieldRootElements, v))
}

// RootElementsLT applies the LT predicate on the "root_elements" field.
func RootElementsLT(v string) predicate.NodeList {
	return predicate.NodeList(sql.FieldLT(FieldRootElements, v))
}

// RootElementsLTE applies the LTE predicate on the "root_elements" field.
func RootElementsLTE(v string) predicate.NodeList {
	return predicate.NodeList(sql.FieldLTE(FieldRootElements, v))
}

// RootElementsContains applies the Contains predicate on the "root_elements" field.
func RootElementsContains(v string) predicate.NodeList {
	return predicate.NodeList(sql.FieldContains(FieldRootElements, v))
}

// RootElementsHasPrefix applies the HasPrefix predicate on the "root_elements" field.
func RootElementsHasPrefix(v string) predicate.NodeList {
	return predicate.NodeList(sql.FieldHasPrefix(FieldRootElements, v))
}

// RootElementsHasSuffix applies the HasSuffix predicate on the "root_elements" field.
func RootElementsHasSuffix(v string) predicate.NodeList {
	return predicate.NodeList(sql.FieldHasSuffix(FieldRootElements, v))
}

// RootElementsEqualFold applies the EqualFold predicate on the "root_elements" field.
func RootElementsEqualFold(v string) predicate.NodeList {
	return predicate.NodeList(sql.FieldEqualFold(FieldRootElements, v))
}

// RootElementsContainsFold applies the ContainsFold predicate on the "root_elements" field.
func RootElementsContainsFold(v string) predicate.NodeList {
	return predicate.NodeList(sql.FieldContainsFold(FieldRootElements, v))
}

// HasNodes applies the HasEdge predicate on the "nodes" edge.
func HasNodes() predicate.NodeList {
	return predicate.NodeList(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NodesTable, NodesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodesWith applies the HasEdge predicate on the "nodes" edge with a given conditions (other predicates).
func HasNodesWith(preds ...predicate.Node) predicate.NodeList {
	return predicate.NodeList(func(s *sql.Selector) {
		step := newNodesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEdges applies the HasEdge predicate on the "edges" edge.
func HasEdges() predicate.NodeList {
	return predicate.NodeList(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EdgesTable, EdgesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEdgesWith applies the HasEdge predicate on the "edges" edge with a given conditions (other predicates).
func HasEdgesWith(preds ...predicate.Edge) predicate.NodeList {
	return predicate.NodeList(func(s *sql.Selector) {
		step := newEdgesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NodeList) predicate.NodeList {
	return predicate.NodeList(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NodeList) predicate.NodeList {
	return predicate.NodeList(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NodeList) predicate.NodeList {
	return predicate.NodeList(sql.NotPredicates(p))
}
