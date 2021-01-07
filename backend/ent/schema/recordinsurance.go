package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Recordinsurance holds the schema definition for the Recordinsurance entity.
type Recordinsurance struct {
	ent.Schema
}

// Fields of the Insurance.
func (Recordinsurance) Fields() []ent.Field {
	return []ent.Field{
		field.Time("recordinsurance_time").Default(time.Now),
	}
}

// Edges of the Recordinsurance.
func (Recordinsurance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Member", Member.Type).Ref("member_recordinsurance").Unique(),
		edge.From("Hospital", Hospital.Type).Ref("hospital_recordinsurance").Unique(),
		edge.From("Officer", Officer.Type).Ref("officer_recordinsurance").Unique(),
		edge.From("Product", Product.Type).Ref("product_recordinsurance").Unique(),
		edge.From("Amountpaid", Amountpaid.Type).Ref("amountpaid_recordinsurance").Unique(),
	}
}
