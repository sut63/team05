package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.String("account_name").NotEmpty(),
		field.String("account_number").NotEmpty(),
		field.Time("transfer_time").Default(time.Now),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Insurance", Insurance.Type).Ref("insurance_payment").Unique(),
		edge.From("MoneyTransfer", MoneyTransfer.Type).Ref("moneytransfer_payment").Unique(),
		edge.From("Bank", Bank.Type).Ref("bank_payment").Unique(),
		edge.From("Member", Member.Type).Ref("member_payment").Unique(),
	}
}
