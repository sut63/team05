package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// MoneyTransfer holds the schema definition for the MoneyTransfer entity.
type MoneyTransfer struct {
	ent.Schema
}

// Fields of the MoneyTransfer.
func (MoneyTransfer) Fields() []ent.Field {
	return []ent.Field{
		field.String("moneytransfer_type").NotEmpty(),
	}
}

// Edges of the MoneyTransfer.
func (MoneyTransfer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("moneytransfer_payment", Payment.Type).StorageKey(edge.Column("moneytransfer_id")),
	}
}
