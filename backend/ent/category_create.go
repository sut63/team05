// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team05/ent/category"
	"github.com/sut63/team05/ent/inquiry"
)

// CategoryCreate is the builder for creating a Category entity.
type CategoryCreate struct {
	config
	mutation *CategoryMutation
	hooks    []Hook
}

// SetCategoryName sets the category_name field.
func (cc *CategoryCreate) SetCategoryName(s string) *CategoryCreate {
	cc.mutation.SetCategoryName(s)
	return cc
}

// AddCategoryInquiryIDs adds the category_inquiry edge to Inquiry by ids.
func (cc *CategoryCreate) AddCategoryInquiryIDs(ids ...int) *CategoryCreate {
	cc.mutation.AddCategoryInquiryIDs(ids...)
	return cc
}

// AddCategoryInquiry adds the category_inquiry edges to Inquiry.
func (cc *CategoryCreate) AddCategoryInquiry(i ...*Inquiry) *CategoryCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cc.AddCategoryInquiryIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cc *CategoryCreate) Mutation() *CategoryMutation {
	return cc.mutation
}

// Save creates the Category in the database.
func (cc *CategoryCreate) Save(ctx context.Context) (*Category, error) {
	if _, ok := cc.mutation.CategoryName(); !ok {
		return nil, &ValidationError{Name: "category_name", err: errors.New("ent: missing required field \"category_name\"")}
	}
	var (
		err  error
		node *Category
	)
	if len(cc.hooks) == 0 {
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CategoryCreate) SaveX(ctx context.Context) *Category {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CategoryCreate) sqlSave(ctx context.Context) (*Category, error) {
	c, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	c.ID = int(id)
	return c, nil
}

func (cc *CategoryCreate) createSpec() (*Category, *sqlgraph.CreateSpec) {
	var (
		c     = &Category{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: category.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: category.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.CategoryName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldCategoryName,
		})
		c.CategoryName = value
	}
	if nodes := cc.mutation.CategoryInquiryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   category.CategoryInquiryTable,
			Columns: []string{category.CategoryInquiryColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return c, _spec
}
