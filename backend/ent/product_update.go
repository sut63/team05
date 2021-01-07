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
func (pu *ProductUpdate) SetProductPaymentOfYear(f float64) *ProductUpdate {
	pu.mutation.ResetProductPaymentOfYear()
	pu.mutation.SetProductPaymentOfYear(f)
	return pu
}

// AddProductPaymentOfYear adds f to product_payment_of_year.
func (pu *ProductUpdate) AddProductPaymentOfYear(f float64) *ProductUpdate {
	pu.mutation.AddProductPaymentOfYear(f)
	return pu
}

// SetProductGenderID sets the product_gender edge to Gender by id.
func (pu *ProductUpdate) SetProductGenderID(id int) *ProductUpdate {
	pu.mutation.SetProductGenderID(id)
	return pu
}

// SetNillableProductGenderID sets the product_gender edge to Gender by id if the given value is not nil.
func (pu *ProductUpdate) SetNillableProductGenderID(id *int) *ProductUpdate {
	if id != nil {
		pu = pu.SetProductGenderID(*id)
	}
	return pu
}

// SetProductGender sets the product_gender edge to Gender.
func (pu *ProductUpdate) SetProductGender(g *Gender) *ProductUpdate {
	return pu.SetProductGenderID(g.ID)
}

// SetProductGroupageID sets the product_groupage edge to GroupOfAge by id.
func (pu *ProductUpdate) SetProductGroupageID(id int) *ProductUpdate {
	pu.mutation.SetProductGroupageID(id)
	return pu
}

// SetNillableProductGroupageID sets the product_groupage edge to GroupOfAge by id if the given value is not nil.
func (pu *ProductUpdate) SetNillableProductGroupageID(id *int) *ProductUpdate {
	if id != nil {
		pu = pu.SetProductGroupageID(*id)
	}
	return pu
}

// SetProductGroupage sets the product_groupage edge to GroupOfAge.
func (pu *ProductUpdate) SetProductGroupage(g *GroupOfAge) *ProductUpdate {
	return pu.SetProductGroupageID(g.ID)
}

// SetProductOfficerID sets the product_officer edge to Officer by id.
func (pu *ProductUpdate) SetProductOfficerID(id int) *ProductUpdate {
	pu.mutation.SetProductOfficerID(id)
	return pu
}

// SetNillableProductOfficerID sets the product_officer edge to Officer by id if the given value is not nil.
func (pu *ProductUpdate) SetNillableProductOfficerID(id *int) *ProductUpdate {
	if id != nil {
		pu = pu.SetProductOfficerID(*id)
	}
	return pu
}

// SetProductOfficer sets the product_officer edge to Officer.
func (pu *ProductUpdate) SetProductOfficer(o *Officer) *ProductUpdate {
	return pu.SetProductOfficerID(o.ID)
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

// Mutation returns the ProductMutation object of the builder.
func (pu *ProductUpdate) Mutation() *ProductMutation {
	return pu.mutation
}

// ClearProductGender clears the product_gender edge to Gender.
func (pu *ProductUpdate) ClearProductGender() *ProductUpdate {
	pu.mutation.ClearProductGender()
	return pu
}

// ClearProductGroupage clears the product_groupage edge to GroupOfAge.
func (pu *ProductUpdate) ClearProductGroupage() *ProductUpdate {
	pu.mutation.ClearProductGroupage()
	return pu
}

// ClearProductOfficer clears the product_officer edge to Officer.
func (pu *ProductUpdate) ClearProductOfficer() *ProductUpdate {
	pu.mutation.ClearProductOfficer()
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

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *ProductUpdate) Save(ctx context.Context) (int, error) {

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
			Type:   field.TypeFloat64,
			Value:  value,
			Column: product.FieldProductPaymentOfYear,
		})
	}
	if value, ok := pu.mutation.AddedProductPaymentOfYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: product.FieldProductPaymentOfYear,
		})
	}
	if pu.mutation.ProductGenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.ProductGenderTable,
			Columns: []string{product.ProductGenderColumn},
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
	if nodes := pu.mutation.ProductGenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.ProductGenderTable,
			Columns: []string{product.ProductGenderColumn},
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
	if pu.mutation.ProductGroupageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.ProductGroupageTable,
			Columns: []string{product.ProductGroupageColumn},
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
	if nodes := pu.mutation.ProductGroupageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.ProductGroupageTable,
			Columns: []string{product.ProductGroupageColumn},
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
	if pu.mutation.ProductOfficerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.ProductOfficerTable,
			Columns: []string{product.ProductOfficerColumn},
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
	if nodes := pu.mutation.ProductOfficerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.ProductOfficerTable,
			Columns: []string{product.ProductOfficerColumn},
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
func (puo *ProductUpdateOne) SetProductPaymentOfYear(f float64) *ProductUpdateOne {
	puo.mutation.ResetProductPaymentOfYear()
	puo.mutation.SetProductPaymentOfYear(f)
	return puo
}

