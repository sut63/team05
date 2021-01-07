package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Hospital holds the schema definition for the Hospital entity.
type Hospital struct {
	ent.Schema
}

// Fields of the Hospital.
func (Hospital) Fields() []ent.Field {
	return []ent.Field{
		field.String("hospital_name").Unique().NotEmpty(),
	}
}

// Edges of the Hospital.
func (Hospital) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hospital_insurance", Insurance.Type).StorageKey(edge.Column("hospital_id")),
		edge.To("hospital_recordinsurance", Recordinsurance.Type).StorageKey(edge.Column("hospital_id")),
	}
}
