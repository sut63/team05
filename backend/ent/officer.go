// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/position"
)

// Officer is the model entity for the Officer schema.
type Officer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OfficerEmail holds the value of the "officer_email" field.
	OfficerEmail string `json:"officer_email,omitempty"`
	// OfficerName holds the value of the "officer_name" field.
	OfficerName string `json:"officer_name,omitempty"`
	// OfficerPassword holds the value of the "officer_password" field.
	OfficerPassword string `json:"officer_password,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OfficerQuery when eager-loading is set.
	Edges       OfficerEdges `json:"edges"`
	position_id *int
}

// OfficerEdges holds the relations/edges for other nodes in the graph.
type OfficerEdges struct {
	// OfficerProduct holds the value of the officer_product edge.
	OfficerProduct []*Product
	// OfficerInsurance holds the value of the officer_insurance edge.
	OfficerInsurance []*Insurance
	// OfficerInquiry holds the value of the officer_inquiry edge.
	OfficerInquiry []*Inquiry
	// OfficerPayback holds the value of the officer_payback edge.
	OfficerPayback []*Payback
	// OfficerRecordinsurance holds the value of the officer_recordinsurance edge.
	OfficerRecordinsurance []*Recordinsurance
	// Position holds the value of the position edge.
	Position *Position
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// OfficerProductOrErr returns the OfficerProduct value or an error if the edge
// was not loaded in eager-loading.
func (e OfficerEdges) OfficerProductOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.OfficerProduct, nil
	}
	return nil, &NotLoadedError{edge: "officer_product"}
}

// OfficerInsuranceOrErr returns the OfficerInsurance value or an error if the edge
// was not loaded in eager-loading.
func (e OfficerEdges) OfficerInsuranceOrErr() ([]*Insurance, error) {
	if e.loadedTypes[1] {
		return e.OfficerInsurance, nil
	}
	return nil, &NotLoadedError{edge: "officer_insurance"}
}

// OfficerInquiryOrErr returns the OfficerInquiry value or an error if the edge
// was not loaded in eager-loading.
func (e OfficerEdges) OfficerInquiryOrErr() ([]*Inquiry, error) {
	if e.loadedTypes[2] {
		return e.OfficerInquiry, nil
	}
	return nil, &NotLoadedError{edge: "officer_inquiry"}
}

// OfficerPaybackOrErr returns the OfficerPayback value or an error if the edge
// was not loaded in eager-loading.
func (e OfficerEdges) OfficerPaybackOrErr() ([]*Payback, error) {
	if e.loadedTypes[3] {
		return e.OfficerPayback, nil
	}
	return nil, &NotLoadedError{edge: "officer_payback"}
}

// OfficerRecordinsuranceOrErr returns the OfficerRecordinsurance value or an error if the edge
// was not loaded in eager-loading.
func (e OfficerEdges) OfficerRecordinsuranceOrErr() ([]*Recordinsurance, error) {
	if e.loadedTypes[4] {
		return e.OfficerRecordinsurance, nil
	}
	return nil, &NotLoadedError{edge: "officer_recordinsurance"}
}

// PositionOrErr returns the Position value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OfficerEdges) PositionOrErr() (*Position, error) {
	if e.loadedTypes[5] {
		if e.Position == nil {
			// The edge position was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: position.Label}
		}
		return e.Position, nil
	}
	return nil, &NotLoadedError{edge: "position"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Officer) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // officer_email
		&sql.NullString{}, // officer_name
		&sql.NullString{}, // officer_password
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Officer) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // position_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Officer fields.
func (o *Officer) assignValues(values ...interface{}) error {
	if m, n := len(values), len(officer.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	o.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field officer_email", values[0])
	} else if value.Valid {
		o.OfficerEmail = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field officer_name", values[1])
	} else if value.Valid {
		o.OfficerName = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field officer_password", values[2])
	} else if value.Valid {
		o.OfficerPassword = value.String
	}
	values = values[3:]
	if len(values) == len(officer.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field position_id", value)
		} else if value.Valid {
			o.position_id = new(int)
			*o.position_id = int(value.Int64)
		}
	}
	return nil
}

// QueryOfficerProduct queries the officer_product edge of the Officer.
func (o *Officer) QueryOfficerProduct() *ProductQuery {
	return (&OfficerClient{config: o.config}).QueryOfficerProduct(o)
}

// QueryOfficerInsurance queries the officer_insurance edge of the Officer.
func (o *Officer) QueryOfficerInsurance() *InsuranceQuery {
	return (&OfficerClient{config: o.config}).QueryOfficerInsurance(o)
}

// QueryOfficerInquiry queries the officer_inquiry edge of the Officer.
func (o *Officer) QueryOfficerInquiry() *InquiryQuery {
	return (&OfficerClient{config: o.config}).QueryOfficerInquiry(o)
}

// QueryOfficerPayback queries the officer_payback edge of the Officer.
func (o *Officer) QueryOfficerPayback() *PaybackQuery {
	return (&OfficerClient{config: o.config}).QueryOfficerPayback(o)
}

// QueryOfficerRecordinsurance queries the officer_recordinsurance edge of the Officer.
func (o *Officer) QueryOfficerRecordinsurance() *RecordinsuranceQuery {
	return (&OfficerClient{config: o.config}).QueryOfficerRecordinsurance(o)
}

// QueryPosition queries the position edge of the Officer.
func (o *Officer) QueryPosition() *PositionQuery {
	return (&OfficerClient{config: o.config}).QueryPosition(o)
}

// Update returns a builder for updating this Officer.
// Note that, you need to call Officer.Unwrap() before calling this method, if this Officer
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Officer) Update() *OfficerUpdateOne {
	return (&OfficerClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (o *Officer) Unwrap() *Officer {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Officer is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Officer) String() string {
	var builder strings.Builder
	builder.WriteString("Officer(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", officer_email=")
	builder.WriteString(o.OfficerEmail)
	builder.WriteString(", officer_name=")
	builder.WriteString(o.OfficerName)
	builder.WriteString(", officer_password=")
	builder.WriteString(o.OfficerPassword)
	builder.WriteByte(')')
	return builder.String()
}

// Officers is a parsable slice of Officer.
type Officers []*Officer

func (o Officers) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
