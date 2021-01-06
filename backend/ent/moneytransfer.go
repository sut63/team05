// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team05/ent/moneytransfer"
)

// MoneyTransfer is the model entity for the MoneyTransfer schema.
type MoneyTransfer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// MoneytransferType holds the value of the "moneytransfer_type" field.
	MoneytransferType string `json:"moneytransfer_type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MoneyTransferQuery when eager-loading is set.
	Edges MoneyTransferEdges `json:"edges"`
}

// MoneyTransferEdges holds the relations/edges for other nodes in the graph.
type MoneyTransferEdges struct {
	// MoneytransferPayment holds the value of the moneytransfer_payment edge.
	MoneytransferPayment []*Payment
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MoneytransferPaymentOrErr returns the MoneytransferPayment value or an error if the edge
// was not loaded in eager-loading.
func (e MoneyTransferEdges) MoneytransferPaymentOrErr() ([]*Payment, error) {
	if e.loadedTypes[0] {
		return e.MoneytransferPayment, nil
	}
	return nil, &NotLoadedError{edge: "moneytransfer_payment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MoneyTransfer) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // moneytransfer_type
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MoneyTransfer fields.
func (mt *MoneyTransfer) assignValues(values ...interface{}) error {
	if m, n := len(values), len(moneytransfer.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	mt.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field moneytransfer_type", values[0])
	} else if value.Valid {
		mt.MoneytransferType = value.String
	}
	return nil
}

// QueryMoneytransferPayment queries the moneytransfer_payment edge of the MoneyTransfer.
func (mt *MoneyTransfer) QueryMoneytransferPayment() *PaymentQuery {
	return (&MoneyTransferClient{config: mt.config}).QueryMoneytransferPayment(mt)
}

// Update returns a builder for updating this MoneyTransfer.
// Note that, you need to call MoneyTransfer.Unwrap() before calling this method, if this MoneyTransfer
// was returned from a transaction, and the transaction was committed or rolled back.
func (mt *MoneyTransfer) Update() *MoneyTransferUpdateOne {
	return (&MoneyTransferClient{config: mt.config}).UpdateOne(mt)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (mt *MoneyTransfer) Unwrap() *MoneyTransfer {
	tx, ok := mt.config.driver.(*txDriver)
	if !ok {
		panic("ent: MoneyTransfer is not a transactional entity")
	}
	mt.config.driver = tx.drv
	return mt
}

// String implements the fmt.Stringer.
func (mt *MoneyTransfer) String() string {
	var builder strings.Builder
	builder.WriteString("MoneyTransfer(")
	builder.WriteString(fmt.Sprintf("id=%v", mt.ID))
	builder.WriteString(", moneytransfer_type=")
	builder.WriteString(mt.MoneytransferType)
	builder.WriteByte(')')
	return builder.String()
}

// MoneyTransfers is a parsable slice of MoneyTransfer.
type MoneyTransfers []*MoneyTransfer

func (mt MoneyTransfers) config(cfg config) {
	for _i := range mt {
		mt[_i].config = cfg
	}
}
