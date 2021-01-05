// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team05/ent/hospital"
	"github.com/sut63/team05/ent/insurance"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/product"
)

// InsuranceCreate is the builder for creating a Insurance entity.
type InsuranceCreate struct {
	config
	mutation *InsuranceMutation
	hooks    []Hook
}

// SetInsuranceAddress sets the insurance_address field.
func (ic *InsuranceCreate) SetInsuranceAddress(s string) *InsuranceCreate {
	ic.mutation.SetInsuranceAddress(s)
	return ic
}

// SetInsuranceInsurer sets the insurance_insurer field.
func (ic *InsuranceCreate) SetInsuranceInsurer(s string) *InsuranceCreate {
	ic.mutation.SetInsuranceInsurer(s)
	return ic
}

// SetInsuranceTimeBuy sets the insurance_time_buy field.
func (ic *InsuranceCreate) SetInsuranceTimeBuy(t time.Time) *InsuranceCreate {
	ic.mutation.SetInsuranceTimeBuy(t)
	return ic
}

// SetNillableInsuranceTimeBuy sets the insurance_time_buy field if the given value is not nil.
func (ic *InsuranceCreate) SetNillableInsuranceTimeBuy(t *time.Time) *InsuranceCreate {
	if t != nil {
		ic.SetInsuranceTimeBuy(*t)
	}
	return ic
}

// SetInsuranceTimeFirstpay sets the insurance_time_firstpay field.
func (ic *InsuranceCreate) SetInsuranceTimeFirstpay(t time.Time) *InsuranceCreate {
	ic.mutation.SetInsuranceTimeFirstpay(t)
	return ic
}

// SetNillableInsuranceTimeFirstpay sets the insurance_time_firstpay field if the given value is not nil.
func (ic *InsuranceCreate) SetNillableInsuranceTimeFirstpay(t *time.Time) *InsuranceCreate {
	if t != nil {
		ic.SetInsuranceTimeFirstpay(*t)
	}
	return ic
}

// SetMemberID sets the Member edge to Member by id.
func (ic *InsuranceCreate) SetMemberID(id int) *InsuranceCreate {
	ic.mutation.SetMemberID(id)
	return ic
}

// SetNillableMemberID sets the Member edge to Member by id if the given value is not nil.
func (ic *InsuranceCreate) SetNillableMemberID(id *int) *InsuranceCreate {
	if id != nil {
		ic = ic.SetMemberID(*id)
	}
	return ic
}

// SetMember sets the Member edge to Member.
func (ic *InsuranceCreate) SetMember(m *Member) *InsuranceCreate {
	return ic.SetMemberID(m.ID)
}

// SetHospitalID sets the Hospital edge to Hospital by id.
func (ic *InsuranceCreate) SetHospitalID(id int) *InsuranceCreate {
	ic.mutation.SetHospitalID(id)
	return ic
}

// SetNillableHospitalID sets the Hospital edge to Hospital by id if the given value is not nil.
func (ic *InsuranceCreate) SetNillableHospitalID(id *int) *InsuranceCreate {
	if id != nil {
		ic = ic.SetHospitalID(*id)
	}
	return ic
}

// SetHospital sets the Hospital edge to Hospital.
func (ic *InsuranceCreate) SetHospital(h *Hospital) *InsuranceCreate {
	return ic.SetHospitalID(h.ID)
}

// SetOfficerID sets the Officer edge to Officer by id.
func (ic *InsuranceCreate) SetOfficerID(id int) *InsuranceCreate {
	ic.mutation.SetOfficerID(id)
	return ic
}

// SetNillableOfficerID sets the Officer edge to Officer by id if the given value is not nil.
func (ic *InsuranceCreate) SetNillableOfficerID(id *int) *InsuranceCreate {
	if id != nil {
		ic = ic.SetOfficerID(*id)
	}
	return ic
}

// SetOfficer sets the Officer edge to Officer.
func (ic *InsuranceCreate) SetOfficer(o *Officer) *InsuranceCreate {
	return ic.SetOfficerID(o.ID)
}

