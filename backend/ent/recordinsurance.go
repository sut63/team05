// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team05/ent/amountpaid"
	"github.com/sut63/team05/ent/hospital"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/product"
	"github.com/sut63/team05/ent/recordinsurance"
)

// Recordinsurance is the model entity for the Recordinsurance schema.
type Recordinsurance struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// NumberOfDaysOfTreat holds the value of the "number_of_days_of_treat" field.
	NumberOfDaysOfTreat int `json:"number_of_days_of_treat,omitempty"`
	// RecordinsuranceContact holds the value of the "recordinsurance_contact" field.
	RecordinsuranceContact string `json:"recordinsurance_contact,omitempty"`
	// RecordinsuranceAddress holds the value of the "recordinsurance_address" field.
	RecordinsuranceAddress string `json:"recordinsurance_address,omitempty"`
	// RecordinsuranceTime holds the value of the "recordinsurance_time" field.
	RecordinsuranceTime time.Time `json:"recordinsurance_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RecordinsuranceQuery when eager-loading is set.
	Edges         RecordinsuranceEdges `json:"edges"`
	amountpaid_id *int
	hospital_id   *int
	member_id     *int
	officer_id    *int
	product_id    *int
}

// RecordinsuranceEdges holds the relations/edges for other nodes in the graph.
type RecordinsuranceEdges struct {
	// Member holds the value of the Member edge.
	Member *Member
	// Hospital holds the value of the Hospital edge.
	Hospital *Hospital
	// Officer holds the value of the Officer edge.
	Officer *Officer
	// Product holds the value of the Product edge.
	Product *Product
	// Amountpaid holds the value of the Amountpaid edge.
	Amountpaid *Amountpaid
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// MemberOrErr returns the Member value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RecordinsuranceEdges) MemberOrErr() (*Member, error) {
	if e.loadedTypes[0] {
		if e.Member == nil {
			// The edge Member was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: member.Label}
		}
		return e.Member, nil
	}
	return nil, &NotLoadedError{edge: "Member"}
}

// HospitalOrErr returns the Hospital value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RecordinsuranceEdges) HospitalOrErr() (*Hospital, error) {
	if e.loadedTypes[1] {
		if e.Hospital == nil {
			// The edge Hospital was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: hospital.Label}
		}
		return e.Hospital, nil
	}
	return nil, &NotLoadedError{edge: "Hospital"}
}

// OfficerOrErr returns the Officer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RecordinsuranceEdges) OfficerOrErr() (*Officer, error) {
	if e.loadedTypes[2] {
		if e.Officer == nil {
			// The edge Officer was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: officer.Label}
		}
		return e.Officer, nil
	}
	return nil, &NotLoadedError{edge: "Officer"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RecordinsuranceEdges) ProductOrErr() (*Product, error) {
	if e.loadedTypes[3] {
		if e.Product == nil {
			// The edge Product was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: product.Label}
		}
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "Product"}
}

// AmountpaidOrErr returns the Amountpaid value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RecordinsuranceEdges) AmountpaidOrErr() (*Amountpaid, error) {
	if e.loadedTypes[4] {
		if e.Amountpaid == nil {
			// The edge Amountpaid was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: amountpaid.Label}
		}
		return e.Amountpaid, nil
	}
	return nil, &NotLoadedError{edge: "Amountpaid"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Recordinsurance) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // number_of_days_of_treat
		&sql.NullString{}, // recordinsurance_contact
		&sql.NullString{}, // recordinsurance_address
		&sql.NullTime{},   // recordinsurance_time
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Recordinsurance) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // amountpaid_id
		&sql.NullInt64{}, // hospital_id
		&sql.NullInt64{}, // member_id
		&sql.NullInt64{}, // officer_id
		&sql.NullInt64{}, // product_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Recordinsurance fields.
func (r *Recordinsurance) assignValues(values ...interface{}) error {
	if m, n := len(values), len(recordinsurance.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	r.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field number_of_days_of_treat", values[0])
	} else if value.Valid {
		r.NumberOfDaysOfTreat = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field recordinsurance_contact", values[1])
	} else if value.Valid {
		r.RecordinsuranceContact = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field recordinsurance_address", values[2])
	} else if value.Valid {
		r.RecordinsuranceAddress = value.String
	}
	if value, ok := values[3].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field recordinsurance_time", values[3])
	} else if value.Valid {
		r.RecordinsuranceTime = value.Time
	}
	values = values[4:]
	if len(values) == len(recordinsurance.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field amountpaid_id", value)
		} else if value.Valid {
			r.amountpaid_id = new(int)
			*r.amountpaid_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field hospital_id", value)
		} else if value.Valid {
			r.hospital_id = new(int)
			*r.hospital_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field member_id", value)
		} else if value.Valid {
			r.member_id = new(int)
			*r.member_id = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field officer_id", value)
		} else if value.Valid {
			r.officer_id = new(int)
			*r.officer_id = int(value.Int64)
		}
		if value, ok := values[4].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field product_id", value)
		} else if value.Valid {
			r.product_id = new(int)
			*r.product_id = int(value.Int64)
		}
	}
	return nil
}

// QueryMember queries the Member edge of the Recordinsurance.
func (r *Recordinsurance) QueryMember() *MemberQuery {
	return (&RecordinsuranceClient{config: r.config}).QueryMember(r)
}

// QueryHospital queries the Hospital edge of the Recordinsurance.
func (r *Recordinsurance) QueryHospital() *HospitalQuery {
	return (&RecordinsuranceClient{config: r.config}).QueryHospital(r)
}

// QueryOfficer queries the Officer edge of the Recordinsurance.
func (r *Recordinsurance) QueryOfficer() *OfficerQuery {
	return (&RecordinsuranceClient{config: r.config}).QueryOfficer(r)
}

// QueryProduct queries the Product edge of the Recordinsurance.
func (r *Recordinsurance) QueryProduct() *ProductQuery {
	return (&RecordinsuranceClient{config: r.config}).QueryProduct(r)
}

// QueryAmountpaid queries the Amountpaid edge of the Recordinsurance.
func (r *Recordinsurance) QueryAmountpaid() *AmountpaidQuery {
	return (&RecordinsuranceClient{config: r.config}).QueryAmountpaid(r)
}

// Update returns a builder for updating this Recordinsurance.
// Note that, you need to call Recordinsurance.Unwrap() before calling this method, if this Recordinsurance
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Recordinsurance) Update() *RecordinsuranceUpdateOne {
	return (&RecordinsuranceClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (r *Recordinsurance) Unwrap() *Recordinsurance {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Recordinsurance is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Recordinsurance) String() string {
	var builder strings.Builder
	builder.WriteString("Recordinsurance(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", number_of_days_of_treat=")
	builder.WriteString(fmt.Sprintf("%v", r.NumberOfDaysOfTreat))
	builder.WriteString(", recordinsurance_contact=")
	builder.WriteString(r.RecordinsuranceContact)
	builder.WriteString(", recordinsurance_address=")
	builder.WriteString(r.RecordinsuranceAddress)
	builder.WriteString(", recordinsurance_time=")
	builder.WriteString(r.RecordinsuranceTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Recordinsurances is a parsable slice of Recordinsurance.
type Recordinsurances []*Recordinsurance

func (r Recordinsurances) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
