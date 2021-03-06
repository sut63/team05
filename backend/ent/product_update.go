// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team05/ent/gender"
	"github.com/sut63/team05/ent/groupofage"
	"github.com/sut63/team05/ent/inquiry"
	"github.com/sut63/team05/ent/insurance"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/payback"
	"github.com/sut63/team05/ent/predicate"
	"github.com/sut63/team05/ent/product"
	"github.com/sut63/team05/ent/recordinsurance"
)

// ProductUpdate is the builder for updating Product entities.
type ProductUpdate struct {
	config
	hooks      []Hook
	mutation   *ProductMutation
	predicates []predicate.Product
}

// Where adds a new predicate for the builder.
func (pu *ProductUpdate) Where(ps ...predicate.Product) *ProductUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetProductName sets the product_name field.
func (pu *ProductUpdate) SetProductName(s string) *ProductUpdate {
	pu.mutation.SetProductName(s)
	return pu
}

// SetProductPrice sets the product_price field.
func (pu *ProductUpdate) SetProductPrice(i int) *ProductUpdate {
	pu.mutation.ResetProductPrice()
	pu.mutation.SetProductPrice(i)
	return pu
}

// AddProductPrice adds i to product_price.
func (pu *ProductUpdate) AddProductPrice(i int) *ProductUpdate {
	pu.mutation.AddProductPrice(i)
	return pu
}

// SetProductTime sets the product_time field.
func (pu *ProductUpdate) SetProductTime(i int) *ProductUpdate {
	pu.mutation.ResetProductTime()
	pu.mutation.SetProductTime(i)
	return pu
}

// AddProductTime adds i to product_time.
func (pu *ProductUpdate) AddProductTime(i int) *ProductUpdate {
	pu.mutation.AddProductTime(i)
	return pu
}

// SetProductPaymentOfYear sets the product_payment_of_year field.
func (pu *ProductUpdate) SetProductPaymentOfYear(i int) *ProductUpdate {
	pu.mutation.ResetProductPaymentOfYear()
	pu.mutation.SetProductPaymentOfYear(i)
	return pu
}

// AddProductPaymentOfYear adds i to product_payment_of_year.
func (pu *ProductUpdate) AddProductPaymentOfYear(i int) *ProductUpdate {
	pu.mutation.AddProductPaymentOfYear(i)
	return pu
}

// SetGenderID sets the gender edge to Gender by id.
func (pu *ProductUpdate) SetGenderID(id int) *ProductUpdate {
	pu.mutation.SetGenderID(id)
	return pu
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (pu *ProductUpdate) SetNillableGenderID(id *int) *ProductUpdate {
	if id != nil {
		pu = pu.SetGenderID(*id)
	}
	return pu
}

// SetGender sets the gender edge to Gender.
func (pu *ProductUpdate) SetGender(g *Gender) *ProductUpdate {
	return pu.SetGenderID(g.ID)
}

// SetGroupofageID sets the groupofage edge to GroupOfAge by id.
func (pu *ProductUpdate) SetGroupofageID(id int) *ProductUpdate {
	pu.mutation.SetGroupofageID(id)
	return pu
}

// SetNillableGroupofageID sets the groupofage edge to GroupOfAge by id if the given value is not nil.
func (pu *ProductUpdate) SetNillableGroupofageID(id *int) *ProductUpdate {
	if id != nil {
		pu = pu.SetGroupofageID(*id)
	}
	return pu
}

// SetGroupofage sets the groupofage edge to GroupOfAge.
func (pu *ProductUpdate) SetGroupofage(g *GroupOfAge) *ProductUpdate {
	return pu.SetGroupofageID(g.ID)
}

// SetOfficerID sets the officer edge to Officer by id.
func (pu *ProductUpdate) SetOfficerID(id int) *ProductUpdate {
	pu.mutation.SetOfficerID(id)
	return pu
}

// SetNillableOfficerID sets the officer edge to Officer by id if the given value is not nil.
func (pu *ProductUpdate) SetNillableOfficerID(id *int) *ProductUpdate {
	if id != nil {
		pu = pu.SetOfficerID(*id)
	}
	return pu
}

// SetOfficer sets the officer edge to Officer.
func (pu *ProductUpdate) SetOfficer(o *Officer) *ProductUpdate {
	return pu.SetOfficerID(o.ID)
}

// AddProductInsuranceIDs adds the product_insurance edge to Insurance by ids.
func (pu *ProductUpdate) AddProductInsuranceIDs(ids ...int) *ProductUpdate {
	pu.mutation.AddProductInsuranceIDs(ids...)
	return pu
}

// AddProductInsurance adds the product_insurance edges to Insurance.
func (pu *ProductUpdate) AddProductInsurance(i ...*Insurance) *ProductUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pu.AddProductInsuranceIDs(ids...)
}

