package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.String("member_email").Unique().NotEmpty(),
		field.String("member_name").Unique().NotEmpty(),
		field.String("member_password").NotEmpty(),
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("member_insurance", Insurance.Type).StorageKey(edge.Column("member_id")),
		edge.To("member_payment", Payment.Type).StorageKey(edge.Column("member_id")),
		edge.To("member_inquiry", Inquiry.Type).StorageKey(edge.Column("member_id")),
		edge.To("member_payback", Payback.Type).StorageKey(edge.Column("member_id")),
		edge.To("member_recordinsurance", Recordinsurance.Type).StorageKey(edge.Column("member_id")),
	}
}
