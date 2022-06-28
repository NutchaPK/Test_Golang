// Code generated by entc, DO NOT EDIT.

package style

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/NutchaPK/test/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Style applies equality check predicate on the "Style" field. It's identical to StyleEQ.
func Style(v string) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStyle), v))
	})
}

// StyleEQ applies the EQ predicate on the "Style" field.
func StyleEQ(v string) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStyle), v))
	})
}

// StyleNEQ applies the NEQ predicate on the "Style" field.
func StyleNEQ(v string) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStyle), v))
	})
}

// StyleIn applies the In predicate on the "Style" field.
func StyleIn(vs ...string) predicate.Style {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Style(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStyle), v...))
	})
}

// StyleNotIn applies the NotIn predicate on the "Style" field.
func StyleNotIn(vs ...string) predicate.Style {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Style(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStyle), v...))
	})
}

// StyleGT applies the GT predicate on the "Style" field.
func StyleGT(v string) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStyle), v))
	})
}

// StyleGTE applies the GTE predicate on the "Style" field.
func StyleGTE(v string) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStyle), v))
	})
}

// StyleLT applies the LT predicate on the "Style" field.
func StyleLT(v string) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStyle), v))
	})
}

// StyleLTE applies the LTE predicate on the "Style" field.
func StyleLTE(v string) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStyle), v))
	})
}

// StyleContains applies the Contains predicate on the "Style" field.
func StyleContains(v string) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStyle), v))
	})
}

// StyleHasPrefix applies the HasPrefix predicate on the "Style" field.
func StyleHasPrefix(v string) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStyle), v))
	})
}

// StyleHasSuffix applies the HasSuffix predicate on the "Style" field.
func StyleHasSuffix(v string) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStyle), v))
	})
}

// StyleEqualFold applies the EqualFold predicate on the "Style" field.
func StyleEqualFold(v string) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStyle), v))
	})
}

// StyleContainsFold applies the ContainsFold predicate on the "Style" field.
func StyleContainsFold(v string) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStyle), v))
	})
}

// HasProducts applies the HasEdge predicate on the "products" edge.
func HasProducts() predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductsTable, ProductsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductsWith applies the HasEdge predicate on the "products" edge with a given conditions (other predicates).
func HasProductsWith(preds ...predicate.Product) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductsTable, ProductsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Style) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Style) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Style) predicate.Style {
	return predicate.Style(func(s *sql.Selector) {
		p(s.Not())
	})
}
