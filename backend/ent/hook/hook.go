// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/sut63/team05/ent"
)

// The AmountpaidFunc type is an adapter to allow the use of ordinary
// function as Amountpaid mutator.
type AmountpaidFunc func(context.Context, *ent.AmountpaidMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AmountpaidFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AmountpaidMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AmountpaidMutation", m)
	}
	return f(ctx, mv)
}

// The BankFunc type is an adapter to allow the use of ordinary
// function as Bank mutator.
type BankFunc func(context.Context, *ent.BankMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BankFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BankMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BankMutation", m)
	}
	return f(ctx, mv)
}

// The CategoryFunc type is an adapter to allow the use of ordinary
// function as Category mutator.
type CategoryFunc func(context.Context, *ent.CategoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CategoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CategoryMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CategoryMutation", m)
	}
	return f(ctx, mv)
}

// The GenderFunc type is an adapter to allow the use of ordinary
// function as Gender mutator.
type GenderFunc func(context.Context, *ent.GenderMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GenderFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GenderMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GenderMutation", m)
	}
	return f(ctx, mv)
}

// The GroupOfAgeFunc type is an adapter to allow the use of ordinary
// function as GroupOfAge mutator.
type GroupOfAgeFunc func(context.Context, *ent.GroupOfAgeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GroupOfAgeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GroupOfAgeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GroupOfAgeMutation", m)
	}
	return f(ctx, mv)
}

// The HospitalFunc type is an adapter to allow the use of ordinary
// function as Hospital mutator.
type HospitalFunc func(context.Context, *ent.HospitalMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f HospitalFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.HospitalMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.HospitalMutation", m)
	}
	return f(ctx, mv)
}

// The InquiryFunc type is an adapter to allow the use of ordinary
// function as Inquiry mutator.
type InquiryFunc func(context.Context, *ent.InquiryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f InquiryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.InquiryMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.InquiryMutation", m)
	}
	return f(ctx, mv)
}

// The InsuranceFunc type is an adapter to allow the use of ordinary
// function as Insurance mutator.
type InsuranceFunc func(context.Context, *ent.InsuranceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f InsuranceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.InsuranceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.InsuranceMutation", m)
	}
	return f(ctx, mv)
}

// The MemberFunc type is an adapter to allow the use of ordinary
// function as Member mutator.
type MemberFunc func(context.Context, *ent.MemberMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MemberFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MemberMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MemberMutation", m)
	}
	return f(ctx, mv)
}

// The MoneyTransferFunc type is an adapter to allow the use of ordinary
// function as MoneyTransfer mutator.
type MoneyTransferFunc func(context.Context, *ent.MoneyTransferMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MoneyTransferFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MoneyTransferMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MoneyTransferMutation", m)
	}
	return f(ctx, mv)
}

// The OfficerFunc type is an adapter to allow the use of ordinary
// function as Officer mutator.
type OfficerFunc func(context.Context, *ent.OfficerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OfficerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OfficerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OfficerMutation", m)
	}
	return f(ctx, mv)
}

// The PaybackFunc type is an adapter to allow the use of ordinary
// function as Payback mutator.
type PaybackFunc func(context.Context, *ent.PaybackMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PaybackFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PaybackMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PaybackMutation", m)
	}
	return f(ctx, mv)
}

// The PaymentFunc type is an adapter to allow the use of ordinary
// function as Payment mutator.
type PaymentFunc func(context.Context, *ent.PaymentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PaymentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PaymentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PaymentMutation", m)
	}
	return f(ctx, mv)
}

// The ProductFunc type is an adapter to allow the use of ordinary
// function as Product mutator.
type ProductFunc func(context.Context, *ent.ProductMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ProductFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.ProductMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ProductMutation", m)
	}
	return f(ctx, mv)
}

// The RecordinsuranceFunc type is an adapter to allow the use of ordinary
// function as Recordinsurance mutator.
type RecordinsuranceFunc func(context.Context, *ent.RecordinsuranceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RecordinsuranceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RecordinsuranceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RecordinsuranceMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	Hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(_ context.Context, m ent.Mutation) (ent.Value, error) {
			return nil, fmt.Errorf("%s operation is not allowed", m.Op())
		})
	}
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
