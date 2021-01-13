package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("product_name").NotEmpty(),
		field.Int("product_price"),
		field.Int("product_time"),
		field.Int("product_payment_of_year"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("gender", Gender.Type).Ref("gender_product").Unique(),
		edge.From("groupofage", GroupOfAge.Type).Ref("groupofage_product").Unique(),
		edge.From("officer", Officer.Type).Ref("officer_product").Unique(),

		edge.To("product_insurance", Insurance.Type).StorageKey(edge.Column("product_id")),
		edge.To("product_inquiry", Inquiry.Type).StorageKey(edge.Column("product_id")),
		edge.To("product_payback", Payback.Type).StorageKey(edge.Column("product_id")),
		edge.To("product_recordinsurance", Recordinsurance.Type).StorageKey(edge.Column("product_id")),
	}
}
