// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/position"
)

// PositionCreate is the builder for creating a Position entity.
type PositionCreate struct {
	config
	mutation *PositionMutation
	hooks    []Hook
}

// SetPositionName sets the position_name field.
func (pc *PositionCreate) SetPositionName(s string) *PositionCreate {
	pc.mutation.SetPositionName(s)
	return pc
}

// AddMemberIDs adds the members edge to Member by ids.
func (pc *PositionCreate) AddMemberIDs(ids ...int) *PositionCreate {
	pc.mutation.AddMemberIDs(ids...)
	return pc
}

// AddMembers adds the members edges to Member.
func (pc *PositionCreate) AddMembers(m ...*Member) *PositionCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pc.AddMemberIDs(ids...)
}

// AddOfficerIDs adds the officers edge to Officer by ids.
func (pc *PositionCreate) AddOfficerIDs(ids ...int) *PositionCreate {
	pc.mutation.AddOfficerIDs(ids...)
	return pc
}

// AddOfficers adds the officers edges to Officer.
func (pc *PositionCreate) AddOfficers(o ...*Officer) *PositionCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pc.AddOfficerIDs(ids...)
}

// Mutation returns the PositionMutation object of the builder.
func (pc *PositionCreate) Mutation() *PositionMutation {
	return pc.mutation
}

// Save creates the Position in the database.
func (pc *PositionCreate) Save(ctx context.Context) (*Position, error) {
	if _, ok := pc.mutation.PositionName(); !ok {
		return nil, &ValidationError{Name: "position_name", err: errors.New("ent: missing required field \"position_name\"")}
	}
	if v, ok := pc.mutation.PositionName(); ok {
		if err := position.PositionNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "position_name", err: fmt.Errorf("ent: validator failed for field \"position_name\": %w", err)}
		}
	}
	var (
		err  error
		node *Position
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PositionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PositionCreate) SaveX(ctx context.Context) *Position {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PositionCreate) sqlSave(ctx context.Context) (*Position, error) {
	po, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	po.ID = int(id)
	return po, nil
}

func (pc *PositionCreate) createSpec() (*Position, *sqlgraph.CreateSpec) {
	var (
		po    = &Position{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: position.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: position.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.PositionName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: position.FieldPositionName,
		})
		po.PositionName = value
	}
	if nodes := pc.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.MembersTable,
			Columns: []string{position.MembersColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.OfficersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   position.OfficersTable,
			Columns: []string{position.OfficersColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return po, _spec
}
