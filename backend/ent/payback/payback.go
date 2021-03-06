// Code generated by entc, DO NOT EDIT.

package payback

import (
	"time"
)

const (
	// Label holds the string label denoting the payback type in the database.
	Label = "payback"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPaybackTransfertime holds the string denoting the payback_transfertime field in the database.
	FieldPaybackTransfertime = "payback_transfertime"
	// FieldPaybackAccountnumber holds the string denoting the payback_accountnumber field in the database.
	FieldPaybackAccountnumber = "payback_accountnumber"
	// FieldPaybackAccountname holds the string denoting the payback_accountname field in the database.
	FieldPaybackAccountname = "payback_accountname"
	// FieldPaybackAccountiden holds the string denoting the payback_accountiden field in the database.
	FieldPaybackAccountiden = "payback_accountiden"

	// EdgeOfficer holds the string denoting the officer edge name in mutations.
	EdgeOfficer = "Officer"
	// EdgeMember holds the string denoting the member edge name in mutations.
	EdgeMember = "Member"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "Product"
	// EdgeBank holds the string denoting the bank edge name in mutations.
	EdgeBank = "Bank"

	// Table holds the table name of the payback in the database.
	Table = "paybacks"
	// OfficerTable is the table the holds the Officer relation/edge.
	OfficerTable = "paybacks"
	// OfficerInverseTable is the table name for the Officer entity.
	// It exists in this package in order to avoid circular dependency with the "officer" package.
	OfficerInverseTable = "officers"
	// OfficerColumn is the table column denoting the Officer relation/edge.
	OfficerColumn = "officer_id"
	// MemberTable is the table the holds the Member relation/edge.
	MemberTable = "paybacks"
	// MemberInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	MemberInverseTable = "members"
	// MemberColumn is the table column denoting the Member relation/edge.
	MemberColumn = "member_id"
	// ProductTable is the table the holds the Product relation/edge.
	ProductTable = "paybacks"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the Product relation/edge.
	ProductColumn = "product_id"
	// BankTable is the table the holds the Bank relation/edge.
	BankTable = "paybacks"
	// BankInverseTable is the table name for the Bank entity.
	// It exists in this package in order to avoid circular dependency with the "bank" package.
	BankInverseTable = "banks"
	// BankColumn is the table column denoting the Bank relation/edge.
	BankColumn = "bank_id"
)

// Columns holds all SQL columns for payback fields.
var Columns = []string{
	FieldID,
	FieldPaybackTransfertime,
	FieldPaybackAccountnumber,
	FieldPaybackAccountname,
	FieldPaybackAccountiden,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Payback type.
var ForeignKeys = []string{
	"bank_id",
	"member_id",
	"officer_id",
	"product_id",
}

var (
	// DefaultPaybackTransfertime holds the default value on creation for the payback_transfertime field.
	DefaultPaybackTransfertime func() time.Time
	// PaybackAccountnumberValidator is a validator for the "payback_accountnumber" field. It is called by the builders before save.
	PaybackAccountnumberValidator func(string) error
	// PaybackAccountnameValidator is a validator for the "payback_accountname" field. It is called by the builders before save.
	PaybackAccountnameValidator func(string) error
	// PaybackAccountidenValidator is a validator for the "payback_accountiden" field. It is called by the builders before save.
	PaybackAccountidenValidator func(string) error
)
