// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team05/ent/moneytransfer"
)

// Moneytransfer is the model entity for the Moneytransfer schema.
type Moneytransfer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// MoneytransferType holds the value of the "moneytransfer_type" field.
	MoneytransferType string `json:"moneytransfer_type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MoneytransferQuery when eager-loading is set.
	Edges MoneytransferEdges `json:"edges"`
}

// MoneytransferEdges holds the relations/edges for other nodes in the graph.
type MoneytransferEdges struct {
	// MoneytransferPayment holds the value of the moneytransfer_payment edge.
	MoneytransferPayment []*Payment
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MoneytransferPaymentOrErr returns the MoneytransferPayment value or an error if the edge
// was not loaded in eager-loading.
func (e MoneytransferEdges) MoneytransferPaymentOrErr() ([]*Payment, error) {
	if e.loadedTypes[0] {
		return e.MoneytransferPayment, nil
	}
	return nil, &NotLoadedError{edge: "moneytransfer_payment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Moneytransfer) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // moneytransfer_type
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Moneytransfer fields.
func (m *Moneytransfer) assignValues(values ...interface{}) error {
	if m, n := len(values), len(moneytransfer.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field moneytransfer_type", values[0])
	} else if value.Valid {
		m.MoneytransferType = value.String
	}
	return nil
}

// QueryMoneytransferPayment queries the moneytransfer_payment edge of the Moneytransfer.
func (m *Moneytransfer) QueryMoneytransferPayment() *PaymentQuery {
	return (&MoneytransferClient{config: m.config}).QueryMoneytransferPayment(m)
}

// Update returns a builder for updating this Moneytransfer.
// Note that, you need to call Moneytransfer.Unwrap() before calling this method, if this Moneytransfer
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Moneytransfer) Update() *MoneytransferUpdateOne {
	return (&MoneytransferClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Moneytransfer) Unwrap() *Moneytransfer {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Moneytransfer is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Moneytransfer) String() string {
	var builder strings.Builder
	builder.WriteString("Moneytransfer(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", moneytransfer_type=")
	builder.WriteString(m.MoneytransferType)
	builder.WriteByte(')')
	return builder.String()
}

// Moneytransfers is a parsable slice of Moneytransfer.
type Moneytransfers []*Moneytransfer

func (m Moneytransfers) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
