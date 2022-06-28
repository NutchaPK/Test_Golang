package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Size holds the schema definition for the Size entity.
type Size struct {
	ent.Schema
}


// Fields of the Size.
func (Size) Fields() []ent.Field {
	return []ent.Field{
		field.String("Size"),
	}
}

// Edges of the Size.
func (Size) Edges() []ent.Edge {
	return []ent.Edge{
			
		edge.To("products" , Product.Type).StorageKey(edge.Column("size_id")),
	}
}

