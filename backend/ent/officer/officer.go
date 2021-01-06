// Code generated by entc, DO NOT EDIT.

package officer

const (
	// Label holds the string label denoting the officer type in the database.
	Label = "officer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOfficerEmail holds the string denoting the officer_email field in the database.
	FieldOfficerEmail = "officer_email"
	// FieldOfficerName holds the string denoting the officer_name field in the database.
	FieldOfficerName = "officer_name"
	// FieldOfficerPassword holds the string denoting the officer_password field in the database.
	FieldOfficerPassword = "officer_password"

	// EdgeOfficers holds the string denoting the officers edge name in mutations.
	EdgeOfficers = "officers"
	// EdgeOfficerInsurance holds the string denoting the officer_insurance edge name in mutations.
	EdgeOfficerInsurance = "officer_insurance"
	// EdgeOfficerInquiry holds the string denoting the officer_inquiry edge name in mutations.
	EdgeOfficerInquiry = "officer_inquiry"

	// Table holds the table name of the officer in the database.
	Table = "officers"
	// OfficersTable is the table the holds the officers relation/edge.
	OfficersTable = "products"
	// OfficersInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	OfficersInverseTable = "products"
	// OfficersColumn is the table column denoting the officers relation/edge.
	OfficersColumn = "officer_id"
	// OfficerInsuranceTable is the table the holds the officer_insurance relation/edge.
	OfficerInsuranceTable = "insurances"
	// OfficerInsuranceInverseTable is the table name for the Insurance entity.
	// It exists in this package in order to avoid circular dependency with the "insurance" package.
	OfficerInsuranceInverseTable = "insurances"
	// OfficerInsuranceColumn is the table column denoting the officer_insurance relation/edge.
	OfficerInsuranceColumn = "officer_id"
	// OfficerInquiryTable is the table the holds the officer_inquiry relation/edge.
	OfficerInquiryTable = "inquiries"
	// OfficerInquiryInverseTable is the table name for the Inquiry entity.
	// It exists in this package in order to avoid circular dependency with the "inquiry" package.
	OfficerInquiryInverseTable = "inquiries"
	// OfficerInquiryColumn is the table column denoting the officer_inquiry relation/edge.
	OfficerInquiryColumn = "officer_id"
)

// Columns holds all SQL columns for officer fields.
var Columns = []string{
	FieldID,
	FieldOfficerEmail,
	FieldOfficerName,
	FieldOfficerPassword,
}

var (
	// OfficerEmailValidator is a validator for the "officer_email" field. It is called by the builders before save.
	OfficerEmailValidator func(string) error
	// OfficerNameValidator is a validator for the "officer_name" field. It is called by the builders before save.
	OfficerNameValidator func(string) error
	// OfficerPasswordValidator is a validator for the "officer_password" field. It is called by the builders before save.
	OfficerPasswordValidator func(string) error
)
