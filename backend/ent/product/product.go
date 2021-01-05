// Code generated by entc, DO NOT EDIT.

package product

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProductName holds the string denoting the product_name field in the database.
	FieldProductName = "product_name"
	// FieldProductPrice holds the string denoting the product_price field in the database.
	FieldProductPrice = "product_price"
	// FieldProductTime holds the string denoting the product_time field in the database.
	FieldProductTime = "product_time"
	// FieldProductPaymentOfYear holds the string denoting the product_payment_of_year field in the database.
	FieldProductPaymentOfYear = "product_payment_of_year"

	// EdgeGender holds the string denoting the gender edge name in mutations.
	EdgeGender = "Gender"
	// EdgeGoupOfAge holds the string denoting the goup_of_age edge name in mutations.
	EdgeGoupOfAge = "Goup_Of_Age"
	// EdgeOfficer holds the string denoting the officer edge name in mutations.
	EdgeOfficer = "Officer"
	// EdgeProductInsurance holds the string denoting the product_insurance edge name in mutations.
	EdgeProductInsurance = "product_insurance"

	// Table holds the table name of the product in the database.
	Table = "products"
	// GenderTable is the table the holds the Gender relation/edge.
	GenderTable = "products"
	// GenderInverseTable is the table name for the Gender entity.
	// It exists in this package in order to avoid circular dependency with the "gender" package.
	GenderInverseTable = "genders"
	// GenderColumn is the table column denoting the Gender relation/edge.
	GenderColumn = "gender_id"
	// GoupOfAgeTable is the table the holds the Goup_Of_Age relation/edge.
	GoupOfAgeTable = "products"
	// GoupOfAgeInverseTable is the table name for the GroupOfAge entity.
	// It exists in this package in order to avoid circular dependency with the "groupofage" package.
	GoupOfAgeInverseTable = "group_of_ages"
	// GoupOfAgeColumn is the table column denoting the Goup_Of_Age relation/edge.
	GoupOfAgeColumn = "group_of_age_id"
	// OfficerTable is the table the holds the Officer relation/edge.
	OfficerTable = "products"
	// OfficerInverseTable is the table name for the Officer entity.
	// It exists in this package in order to avoid circular dependency with the "officer" package.
	OfficerInverseTable = "officers"
	// OfficerColumn is the table column denoting the Officer relation/edge.
	OfficerColumn = "officer_id"
	// ProductInsuranceTable is the table the holds the product_insurance relation/edge.
	ProductInsuranceTable = "insurances"
	// ProductInsuranceInverseTable is the table name for the Insurance entity.
	// It exists in this package in order to avoid circular dependency with the "insurance" package.
	ProductInsuranceInverseTable = "insurances"
	// ProductInsuranceColumn is the table column denoting the product_insurance relation/edge.
	ProductInsuranceColumn = "product_id"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldProductName,
	FieldProductPrice,
	FieldProductTime,
	FieldProductPaymentOfYear,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Product type.
var ForeignKeys = []string{
	"gender_id",
	"group_of_age_id",
	"officer_id",
}
