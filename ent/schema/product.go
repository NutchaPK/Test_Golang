package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Price"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("style" , Style.Type).Ref("products").Unique(),	
		edge.From("size" , Size.Type).Ref("products").Unique(),	
		edge.From("gender" , Gender.Type).Ref("products").Unique(),	
		edge.To("orders" , Order.Type).StorageKey(edge.Column("product_id")),

	}
}