// AddProductPaymentOfYear adds f to product_payment_of_year.
func (puo *ProductUpdateOne) AddProductPaymentOfYear(f float64) *ProductUpdateOne {
	puo.mutation.AddProductPaymentOfYear(f)
	return puo
}

// SetProductGenderID sets the product_gender edge to Gender by id.
func (puo *ProductUpdateOne) SetProductGenderID(id int) *ProductUpdateOne {
	puo.mutation.SetProductGenderID(id)
	return puo
}

// SetNillableProductGenderID sets the product_gender edge to Gender by id if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableProductGenderID(id *int) *ProductUpdateOne {
	if id != nil {
		puo = puo.SetProductGenderID(*id)
	}
	return puo
}

// SetProductGender sets the product_gender edge to Gender.
func (puo *ProductUpdateOne) SetProductGender(g *Gender) *ProductUpdateOne {
	return puo.SetProductGenderID(g.ID)
}

// SetProductGroupageID sets the product_groupage edge to GroupOfAge by id.
func (puo *ProductUpdateOne) SetProductGroupageID(id int) *ProductUpdateOne {
	puo.mutation.SetProductGroupageID(id)
	return puo
}

// SetNillableProductGroupageID sets the product_groupage edge to GroupOfAge by id if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableProductGroupageID(id *int) *ProductUpdateOne {
	if id != nil {
		puo = puo.SetProductGroupageID(*id)
	}
	return puo
}

// SetProductGroupage sets the product_groupage edge to GroupOfAge.
func (puo *ProductUpdateOne) SetProductGroupage(g *GroupOfAge) *ProductUpdateOne {
	return puo.SetProductGroupageID(g.ID)
}

// SetProductOfficerID sets the product_officer edge to Officer by id.
func (puo *ProductUpdateOne) SetProductOfficerID(id int) *ProductUpdateOne {
	puo.mutation.SetProductOfficerID(id)
	return puo
}

// SetNillableProductOfficerID sets the product_officer edge to Officer by id if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableProductOfficerID(id *int) *ProductUpdateOne {
	if id != nil {
		puo = puo.SetProductOfficerID(*id)
	}
	return puo
}

// SetProductOfficer sets the product_officer edge to Officer.
func (puo *ProductUpdateOne) SetProductOfficer(o *Officer) *ProductUpdateOne {
	return puo.SetProductOfficerID(o.ID)
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

// Mutation returns the ProductMutation object of the builder.
func (puo *ProductUpdateOne) Mutation() *ProductMutation {
	return puo.mutation
}

// ClearProductGender clears the product_gender edge to Gender.
func (puo *ProductUpdateOne) ClearProductGender() *ProductUpdateOne {
	puo.mutation.ClearProductGender()
	return puo
}

// ClearProductGroupage clears the product_groupage edge to GroupOfAge.
func (puo *ProductUpdateOne) ClearProductGroupage() *ProductUpdateOne {
	puo.mutation.ClearProductGroupage()
	return puo
}

// ClearProductOfficer clears the product_officer edge to Officer.
func (puo *ProductUpdateOne) ClearProductOfficer() *ProductUpdateOne {
	puo.mutation.ClearProductOfficer()
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

// Save executes the query and returns the updated entity.
func (puo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {

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
			Type:   field.TypeFloat64,
			Value:  value,
			Column: product.FieldProductPaymentOfYear,
		})
	}
	if value, ok := puo.mutation.AddedProductPaymentOfYear(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: product.FieldProductPaymentOfYear,
		})
	}
	if puo.mutation.ProductGenderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.ProductGenderTable,
			Columns: []string{product.ProductGenderColumn},
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
	if nodes := puo.mutation.ProductGenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.ProductGenderTable,
			Columns: []string{product.ProductGenderColumn},
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
	if puo.mutation.ProductGroupageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.ProductGroupageTable,
			Columns: []string{product.ProductGroupageColumn},
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
	if nodes := puo.mutation.ProductGroupageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.ProductGroupageTable,
			Columns: []string{product.ProductGroupageColumn},
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
	if puo.mutation.ProductOfficerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.ProductOfficerTable,
			Columns: []string{product.ProductOfficerColumn},
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
	if nodes := puo.mutation.ProductOfficerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.ProductOfficerTable,
			Columns: []string{product.ProductOfficerColumn},
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
