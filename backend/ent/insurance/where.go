// Code generated by entc, DO NOT EDIT.

package insurance

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team05/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// InsuranceIdentification applies equality check predicate on the "insurance_identification" field. It's identical to InsuranceIdentificationEQ.
func InsuranceIdentification(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInsuranceIdentification), v))
	})
}

// InsuranceInsurer applies equality check predicate on the "insurance_insurer" field. It's identical to InsuranceInsurerEQ.
func InsuranceInsurer(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInsuranceInsurer), v))
	})
}

// InsuranceAddress applies equality check predicate on the "insurance_address" field. It's identical to InsuranceAddressEQ.
func InsuranceAddress(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInsuranceAddress), v))
	})
}

// InsuranceTimeBuy applies equality check predicate on the "insurance_time_buy" field. It's identical to InsuranceTimeBuyEQ.
func InsuranceTimeBuy(v time.Time) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInsuranceTimeBuy), v))
	})
}

// InsuranceTimeFirstpay applies equality check predicate on the "insurance_time_firstpay" field. It's identical to InsuranceTimeFirstpayEQ.
func InsuranceTimeFirstpay(v time.Time) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInsuranceTimeFirstpay), v))
	})
}

// InsuranceIdentificationEQ applies the EQ predicate on the "insurance_identification" field.
func InsuranceIdentificationEQ(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInsuranceIdentification), v))
	})
}

// InsuranceIdentificationNEQ applies the NEQ predicate on the "insurance_identification" field.
func InsuranceIdentificationNEQ(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInsuranceIdentification), v))
	})
}

// InsuranceIdentificationIn applies the In predicate on the "insurance_identification" field.
func InsuranceIdentificationIn(vs ...string) predicate.Insurance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Insurance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInsuranceIdentification), v...))
	})
}

// InsuranceIdentificationNotIn applies the NotIn predicate on the "insurance_identification" field.
func InsuranceIdentificationNotIn(vs ...string) predicate.Insurance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Insurance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInsuranceIdentification), v...))
	})
}

// InsuranceIdentificationGT applies the GT predicate on the "insurance_identification" field.
func InsuranceIdentificationGT(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInsuranceIdentification), v))
	})
}

// InsuranceIdentificationGTE applies the GTE predicate on the "insurance_identification" field.
func InsuranceIdentificationGTE(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInsuranceIdentification), v))
	})
}

// InsuranceIdentificationLT applies the LT predicate on the "insurance_identification" field.
func InsuranceIdentificationLT(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInsuranceIdentification), v))
	})
}

// InsuranceIdentificationLTE applies the LTE predicate on the "insurance_identification" field.
func InsuranceIdentificationLTE(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInsuranceIdentification), v))
	})
}

// InsuranceIdentificationContains applies the Contains predicate on the "insurance_identification" field.
func InsuranceIdentificationContains(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInsuranceIdentification), v))
	})
}

// InsuranceIdentificationHasPrefix applies the HasPrefix predicate on the "insurance_identification" field.
func InsuranceIdentificationHasPrefix(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInsuranceIdentification), v))
	})
}

// InsuranceIdentificationHasSuffix applies the HasSuffix predicate on the "insurance_identification" field.
func InsuranceIdentificationHasSuffix(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInsuranceIdentification), v))
	})
}

// InsuranceIdentificationEqualFold applies the EqualFold predicate on the "insurance_identification" field.
func InsuranceIdentificationEqualFold(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInsuranceIdentification), v))
	})
}

// InsuranceIdentificationContainsFold applies the ContainsFold predicate on the "insurance_identification" field.
func InsuranceIdentificationContainsFold(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInsuranceIdentification), v))
	})
}

// InsuranceInsurerEQ applies the EQ predicate on the "insurance_insurer" field.
func InsuranceInsurerEQ(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInsuranceInsurer), v))
	})
}

// InsuranceInsurerNEQ applies the NEQ predicate on the "insurance_insurer" field.
func InsuranceInsurerNEQ(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInsuranceInsurer), v))
	})
}

// InsuranceInsurerIn applies the In predicate on the "insurance_insurer" field.
func InsuranceInsurerIn(vs ...string) predicate.Insurance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Insurance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInsuranceInsurer), v...))
	})
}

// InsuranceInsurerNotIn applies the NotIn predicate on the "insurance_insurer" field.
func InsuranceInsurerNotIn(vs ...string) predicate.Insurance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Insurance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInsuranceInsurer), v...))
	})
}

// InsuranceInsurerGT applies the GT predicate on the "insurance_insurer" field.
func InsuranceInsurerGT(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInsuranceInsurer), v))
	})
}

