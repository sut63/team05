package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Bank holds the schema definition for the Bank entity.
type Bank struct {
	ent.Schema
}

// Fields of the Bank.
func (Bank) Fields() []ent.Field {
	return []ent.Field{
		field.String("bank_type").NotEmpty(),
	}
}

// Edges of the Bank.
func (Bank) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bank_payment", Payment.Type).StorageKey(edge.Column("bank_id")),
		edge.To("bank_payback", Payback.Type).StorageKey(edge.Column("bank_id")),
	}
}
