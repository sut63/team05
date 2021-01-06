// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team05/ent/category"
	"github.com/sut63/team05/ent/inquiry"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/predicate"
	"github.com/sut63/team05/ent/product"
)

// InquiryUpdate is the builder for updating Inquiry entities.
type InquiryUpdate struct {
	config
	hooks      []Hook
	mutation   *InquiryMutation
	predicates []predicate.Inquiry
}

// Where adds a new predicate for the builder.
func (iu *InquiryUpdate) Where(ps ...predicate.Inquiry) *InquiryUpdate {
	iu.predicates = append(iu.predicates, ps...)
	return iu
}

// SetInquiryInguiryMessages sets the Inquiry_inguiry_messages field.
func (iu *InquiryUpdate) SetInquiryInguiryMessages(s string) *InquiryUpdate {
	iu.mutation.SetInquiryInguiryMessages(s)
	return iu
}

// SetInquiryTimeMessages sets the Inquiry_time_messages field.
func (iu *InquiryUpdate) SetInquiryTimeMessages(t time.Time) *InquiryUpdate {
	iu.mutation.SetInquiryTimeMessages(t)
	return iu
}

// SetNillableInquiryTimeMessages sets the Inquiry_time_messages field if the given value is not nil.
func (iu *InquiryUpdate) SetNillableInquiryTimeMessages(t *time.Time) *InquiryUpdate {
	if t != nil {
		iu.SetInquiryTimeMessages(*t)
	}
	return iu
}

// SetMemberID sets the Member edge to Member by id.
func (iu *InquiryUpdate) SetMemberID(id int) *InquiryUpdate {
	iu.mutation.SetMemberID(id)
	return iu
}

// SetNillableMemberID sets the Member edge to Member by id if the given value is not nil.
func (iu *InquiryUpdate) SetNillableMemberID(id *int) *InquiryUpdate {
	if id != nil {
		iu = iu.SetMemberID(*id)
	}
	return iu
}

// SetMember sets the Member edge to Member.
func (iu *InquiryUpdate) SetMember(m *Member) *InquiryUpdate {
	return iu.SetMemberID(m.ID)
}

// SetCategoryID sets the Category edge to Category by id.
func (iu *InquiryUpdate) SetCategoryID(id int) *InquiryUpdate {
	iu.mutation.SetCategoryID(id)
	return iu
}

// SetNillableCategoryID sets the Category edge to Category by id if the given value is not nil.
func (iu *InquiryUpdate) SetNillableCategoryID(id *int) *InquiryUpdate {
	if id != nil {
		iu = iu.SetCategoryID(*id)
	}
	return iu
}

// SetCategory sets the Category edge to Category.
func (iu *InquiryUpdate) SetCategory(c *Category) *InquiryUpdate {
	return iu.SetCategoryID(c.ID)
}

// SetOfficerID sets the Officer edge to Officer by id.
func (iu *InquiryUpdate) SetOfficerID(id int) *InquiryUpdate {
	iu.mutation.SetOfficerID(id)
	return iu
}

// SetNillableOfficerID sets the Officer edge to Officer by id if the given value is not nil.
func (iu *InquiryUpdate) SetNillableOfficerID(id *int) *InquiryUpdate {
	if id != nil {
		iu = iu.SetOfficerID(*id)
	}
	return iu
}

// SetOfficer sets the Officer edge to Officer.
func (iu *InquiryUpdate) SetOfficer(o *Officer) *InquiryUpdate {
	return iu.SetOfficerID(o.ID)
}

// SetProductID sets the Product edge to Product by id.
func (iu *InquiryUpdate) SetProductID(id int) *InquiryUpdate {
	iu.mutation.SetProductID(id)
	return iu
}

// SetNillableProductID sets the Product edge to Product by id if the given value is not nil.
func (iu *InquiryUpdate) SetNillableProductID(id *int) *InquiryUpdate {
	if id != nil {
		iu = iu.SetProductID(*id)
	}
	return iu
}

// SetProduct sets the Product edge to Product.
func (iu *InquiryUpdate) SetProduct(p *Product) *InquiryUpdate {
	return iu.SetProductID(p.ID)
}

// Mutation returns the InquiryMutation object of the builder.
func (iu *InquiryUpdate) Mutation() *InquiryMutation {
	return iu.mutation
}

// ClearMember clears the Member edge to Member.
func (iu *InquiryUpdate) ClearMember() *InquiryUpdate {
	iu.mutation.ClearMember()
	return iu
}

