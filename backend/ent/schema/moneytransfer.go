package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Moneytransfer holds the schema definition for the Moneytransfer entity.
type Moneytransfer struct {
	ent.Schema
}

// Fields of the Moneytransfer.
func (Moneytransfer) Fields() []ent.Field {
	return []ent.Field{
		field.String("moneytransfer_type").NotEmpty(),
	}
}

// Edges of the Moneytransfer.
func (Moneytransfer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("moneytransfer_payment", Payment.Type).StorageKey(edge.Column("moneytransfer_id")),
	}
}
