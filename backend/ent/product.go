// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team05/ent/gender"
	"github.com/sut63/team05/ent/groupofage"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/product"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ProductName holds the value of the "product_name" field.
	ProductName string `json:"product_name,omitempty"`
	// ProductPrice holds the value of the "product_price" field.
	ProductPrice int `json:"product_price,omitempty"`
	// ProductTime holds the value of the "product_time" field.
	ProductTime int `json:"product_time,omitempty"`
	// ProductPaymentOfYear holds the value of the "product_payment_of_year" field.
	ProductPaymentOfYear float64 `json:"product_payment_of_year,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges           ProductEdges `json:"edges"`
	gender_id       *int
	group_of_age_id *int
	officer_id      *int
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
	// ProductGender holds the value of the product_gender edge.
	ProductGender *Gender
	// ProductGroupage holds the value of the product_groupage edge.
	ProductGroupage *GroupOfAge
	// ProductOfficer holds the value of the product_officer edge.
	ProductOfficer *Officer
	// ProductInsurance holds the value of the product_insurance edge.
	ProductInsurance []*Insurance
	// ProductInquiry holds the value of the product_inquiry edge.
	ProductInquiry []*Inquiry
<<<<<<< HEAD
	// ProductPayback holds the value of the product_payback edge.
	ProductPayback []*Payback
=======
	// ProductRecordinsurance holds the value of the product_recordinsurance edge.
	ProductRecordinsurance []*Recordinsurance
>>>>>>> 4637a9d (ทำ Entity สำหรับเก็บข้อมูลสิทธิประกันสุขภาพ - fix #53)
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// ProductGenderOrErr returns the ProductGender value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) ProductGenderOrErr() (*Gender, error) {
	if e.loadedTypes[0] {
		if e.ProductGender == nil {
			// The edge product_gender was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: gender.Label}
		}
		return e.ProductGender, nil
	}
	return nil, &NotLoadedError{edge: "product_gender"}
}

// ProductGroupageOrErr returns the ProductGroupage value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) ProductGroupageOrErr() (*GroupOfAge, error) {
	if e.loadedTypes[1] {
		if e.ProductGroupage == nil {
			// The edge product_groupage was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: groupofage.Label}
		}
		return e.ProductGroupage, nil
	}
	return nil, &NotLoadedError{edge: "product_groupage"}
}

// ProductOfficerOrErr returns the ProductOfficer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductEdges) ProductOfficerOrErr() (*Officer, error) {
	if e.loadedTypes[2] {
		if e.ProductOfficer == nil {
			// The edge product_officer was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: officer.Label}
		}
		return e.ProductOfficer, nil
	}
	return nil, &NotLoadedError{edge: "product_officer"}
}

// ProductInsuranceOrErr returns the ProductInsurance value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ProductInsuranceOrErr() ([]*Insurance, error) {
	if e.loadedTypes[3] {
		return e.ProductInsurance, nil
	}
	return nil, &NotLoadedError{edge: "product_insurance"}
}

// ProductInquiryOrErr returns the ProductInquiry value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ProductInquiryOrErr() ([]*Inquiry, error) {
	if e.loadedTypes[4] {
		return e.ProductInquiry, nil
	}
	return nil, &NotLoadedError{edge: "product_inquiry"}
}

<<<<<<< HEAD
// ProductPaybackOrErr returns the ProductPayback value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ProductPaybackOrErr() ([]*Payback, error) {
	if e.loadedTypes[5] {
		return e.ProductPayback, nil
	}
	return nil, &NotLoadedError{edge: "product_payback"}
=======
// ProductRecordinsuranceOrErr returns the ProductRecordinsurance value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ProductRecordinsuranceOrErr() ([]*Recordinsurance, error) {
	if e.loadedTypes[5] {
		return e.ProductRecordinsurance, nil
	}
	return nil, &NotLoadedError{edge: "product_recordinsurance"}
>>>>>>> 4637a9d (ทำ Entity สำหรับเก็บข้อมูลสิทธิประกันสุขภาพ - fix #53)
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},   // id
		&sql.NullString{},  // product_name
		&sql.NullInt64{},   // product_price
		&sql.NullInt64{},   // product_time
		&sql.NullFloat64{}, // product_payment_of_year
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Product) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // gender_id
		&sql.NullInt64{}, // group_of_age_id
		&sql.NullInt64{}, // officer_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(values ...interface{}) error {
	if m, n := len(values), len(product.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field product_name", values[0])
	} else if value.Valid {
		pr.ProductName = value.String
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field product_price", values[1])
	} else if value.Valid {
		pr.ProductPrice = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field product_time", values[2])
	} else if value.Valid {
		pr.ProductTime = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field product_payment_of_year", values[3])
	} else if value.Valid {
		pr.ProductPaymentOfYear = value.Float64
	}
	values = values[4:]
	if len(values) == len(product.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field gender_id", value)
		} else if value.Valid {
			pr.gender_id = new(int)
			*pr.gender_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field group_of_age_id", value)
		} else if value.Valid {
			pr.group_of_age_id = new(int)
			*pr.group_of_age_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field officer_id", value)
		} else if value.Valid {
			pr.officer_id = new(int)
			*pr.officer_id = int(value.Int64)
		}
	}
	return nil
}

// QueryProductGender queries the product_gender edge of the Product.
func (pr *Product) QueryProductGender() *GenderQuery {
	return (&ProductClient{config: pr.config}).QueryProductGender(pr)
}

// QueryProductGroupage queries the product_groupage edge of the Product.
func (pr *Product) QueryProductGroupage() *GroupOfAgeQuery {
	return (&ProductClient{config: pr.config}).QueryProductGroupage(pr)
}

// QueryProductOfficer queries the product_officer edge of the Product.
func (pr *Product) QueryProductOfficer() *OfficerQuery {
	return (&ProductClient{config: pr.config}).QueryProductOfficer(pr)
}

// QueryProductInsurance queries the product_insurance edge of the Product.
func (pr *Product) QueryProductInsurance() *InsuranceQuery {
	return (&ProductClient{config: pr.config}).QueryProductInsurance(pr)
}

// QueryProductInquiry queries the product_inquiry edge of the Product.
func (pr *Product) QueryProductInquiry() *InquiryQuery {
	return (&ProductClient{config: pr.config}).QueryProductInquiry(pr)
}

<<<<<<< HEAD
// QueryProductPayback queries the product_payback edge of the Product.
func (pr *Product) QueryProductPayback() *PaybackQuery {
	return (&ProductClient{config: pr.config}).QueryProductPayback(pr)
=======
// QueryProductRecordinsurance queries the product_recordinsurance edge of the Product.
func (pr *Product) QueryProductRecordinsurance() *RecordinsuranceQuery {
	return (&ProductClient{config: pr.config}).QueryProductRecordinsurance(pr)
>>>>>>> 4637a9d (ทำ Entity สำหรับเก็บข้อมูลสิทธิประกันสุขภาพ - fix #53)
}

// Update returns a builder for updating this Product.
// Note that, you need to call Product.Unwrap() before calling this method, if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return (&ProductClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", product_name=")
	builder.WriteString(pr.ProductName)
	builder.WriteString(", product_price=")
	builder.WriteString(fmt.Sprintf("%v", pr.ProductPrice))
	builder.WriteString(", product_time=")
	builder.WriteString(fmt.Sprintf("%v", pr.ProductTime))
	builder.WriteString(", product_payment_of_year=")
	builder.WriteString(fmt.Sprintf("%v", pr.ProductPaymentOfYear))
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product

func (pr Products) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
