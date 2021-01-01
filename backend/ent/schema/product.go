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
		edge.From("Gender", Gender.Type).Ref("genders").Unique(),
		edge.From("Goup_Of_Age", GroupOfAge.Type).Ref("gropages").Unique(),
		edge.From("Officer", Officer.Type).Ref("officers").Unique(),
	}
}
