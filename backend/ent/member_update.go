// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team05/ent/insurance"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/payment"
	"github.com/sut63/team05/ent/predicate"
)

// MemberUpdate is the builder for updating Member entities.
type MemberUpdate struct {
	config
	hooks      []Hook
	mutation   *MemberMutation
	predicates []predicate.Member
}

// Where adds a new predicate for the builder.
func (mu *MemberUpdate) Where(ps ...predicate.Member) *MemberUpdate {
	mu.predicates = append(mu.predicates, ps...)
	return mu
}

// SetMemberEmail sets the member_email field.
func (mu *MemberUpdate) SetMemberEmail(s string) *MemberUpdate {
	mu.mutation.SetMemberEmail(s)
	return mu
}

// SetMemberName sets the member_name field.
func (mu *MemberUpdate) SetMemberName(s string) *MemberUpdate {
	mu.mutation.SetMemberName(s)
	return mu
}

// SetMemberPassword sets the member_password field.
func (mu *MemberUpdate) SetMemberPassword(s string) *MemberUpdate {
	mu.mutation.SetMemberPassword(s)
	return mu
}

// AddMemberInsuranceIDs adds the member_insurance edge to Insurance by ids.
func (mu *MemberUpdate) AddMemberInsuranceIDs(ids ...int) *MemberUpdate {
	mu.mutation.AddMemberInsuranceIDs(ids...)
	return mu
}

// AddMemberInsurance adds the member_insurance edges to Insurance.
func (mu *MemberUpdate) AddMemberInsurance(i ...*Insurance) *MemberUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return mu.AddMemberInsuranceIDs(ids...)
}

// AddMemberPaymentIDs adds the member_payment edge to Payment by ids.
func (mu *MemberUpdate) AddMemberPaymentIDs(ids ...int) *MemberUpdate {
	mu.mutation.AddMemberPaymentIDs(ids...)
	return mu
}

// AddMemberPayment adds the member_payment edges to Payment.
func (mu *MemberUpdate) AddMemberPayment(p ...*Payment) *MemberUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.AddMemberPaymentIDs(ids...)
}

// Mutation returns the MemberMutation object of the builder.
func (mu *MemberUpdate) Mutation() *MemberMutation {
	return mu.mutation
}

// RemoveMemberInsuranceIDs removes the member_insurance edge to Insurance by ids.
func (mu *MemberUpdate) RemoveMemberInsuranceIDs(ids ...int) *MemberUpdate {
	mu.mutation.RemoveMemberInsuranceIDs(ids...)
	return mu
}

// RemoveMemberInsurance removes member_insurance edges to Insurance.
func (mu *MemberUpdate) RemoveMemberInsurance(i ...*Insurance) *MemberUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return mu.RemoveMemberInsuranceIDs(ids...)
}

// RemoveMemberPaymentIDs removes the member_payment edge to Payment by ids.
func (mu *MemberUpdate) RemoveMemberPaymentIDs(ids ...int) *MemberUpdate {
	mu.mutation.RemoveMemberPaymentIDs(ids...)
	return mu
}