// InsuranceInsurerGTE applies the GTE predicate on the "insurance_insurer" field.
func InsuranceInsurerGTE(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInsuranceInsurer), v))
	})
}

// InsuranceInsurerLT applies the LT predicate on the "insurance_insurer" field.
func InsuranceInsurerLT(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInsuranceInsurer), v))
	})
}

// InsuranceInsurerLTE applies the LTE predicate on the "insurance_insurer" field.
func InsuranceInsurerLTE(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInsuranceInsurer), v))
	})
}

// InsuranceInsurerContains applies the Contains predicate on the "insurance_insurer" field.
func InsuranceInsurerContains(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInsuranceInsurer), v))
	})
}

// InsuranceInsurerHasPrefix applies the HasPrefix predicate on the "insurance_insurer" field.
func InsuranceInsurerHasPrefix(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInsuranceInsurer), v))
	})
}

// InsuranceInsurerHasSuffix applies the HasSuffix predicate on the "insurance_insurer" field.
func InsuranceInsurerHasSuffix(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInsuranceInsurer), v))
	})
}

// InsuranceInsurerEqualFold applies the EqualFold predicate on the "insurance_insurer" field.
func InsuranceInsurerEqualFold(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInsuranceInsurer), v))
	})
}

// InsuranceInsurerContainsFold applies the ContainsFold predicate on the "insurance_insurer" field.
func InsuranceInsurerContainsFold(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInsuranceInsurer), v))
	})
}

// InsuranceAddressEQ applies the EQ predicate on the "insurance_address" field.
func InsuranceAddressEQ(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInsuranceAddress), v))
	})
}

// InsuranceAddressNEQ applies the NEQ predicate on the "insurance_address" field.
func InsuranceAddressNEQ(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInsuranceAddress), v))
	})
}

// InsuranceAddressIn applies the In predicate on the "insurance_address" field.
func InsuranceAddressIn(vs ...string) predicate.Insurance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Insurance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInsuranceAddress), v...))
	})
}

// InsuranceAddressNotIn applies the NotIn predicate on the "insurance_address" field.
func InsuranceAddressNotIn(vs ...string) predicate.Insurance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Insurance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInsuranceAddress), v...))
	})
}

// InsuranceAddressGT applies the GT predicate on the "insurance_address" field.
func InsuranceAddressGT(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInsuranceAddress), v))
	})
}

// InsuranceAddressGTE applies the GTE predicate on the "insurance_address" field.
func InsuranceAddressGTE(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInsuranceAddress), v))
	})
}

// InsuranceAddressLT applies the LT predicate on the "insurance_address" field.
func InsuranceAddressLT(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInsuranceAddress), v))
	})
}

// InsuranceAddressLTE applies the LTE predicate on the "insurance_address" field.
func InsuranceAddressLTE(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInsuranceAddress), v))
	})
}

// InsuranceAddressContains applies the Contains predicate on the "insurance_address" field.
func InsuranceAddressContains(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInsuranceAddress), v))
	})
}

// InsuranceAddressHasPrefix applies the HasPrefix predicate on the "insurance_address" field.
func InsuranceAddressHasPrefix(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInsuranceAddress), v))
	})
}

// InsuranceAddressHasSuffix applies the HasSuffix predicate on the "insurance_address" field.
func InsuranceAddressHasSuffix(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInsuranceAddress), v))
	})
}

// InsuranceAddressEqualFold applies the EqualFold predicate on the "insurance_address" field.
func InsuranceAddressEqualFold(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInsuranceAddress), v))
	})
}

// InsuranceAddressContainsFold applies the ContainsFold predicate on the "insurance_address" field.
func InsuranceAddressContainsFold(v string) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInsuranceAddress), v))
	})
}

// InsuranceTimeBuyEQ applies the EQ predicate on the "insurance_time_buy" field.
func InsuranceTimeBuyEQ(v time.Time) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInsuranceTimeBuy), v))
	})
}

// InsuranceTimeBuyNEQ applies the NEQ predicate on the "insurance_time_buy" field.
func InsuranceTimeBuyNEQ(v time.Time) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInsuranceTimeBuy), v))
	})
}

// InsuranceTimeBuyIn applies the In predicate on the "insurance_time_buy" field.
func InsuranceTimeBuyIn(vs ...time.Time) predicate.Insurance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Insurance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInsuranceTimeBuy), v...))
	})
}