// SetProductID sets the Product edge to Product by id.
func (ic *InsuranceCreate) SetProductID(id int) *InsuranceCreate {
	ic.mutation.SetProductID(id)
	return ic
}

// SetNillableProductID sets the Product edge to Product by id if the given value is not nil.
func (ic *InsuranceCreate) SetNillableProductID(id *int) *InsuranceCreate {
	if id != nil {
		ic = ic.SetProductID(*id)
	}
	return ic
}

// SetProduct sets the Product edge to Product.
func (ic *InsuranceCreate) SetProduct(p *Product) *InsuranceCreate {
	return ic.SetProductID(p.ID)
}

// Mutation returns the InsuranceMutation object of the builder.
func (ic *InsuranceCreate) Mutation() *InsuranceMutation {
	return ic.mutation
}

// Save creates the Insurance in the database.
func (ic *InsuranceCreate) Save(ctx context.Context) (*Insurance, error) {
	if _, ok := ic.mutation.InsuranceAddress(); !ok {
		return nil, &ValidationError{Name: "insurance_address", err: errors.New("ent: missing required field \"insurance_address\"")}
	}
	if v, ok := ic.mutation.InsuranceAddress(); ok {
		if err := insurance.InsuranceAddressValidator(v); err != nil {
			return nil, &ValidationError{Name: "insurance_address", err: fmt.Errorf("ent: validator failed for field \"insurance_address\": %w", err)}
		}
	}
	if _, ok := ic.mutation.InsuranceInsurer(); !ok {
		return nil, &ValidationError{Name: "insurance_insurer", err: errors.New("ent: missing required field \"insurance_insurer\"")}
	}
	if v, ok := ic.mutation.InsuranceInsurer(); ok {
		if err := insurance.InsuranceInsurerValidator(v); err != nil {
			return nil, &ValidationError{Name: "insurance_insurer", err: fmt.Errorf("ent: validator failed for field \"insurance_insurer\": %w", err)}
		}
	}
	if _, ok := ic.mutation.InsuranceTimeBuy(); !ok {
		v := insurance.DefaultInsuranceTimeBuy()
		ic.mutation.SetInsuranceTimeBuy(v)
	}
	if _, ok := ic.mutation.InsuranceTimeFirstpay(); !ok {
		v := insurance.DefaultInsuranceTimeFirstpay()
		ic.mutation.SetInsuranceTimeFirstpay(v)
	}
	var (
		err  error
		node *Insurance
	)
	if len(ic.hooks) == 0 {
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InsuranceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ic.mutation = mutation
			node, err = ic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			mut = ic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InsuranceCreate) SaveX(ctx context.Context) *Insurance {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ic *InsuranceCreate) sqlSave(ctx context.Context) (*Insurance, error) {
	i, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	i.ID = int(id)
	return i, nil
}

func (ic *InsuranceCreate) createSpec() (*Insurance, *sqlgraph.CreateSpec) {
	var (
		i     = &Insurance{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: insurance.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: insurance.FieldID,
			},
		}
	)
	if value, ok := ic.mutation.InsuranceAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: insurance.FieldInsuranceAddress,
		})
		i.InsuranceAddress = value
	}
	if value, ok := ic.mutation.InsuranceInsurer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: insurance.FieldInsuranceInsurer,
		})
		i.InsuranceInsurer = value
	}
	if value, ok := ic.mutation.InsuranceTimeBuy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: insurance.FieldInsuranceTimeBuy,
		})
		i.InsuranceTimeBuy = value
	}
	if value, ok := ic.mutation.InsuranceTimeFirstpay(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: insurance.FieldInsuranceTimeFirstpay,
		})
		i.InsuranceTimeFirstpay = value
	}
	if nodes := ic.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   insurance.MemberTable,
			Columns: []string{insurance.MemberColumn},
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
	if nodes := ic.mutation.HospitalIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   insurance.HospitalTable,
			Columns: []string{insurance.HospitalColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: hospital.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.OfficerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   insurance.OfficerTable,
			Columns: []string{insurance.OfficerColumn},
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
	if nodes := ic.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   insurance.ProductTable,
			Columns: []string{insurance.ProductColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return i, _spec
}
