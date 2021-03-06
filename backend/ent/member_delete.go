// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/predicate"
)

// MemberDelete is the builder for deleting a Member entity.
type MemberDelete struct {
	config
	hooks      []Hook
	mutation   *MemberMutation
	predicates []predicate.Member
}

// Where adds a new predicate to the delete builder.
func (md *MemberDelete) Where(ps ...predicate.Member) *MemberDelete {
	md.predicates = append(md.predicates, ps...)
	return md
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (md *MemberDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(md.hooks) == 0 {
		affected, err = md.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MemberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			md.mutation = mutation
			affected, err = md.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(md.hooks) - 1; i >= 0; i-- {
			mut = md.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, md.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (md *MemberDelete) ExecX(ctx context.Context) int {
	n, err := md.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (md *MemberDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: member.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: member.FieldID,
			},
		},
	}
	if ps := md.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, md.driver, _spec)
}

// MemberDeleteOne is the builder for deleting a single Member entity.
type MemberDeleteOne struct {
	md *MemberDelete
}

// Exec executes the deletion query.
func (mdo *MemberDeleteOne) Exec(ctx context.Context) error {
	n, err := mdo.md.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{member.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mdo *MemberDeleteOne) ExecX(ctx context.Context) {
	mdo.md.ExecX(ctx)
}
