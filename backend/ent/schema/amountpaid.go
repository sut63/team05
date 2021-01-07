package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Amountpaid holds the schema definition for the Amountpaid entity.
type Amountpaid struct {
	ent.Schema
}

// Fields of the Amountpaid.
func (Amountpaid) Fields() []ent.Field {
	return []ent.Field{
		field.Float("amountpaid_money").Unique(),
	}
}

// Edges of the Amountpaid.
func (Amountpaid) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("amountpaid_recordinsurance", Recordinsurance.Type).StorageKey(edge.Column("amountpaid_id")),
	}
}