// InsuranceTimeBuyNotIn applies the NotIn predicate on the "insurance_time_buy" field.
func InsuranceTimeBuyNotIn(vs ...time.Time) predicate.Insurance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Insurance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInsuranceTimeBuy), v...))
	})
}

// InsuranceTimeBuyGT applies the GT predicate on the "insurance_time_buy" field.
func InsuranceTimeBuyGT(v time.Time) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInsuranceTimeBuy), v))
	})
}

// InsuranceTimeBuyGTE applies the GTE predicate on the "insurance_time_buy" field.
func InsuranceTimeBuyGTE(v time.Time) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInsuranceTimeBuy), v))
	})
}

// InsuranceTimeBuyLT applies the LT predicate on the "insurance_time_buy" field.
func InsuranceTimeBuyLT(v time.Time) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInsuranceTimeBuy), v))
	})
}

// InsuranceTimeBuyLTE applies the LTE predicate on the "insurance_time_buy" field.
func InsuranceTimeBuyLTE(v time.Time) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInsuranceTimeBuy), v))
	})
}

// InsuranceTimeFirstpayEQ applies the EQ predicate on the "insurance_time_firstpay" field.
func InsuranceTimeFirstpayEQ(v time.Time) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInsuranceTimeFirstpay), v))
	})
}

// InsuranceTimeFirstpayNEQ applies the NEQ predicate on the "insurance_time_firstpay" field.
func InsuranceTimeFirstpayNEQ(v time.Time) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInsuranceTimeFirstpay), v))
	})
}

// InsuranceTimeFirstpayIn applies the In predicate on the "insurance_time_firstpay" field.
func InsuranceTimeFirstpayIn(vs ...time.Time) predicate.Insurance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Insurance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInsuranceTimeFirstpay), v...))
	})
}

// InsuranceTimeFirstpayNotIn applies the NotIn predicate on the "insurance_time_firstpay" field.
func InsuranceTimeFirstpayNotIn(vs ...time.Time) predicate.Insurance {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Insurance(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInsuranceTimeFirstpay), v...))
	})
}

// InsuranceTimeFirstpayGT applies the GT predicate on the "insurance_time_firstpay" field.
func InsuranceTimeFirstpayGT(v time.Time) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInsuranceTimeFirstpay), v))
	})
}

// InsuranceTimeFirstpayGTE applies the GTE predicate on the "insurance_time_firstpay" field.
func InsuranceTimeFirstpayGTE(v time.Time) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInsuranceTimeFirstpay), v))
	})
}

// InsuranceTimeFirstpayLT applies the LT predicate on the "insurance_time_firstpay" field.
func InsuranceTimeFirstpayLT(v time.Time) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInsuranceTimeFirstpay), v))
	})
}

// InsuranceTimeFirstpayLTE applies the LTE predicate on the "insurance_time_firstpay" field.
func InsuranceTimeFirstpayLTE(v time.Time) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInsuranceTimeFirstpay), v))
	})
}

// HasMember applies the HasEdge predicate on the "Member" edge.
func HasMember() predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MemberTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MemberTable, MemberColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemberWith applies the HasEdge predicate on the "Member" edge with a given conditions (other predicates).
func HasMemberWith(preds ...predicate.Member) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MemberInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MemberTable, MemberColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHospital applies the HasEdge predicate on the "Hospital" edge.
func HasHospital() predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HospitalTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HospitalTable, HospitalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHospitalWith applies the HasEdge predicate on the "Hospital" edge with a given conditions (other predicates).
func HasHospitalWith(preds ...predicate.Hospital) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HospitalInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HospitalTable, HospitalColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOfficer applies the HasEdge predicate on the "Officer" edge.
func HasOfficer() predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OfficerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OfficerTable, OfficerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOfficerWith applies the HasEdge predicate on the "Officer" edge with a given conditions (other predicates).
func HasOfficerWith(preds ...predicate.Officer) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OfficerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OfficerTable, OfficerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProduct applies the HasEdge predicate on the "Product" edge.
func HasProduct() predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "Product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.Product) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInsurancePayment applies the HasEdge predicate on the "insurance_payment" edge.
func HasInsurancePayment() predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InsurancePaymentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InsurancePaymentTable, InsurancePaymentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInsurancePaymentWith applies the HasEdge predicate on the "insurance_payment" edge with a given conditions (other predicates).
func HasInsurancePaymentWith(preds ...predicate.Payment) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InsurancePaymentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InsurancePaymentTable, InsurancePaymentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Insurance) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Insurance) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Insurance) predicate.Insurance {
	return predicate.Insurance(func(s *sql.Selector) {
		p(s.Not())
	})
}
