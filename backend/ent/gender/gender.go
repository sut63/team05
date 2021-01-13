// Code generated by entc, DO NOT EDIT.

package gender

const (
	// Label holds the string label denoting the gender type in the database.
	Label = "gender"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGenderName holds the string denoting the gender_name field in the database.
	FieldGenderName = "gender_name"

	// EdgeGenderProduct holds the string denoting the gender_product edge name in mutations.
	EdgeGenderProduct = "gender_product"

	// Table holds the table name of the gender in the database.
	Table = "genders"
	// GenderProductTable is the table the holds the gender_product relation/edge.
	GenderProductTable = "products"
	// GenderProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	GenderProductInverseTable = "products"
	// GenderProductColumn is the table column denoting the gender_product relation/edge.
	GenderProductColumn = "gender_id"
)

// Columns holds all SQL columns for gender fields.
var Columns = []string{
	FieldID,
	FieldGenderName,
}

var (
	// GenderNameValidator is a validator for the "gender_name" field. It is called by the builders before save.
	GenderNameValidator func(string) error
)
