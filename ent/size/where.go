// Code generated by entc, DO NOT EDIT.

package size

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/NutchaPK/test/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Size applies equality check predicate on the "Size" field. It's identical to SizeEQ.
func Size(v string) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSize), v))
	})
}

// SizeEQ applies the EQ predicate on the "Size" field.
func SizeEQ(v string) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSize), v))
	})
}

// SizeNEQ applies the NEQ predicate on the "Size" field.
func SizeNEQ(v string) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSize), v))
	})
}

// SizeIn applies the In predicate on the "Size" field.
func SizeIn(vs ...string) predicate.Size {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Size(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSize), v...))
	})
}

// SizeNotIn applies the NotIn predicate on the "Size" field.
func SizeNotIn(vs ...string) predicate.Size {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Size(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSize), v...))
	})
}

// SizeGT applies the GT predicate on the "Size" field.
func SizeGT(v string) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSize), v))
	})
}

// SizeGTE applies the GTE predicate on the "Size" field.
func SizeGTE(v string) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSize), v))
	})
}

// SizeLT applies the LT predicate on the "Size" field.
func SizeLT(v string) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSize), v))
	})
}

// SizeLTE applies the LTE predicate on the "Size" field.
func SizeLTE(v string) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSize), v))
	})
}

// SizeContains applies the Contains predicate on the "Size" field.
func SizeContains(v string) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSize), v))
	})
}

// SizeHasPrefix applies the HasPrefix predicate on the "Size" field.
func SizeHasPrefix(v string) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSize), v))
	})
}

// SizeHasSuffix applies the HasSuffix predicate on the "Size" field.
func SizeHasSuffix(v string) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSize), v))
	})
}

// SizeEqualFold applies the EqualFold predicate on the "Size" field.
func SizeEqualFold(v string) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSize), v))
	})
}

// SizeContainsFold applies the ContainsFold predicate on the "Size" field.
func SizeContainsFold(v string) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSize), v))
	})
}

// HasProducts applies the HasEdge predicate on the "products" edge.
func HasProducts() predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductsTable, ProductsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductsWith applies the HasEdge predicate on the "products" edge with a given conditions (other predicates).
func HasProductsWith(preds ...predicate.Product) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
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
func And(predicates ...predicate.Size) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Size) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
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
func Not(p predicate.Size) predicate.Size {
	return predicate.Size(func(s *sql.Selector) {
		p(s.Not())
	})
}
