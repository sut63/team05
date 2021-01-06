// Code generated by entc, DO NOT EDIT.

package payment

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team05/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AccountName applies equality check predicate on the "account_name" field. It's identical to AccountNameEQ.
func AccountName(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountName), v))
	})
}

// AccountNumber applies equality check predicate on the "account_number" field. It's identical to AccountNumberEQ.
func AccountNumber(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountNumber), v))
	})
}

// TransferTime applies equality check predicate on the "transfer_time" field. It's identical to TransferTimeEQ.
func TransferTime(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransferTime), v))
	})
}

// AccountNameEQ applies the EQ predicate on the "account_name" field.
func AccountNameEQ(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountName), v))
	})
}

// AccountNameNEQ applies the NEQ predicate on the "account_name" field.
func AccountNameNEQ(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccountName), v))
	})
}

// AccountNameIn applies the In predicate on the "account_name" field.
func AccountNameIn(vs ...string) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAccountName), v...))
	})
}

// AccountNameNotIn applies the NotIn predicate on the "account_name" field.
func AccountNameNotIn(vs ...string) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAccountName), v...))
	})
}

// AccountNameGT applies the GT predicate on the "account_name" field.
func AccountNameGT(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccountName), v))
	})
}

// AccountNameGTE applies the GTE predicate on the "account_name" field.
func AccountNameGTE(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccountName), v))
	})
}

// AccountNameLT applies the LT predicate on the "account_name" field.
func AccountNameLT(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccountName), v))
	})
}

// AccountNameLTE applies the LTE predicate on the "account_name" field.
func AccountNameLTE(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccountName), v))
	})
}

// AccountNameContains applies the Contains predicate on the "account_name" field.
func AccountNameContains(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccountName), v))
	})
}

// AccountNameHasPrefix applies the HasPrefix predicate on the "account_name" field.
func AccountNameHasPrefix(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccountName), v))
	})
}

// AccountNameHasSuffix applies the HasSuffix predicate on the "account_name" field.
func AccountNameHasSuffix(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccountName), v))
	})
}

// AccountNameEqualFold applies the EqualFold predicate on the "account_name" field.
func AccountNameEqualFold(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccountName), v))
	})
}

// AccountNameContainsFold applies the ContainsFold predicate on the "account_name" field.
func AccountNameContainsFold(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccountName), v))
	})
}

// AccountNumberEQ applies the EQ predicate on the "account_number" field.
func AccountNumberEQ(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberNEQ applies the NEQ predicate on the "account_number" field.
func AccountNumberNEQ(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberIn applies the In predicate on the "account_number" field.
func AccountNumberIn(vs ...string) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAccountNumber), v...))
	})
}

// AccountNumberNotIn applies the NotIn predicate on the "account_number" field.
func AccountNumberNotIn(vs ...string) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAccountNumber), v...))
	})
}

// AccountNumberGT applies the GT predicate on the "account_number" field.
func AccountNumberGT(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberGTE applies the GTE predicate on the "account_number" field.
func AccountNumberGTE(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberLT applies the LT predicate on the "account_number" field.
func AccountNumberLT(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberLTE applies the LTE predicate on the "account_number" field.
func AccountNumberLTE(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberContains applies the Contains predicate on the "account_number" field.
func AccountNumberContains(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberHasPrefix applies the HasPrefix predicate on the "account_number" field.
func AccountNumberHasPrefix(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberHasSuffix applies the HasSuffix predicate on the "account_number" field.
func AccountNumberHasSuffix(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberEqualFold applies the EqualFold predicate on the "account_number" field.
func AccountNumberEqualFold(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAccountNumber), v))
	})
}

// AccountNumberContainsFold applies the ContainsFold predicate on the "account_number" field.
func AccountNumberContainsFold(v string) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAccountNumber), v))
	})
}

// TransferTimeEQ applies the EQ predicate on the "transfer_time" field.
func TransferTimeEQ(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransferTime), v))
	})
}

// TransferTimeNEQ applies the NEQ predicate on the "transfer_time" field.
func TransferTimeNEQ(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTransferTime), v))
	})
}

// TransferTimeIn applies the In predicate on the "transfer_time" field.
func TransferTimeIn(vs ...time.Time) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTransferTime), v...))
	})
}

// TransferTimeNotIn applies the NotIn predicate on the "transfer_time" field.
func TransferTimeNotIn(vs ...time.Time) predicate.Payment {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Payment(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTransferTime), v...))
	})
}

// TransferTimeGT applies the GT predicate on the "transfer_time" field.
func TransferTimeGT(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTransferTime), v))
	})
}

// TransferTimeGTE applies the GTE predicate on the "transfer_time" field.
func TransferTimeGTE(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTransferTime), v))
	})
}

// TransferTimeLT applies the LT predicate on the "transfer_time" field.
func TransferTimeLT(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTransferTime), v))
	})
}

// TransferTimeLTE applies the LTE predicate on the "transfer_time" field.
func TransferTimeLTE(v time.Time) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTransferTime), v))
	})
}

// HasInsurance applies the HasEdge predicate on the "Insurance" edge.
func HasInsurance() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InsuranceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InsuranceTable, InsuranceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInsuranceWith applies the HasEdge predicate on the "Insurance" edge with a given conditions (other predicates).
func HasInsuranceWith(preds ...predicate.Insurance) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InsuranceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InsuranceTable, InsuranceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMoneyTransfer applies the HasEdge predicate on the "MoneyTransfer" edge.
func HasMoneyTransfer() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MoneyTransferTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MoneyTransferTable, MoneyTransferColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMoneyTransferWith applies the HasEdge predicate on the "MoneyTransfer" edge with a given conditions (other predicates).
func HasMoneyTransferWith(preds ...predicate.MoneyTransfer) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MoneyTransferInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MoneyTransferTable, MoneyTransferColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBank applies the HasEdge predicate on the "Bank" edge.
func HasBank() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BankTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BankTable, BankColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBankWith applies the HasEdge predicate on the "Bank" edge with a given conditions (other predicates).
func HasBankWith(preds ...predicate.Bank) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BankInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BankTable, BankColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMember applies the HasEdge predicate on the "Member" edge.
func HasMember() predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MemberTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MemberTable, MemberColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemberWith applies the HasEdge predicate on the "Member" edge with a given conditions (other predicates).
func HasMemberWith(preds ...predicate.Member) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
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

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
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
func Not(p predicate.Payment) predicate.Payment {
	return predicate.Payment(func(s *sql.Selector) {
		p(s.Not())
	})
}
