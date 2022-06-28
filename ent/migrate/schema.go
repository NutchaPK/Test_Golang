// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// GendersColumns holds the columns for the "genders" table.
	GendersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gender", Type: field.TypeString},
	}
	// GendersTable holds the schema information for the "genders" table.
	GendersTable = &schema.Table{
		Name:       "genders",
		Columns:    GendersColumns,
		PrimaryKey: []*schema.Column{GendersColumns[0]},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeString},
		{Name: "shipping_address", Type: field.TypeString},
		{Name: "paid_date", Type: field.TypeTime},
		{Name: "product_id", Type: field.TypeInt, Nullable: true},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "orders_products_orders",
				Columns:    []*schema.Column{OrdersColumns[4]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "price", Type: field.TypeInt},
		{Name: "gender_id", Type: field.TypeInt, Nullable: true},
		{Name: "size_id", Type: field.TypeInt, Nullable: true},
		{Name: "style_id", Type: field.TypeInt, Nullable: true},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "products_genders_products",
				Columns:    []*schema.Column{ProductsColumns[2]},
				RefColumns: []*schema.Column{GendersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "products_sizes_products",
				Columns:    []*schema.Column{ProductsColumns[3]},
				RefColumns: []*schema.Column{SizesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "products_styles_products",
				Columns:    []*schema.Column{ProductsColumns[4]},
				RefColumns: []*schema.Column{StylesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SizesColumns holds the columns for the "sizes" table.
	SizesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "size", Type: field.TypeString},
	}
	// SizesTable holds the schema information for the "sizes" table.
	SizesTable = &schema.Table{
		Name:       "sizes",
		Columns:    SizesColumns,
		PrimaryKey: []*schema.Column{SizesColumns[0]},
	}
	// StylesColumns holds the columns for the "styles" table.
	StylesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "style", Type: field.TypeString},
	}
	// StylesTable holds the schema information for the "styles" table.
	StylesTable = &schema.Table{
		Name:       "styles",
		Columns:    StylesColumns,
		PrimaryKey: []*schema.Column{StylesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		GendersTable,
		OrdersTable,
		ProductsTable,
		SizesTable,
		StylesTable,
	}
)

func init() {
	OrdersTable.ForeignKeys[0].RefTable = ProductsTable
	ProductsTable.ForeignKeys[0].RefTable = GendersTable
	ProductsTable.ForeignKeys[1].RefTable = SizesTable
	ProductsTable.ForeignKeys[2].RefTable = StylesTable
}