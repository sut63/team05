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
		field.String("product_name").Unique(),
		field.Int("product_price"),
		field.Int("product_time"),
		field.Float("product_payment_of_year"),
		/*field.Time("created_at").
		  Default(time.Now),*/
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product_gender", Gender.Type).Ref("genders").Unique(),
		edge.From("product_groupage", GroupOfAge.Type).Ref("groupage_product").Unique(),
		edge.From("product_officer", Officer.Type).Ref("officer_product").Unique(),
		edge.To("product_insurance", Insurance.Type).StorageKey(edge.Column("product_id")),
		edge.To("product_inquiry", Inquiry.Type).StorageKey(edge.Column("product_id")),
		edge.To("product_payback", Payback.Type).StorageKey(edge.Column("product_id")),
		edge.To("product_recordinsurance", Recordinsurance.Type).StorageKey(edge.Column("product_id")),
	}
}
