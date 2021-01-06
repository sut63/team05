// Code generated by entc, DO NOT EDIT.

package inquiry

import (
	"time"
)

const (
	// Label holds the string label denoting the inquiry type in the database.
	Label = "inquiry"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldInquiryInguiryMessages holds the string denoting the inquiry_inguiry_messages field in the database.
	FieldInquiryInguiryMessages = "inquiry_inguiry_messages"
	// FieldInquiryTimeMessages holds the string denoting the inquiry_time_messages field in the database.
	FieldInquiryTimeMessages = "inquiry_time_messages"

	// EdgeMember holds the string denoting the member edge name in mutations.
	EdgeMember = "Member"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "Category"
	// EdgeOfficer holds the string denoting the officer edge name in mutations.
	EdgeOfficer = "Officer"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "Product"

	// Table holds the table name of the inquiry in the database.
	Table = "inquiries"
	// MemberTable is the table the holds the Member relation/edge.
	MemberTable = "inquiries"
	// MemberInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	MemberInverseTable = "members"
	// MemberColumn is the table column denoting the Member relation/edge.
	MemberColumn = "member_id"
	// CategoryTable is the table the holds the Category relation/edge.
	CategoryTable = "inquiries"
	// CategoryInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryInverseTable = "categories"
	// CategoryColumn is the table column denoting the Category relation/edge.
	CategoryColumn = "category_id"
	// OfficerTable is the table the holds the Officer relation/edge.
	OfficerTable = "inquiries"
	// OfficerInverseTable is the table name for the Officer entity.
	// It exists in this package in order to avoid circular dependency with the "officer" package.
	OfficerInverseTable = "officers"
	// OfficerColumn is the table column denoting the Officer relation/edge.
	OfficerColumn = "officer_id"
	// ProductTable is the table the holds the Product relation/edge.
	ProductTable = "inquiries"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the Product relation/edge.
	ProductColumn = "product_id"
)

// Columns holds all SQL columns for inquiry fields.
var Columns = []string{
	FieldID,
	FieldInquiryInguiryMessages,
	FieldInquiryTimeMessages,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Inquiry type.
var ForeignKeys = []string{
	"category_id",
	"member_id",
	"officer_id",
	"product_id",
}

var (
	// InquiryInguiryMessagesValidator is a validator for the "Inquiry_inguiry_messages" field. It is called by the builders before save.
	InquiryInguiryMessagesValidator func(string) error
	// DefaultInquiryTimeMessages holds the default value on creation for the Inquiry_time_messages field.
	DefaultInquiryTimeMessages func() time.Time
)
