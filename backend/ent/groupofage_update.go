// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team05/ent/groupofage"
	"github.com/sut63/team05/ent/predicate"
	"github.com/sut63/team05/ent/product"
)

// GroupOfAgeUpdate is the builder for updating GroupOfAge entities.
type GroupOfAgeUpdate struct {
	config
	hooks      []Hook
	mutation   *GroupOfAgeMutation
	predicates []predicate.GroupOfAge
}

// Where adds a new predicate for the builder.
func (goau *GroupOfAgeUpdate) Where(ps ...predicate.GroupOfAge) *GroupOfAgeUpdate {
	goau.predicates = append(goau.predicates, ps...)
	return goau
}

// SetGroupOfAgeName sets the group_of_age_name field.
func (goau *GroupOfAgeUpdate) SetGroupOfAgeName(s string) *GroupOfAgeUpdate {
	goau.mutation.SetGroupOfAgeName(s)
	return goau
}

// SetGroupOfAgeAge sets the group_of_age_age field.
func (goau *GroupOfAgeUpdate) SetGroupOfAgeAge(s string) *GroupOfAgeUpdate {
	goau.mutation.SetGroupOfAgeAge(s)
	return goau
}

// AddGroupofageProductIDs adds the groupofage_product edge to Product by ids.
func (goau *GroupOfAgeUpdate) AddGroupofageProductIDs(ids ...int) *GroupOfAgeUpdate {
	goau.mutation.AddGroupofageProductIDs(ids...)
	return goau
}

// AddGroupofageProduct adds the groupofage_product edges to Product.
func (goau *GroupOfAgeUpdate) AddGroupofageProduct(p ...*Product) *GroupOfAgeUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return goau.AddGroupofageProductIDs(ids...)
}

// Mutation returns the GroupOfAgeMutation object of the builder.
func (goau *GroupOfAgeUpdate) Mutation() *GroupOfAgeMutation {
	return goau.mutation
}

// RemoveGroupofageProductIDs removes the groupofage_product edge to Product by ids.
func (goau *GroupOfAgeUpdate) RemoveGroupofageProductIDs(ids ...int) *GroupOfAgeUpdate {
	goau.mutation.RemoveGroupofageProductIDs(ids...)
	return goau
}

