package schema

import (
	"errors"
	"regexp"
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

		field.Int("number_of_days_of_treat").Range(0, 30),
		field.String("recordinsurance_contact").NotEmpty().MinLen(10).MaxLen(10).
			Validate(func(s string) error {
				match, _ := regexp.MatchString("^[0]\\d", s)
				if !match {
					return errors.New("เบอร์โทรที่กรอกต้องขึ้นต้นด้วย 0")
				}
				return nil
			}),

		field.String("recordinsurance_address").Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ก-๙0-9a-zA-Z- ./\\s]+$", s)
			if !match {
				return errors.New("ป้อนรูปแบบของที่อยู่ไม่ถูกต้อง")
			}
			return nil
		}).NotEmpty(),
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
