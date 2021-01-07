package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Payback holds the schema definition for the Payback entity.
type Payback struct {
	ent.Schema
}

// Fields of the Payback.
func (Payback) Fields() []ent.Field {
	return []ent.Field{
		field.String("Accountnumber").NotEmpty(),
		field.Time("Transfertime").Default(time.Now),
	}
}

// Edges of the Payback.
func (Payback) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Officer", Officer.Type).Ref("officer_payback").Unique(),
		edge.From("Member", Member.Type).Ref("member_payback").Unique(),
		edge.From("Product", Product.Type).Ref("product_payback").Unique(),
		edge.From("Bank", Bank.Type).Ref("bank_payback").Unique(),
	}
}
