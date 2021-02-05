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

// Inquiry holds the schema definition for the Inquiry entity.
type Inquiry struct {
	ent.Schema
}

// Fields of the Inquiry.
func (Inquiry) Fields() []ent.Field {
	return []ent.Field{
		field.String("Inquiry_name_messages").NotEmpty().Match(regexp.MustCompile("^[a-zA-Z ]+$")).
			Validate(func(s string) error {
				if strings.ToLower(s) == s {
					return errors.New("รูปแบบชื่อไม่ถูกต้อง")
				}
				return nil
			}),

		field.String("Inquiry_phone_messages").NotEmpty().MaxLen(10).MinLen(10).
			NotEmpty().Validate(func(s string) error {
			match, _ := regexp.MatchString("^[0-9]+$", s)
			if !match {
				return errors.New("รูปแบบเบอร์โทรศัพท์ไม่ถูกต้อง กรุณาใส่เลขเบอร์โทร 10 หลัก")
			}
			return nil
		}),

		field.Int("Inquiry_age_messages").Positive(),

		field.String("Inquiry_messages").
			NotEmpty().Validate(func(s string) error {
			match, _ := regexp.MatchString("^[ก-๙0-9a-zA-Z- ./\\s]+$", s)
			if !match {
				return errors.New("รูปแบบรายละเอียดไม่ถูกต้อง")
			}
			return nil
		}),

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
