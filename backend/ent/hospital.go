// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team05/ent/hospital"
)

// Hospital is the model entity for the Hospital schema.
type Hospital struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// HospitalName holds the value of the "hospital_name" field.
	HospitalName string `json:"hospital_name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HospitalQuery when eager-loading is set.
	Edges HospitalEdges `json:"edges"`
}

// HospitalEdges holds the relations/edges for other nodes in the graph.
type HospitalEdges struct {
	// HospitalInsurance holds the value of the hospital_insurance edge.
	HospitalInsurance []*Insurance
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// HospitalInsuranceOrErr returns the HospitalInsurance value or an error if the edge
// was not loaded in eager-loading.
func (e HospitalEdges) HospitalInsuranceOrErr() ([]*Insurance, error) {
	if e.loadedTypes[0] {
		return e.HospitalInsurance, nil
	}
	return nil, &NotLoadedError{edge: "hospital_insurance"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Hospital) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // hospital_name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Hospital fields.
func (h *Hospital) assignValues(values ...interface{}) error {
	if m, n := len(values), len(hospital.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	h.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field hospital_name", values[0])
	} else if value.Valid {
		h.HospitalName = value.String
	}
	return nil
}

// QueryHospitalInsurance queries the hospital_insurance edge of the Hospital.
func (h *Hospital) QueryHospitalInsurance() *InsuranceQuery {
	return (&HospitalClient{config: h.config}).QueryHospitalInsurance(h)
}

// Update returns a builder for updating this Hospital.
// Note that, you need to call Hospital.Unwrap() before calling this method, if this Hospital
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Hospital) Update() *HospitalUpdateOne {
	return (&HospitalClient{config: h.config}).UpdateOne(h)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (h *Hospital) Unwrap() *Hospital {
	tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Hospital is not a transactional entity")
	}
	h.config.driver = tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Hospital) String() string {
	var builder strings.Builder
	builder.WriteString("Hospital(")
	builder.WriteString(fmt.Sprintf("id=%v", h.ID))
	builder.WriteString(", hospital_name=")
	builder.WriteString(h.HospitalName)
	builder.WriteByte(')')
	return builder.String()
}

// Hospitals is a parsable slice of Hospital.
type Hospitals []*Hospital

func (h Hospitals) config(cfg config) {
	for _i := range h {
		h[_i].config = cfg
	}
}
