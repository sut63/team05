// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team05/ent/bank"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/payback"
	"github.com/sut63/team05/ent/product"
)

// Payback is the model entity for the Payback schema.
type Payback struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PaybackTransfertime holds the value of the "payback_transfertime" field.
	PaybackTransfertime time.Time `json:"payback_transfertime,omitempty"`
	// PaybackAccountnumber holds the value of the "payback_accountnumber" field.
	PaybackAccountnumber string `json:"payback_accountnumber,omitempty"`
	// PaybackAccountname holds the value of the "payback_accountname" field.
	PaybackAccountname string `json:"payback_accountname,omitempty"`
	// PaybackAccountiden holds the value of the "payback_accountiden" field.
	PaybackAccountiden string `json:"payback_accountiden,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PaybackQuery when eager-loading is set.
	Edges      PaybackEdges `json:"edges"`
	bank_id    *int
	member_id  *int
	officer_id *int
	product_id *int
}

// PaybackEdges holds the relations/edges for other nodes in the graph.
type PaybackEdges struct {
	// Officer holds the value of the Officer edge.
	Officer *Officer
	// Member holds the value of the Member edge.
	Member *Member
	// Product holds the value of the Product edge.
	Product *Product
	// Bank holds the value of the Bank edge.
	Bank *Bank
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// OfficerOrErr returns the Officer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaybackEdges) OfficerOrErr() (*Officer, error) {
	if e.loadedTypes[0] {
		if e.Officer == nil {
			// The edge Officer was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: officer.Label}
		}
		return e.Officer, nil
	}
	return nil, &NotLoadedError{edge: "Officer"}
}

// MemberOrErr returns the Member value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaybackEdges) MemberOrErr() (*Member, error) {
	if e.loadedTypes[1] {
		if e.Member == nil {
			// The edge Member was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: member.Label}
		}
		return e.Member, nil
	}
	return nil, &NotLoadedError{edge: "Member"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaybackEdges) ProductOrErr() (*Product, error) {
	if e.loadedTypes[2] {
		if e.Product == nil {
			// The edge Product was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: product.Label}
		}
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "Product"}
}

// BankOrErr returns the Bank value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PaybackEdges) BankOrErr() (*Bank, error) {
	if e.loadedTypes[3] {
		if e.Bank == nil {
			// The edge Bank was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: bank.Label}
		}
		return e.Bank, nil
	}
	return nil, &NotLoadedError{edge: "Bank"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Payback) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // payback_transfertime
		&sql.NullString{}, // payback_accountnumber
		&sql.NullString{}, // payback_accountname
		&sql.NullString{}, // payback_accountiden
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Payback) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // bank_id
		&sql.NullInt64{}, // member_id
		&sql.NullInt64{}, // officer_id
		&sql.NullInt64{}, // product_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Payback fields.
func (pa *Payback) assignValues(values ...interface{}) error {
	if m, n := len(values), len(payback.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pa.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field payback_transfertime", values[0])
	} else if value.Valid {
		pa.PaybackTransfertime = value.Time
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field payback_accountnumber", values[1])
	} else if value.Valid {
		pa.PaybackAccountnumber = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field payback_accountname", values[2])
	} else if value.Valid {
		pa.PaybackAccountname = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field payback_accountiden", values[3])
	} else if value.Valid {
		pa.PaybackAccountiden = value.String
	}
	values = values[4:]
	if len(values) == len(payback.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field bank_id", value)
		} else if value.Valid {
			pa.bank_id = new(int)
			*pa.bank_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field member_id", value)
		} else if value.Valid {
			pa.member_id = new(int)
			*pa.member_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field officer_id", value)
		} else if value.Valid {
			pa.officer_id = new(int)
			*pa.officer_id = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field product_id", value)
		} else if value.Valid {
			pa.product_id = new(int)
			*pa.product_id = int(value.Int64)
		}
	}
	return nil
}

// QueryOfficer queries the Officer edge of the Payback.
func (pa *Payback) QueryOfficer() *OfficerQuery {
	return (&PaybackClient{config: pa.config}).QueryOfficer(pa)
}

// QueryMember queries the Member edge of the Payback.
func (pa *Payback) QueryMember() *MemberQuery {
	return (&PaybackClient{config: pa.config}).QueryMember(pa)
}

// QueryProduct queries the Product edge of the Payback.
func (pa *Payback) QueryProduct() *ProductQuery {
	return (&PaybackClient{config: pa.config}).QueryProduct(pa)
}

// QueryBank queries the Bank edge of the Payback.
func (pa *Payback) QueryBank() *BankQuery {
	return (&PaybackClient{config: pa.config}).QueryBank(pa)
}

// Update returns a builder for updating this Payback.
// Note that, you need to call Payback.Unwrap() before calling this method, if this Payback
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Payback) Update() *PaybackUpdateOne {
	return (&PaybackClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pa *Payback) Unwrap() *Payback {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Payback is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Payback) String() string {
	var builder strings.Builder
	builder.WriteString("Payback(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", payback_transfertime=")
	builder.WriteString(pa.PaybackTransfertime.Format(time.ANSIC))
	builder.WriteString(", payback_accountnumber=")
	builder.WriteString(pa.PaybackAccountnumber)
	builder.WriteString(", payback_accountname=")
	builder.WriteString(pa.PaybackAccountname)
	builder.WriteString(", payback_accountiden=")
	builder.WriteString(pa.PaybackAccountiden)
	builder.WriteByte(')')
	return builder.String()
}

// Paybacks is a parsable slice of Payback.
type Paybacks []*Payback

func (pa Paybacks) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
