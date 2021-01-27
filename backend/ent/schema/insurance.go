package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Insurance holds the schema definition for the Insurance entity.
type Insurance struct {
	ent.Schema
}

// Fields of the Insurance.
func (Insurance) Fields() []ent.Field {
	return []ent.Field{
		field.String("insurance_address").NotEmpty(),
		field.String("insurance_insurer").NotEmpty(),
		field.Time("insurance_time_buy").Default(time.Now),
		field.Time("insurance_time_firstpay").Default(time.Now),
	}
}

// Edges of the Insurance.
func (Insurance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Member", Member.Type).Ref("member_insurance").Unique(),
		edge.From("Hospital", Hospital.Type).Ref("hospital_insurance").Unique(),
		edge.From("Officer", Officer.Type).Ref("officer_insurance").Unique(),
		edge.From("Product", Product.Type).Ref("product_insurance").Unique(),
		edge.To("insurance_payment", Payment.Type).StorageKey(edge.Column("insurance_id")),
	}
}