// ClearCategory clears the Category edge to Category.
func (iu *InquiryUpdate) ClearCategory() *InquiryUpdate {
	iu.mutation.ClearCategory()
	return iu
}

// ClearOfficer clears the Officer edge to Officer.
func (iu *InquiryUpdate) ClearOfficer() *InquiryUpdate {
	iu.mutation.ClearOfficer()
	return iu
}

// ClearProduct clears the Product edge to Product.
func (iu *InquiryUpdate) ClearProduct() *InquiryUpdate {
	iu.mutation.ClearProduct()
	return iu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (iu *InquiryUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := iu.mutation.InquiryInguiryMessages(); ok {
		if err := inquiry.InquiryInguiryMessagesValidator(v); err != nil {
			return 0, &ValidationError{Name: "Inquiry_inguiry_messages", err: fmt.Errorf("ent: validator failed for field \"Inquiry_inguiry_messages\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InquiryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InquiryUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InquiryUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InquiryUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iu *InquiryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   inquiry.Table,
			Columns: inquiry.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: inquiry.FieldID,
			},
		},
	}
	if ps := iu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.InquiryInguiryMessages(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inquiry.FieldInquiryInguiryMessages,
		})
	}
	if value, ok := iu.mutation.InquiryTimeMessages(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: inquiry.FieldInquiryTimeMessages,
		})
	}
	if iu.mutation.MemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inquiry.MemberTable,
			Columns: []string{inquiry.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: member.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inquiry.MemberTable,
			Columns: []string{inquiry.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: member.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inquiry.CategoryTable,
			Columns: []string{inquiry.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inquiry.CategoryTable,
			Columns: []string{inquiry.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.OfficerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inquiry.OfficerTable,
			Columns: []string{inquiry.OfficerColumn},
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
	if nodes := iu.mutation.OfficerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inquiry.OfficerTable,
			Columns: []string{inquiry.OfficerColumn},
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
	if iu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inquiry.ProductTable,
			Columns: []string{inquiry.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inquiry.ProductTable,
			Columns: []string{inquiry.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{inquiry.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// InquiryUpdateOne is the builder for updating a single Inquiry entity.
type InquiryUpdateOne struct {
	config
	hooks    []Hook
	mutation *InquiryMutation
}

// SetInquiryInguiryMessages sets the Inquiry_inguiry_messages field.
func (iuo *InquiryUpdateOne) SetInquiryInguiryMessages(s string) *InquiryUpdateOne {
	iuo.mutation.SetInquiryInguiryMessages(s)
	return iuo
}

// SetInquiryTimeMessages sets the Inquiry_time_messages field.
func (iuo *InquiryUpdateOne) SetInquiryTimeMessages(t time.Time) *InquiryUpdateOne {
	iuo.mutation.SetInquiryTimeMessages(t)
	return iuo
}

// SetNillableInquiryTimeMessages sets the Inquiry_time_messages field if the given value is not nil.
func (iuo *InquiryUpdateOne) SetNillableInquiryTimeMessages(t *time.Time) *InquiryUpdateOne {
	if t != nil {
		iuo.SetInquiryTimeMessages(*t)
	}
	return iuo
}

// SetMemberID sets the Member edge to Member by id.
func (iuo *InquiryUpdateOne) SetMemberID(id int) *InquiryUpdateOne {
	iuo.mutation.SetMemberID(id)
	return iuo
}

// SetNillableMemberID sets the Member edge to Member by id if the given value is not nil.
func (iuo *InquiryUpdateOne) SetNillableMemberID(id *int) *InquiryUpdateOne {
	if id != nil {
		iuo = iuo.SetMemberID(*id)
	}
	return iuo
}

// SetMember sets the Member edge to Member.
func (iuo *InquiryUpdateOne) SetMember(m *Member) *InquiryUpdateOne {
	return iuo.SetMemberID(m.ID)
}

// SetCategoryID sets the Category edge to Category by id.
func (iuo *InquiryUpdateOne) SetCategoryID(id int) *InquiryUpdateOne {
	iuo.mutation.SetCategoryID(id)
	return iuo
}

// SetNillableCategoryID sets the Category edge to Category by id if the given value is not nil.
func (iuo *InquiryUpdateOne) SetNillableCategoryID(id *int) *InquiryUpdateOne {
	if id != nil {
		iuo = iuo.SetCategoryID(*id)
	}
	return iuo
}

// SetCategory sets the Category edge to Category.
func (iuo *InquiryUpdateOne) SetCategory(c *Category) *InquiryUpdateOne {
	return iuo.SetCategoryID(c.ID)
}

// SetOfficerID sets the Officer edge to Officer by id.
func (iuo *InquiryUpdateOne) SetOfficerID(id int) *InquiryUpdateOne {
	iuo.mutation.SetOfficerID(id)
	return iuo
}

// SetNillableOfficerID sets the Officer edge to Officer by id if the given value is not nil.
func (iuo *InquiryUpdateOne) SetNillableOfficerID(id *int) *InquiryUpdateOne {
	if id != nil {
		iuo = iuo.SetOfficerID(*id)
	}
	return iuo
}

// SetOfficer sets the Officer edge to Officer.
func (iuo *InquiryUpdateOne) SetOfficer(o *Officer) *InquiryUpdateOne {
	return iuo.SetOfficerID(o.ID)
}

// SetProductID sets the Product edge to Product by id.
func (iuo *InquiryUpdateOne) SetProductID(id int) *InquiryUpdateOne {
	iuo.mutation.SetProductID(id)
	return iuo
}

// SetNillableProductID sets the Product edge to Product by id if the given value is not nil.
func (iuo *InquiryUpdateOne) SetNillableProductID(id *int) *InquiryUpdateOne {
	if id != nil {
		iuo = iuo.SetProductID(*id)
	}
	return iuo
}

// SetProduct sets the Product edge to Product.
func (iuo *InquiryUpdateOne) SetProduct(p *Product) *InquiryUpdateOne {
	return iuo.SetProductID(p.ID)
}

// Mutation returns the InquiryMutation object of the builder.
func (iuo *InquiryUpdateOne) Mutation() *InquiryMutation {
	return iuo.mutation
}

// ClearMember clears the Member edge to Member.
func (iuo *InquiryUpdateOne) ClearMember() *InquiryUpdateOne {
	iuo.mutation.ClearMember()
	return iuo
}

// ClearCategory clears the Category edge to Category.
func (iuo *InquiryUpdateOne) ClearCategory() *InquiryUpdateOne {
	iuo.mutation.ClearCategory()
	return iuo
}

// ClearOfficer clears the Officer edge to Officer.
func (iuo *InquiryUpdateOne) ClearOfficer() *InquiryUpdateOne {
	iuo.mutation.ClearOfficer()
	return iuo
}

// ClearProduct clears the Product edge to Product.
func (iuo *InquiryUpdateOne) ClearProduct() *InquiryUpdateOne {
	iuo.mutation.ClearProduct()
	return iuo
}

// Save executes the query and returns the updated entity.
func (iuo *InquiryUpdateOne) Save(ctx context.Context) (*Inquiry, error) {
	if v, ok := iuo.mutation.InquiryInguiryMessages(); ok {
		if err := inquiry.InquiryInguiryMessagesValidator(v); err != nil {
			return nil, &ValidationError{Name: "Inquiry_inguiry_messages", err: fmt.Errorf("ent: validator failed for field \"Inquiry_inguiry_messages\": %w", err)}
		}
	}

	var (
		err  error
		node *Inquiry
	)
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InquiryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InquiryUpdateOne) SaveX(ctx context.Context) *Inquiry {
	i, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return i
}

// Exec executes the query on the entity.
func (iuo *InquiryUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InquiryUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iuo *InquiryUpdateOne) sqlSave(ctx context.Context) (i *Inquiry, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   inquiry.Table,
			Columns: inquiry.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: inquiry.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Inquiry.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := iuo.mutation.InquiryInguiryMessages(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: inquiry.FieldInquiryInguiryMessages,
		})
	}
	if value, ok := iuo.mutation.InquiryTimeMessages(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: inquiry.FieldInquiryTimeMessages,
		})
	}
	if iuo.mutation.MemberCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inquiry.MemberTable,
			Columns: []string{inquiry.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: member.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inquiry.MemberTable,
			Columns: []string{inquiry.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: member.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inquiry.CategoryTable,
			Columns: []string{inquiry.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inquiry.CategoryTable,
			Columns: []string{inquiry.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.OfficerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inquiry.OfficerTable,
			Columns: []string{inquiry.OfficerColumn},
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
	if nodes := iuo.mutation.OfficerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inquiry.OfficerTable,
			Columns: []string{inquiry.OfficerColumn},
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
	if iuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inquiry.ProductTable,
			Columns: []string{inquiry.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   inquiry.ProductTable,
			Columns: []string{inquiry.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	i = &Inquiry{config: iuo.config}
	_spec.Assign = i.assignValues
	_spec.ScanValues = i.scanValues()
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{inquiry.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return i, nil
}
