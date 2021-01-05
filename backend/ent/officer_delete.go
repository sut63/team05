// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/predicate"
)

// OfficerDelete is the builder for deleting a Officer entity.
type OfficerDelete struct {
	config
	hooks      []Hook
	mutation   *OfficerMutation
	predicates []predicate.Officer
}

// Where adds a new predicate to the delete builder.
func (od *OfficerDelete) Where(ps ...predicate.Officer) *OfficerDelete {
	od.predicates = append(od.predicates, ps...)
	return od
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (od *OfficerDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(od.hooks) == 0 {
		affected, err = od.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OfficerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			od.mutation = mutation
			affected, err = od.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(od.hooks) - 1; i >= 0; i-- {
			mut = od.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, od.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (od *OfficerDelete) ExecX(ctx context.Context) int {
	n, err := od.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (od *OfficerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: officer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: officer.FieldID,
			},
		},
	}
	if ps := od.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, od.driver, _spec)
}

// OfficerDeleteOne is the builder for deleting a single Officer entity.
type OfficerDeleteOne struct {
	od *OfficerDelete
}

// Exec executes the deletion query.
func (odo *OfficerDeleteOne) Exec(ctx context.Context) error {
	n, err := odo.od.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{officer.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (odo *OfficerDeleteOne) ExecX(ctx context.Context) {
	odo.od.ExecX(ctx)
}