// AddProductInquiryIDs adds the product_inquiry edge to Inquiry by ids.
func (pu *ProductUpdate) AddProductInquiryIDs(ids ...int) *ProductUpdate {
	pu.mutation.AddProductInquiryIDs(ids...)
	return pu
}

// AddProductInquiry adds the product_inquiry edges to Inquiry.
func (pu *ProductUpdate) AddProductInquiry(i ...*Inquiry) *ProductUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pu.AddProductInquiryIDs(ids...)
}

// AddProductPaybackIDs adds the product_payback edge to Payback by ids.
func (pu *ProductUpdate) AddProductPaybackIDs(ids ...int) *ProductUpdate {
	pu.mutation.AddProductPaybackIDs(ids...)
	return pu
}

// AddProductPayback adds the product_payback edges to Payback.
func (pu *ProductUpdate) AddProductPayback(p ...*Payback) *ProductUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddProductPaybackIDs(ids...)
}

// AddProductRecordinsuranceIDs adds the product_recordinsurance edge to Recordinsurance by ids.
func (pu *ProductUpdate) AddProductRecordinsuranceIDs(ids ...int) *ProductUpdate {
	pu.mutation.AddProductRecordinsuranceIDs(ids...)
	return pu
}

// AddProductRecordinsurance adds the product_recordinsurance edges to Recordinsurance.
func (pu *ProductUpdate) AddProductRecordinsurance(r ...*Recordinsurance) *ProductUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.AddProductRecordinsuranceIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (pu *ProductUpdate) Mutation() *ProductMutation {
	return pu.mutation
}

// ClearGender clears the gender edge to Gender.
func (pu *ProductUpdate) ClearGender() *ProductUpdate {
	pu.mutation.ClearGender()
	return pu
}

// ClearGroupofage clears the groupofage edge to GroupOfAge.
func (pu *ProductUpdate) ClearGroupofage() *ProductUpdate {
	pu.mutation.ClearGroupofage()
	return pu
}

// ClearOfficer clears the officer edge to Officer.
func (pu *ProductUpdate) ClearOfficer() *ProductUpdate {
	pu.mutation.ClearOfficer()
	return pu
}

// RemoveProductInsuranceIDs removes the product_insurance edge to Insurance by ids.
func (pu *ProductUpdate) RemoveProductInsuranceIDs(ids ...int) *ProductUpdate {
	pu.mutation.RemoveProductInsuranceIDs(ids...)
	return pu
}

// RemoveProductInsurance removes product_insurance edges to Insurance.
func (pu *ProductUpdate) RemoveProductInsurance(i ...*Insurance) *ProductUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pu.RemoveProductInsuranceIDs(ids...)
}

// RemoveProductInquiryIDs removes the product_inquiry edge to Inquiry by ids.
func (pu *ProductUpdate) RemoveProductInquiryIDs(ids ...int) *ProductUpdate {
	pu.mutation.RemoveProductInquiryIDs(ids...)
	return pu
}

// RemoveProductInquiry removes product_inquiry edges to Inquiry.
func (pu *ProductUpdate) RemoveProductInquiry(i ...*Inquiry) *ProductUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pu.RemoveProductInquiryIDs(ids...)
}

