package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Officer holds the schema definition for the Officer entity.
type Officer struct {
	ent.Schema
}

// Fields of the Officer.
func (Officer) Fields() []ent.Field {
	return []ent.Field{
		field.String("officer_email").Unique().NotEmpty(),
		field.String("officer_name").Unique().NotEmpty(),
		field.String("officer_password").NotEmpty(),
		/*field.Float("product_payment_of_year"),
		  field.Time("created_at").
		      Default(time.Now),*/
	}
}

// Edges of the Officer.
func (Officer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("officer_product", Product.Type).StorageKey(edge.Column("officer_id")),
		edge.To("officer_insurance", Insurance.Type).StorageKey(edge.Column("officer_id")),
		edge.To("officer_inquiry", Inquiry.Type).StorageKey(edge.Column("officer_id")),
		edge.To("officer_payback", Payback.Type).StorageKey(edge.Column("officer_id")),
		edge.To("officer_recordinsurance", Recordinsurance.Type).StorageKey(edge.Column("officer_id")),
	}
}
