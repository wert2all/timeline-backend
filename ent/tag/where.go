// Code generated by ent, DO NOT EDIT.

package tag

import (
	"timeline/backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Tag {
	return predicate.Tag(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Tag {
	return predicate.Tag(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Tag {
	return predicate.Tag(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Tag {
	return predicate.Tag(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Tag {
	return predicate.Tag(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Tag {
	return predicate.Tag(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Tag {
	return predicate.Tag(sql.FieldLTE(FieldID, id))
}

// Tag applies equality check predicate on the "tag" field. It's identical to TagEQ.
func Tag(v string) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldTag, v))
}

// TagEQ applies the EQ predicate on the "tag" field.
func TagEQ(v string) predicate.Tag {
	return predicate.Tag(sql.FieldEQ(FieldTag, v))
}

// TagNEQ applies the NEQ predicate on the "tag" field.
func TagNEQ(v string) predicate.Tag {
	return predicate.Tag(sql.FieldNEQ(FieldTag, v))
}

// TagIn applies the In predicate on the "tag" field.
func TagIn(vs ...string) predicate.Tag {
	return predicate.Tag(sql.FieldIn(FieldTag, vs...))
}

// TagNotIn applies the NotIn predicate on the "tag" field.
func TagNotIn(vs ...string) predicate.Tag {
	return predicate.Tag(sql.FieldNotIn(FieldTag, vs...))
}

// TagGT applies the GT predicate on the "tag" field.
func TagGT(v string) predicate.Tag {
	return predicate.Tag(sql.FieldGT(FieldTag, v))
}

// TagGTE applies the GTE predicate on the "tag" field.
func TagGTE(v string) predicate.Tag {
	return predicate.Tag(sql.FieldGTE(FieldTag, v))
}

// TagLT applies the LT predicate on the "tag" field.
func TagLT(v string) predicate.Tag {
	return predicate.Tag(sql.FieldLT(FieldTag, v))
}

// TagLTE applies the LTE predicate on the "tag" field.
func TagLTE(v string) predicate.Tag {
	return predicate.Tag(sql.FieldLTE(FieldTag, v))
}

// TagContains applies the Contains predicate on the "tag" field.
func TagContains(v string) predicate.Tag {
	return predicate.Tag(sql.FieldContains(FieldTag, v))
}

// TagHasPrefix applies the HasPrefix predicate on the "tag" field.
func TagHasPrefix(v string) predicate.Tag {
	return predicate.Tag(sql.FieldHasPrefix(FieldTag, v))
}

// TagHasSuffix applies the HasSuffix predicate on the "tag" field.
func TagHasSuffix(v string) predicate.Tag {
	return predicate.Tag(sql.FieldHasSuffix(FieldTag, v))
}

// TagEqualFold applies the EqualFold predicate on the "tag" field.
func TagEqualFold(v string) predicate.Tag {
	return predicate.Tag(sql.FieldEqualFold(FieldTag, v))
}

// TagContainsFold applies the ContainsFold predicate on the "tag" field.
func TagContainsFold(v string) predicate.Tag {
	return predicate.Tag(sql.FieldContainsFold(FieldTag, v))
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, EventTable, EventPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.Tag {
	return predicate.Tag(func(s *sql.Selector) {
		step := newEventStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tag) predicate.Tag {
	return predicate.Tag(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tag) predicate.Tag {
	return predicate.Tag(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tag) predicate.Tag {
	return predicate.Tag(sql.NotPredicates(p))
}
