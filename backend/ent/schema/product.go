package schema

import (
	"errors"
	"regexp"

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
		field.String("product_name").Validate(func(s string) error {
			match, _ := regexp.MatchString("[A-Z]$", s)
			if !match {
				return errors.New("ชื่อผลิตภัณฑ์ต้องขึ้นต้นด้วยตัวใหญ่เท่านั้น")
			}
			return nil
		}).NotEmpty().Unique(),

		//field.String("product_name").NotEmpty().MinLen(5).Match(regexp.MustCompile("[ABCDEFGHIJKLMNOPQRSTUVWXYZ]")),
		//field.String("product_name").NotEmpty().Unique(),
		field.Int("product_price").Range(100000, 1000000),
		field.Int("product_time").Min(1).Positive(),
		field.Int("product_payment_of_year").Min(10000),
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