// RemoveProductPaybackIDs removes the product_payback edge to Payback by ids.
func (pu *ProductUpdate) RemoveProductPaybackIDs(ids ...int) *ProductUpdate {
	pu.mutation.RemoveProductPaybackIDs(ids...)
	return pu
}

// RemoveProductPayback removes product_payback edges to Payback.
func (pu *ProductUpdate) RemoveProductPayback(p ...*Payback) *ProductUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveProductPaybackIDs(ids...)
}

// RemoveProductRecordinsuranceIDs removes the product_recordinsurance edge to Recordinsurance by ids.
func (pu *ProductUpdate) RemoveProductRecordinsuranceIDs(ids ...int) *ProductUpdate {
	pu.mutation.RemoveProductRecordinsuranceIDs(ids...)
	return pu
}

// RemoveProductRecordinsurance removes product_recordinsurance edges to Recordinsurance.
func (pu *ProductUpdate) RemoveProductRecordinsurance(r ...*Recordinsurance) *ProductUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.RemoveProductRecordinsuranceIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *ProductUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := pu.mutation.ProductName(); ok {
		if err := product.ProductNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "product_name", err: fmt.Errorf("ent: validator failed for field \"product_name\": %w", err)}
		}
	}
	if v, ok := pu.mutation.ProductPrice(); ok {
		if err := product.ProductPriceValidator(v); err != nil {
			return 0, &ValidationError{Name: "product_price", err: fmt.Errorf("ent: validator failed for field \"product_price\": %w", err)}
		}
	}
	if v, ok := pu.mutation.ProductTime(); ok {
		if err := product.ProductTimeValidator(v); err != nil {
			return 0, &ValidationError{Name: "product_time", err: fmt.Errorf("ent: validator failed for field \"product_time\": %w", err)}
		}
	}
	if v, ok := pu.mutation.ProductPaymentOfYear(); ok {
		if err := product.ProductPaymentOfYearValidator(v); err != nil {
			return 0, &ValidationError{Name: "product_payment_of_year", err: fmt.Errorf("ent: validator failed for field \"product_payment_of_year\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProductUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProductUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProductUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: product.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.ProductName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldProductName,
		})
	}
	if value, ok := pu.mutation.ProductPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldProductPrice,
		})
	}
	if value, ok := pu.mutation.AddedProductPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldProductPrice,
		})
	}
	if value, ok := pu.mutation.ProductTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldProductTime,
		})
	}
	if value, ok := pu.mutation.AddedProductTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldProductTime,
		})
	}
	if value, ok := pu.mutation.ProductPaymentOfYear(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldProductPaymentOfYear,
		})
	}
	if value, ok := pu.mutation.AddedProductPaymentOfYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldProductPaymentOfYear,
		})
	}
	if pu.mutation.GenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.GenderTable,
			Columns: []string{product.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.GenderTable,
			Columns: []string{product.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.GroupofageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.GroupofageTable,
			Columns: []string{product.GroupofageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: groupofage.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.GroupofageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.GroupofageTable,
			Columns: []string{product.GroupofageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: groupofage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.OfficerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.OfficerTable,
			Columns: []string{product.OfficerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: officer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OfficerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.OfficerTable,
			Columns: []string{product.OfficerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: officer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := pu.mutation.RemovedProductInsuranceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductInsuranceTable,
			Columns: []string{product.ProductInsuranceColumn},
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
	if nodes := pu.mutation.ProductInsuranceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductInsuranceTable,
			Columns: []string{product.ProductInsuranceColumn},
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
	if nodes := pu.mutation.RemovedProductInquiryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductInquiryTable,
			Columns: []string{product.ProductInquiryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inquiry.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProductInquiryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductInquiryTable,
			Columns: []string{product.ProductInquiryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inquiry.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := pu.mutation.RemovedProductPaybackIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductPaybackTable,
			Columns: []string{product.ProductPaybackColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payback.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProductPaybackIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductPaybackTable,
			Columns: []string{product.ProductPaybackColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payback.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := pu.mutation.RemovedProductRecordinsuranceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductRecordinsuranceTable,
			Columns: []string{product.ProductRecordinsuranceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: recordinsurance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ProductRecordinsuranceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductRecordinsuranceTable,
			Columns: []string{product.ProductRecordinsuranceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: recordinsurance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProductUpdateOne is the builder for updating a single Product entity.
type ProductUpdateOne struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// SetProductName sets the product_name field.
func (puo *ProductUpdateOne) SetProductName(s string) *ProductUpdateOne {
	puo.mutation.SetProductName(s)
	return puo
}

// SetProductPrice sets the product_price field.
func (puo *ProductUpdateOne) SetProductPrice(i int) *ProductUpdateOne {
	puo.mutation.ResetProductPrice()
	puo.mutation.SetProductPrice(i)
	return puo
}

// AddProductPrice adds i to product_price.
func (puo *ProductUpdateOne) AddProductPrice(i int) *ProductUpdateOne {
	puo.mutation.AddProductPrice(i)
	return puo
}

// SetProductTime sets the product_time field.
func (puo *ProductUpdateOne) SetProductTime(i int) *ProductUpdateOne {
	puo.mutation.ResetProductTime()
	puo.mutation.SetProductTime(i)
	return puo
}

// AddProductTime adds i to product_time.
func (puo *ProductUpdateOne) AddProductTime(i int) *ProductUpdateOne {
	puo.mutation.AddProductTime(i)
	return puo
}

// SetProductPaymentOfYear sets the product_payment_of_year field.
func (puo *ProductUpdateOne) SetProductPaymentOfYear(i int) *ProductUpdateOne {
	puo.mutation.ResetProductPaymentOfYear()
	puo.mutation.SetProductPaymentOfYear(i)
	return puo
}

// AddProductPaymentOfYear adds i to product_payment_of_year.
func (puo *ProductUpdateOne) AddProductPaymentOfYear(i int) *ProductUpdateOne {
	puo.mutation.AddProductPaymentOfYear(i)
	return puo
}

// SetGenderID sets the gender edge to Gender by id.
func (puo *ProductUpdateOne) SetGenderID(id int) *ProductUpdateOne {
	puo.mutation.SetGenderID(id)
	return puo
}

// SetNillableGenderID sets the gender edge to Gender by id if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableGenderID(id *int) *ProductUpdateOne {
	if id != nil {
		puo = puo.SetGenderID(*id)
	}
	return puo
}

// SetGender sets the gender edge to Gender.
func (puo *ProductUpdateOne) SetGender(g *Gender) *ProductUpdateOne {
	return puo.SetGenderID(g.ID)
}

// SetGroupofageID sets the groupofage edge to GroupOfAge by id.
func (puo *ProductUpdateOne) SetGroupofageID(id int) *ProductUpdateOne {
	puo.mutation.SetGroupofageID(id)
	return puo
}

// SetNillableGroupofageID sets the groupofage edge to GroupOfAge by id if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableGroupofageID(id *int) *ProductUpdateOne {
	if id != nil {
		puo = puo.SetGroupofageID(*id)
	}
	return puo
}

// SetGroupofage sets the groupofage edge to GroupOfAge.
func (puo *ProductUpdateOne) SetGroupofage(g *GroupOfAge) *ProductUpdateOne {
	return puo.SetGroupofageID(g.ID)
}

// SetOfficerID sets the officer edge to Officer by id.
func (puo *ProductUpdateOne) SetOfficerID(id int) *ProductUpdateOne {
	puo.mutation.SetOfficerID(id)
	return puo
}

// SetNillableOfficerID sets the officer edge to Officer by id if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableOfficerID(id *int) *ProductUpdateOne {
	if id != nil {
		puo = puo.SetOfficerID(*id)
	}
	return puo
}

// SetOfficer sets the officer edge to Officer.
func (puo *ProductUpdateOne) SetOfficer(o *Officer) *ProductUpdateOne {
	return puo.SetOfficerID(o.ID)
}

// AddProductInsuranceIDs adds the product_insurance edge to Insurance by ids.
func (puo *ProductUpdateOne) AddProductInsuranceIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.AddProductInsuranceIDs(ids...)
	return puo
}

// AddProductInsurance adds the product_insurance edges to Insurance.
func (puo *ProductUpdateOne) AddProductInsurance(i ...*Insurance) *ProductUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return puo.AddProductInsuranceIDs(ids...)
}

// AddProductInquiryIDs adds the product_inquiry edge to Inquiry by ids.
func (puo *ProductUpdateOne) AddProductInquiryIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.AddProductInquiryIDs(ids...)
	return puo
}

// AddProductInquiry adds the product_inquiry edges to Inquiry.
func (puo *ProductUpdateOne) AddProductInquiry(i ...*Inquiry) *ProductUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return puo.AddProductInquiryIDs(ids...)
}

// AddProductPaybackIDs adds the product_payback edge to Payback by ids.
func (puo *ProductUpdateOne) AddProductPaybackIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.AddProductPaybackIDs(ids...)
	return puo
}

// AddProductPayback adds the product_payback edges to Payback.
func (puo *ProductUpdateOne) AddProductPayback(p ...*Payback) *ProductUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddProductPaybackIDs(ids...)
}

// AddProductRecordinsuranceIDs adds the product_recordinsurance edge to Recordinsurance by ids.
func (puo *ProductUpdateOne) AddProductRecordinsuranceIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.AddProductRecordinsuranceIDs(ids...)
	return puo
}

// AddProductRecordinsurance adds the product_recordinsurance edges to Recordinsurance.
func (puo *ProductUpdateOne) AddProductRecordinsurance(r ...*Recordinsurance) *ProductUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.AddProductRecordinsuranceIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (puo *ProductUpdateOne) Mutation() *ProductMutation {
	return puo.mutation
}

// ClearGender clears the gender edge to Gender.
func (puo *ProductUpdateOne) ClearGender() *ProductUpdateOne {
	puo.mutation.ClearGender()
	return puo
}

// ClearGroupofage clears the groupofage edge to GroupOfAge.
func (puo *ProductUpdateOne) ClearGroupofage() *ProductUpdateOne {
	puo.mutation.ClearGroupofage()
	return puo
}

// ClearOfficer clears the officer edge to Officer.
func (puo *ProductUpdateOne) ClearOfficer() *ProductUpdateOne {
	puo.mutation.ClearOfficer()
	return puo
}

// RemoveProductInsuranceIDs removes the product_insurance edge to Insurance by ids.
func (puo *ProductUpdateOne) RemoveProductInsuranceIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.RemoveProductInsuranceIDs(ids...)
	return puo
}

// RemoveProductInsurance removes product_insurance edges to Insurance.
func (puo *ProductUpdateOne) RemoveProductInsurance(i ...*Insurance) *ProductUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return puo.RemoveProductInsuranceIDs(ids...)
}