// RemoveMemberPayment removes member_payment edges to Payment.
func (mu *MemberUpdate) RemoveMemberPayment(p ...*Payment) *MemberUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.RemoveMemberPaymentIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (mu *MemberUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := mu.mutation.MemberEmail(); ok {
		if err := member.MemberEmailValidator(v); err != nil {
			return 0, &ValidationError{Name: "member_email", err: fmt.Errorf("ent: validator failed for field \"member_email\": %w", err)}
		}
	}
	if v, ok := mu.mutation.MemberName(); ok {
		if err := member.MemberNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "member_name", err: fmt.Errorf("ent: validator failed for field \"member_name\": %w", err)}
		}
	}
	if v, ok := mu.mutation.MemberPassword(); ok {
		if err := member.MemberPasswordValidator(v); err != nil {
			return 0, &ValidationError{Name: "member_password", err: fmt.Errorf("ent: validator failed for field \"member_password\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MemberUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MemberUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MemberUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: member.FieldID,
			},
		},
	}
	if ps := mu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.MemberEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldMemberEmail,
		})
	}
	if value, ok := mu.mutation.MemberName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldMemberName,
		})
	}
	if value, ok := mu.mutation.MemberPassword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldMemberPassword,
		})
	}
	if nodes := mu.mutation.RemovedMemberInsuranceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MemberInsuranceTable,
			Columns: []string{member.MemberInsuranceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: insurance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MemberInsuranceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MemberInsuranceTable,
			Columns: []string{member.MemberInsuranceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: insurance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := mu.mutation.RemovedMemberPaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MemberPaymentTable,
			Columns: []string{member.MemberPaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MemberPaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MemberPaymentTable,
			Columns: []string{member.MemberPaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MemberUpdateOne is the builder for updating a single Member entity.
type MemberUpdateOne struct {
	config
	hooks    []Hook
	mutation *MemberMutation
}

// SetMemberEmail sets the member_email field.
func (muo *MemberUpdateOne) SetMemberEmail(s string) *MemberUpdateOne {
	muo.mutation.SetMemberEmail(s)
	return muo
}

// SetMemberName sets the member_name field.
func (muo *MemberUpdateOne) SetMemberName(s string) *MemberUpdateOne {
	muo.mutation.SetMemberName(s)
	return muo
}

// SetMemberPassword sets the member_password field.
func (muo *MemberUpdateOne) SetMemberPassword(s string) *MemberUpdateOne {
	muo.mutation.SetMemberPassword(s)
	return muo
}

// AddMemberInsuranceIDs adds the member_insurance edge to Insurance by ids.
func (muo *MemberUpdateOne) AddMemberInsuranceIDs(ids ...int) *MemberUpdateOne {
	muo.mutation.AddMemberInsuranceIDs(ids...)
	return muo
}

// AddMemberInsurance adds the member_insurance edges to Insurance.
func (muo *MemberUpdateOne) AddMemberInsurance(i ...*Insurance) *MemberUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return muo.AddMemberInsuranceIDs(ids...)
}

// AddMemberPaymentIDs adds the member_payment edge to Payment by ids.
func (muo *MemberUpdateOne) AddMemberPaymentIDs(ids ...int) *MemberUpdateOne {
	muo.mutation.AddMemberPaymentIDs(ids...)
	return muo
}

// AddMemberPayment adds the member_payment edges to Payment.
func (muo *MemberUpdateOne) AddMemberPayment(p ...*Payment) *MemberUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.AddMemberPaymentIDs(ids...)
}

// Mutation returns the MemberMutation object of the builder.
func (muo *MemberUpdateOne) Mutation() *MemberMutation {
	return muo.mutation
}

// RemoveMemberInsuranceIDs removes the member_insurance edge to Insurance by ids.
func (muo *MemberUpdateOne) RemoveMemberInsuranceIDs(ids ...int) *MemberUpdateOne {
	muo.mutation.RemoveMemberInsuranceIDs(ids...)
	return muo
}

// RemoveMemberInsurance removes member_insurance edges to Insurance.
func (muo *MemberUpdateOne) RemoveMemberInsurance(i ...*Insurance) *MemberUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return muo.RemoveMemberInsuranceIDs(ids...)
}

// RemoveMemberPaymentIDs removes the member_payment edge to Payment by ids.
func (muo *MemberUpdateOne) RemoveMemberPaymentIDs(ids ...int) *MemberUpdateOne {
	muo.mutation.RemoveMemberPaymentIDs(ids...)
	return muo
}

// RemoveMemberPayment removes member_payment edges to Payment.
func (muo *MemberUpdateOne) RemoveMemberPayment(p ...*Payment) *MemberUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.RemoveMemberPaymentIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (muo *MemberUpdateOne) Save(ctx context.Context) (*Member, error) {
	if v, ok := muo.mutation.MemberEmail(); ok {
		if err := member.MemberEmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "member_email", err: fmt.Errorf("ent: validator failed for field \"member_email\": %w", err)}
		}
	}
	if v, ok := muo.mutation.MemberName(); ok {
		if err := member.MemberNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "member_name", err: fmt.Errorf("ent: validator failed for field \"member_name\": %w", err)}
		}
	}
	if v, ok := muo.mutation.MemberPassword(); ok {
		if err := member.MemberPasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "member_password", err: fmt.Errorf("ent: validator failed for field \"member_password\": %w", err)}
		}
	}

	var (
		err  error
		node *Member
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MemberUpdateOne) SaveX(ctx context.Context) *Member {
	m, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// Exec executes the query on the entity.
func (muo *MemberUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MemberUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MemberUpdateOne) sqlSave(ctx context.Context) (m *Member, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: member.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Member.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := muo.mutation.MemberEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldMemberEmail,
		})
	}
	if value, ok := muo.mutation.MemberName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldMemberName,
		})
	}
	if value, ok := muo.mutation.MemberPassword(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: member.FieldMemberPassword,
		})
	}
	if nodes := muo.mutation.RemovedMemberInsuranceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MemberInsuranceTable,
			Columns: []string{member.MemberInsuranceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: insurance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MemberInsuranceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MemberInsuranceTable,
			Columns: []string{member.MemberInsuranceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: insurance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := muo.mutation.RemovedMemberPaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MemberPaymentTable,
			Columns: []string{member.MemberPaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MemberPaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MemberPaymentTable,
			Columns: []string{member.MemberPaymentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	m = &Member{config: muo.config}
	_spec.Assign = m.assignValues
	_spec.ScanValues = m.scanValues()
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return m, nil
}
