package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Style holds the schema definition for the Style entity.
type Style struct {
	ent.Schema
}

// Fields of the Style.
func (Style) Fields() []ent.Field {
	return []ent.Field{
		field.String("Style"),
	}
}

// Edges of the Style.
func (Style) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products" , Product.Type).StorageKey(edge.Column("style_id")),
	}
}