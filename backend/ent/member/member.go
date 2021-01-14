// Code generated by entc, DO NOT EDIT.

package member

const (
	// Label holds the string label denoting the member type in the database.
	Label = "member"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMemberEmail holds the string denoting the member_email field in the database.
	FieldMemberEmail = "member_email"
	// FieldMemberName holds the string denoting the member_name field in the database.
	FieldMemberName = "member_name"
	// FieldMemberPassword holds the string denoting the member_password field in the database.
	FieldMemberPassword = "member_password"

	// EdgeMemberInsurance holds the string denoting the member_insurance edge name in mutations.
	EdgeMemberInsurance = "member_insurance"
	// EdgeMemberPayment holds the string denoting the member_payment edge name in mutations.
	EdgeMemberPayment = "member_payment"
	// EdgeMemberInquiry holds the string denoting the member_inquiry edge name in mutations.
	EdgeMemberInquiry = "member_inquiry"
	// EdgeMemberPayback holds the string denoting the member_payback edge name in mutations.
	EdgeMemberPayback = "member_payback"
	// EdgeMemberRecordinsurance holds the string denoting the member_recordinsurance edge name in mutations.
	EdgeMemberRecordinsurance = "member_recordinsurance"
	// EdgePosition holds the string denoting the position edge name in mutations.
	EdgePosition = "position"

	// Table holds the table name of the member in the database.
	Table = "members"
	// MemberInsuranceTable is the table the holds the member_insurance relation/edge.
	MemberInsuranceTable = "insurances"
	// MemberInsuranceInverseTable is the table name for the Insurance entity.
	// It exists in this package in order to avoid circular dependency with the "insurance" package.
	MemberInsuranceInverseTable = "insurances"
	// MemberInsuranceColumn is the table column denoting the member_insurance relation/edge.
	MemberInsuranceColumn = "member_id"
	// MemberPaymentTable is the table the holds the member_payment relation/edge.
	MemberPaymentTable = "payments"
	// MemberPaymentInverseTable is the table name for the Payment entity.
	// It exists in this package in order to avoid circular dependency with the "payment" package.
	MemberPaymentInverseTable = "payments"
	// MemberPaymentColumn is the table column denoting the member_payment relation/edge.
	MemberPaymentColumn = "member_id"
	// MemberInquiryTable is the table the holds the member_inquiry relation/edge.
	MemberInquiryTable = "inquiries"
	// MemberInquiryInverseTable is the table name for the Inquiry entity.
	// It exists in this package in order to avoid circular dependency with the "inquiry" package.
	MemberInquiryInverseTable = "inquiries"
	// MemberInquiryColumn is the table column denoting the member_inquiry relation/edge.
	MemberInquiryColumn = "member_id"
	// MemberPaybackTable is the table the holds the member_payback relation/edge.
	MemberPaybackTable = "paybacks"
	// MemberPaybackInverseTable is the table name for the Payback entity.
	// It exists in this package in order to avoid circular dependency with the "payback" package.
	MemberPaybackInverseTable = "paybacks"
	// MemberPaybackColumn is the table column denoting the member_payback relation/edge.
	MemberPaybackColumn = "member_id"
	// MemberRecordinsuranceTable is the table the holds the member_recordinsurance relation/edge.
	MemberRecordinsuranceTable = "recordinsurances"
	// MemberRecordinsuranceInverseTable is the table name for the Recordinsurance entity.
	// It exists in this package in order to avoid circular dependency with the "recordinsurance" package.
	MemberRecordinsuranceInverseTable = "recordinsurances"
	// MemberRecordinsuranceColumn is the table column denoting the member_recordinsurance relation/edge.
	MemberRecordinsuranceColumn = "member_id"
	// PositionTable is the table the holds the position relation/edge.
	PositionTable = "members"
	// PositionInverseTable is the table name for the Position entity.
	// It exists in this package in order to avoid circular dependency with the "position" package.
	PositionInverseTable = "positions"
	// PositionColumn is the table column denoting the position relation/edge.
	PositionColumn = "position_id"
)

// Columns holds all SQL columns for member fields.
var Columns = []string{
	FieldID,
	FieldMemberEmail,
	FieldMemberName,
	FieldMemberPassword,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Member type.
var ForeignKeys = []string{
	"position_id",
}

var (
	// MemberEmailValidator is a validator for the "member_email" field. It is called by the builders before save.
	MemberEmailValidator func(string) error
	// MemberNameValidator is a validator for the "member_name" field. It is called by the builders before save.
	MemberNameValidator func(string) error
	// MemberPasswordValidator is a validator for the "member_password" field. It is called by the builders before save.
	MemberPasswordValidator func(string) error
)
