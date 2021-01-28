package schema

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Insurance holds the schema definition for the Insurance entity.
type Insurance struct {
	ent.Schema
}

// Fields of the Insurance.
func (Insurance) Fields() []ent.Field {
	return []ent.Field{
		field.String("insurance_identification").NotEmpty().MinLen(13).MaxLen(13).
			Validate(func(s string) error {
				match, _ := regexp.MatchString("^[0-9]+$", s)
				if !match {
					return errors.New("Identification must be Number")
				}
				return nil
			}),
		field.String("insurance_insurer").NotEmpty().Match(regexp.MustCompile("^[a-zA-Z ]+$")).
			Validate(func(s string) error {
				if strings.ToLower(s) == s {
					return errors.New("Insurance name must begin with uppercase")
				}
				return nil
			}),

		field.String("insurance_address").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ก-๙0-9a-zA-Z- ./\\s]+$", s)
			if !match {
				return errors.New("Invalid address format")
			}
			return nil
		}).NotEmpty(),
		field.Time("insurance_time_buy").Immutable().Default(time.Now),
		field.Time("insurance_time_firstpay"),
	}
}

// Edges of the Insurance.
func (Insurance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Member", Member.Type).Ref("member_insurance").Unique(),
		edge.From("Hospital", Hospital.Type).Ref("hospital_insurance").Unique(),
		edge.From("Officer", Officer.Type).Ref("officer_insurance").Unique(),
		edge.From("Product", Product.Type).Ref("product_insurance").Unique(),
		edge.To("insurance_payment", Payment.Type).StorageKey(edge.Column("insurance_id")),
	}
}
