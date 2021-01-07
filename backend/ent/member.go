// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team05/ent/member"
)

// Member is the model entity for the Member schema.
type Member struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// MemberEmail holds the value of the "member_email" field.
	MemberEmail string `json:"member_email,omitempty"`
	// MemberName holds the value of the "member_name" field.
	MemberName string `json:"member_name,omitempty"`
	// MemberPassword holds the value of the "member_password" field.
	MemberPassword string `json:"member_password,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MemberQuery when eager-loading is set.
	Edges MemberEdges `json:"edges"`
}

// MemberEdges holds the relations/edges for other nodes in the graph.
type MemberEdges struct {
	// MemberInsurance holds the value of the member_insurance edge.
	MemberInsurance []*Insurance
	// MemberPayment holds the value of the member_payment edge.
	MemberPayment []*Payment
	// MemberInquiry holds the value of the member_inquiry edge.
	MemberInquiry []*Inquiry
<<<<<<< HEAD
	// MemberPayback holds the value of the member_payback edge.
	MemberPayback []*Payback
=======
	// MemberRecordinsurance holds the value of the member_recordinsurance edge.
	MemberRecordinsurance []*Recordinsurance
>>>>>>> 4637a9d (ทำ Entity สำหรับเก็บข้อมูลสิทธิประกันสุขภาพ - fix #53)
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// MemberInsuranceOrErr returns the MemberInsurance value or an error if the edge
// was not loaded in eager-loading.
func (e MemberEdges) MemberInsuranceOrErr() ([]*Insurance, error) {
	if e.loadedTypes[0] {
		return e.MemberInsurance, nil
	}
	return nil, &NotLoadedError{edge: "member_insurance"}
}

// MemberPaymentOrErr returns the MemberPayment value or an error if the edge
// was not loaded in eager-loading.
func (e MemberEdges) MemberPaymentOrErr() ([]*Payment, error) {
	if e.loadedTypes[1] {
		return e.MemberPayment, nil
	}
	return nil, &NotLoadedError{edge: "member_payment"}
}

// MemberInquiryOrErr returns the MemberInquiry value or an error if the edge
// was not loaded in eager-loading.
func (e MemberEdges) MemberInquiryOrErr() ([]*Inquiry, error) {
	if e.loadedTypes[2] {
		return e.MemberInquiry, nil
	}
	return nil, &NotLoadedError{edge: "member_inquiry"}
}

<<<<<<< HEAD
// MemberPaybackOrErr returns the MemberPayback value or an error if the edge
// was not loaded in eager-loading.
func (e MemberEdges) MemberPaybackOrErr() ([]*Payback, error) {
	if e.loadedTypes[3] {
		return e.MemberPayback, nil
	}
	return nil, &NotLoadedError{edge: "member_payback"}
=======
// MemberRecordinsuranceOrErr returns the MemberRecordinsurance value or an error if the edge
// was not loaded in eager-loading.
func (e MemberEdges) MemberRecordinsuranceOrErr() ([]*Recordinsurance, error) {
	if e.loadedTypes[3] {
		return e.MemberRecordinsurance, nil
	}
	return nil, &NotLoadedError{edge: "member_recordinsurance"}
>>>>>>> 4637a9d (ทำ Entity สำหรับเก็บข้อมูลสิทธิประกันสุขภาพ - fix #53)
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Member) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // member_email
		&sql.NullString{}, // member_name
		&sql.NullString{}, // member_password
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Member fields.
func (m *Member) assignValues(values ...interface{}) error {
	if m, n := len(values), len(member.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field member_email", values[0])
	} else if value.Valid {
		m.MemberEmail = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field member_name", values[1])
	} else if value.Valid {
		m.MemberName = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field member_password", values[2])
	} else if value.Valid {
		m.MemberPassword = value.String
	}
	return nil
}

// QueryMemberInsurance queries the member_insurance edge of the Member.
func (m *Member) QueryMemberInsurance() *InsuranceQuery {
	return (&MemberClient{config: m.config}).QueryMemberInsurance(m)
}

// QueryMemberPayment queries the member_payment edge of the Member.
func (m *Member) QueryMemberPayment() *PaymentQuery {
	return (&MemberClient{config: m.config}).QueryMemberPayment(m)
}

// QueryMemberInquiry queries the member_inquiry edge of the Member.
func (m *Member) QueryMemberInquiry() *InquiryQuery {
	return (&MemberClient{config: m.config}).QueryMemberInquiry(m)
}

<<<<<<< HEAD
// QueryMemberPayback queries the member_payback edge of the Member.
func (m *Member) QueryMemberPayback() *PaybackQuery {
	return (&MemberClient{config: m.config}).QueryMemberPayback(m)
=======
// QueryMemberRecordinsurance queries the member_recordinsurance edge of the Member.
func (m *Member) QueryMemberRecordinsurance() *RecordinsuranceQuery {
	return (&MemberClient{config: m.config}).QueryMemberRecordinsurance(m)
>>>>>>> 4637a9d (ทำ Entity สำหรับเก็บข้อมูลสิทธิประกันสุขภาพ - fix #53)
}

// Update returns a builder for updating this Member.
// Note that, you need to call Member.Unwrap() before calling this method, if this Member
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Member) Update() *MemberUpdateOne {
	return (&MemberClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Member) Unwrap() *Member {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Member is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Member) String() string {
	var builder strings.Builder
	builder.WriteString("Member(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", member_email=")
	builder.WriteString(m.MemberEmail)
	builder.WriteString(", member_name=")
	builder.WriteString(m.MemberName)
	builder.WriteString(", member_password=")
	builder.WriteString(m.MemberPassword)
	builder.WriteByte(')')
	return builder.String()
}

// Members is a parsable slice of Member.
type Members []*Member

func (m Members) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
