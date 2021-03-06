// Code generated by entc, DO NOT EDIT.

package position

const (
	// Label holds the string label denoting the position type in the database.
	Label = "position"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPositionName holds the string denoting the position_name field in the database.
	FieldPositionName = "position_name"

	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// EdgeOfficers holds the string denoting the officers edge name in mutations.
	EdgeOfficers = "officers"

	// Table holds the table name of the position in the database.
	Table = "positions"
	// MembersTable is the table the holds the members relation/edge.
	MembersTable = "members"
	// MembersInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	MembersInverseTable = "members"
	// MembersColumn is the table column denoting the members relation/edge.
	MembersColumn = "position_id"
	// OfficersTable is the table the holds the officers relation/edge.
	OfficersTable = "officers"
	// OfficersInverseTable is the table name for the Officer entity.
	// It exists in this package in order to avoid circular dependency with the "officer" package.
	OfficersInverseTable = "officers"
	// OfficersColumn is the table column denoting the officers relation/edge.
	OfficersColumn = "position_id"
)

// Columns holds all SQL columns for position fields.
var Columns = []string{
	FieldID,
	FieldPositionName,
}

var (
	// PositionNameValidator is a validator for the "position_name" field. It is called by the builders before save.
	PositionNameValidator func(string) error
)