// RemoveProductInquiryIDs removes the product_inquiry edge to Inquiry by ids.
func (puo *ProductUpdateOne) RemoveProductInquiryIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.RemoveProductInquiryIDs(ids...)
	return puo
}

// RemoveProductInquiry removes product_inquiry edges to Inquiry.
func (puo *ProductUpdateOne) RemoveProductInquiry(i ...*Inquiry) *ProductUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return puo.RemoveProductInquiryIDs(ids...)
}

// RemoveProductPaybackIDs removes the product_payback edge to Payback by ids.
func (puo *ProductUpdateOne) RemoveProductPaybackIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.RemoveProductPaybackIDs(ids...)
	return puo
}

// RemoveProductPayback removes product_payback edges to Payback.
func (puo *ProductUpdateOne) RemoveProductPayback(p ...*Payback) *ProductUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveProductPaybackIDs(ids...)
}

// RemoveProductRecordinsuranceIDs removes the product_recordinsurance edge to Recordinsurance by ids.
func (puo *ProductUpdateOne) RemoveProductRecordinsuranceIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.RemoveProductRecordinsuranceIDs(ids...)
	return puo
}

// RemoveProductRecordinsurance removes product_recordinsurance edges to Recordinsurance.
func (puo *ProductUpdateOne) RemoveProductRecordinsurance(r ...*Recordinsurance) *ProductUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.RemoveProductRecordinsuranceIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {
	if v, ok := puo.mutation.ProductName(); ok {
		if err := product.ProductNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "product_name", err: fmt.Errorf("ent: validator failed for field \"product_name\": %w", err)}
		}
	}
	if v, ok := puo.mutation.ProductPrice(); ok {
		if err := product.ProductPriceValidator(v); err != nil {
			return nil, &ValidationError{Name: "product_price", err: fmt.Errorf("ent: validator failed for field \"product_price\": %w", err)}
		}
	}
	if v, ok := puo.mutation.ProductTime(); ok {
		if err := product.ProductTimeValidator(v); err != nil {
			return nil, &ValidationError{Name: "product_time", err: fmt.Errorf("ent: validator failed for field \"product_time\": %w", err)}
		}
	}
	if v, ok := puo.mutation.ProductPaymentOfYear(); ok {
		if err := product.ProductPaymentOfYearValidator(v); err != nil {
			return nil, &ValidationError{Name: "product_payment_of_year", err: fmt.Errorf("ent: validator failed for field \"product_payment_of_year\": %w", err)}
		}
	}

	var (
		err  error
		node *Product
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProductUpdateOne) SaveX(ctx context.Context) *Product {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *ProductUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProductUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProductUpdateOne) sqlSave(ctx context.Context) (pr *Product, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: product.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Product.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.ProductName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldProductName,
		})
	}
	if value, ok := puo.mutation.ProductPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldProductPrice,
		})
	}
	if value, ok := puo.mutation.AddedProductPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldProductPrice,
		})
	}
	if value, ok := puo.mutation.ProductTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldProductTime,
		})
	}
	if value, ok := puo.mutation.AddedProductTime(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldProductTime,
		})
	}
	if value, ok := puo.mutation.ProductPaymentOfYear(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldProductPaymentOfYear,
		})
	}
	if value, ok := puo.mutation.AddedProductPaymentOfYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldProductPaymentOfYear,
		})
	}
	if puo.mutation.GenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.GenderTable,
			Columns: []string{product.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.GenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.GenderTable,
			Columns: []string{product.GenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: gender.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.GroupofageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.GroupofageTable,
			Columns: []string{product.GroupofageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: groupofage.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.GroupofageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.GroupofageTable,
			Columns: []string{product.GroupofageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: groupofage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.OfficerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.OfficerTable,
			Columns: []string{product.OfficerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: officer.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OfficerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.OfficerTable,
			Columns: []string{product.OfficerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: officer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := puo.mutation.RemovedProductInsuranceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductInsuranceTable,
			Columns: []string{product.ProductInsuranceColumn},
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
	if nodes := puo.mutation.ProductInsuranceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductInsuranceTable,
			Columns: []string{product.ProductInsuranceColumn},
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
	if nodes := puo.mutation.RemovedProductInquiryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductInquiryTable,
			Columns: []string{product.ProductInquiryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inquiry.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProductInquiryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductInquiryTable,
			Columns: []string{product.ProductInquiryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: inquiry.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := puo.mutation.RemovedProductPaybackIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductPaybackTable,
			Columns: []string{product.ProductPaybackColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payback.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProductPaybackIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductPaybackTable,
			Columns: []string{product.ProductPaybackColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payback.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := puo.mutation.RemovedProductRecordinsuranceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductRecordinsuranceTable,
			Columns: []string{product.ProductRecordinsuranceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: recordinsurance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ProductRecordinsuranceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ProductRecordinsuranceTable,
			Columns: []string{product.ProductRecordinsuranceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: recordinsurance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Product{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
