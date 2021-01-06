// Code generated by entc, DO NOT EDIT.

package payment

import (
	"time"
)

const (
	// Label holds the string label denoting the payment type in the database.
	Label = "payment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAccountName holds the string denoting the account_name field in the database.
	FieldAccountName = "account_name"
	// FieldAccountNumber holds the string denoting the account_number field in the database.
	FieldAccountNumber = "account_number"
	// FieldTransferTime holds the string denoting the transfer_time field in the database.
	FieldTransferTime = "transfer_time"

	// EdgeInsurance holds the string denoting the insurance edge name in mutations.
	EdgeInsurance = "Insurance"
	// EdgeMoneyTransfer holds the string denoting the moneytransfer edge name in mutations.
	EdgeMoneyTransfer = "MoneyTransfer"
	// EdgeBank holds the string denoting the bank edge name in mutations.
	EdgeBank = "Bank"
	// EdgeMember holds the string denoting the member edge name in mutations.
	EdgeMember = "Member"

	// Table holds the table name of the payment in the database.
	Table = "payments"
	// InsuranceTable is the table the holds the Insurance relation/edge.
	InsuranceTable = "payments"
	// InsuranceInverseTable is the table name for the Insurance entity.
	// It exists in this package in order to avoid circular dependency with the "insurance" package.
	InsuranceInverseTable = "insurances"
	// InsuranceColumn is the table column denoting the Insurance relation/edge.
	InsuranceColumn = "insurance_id"
	// MoneyTransferTable is the table the holds the MoneyTransfer relation/edge.
	MoneyTransferTable = "payments"
	// MoneyTransferInverseTable is the table name for the MoneyTransfer entity.
	// It exists in this package in order to avoid circular dependency with the "moneytransfer" package.
	MoneyTransferInverseTable = "money_transfers"
	// MoneyTransferColumn is the table column denoting the MoneyTransfer relation/edge.
	MoneyTransferColumn = "moneytransfer_id"
	// BankTable is the table the holds the Bank relation/edge.
	BankTable = "payments"
	// BankInverseTable is the table name for the Bank entity.
	// It exists in this package in order to avoid circular dependency with the "bank" package.
	BankInverseTable = "banks"
	// BankColumn is the table column denoting the Bank relation/edge.
	BankColumn = "bank_id"
	// MemberTable is the table the holds the Member relation/edge.
	MemberTable = "payments"
	// MemberInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	MemberInverseTable = "members"
	// MemberColumn is the table column denoting the Member relation/edge.
	MemberColumn = "member_id"
)

// Columns holds all SQL columns for payment fields.
var Columns = []string{
	FieldID,
	FieldAccountName,
	FieldAccountNumber,
	FieldTransferTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Payment type.
var ForeignKeys = []string{
	"bank_id",
	"insurance_id",
	"member_id",
	"moneytransfer_id",
}

var (
	// AccountNameValidator is a validator for the "account_name" field. It is called by the builders before save.
	AccountNameValidator func(string) error
	// AccountNumberValidator is a validator for the "account_number" field. It is called by the builders before save.
	AccountNumberValidator func(string) error
	// DefaultTransferTime holds the default value on creation for the transfer_time field.
	DefaultTransferTime func() time.Time
)