// RemoveGroupofageProduct removes groupofage_product edges to Product.
func (goau *GroupOfAgeUpdate) RemoveGroupofageProduct(p ...*Product) *GroupOfAgeUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return goau.RemoveGroupofageProductIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (goau *GroupOfAgeUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := goau.mutation.GroupOfAgeName(); ok {
		if err := groupofage.GroupOfAgeNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "group_of_age_name", err: fmt.Errorf("ent: validator failed for field \"group_of_age_name\": %w", err)}
		}
	}
	if v, ok := goau.mutation.GroupOfAgeAge(); ok {
		if err := groupofage.GroupOfAgeAgeValidator(v); err != nil {
			return 0, &ValidationError{Name: "group_of_age_age", err: fmt.Errorf("ent: validator failed for field \"group_of_age_age\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(goau.hooks) == 0 {
		affected, err = goau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupOfAgeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			goau.mutation = mutation
			affected, err = goau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(goau.hooks) - 1; i >= 0; i-- {
			mut = goau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, goau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (goau *GroupOfAgeUpdate) SaveX(ctx context.Context) int {
	affected, err := goau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (goau *GroupOfAgeUpdate) Exec(ctx context.Context) error {
	_, err := goau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (goau *GroupOfAgeUpdate) ExecX(ctx context.Context) {
	if err := goau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (goau *GroupOfAgeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupofage.Table,
			Columns: groupofage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupofage.FieldID,
			},
		},
	}
	if ps := goau.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := goau.mutation.GroupOfAgeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: groupofage.FieldGroupOfAgeName,
		})
	}
	if value, ok := goau.mutation.GroupOfAgeAge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: groupofage.FieldGroupOfAgeAge,
		})
	}
	if nodes := goau.mutation.RemovedGroupofageProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   groupofage.GroupofageProductTable,
			Columns: []string{groupofage.GroupofageProductColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := goau.mutation.GroupofageProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   groupofage.GroupofageProductTable,
			Columns: []string{groupofage.GroupofageProductColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, goau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupofage.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GroupOfAgeUpdateOne is the builder for updating a single GroupOfAge entity.
type GroupOfAgeUpdateOne struct {
	config
	hooks    []Hook
	mutation *GroupOfAgeMutation
}

// SetGroupOfAgeName sets the group_of_age_name field.
func (goauo *GroupOfAgeUpdateOne) SetGroupOfAgeName(s string) *GroupOfAgeUpdateOne {
	goauo.mutation.SetGroupOfAgeName(s)
	return goauo
}

// SetGroupOfAgeAge sets the group_of_age_age field.
func (goauo *GroupOfAgeUpdateOne) SetGroupOfAgeAge(s string) *GroupOfAgeUpdateOne {
	goauo.mutation.SetGroupOfAgeAge(s)
	return goauo
}

// AddGroupofageProductIDs adds the groupofage_product edge to Product by ids.
func (goauo *GroupOfAgeUpdateOne) AddGroupofageProductIDs(ids ...int) *GroupOfAgeUpdateOne {
	goauo.mutation.AddGroupofageProductIDs(ids...)
	return goauo
}

// AddGroupofageProduct adds the groupofage_product edges to Product.
func (goauo *GroupOfAgeUpdateOne) AddGroupofageProduct(p ...*Product) *GroupOfAgeUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return goauo.AddGroupofageProductIDs(ids...)
}

// Mutation returns the GroupOfAgeMutation object of the builder.
func (goauo *GroupOfAgeUpdateOne) Mutation() *GroupOfAgeMutation {
	return goauo.mutation
}

// RemoveGroupofageProductIDs removes the groupofage_product edge to Product by ids.
func (goauo *GroupOfAgeUpdateOne) RemoveGroupofageProductIDs(ids ...int) *GroupOfAgeUpdateOne {
	goauo.mutation.RemoveGroupofageProductIDs(ids...)
	return goauo
}

// RemoveGroupofageProduct removes groupofage_product edges to Product.
func (goauo *GroupOfAgeUpdateOne) RemoveGroupofageProduct(p ...*Product) *GroupOfAgeUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return goauo.RemoveGroupofageProductIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (goauo *GroupOfAgeUpdateOne) Save(ctx context.Context) (*GroupOfAge, error) {
	if v, ok := goauo.mutation.GroupOfAgeName(); ok {
		if err := groupofage.GroupOfAgeNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "group_of_age_name", err: fmt.Errorf("ent: validator failed for field \"group_of_age_name\": %w", err)}
		}
	}
	if v, ok := goauo.mutation.GroupOfAgeAge(); ok {
		if err := groupofage.GroupOfAgeAgeValidator(v); err != nil {
			return nil, &ValidationError{Name: "group_of_age_age", err: fmt.Errorf("ent: validator failed for field \"group_of_age_age\": %w", err)}
		}
	}

	var (
		err  error
		node *GroupOfAge
	)
	if len(goauo.hooks) == 0 {
		node, err = goauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupOfAgeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			goauo.mutation = mutation
			node, err = goauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(goauo.hooks) - 1; i >= 0; i-- {
			mut = goauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, goauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (goauo *GroupOfAgeUpdateOne) SaveX(ctx context.Context) *GroupOfAge {
	goa, err := goauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return goa
}

// Exec executes the query on the entity.
func (goauo *GroupOfAgeUpdateOne) Exec(ctx context.Context) error {
	_, err := goauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (goauo *GroupOfAgeUpdateOne) ExecX(ctx context.Context) {
	if err := goauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (goauo *GroupOfAgeUpdateOne) sqlSave(ctx context.Context) (goa *GroupOfAge, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupofage.Table,
			Columns: groupofage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupofage.FieldID,
			},
		},
	}
	id, ok := goauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing GroupOfAge.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := goauo.mutation.GroupOfAgeName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: groupofage.FieldGroupOfAgeName,
		})
	}
	if value, ok := goauo.mutation.GroupOfAgeAge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: groupofage.FieldGroupOfAgeAge,
		})
	}
	if nodes := goauo.mutation.RemovedGroupofageProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   groupofage.GroupofageProductTable,
			Columns: []string{groupofage.GroupofageProductColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := goauo.mutation.GroupofageProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   groupofage.GroupofageProductTable,
			Columns: []string{groupofage.GroupofageProductColumn},
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
	goa = &GroupOfAge{config: goauo.config}
	_spec.Assign = goa.assignValues
	_spec.ScanValues = goa.scanValues()
	if err = sqlgraph.UpdateNode(ctx, goauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupofage.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return goa, nil
}
