// Code generated by entc, DO NOT EDIT.

package groupofage

const (
	// Label holds the string label denoting the groupofage type in the database.
	Label = "group_of_age"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGroupOfAgeName holds the string denoting the group_of_age_name field in the database.
	FieldGroupOfAgeName = "group_of_age_name"
	// FieldGroupOfAgeAge holds the string denoting the group_of_age_age field in the database.
	FieldGroupOfAgeAge = "group_of_age_age"

	// EdgeGroupofageProduct holds the string denoting the groupofage_product edge name in mutations.
	EdgeGroupofageProduct = "groupofage_product"

	// Table holds the table name of the groupofage in the database.
	Table = "group_of_ages"
	// GroupofageProductTable is the table the holds the groupofage_product relation/edge.
	GroupofageProductTable = "products"
	// GroupofageProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	GroupofageProductInverseTable = "products"
	// GroupofageProductColumn is the table column denoting the groupofage_product relation/edge.
	GroupofageProductColumn = "group_of_age_id"
)

// Columns holds all SQL columns for groupofage fields.
var Columns = []string{
	FieldID,
	FieldGroupOfAgeName,
	FieldGroupOfAgeAge,
}

var (
	// GroupOfAgeNameValidator is a validator for the "group_of_age_name" field. It is called by the builders before save.
	GroupOfAgeNameValidator func(string) error
	// GroupOfAgeAgeValidator is a validator for the "group_of_age_age" field. It is called by the builders before save.
	GroupOfAgeAgeValidator func(string) error
)
