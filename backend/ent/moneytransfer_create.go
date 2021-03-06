// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team05/ent/moneytransfer"
	"github.com/sut63/team05/ent/payment"
)

// MoneytransferCreate is the builder for creating a Moneytransfer entity.
type MoneytransferCreate struct {
	config
	mutation *MoneytransferMutation
	hooks    []Hook
}

// SetMoneytransferType sets the moneytransfer_type field.
func (mc *MoneytransferCreate) SetMoneytransferType(s string) *MoneytransferCreate {
	mc.mutation.SetMoneytransferType(s)
	return mc
}

// AddMoneytransferPaymentIDs adds the moneytransfer_payment edge to Payment by ids.
func (mc *MoneytransferCreate) AddMoneytransferPaymentIDs(ids ...int) *MoneytransferCreate {
	mc.mutation.AddMoneytransferPaymentIDs(ids...)
	return mc
}

// AddMoneytransferPayment adds the moneytransfer_payment edges to Payment.
func (mc *MoneytransferCreate) AddMoneytransferPayment(p ...*Payment) *MoneytransferCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mc.AddMoneytransferPaymentIDs(ids...)
}

// Mutation returns the MoneytransferMutation object of the builder.
func (mc *MoneytransferCreate) Mutation() *MoneytransferMutation {
	return mc.mutation
}

// Save creates the Moneytransfer in the database.
func (mc *MoneytransferCreate) Save(ctx context.Context) (*Moneytransfer, error) {
	if _, ok := mc.mutation.MoneytransferType(); !ok {
		return nil, &ValidationError{Name: "moneytransfer_type", err: errors.New("ent: missing required field \"moneytransfer_type\"")}
	}
	if v, ok := mc.mutation.MoneytransferType(); ok {
		if err := moneytransfer.MoneytransferTypeValidator(v); err != nil {
			return nil, &ValidationError{Name: "moneytransfer_type", err: fmt.Errorf("ent: validator failed for field \"moneytransfer_type\": %w", err)}
		}
	}
	var (
		err  error
		node *Moneytransfer
	)
	if len(mc.hooks) == 0 {
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MoneytransferMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MoneytransferCreate) SaveX(ctx context.Context) *Moneytransfer {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *MoneytransferCreate) sqlSave(ctx context.Context) (*Moneytransfer, error) {
	m, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	m.ID = int(id)
	return m, nil
}

func (mc *MoneytransferCreate) createSpec() (*Moneytransfer, *sqlgraph.CreateSpec) {
	var (
		m     = &Moneytransfer{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: moneytransfer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: moneytransfer.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.MoneytransferType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: moneytransfer.FieldMoneytransferType,
		})
		m.MoneytransferType = value
	}
	if nodes := mc.mutation.MoneytransferPaymentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   moneytransfer.MoneytransferPaymentTable,
			Columns: []string{moneytransfer.MoneytransferPaymentColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return m, _spec
}
