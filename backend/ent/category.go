// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team05/ent/category"
)

// Category is the model entity for the Category schema.
type Category struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CategoryName holds the value of the "category_name" field.
	CategoryName string `json:"category_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CategoryQuery when eager-loading is set.
	Edges CategoryEdges `json:"edges"`
}

// CategoryEdges holds the relations/edges for other nodes in the graph.
type CategoryEdges struct {
	// CategoryInquiry holds the value of the category_inquiry edge.
	CategoryInquiry []*Inquiry
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CategoryInquiryOrErr returns the CategoryInquiry value or an error if the edge
// was not loaded in eager-loading.
func (e CategoryEdges) CategoryInquiryOrErr() ([]*Inquiry, error) {
	if e.loadedTypes[0] {
		return e.CategoryInquiry, nil
	}
	return nil, &NotLoadedError{edge: "category_inquiry"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Category) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // category_name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Category fields.
func (c *Category) assignValues(values ...interface{}) error {
	if m, n := len(values), len(category.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field category_name", values[0])
	} else if value.Valid {
		c.CategoryName = value.String
	}
	return nil
}

// QueryCategoryInquiry queries the category_inquiry edge of the Category.
func (c *Category) QueryCategoryInquiry() *InquiryQuery {
	return (&CategoryClient{config: c.config}).QueryCategoryInquiry(c)
}

// Update returns a builder for updating this Category.
// Note that, you need to call Category.Unwrap() before calling this method, if this Category
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Category) Update() *CategoryUpdateOne {
	return (&CategoryClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Category) Unwrap() *Category {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Category is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Category) String() string {
	var builder strings.Builder
	builder.WriteString("Category(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", category_name=")
	builder.WriteString(c.CategoryName)
	builder.WriteByte(')')
	return builder.String()
}

// Categories is a parsable slice of Category.
type Categories []*Category

func (c Categories) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
