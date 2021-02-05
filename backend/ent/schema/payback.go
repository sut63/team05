package schema

import (
	"errors"
	"regexp"
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Payback holds the schema definition for the Payback entity.
type Payback struct {
	ent.Schema
}

// Fields of the Payback.
func (Payback) Fields() []ent.Field {
	return []ent.Field{

		field.Time("payback_transfertime").Default(time.Now),

		field.String("payback_accountnumber").NotEmpty().MinLen(10).MaxLen(10).
			Validate(func(s string) error {
				match, _ := regexp.MatchString("[0123456789]\\d{9}", s)
				if !match {
					return errors.New("เลขบัญชีที่กรอกไม่ถูกต้อง")
				}
				return nil
			}),

		field.String("payback_accountname").NotEmpty().MinLen(10).MaxLen(10).
			Validate(func(s string) error {
				match, _ := regexp.MatchString("[0]\\d", s)
				if !match {
					return errors.New("เบอร์โทรที่กรอกไม่ถูกต้อง")
				}
				return nil
			}),

		field.String("payback_accountiden").NotEmpty().MinLen(13).MaxLen(13).
			Validate(func(s string) error {
				match, _ := regexp.MatchString("[0123456789]\\d{12}", s)
				if !match {
					return errors.New("เลขประจำตัวประชาชนไม่ถูกต้อง")
				}
				return nil
			}),
	}
}

// Edges of the Payback.
func (Payback) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Officer", Officer.Type).Ref("officer_payback").Unique(),
		edge.From("Member", Member.Type).Ref("member_payback").Unique(),
		edge.From("Product", Product.Type).Ref("product_payback").Unique(),
		edge.From("Bank", Bank.Type).Ref("bank_payback").Unique(),
	}
}