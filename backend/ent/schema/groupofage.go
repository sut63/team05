package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// GroupOfAge holds the schema definition for the GroupOfAge entity.
type GroupOfAge struct {
	ent.Schema
}

// Fields of the GroupOfAge.
func (GroupOfAge) Fields() []ent.Field {
	return []ent.Field{
		field.String("group_of_age_name").Unique(),
		field.String("group_of_age_age"),
	}
}

// Edges of the GroupOfAge.
func (GroupOfAge) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("groupage_product", Product.Type).StorageKey(edge.Column("group_of_age_id")),
	}
}
