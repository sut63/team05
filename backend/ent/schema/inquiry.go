package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Inquiry holds the schema definition for the Inquiry entity.
type Inquiry struct {
	ent.Schema
}

// Fields of the Inquiry.
func (Inquiry) Fields() []ent.Field {
	return []ent.Field{
		field.String("Inquiry_messages").NotEmpty(),
		field.Time("Inquiry_time_messages").Default(time.Now),
	}
}

// Edges of the Inquiry.
func (Inquiry) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Member", Member.Type).Ref("member_inquiry").Unique(),
		edge.From("Category", Category.Type).Ref("category_inquiry").Unique(),
		edge.From("Officer", Officer.Type).Ref("officer_inquiry").Unique(),
		edge.From("Product", Product.Type).Ref("product_inquiry").Unique(),
	}
}
