// Code generated by entc, DO NOT EDIT.

package bank

const (
	// Label holds the string label denoting the bank type in the database.
	Label = "bank"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBankType holds the string denoting the bank_type field in the database.
	FieldBankType = "bank_type"

	// EdgeBankPayment holds the string denoting the bank_payment edge name in mutations.
	EdgeBankPayment = "bank_payment"
	// EdgeBankPayback holds the string denoting the bank_payback edge name in mutations.
	EdgeBankPayback = "bank_payback"

	// Table holds the table name of the bank in the database.
	Table = "banks"
	// BankPaymentTable is the table the holds the bank_payment relation/edge.
	BankPaymentTable = "payments"
	// BankPaymentInverseTable is the table name for the Payment entity.
	// It exists in this package in order to avoid circular dependency with the "payment" package.
	BankPaymentInverseTable = "payments"
	// BankPaymentColumn is the table column denoting the bank_payment relation/edge.
	BankPaymentColumn = "bank_id"
	// BankPaybackTable is the table the holds the bank_payback relation/edge.
	BankPaybackTable = "paybacks"
	// BankPaybackInverseTable is the table name for the Payback entity.
	// It exists in this package in order to avoid circular dependency with the "payback" package.
	BankPaybackInverseTable = "paybacks"
	// BankPaybackColumn is the table column denoting the bank_payback relation/edge.
	BankPaybackColumn = "bank_id"
)

// Columns holds all SQL columns for bank fields.
var Columns = []string{
	FieldID,
	FieldBankType,
}

var (
	// BankTypeValidator is a validator for the "bank_type" field. It is called by the builders before save.
	BankTypeValidator func(string) error
)
