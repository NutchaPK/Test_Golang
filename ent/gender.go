// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NutchaPK/test/ent/gender"
)

// Gender is the model entity for the Gender schema.
type Gender struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Gender holds the value of the "Gender" field.
	Gender string `json:"Gender,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GenderQuery when eager-loading is set.
	Edges GenderEdges `json:"edges"`
}

// GenderEdges holds the relations/edges for other nodes in the graph.
type GenderEdges struct {
	// Products holds the value of the products edge.
	Products []*Product `json:"products,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e GenderEdges) ProductsOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Gender) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case gender.FieldID:
			values[i] = new(sql.NullInt64)
		case gender.FieldGender:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Gender", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Gender fields.
func (ge *Gender) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gender.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ge.ID = int(value.Int64)
		case gender.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Gender", values[i])
			} else if value.Valid {
				ge.Gender = value.String
			}
		}
	}
	return nil
}

// QueryProducts queries the "products" edge of the Gender entity.
func (ge *Gender) QueryProducts() *ProductQuery {
	return (&GenderClient{config: ge.config}).QueryProducts(ge)
}

// Update returns a builder for updating this Gender.
// Note that you need to call Gender.Unwrap() before calling this method if this Gender
// was returned from a transaction, and the transaction was committed or rolled back.
func (ge *Gender) Update() *GenderUpdateOne {
	return (&GenderClient{config: ge.config}).UpdateOne(ge)
}

// Unwrap unwraps the Gender entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ge *Gender) Unwrap() *Gender {
	tx, ok := ge.config.driver.(*txDriver)
	if !ok {
		panic("ent: Gender is not a transactional entity")
	}
	ge.config.driver = tx.drv
	return ge
}

// String implements the fmt.Stringer.
func (ge *Gender) String() string {
	var builder strings.Builder
	builder.WriteString("Gender(")
	builder.WriteString(fmt.Sprintf("id=%v", ge.ID))
	builder.WriteString(", Gender=")
	builder.WriteString(ge.Gender)
	builder.WriteByte(')')
	return builder.String()
}

// Genders is a parsable slice of Gender.
type Genders []*Gender

func (ge Genders) config(cfg config) {
	for _i := range ge {
		ge[_i].config = cfg
	}
}
