// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

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

// InquiryQuery is the builder for querying Inquiry entities.
type InquiryQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Inquiry
	// eager-loading edges.
	withMember   *MemberQuery
	withCategory *CategoryQuery
	withOfficer  *OfficerQuery
	withProduct  *ProductQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (iq *InquiryQuery) Where(ps ...predicate.Inquiry) *InquiryQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit adds a limit step to the query.
func (iq *InquiryQuery) Limit(limit int) *InquiryQuery {
	iq.limit = &limit
	return iq
}

// Offset adds an offset step to the query.
func (iq *InquiryQuery) Offset(offset int) *InquiryQuery {
	iq.offset = &offset
	return iq
}

// Order adds an order step to the query.
func (iq *InquiryQuery) Order(o ...OrderFunc) *InquiryQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QueryMember chains the current query on the Member edge.
func (iq *InquiryQuery) QueryMember() *MemberQuery {
	query := &MemberQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(inquiry.Table, inquiry.FieldID, iq.sqlQuery()),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, inquiry.MemberTable, inquiry.MemberColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCategory chains the current query on the Category edge.
func (iq *InquiryQuery) QueryCategory() *CategoryQuery {
	query := &CategoryQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(inquiry.Table, inquiry.FieldID, iq.sqlQuery()),
			sqlgraph.To(category.Table, category.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, inquiry.CategoryTable, inquiry.CategoryColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOfficer chains the current query on the Officer edge.
func (iq *InquiryQuery) QueryOfficer() *OfficerQuery {
	query := &OfficerQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(inquiry.Table, inquiry.FieldID, iq.sqlQuery()),
			sqlgraph.To(officer.Table, officer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, inquiry.OfficerTable, inquiry.OfficerColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProduct chains the current query on the Product edge.
func (iq *InquiryQuery) QueryProduct() *ProductQuery {
	query := &ProductQuery{config: iq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(inquiry.Table, inquiry.FieldID, iq.sqlQuery()),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, inquiry.ProductTable, inquiry.ProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Inquiry entity in the query. Returns *NotFoundError when no inquiry was found.
func (iq *InquiryQuery) First(ctx context.Context) (*Inquiry, error) {
	is, err := iq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(is) == 0 {
		return nil, &NotFoundError{inquiry.Label}
	}
	return is[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InquiryQuery) FirstX(ctx context.Context) *Inquiry {
	i, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return i
}

// FirstID returns the first Inquiry id in the query. Returns *NotFoundError when no id was found.
func (iq *InquiryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{inquiry.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (iq *InquiryQuery) FirstXID(ctx context.Context) int {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Inquiry entity in the query, returns an error if not exactly one entity was returned.
func (iq *InquiryQuery) Only(ctx context.Context) (*Inquiry, error) {
	is, err := iq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(is) {
	case 1:
		return is[0], nil
	case 0:
		return nil, &NotFoundError{inquiry.Label}
	default:
		return nil, &NotSingularError{inquiry.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InquiryQuery) OnlyX(ctx context.Context) *Inquiry {
	i, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return i
}

// OnlyID returns the only Inquiry id in the query, returns an error if not exactly one id was returned.
func (iq *InquiryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{inquiry.Label}
	default:
		err = &NotSingularError{inquiry.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *InquiryQuery) OnlyIDX(ctx context.Context) int {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Inquiries.
func (iq *InquiryQuery) All(ctx context.Context) ([]*Inquiry, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return iq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (iq *InquiryQuery) AllX(ctx context.Context) []*Inquiry {
	is, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return is
}

// IDs executes the query and returns a list of Inquiry ids.
func (iq *InquiryQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := iq.Select(inquiry.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InquiryQuery) IDsX(ctx context.Context) []int {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InquiryQuery) Count(ctx context.Context) (int, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return iq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InquiryQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InquiryQuery) Exist(ctx context.Context) (bool, error) {
	if err := iq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return iq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InquiryQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InquiryQuery) Clone() *InquiryQuery {
	return &InquiryQuery{
		config:     iq.config,
		limit:      iq.limit,
		offset:     iq.offset,
		order:      append([]OrderFunc{}, iq.order...),
		unique:     append([]string{}, iq.unique...),
		predicates: append([]predicate.Inquiry{}, iq.predicates...),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

//  WithMember tells the query-builder to eager-loads the nodes that are connected to
// the "Member" edge. The optional arguments used to configure the query builder of the edge.
func (iq *InquiryQuery) WithMember(opts ...func(*MemberQuery)) *InquiryQuery {
	query := &MemberQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withMember = query
	return iq
}

//  WithCategory tells the query-builder to eager-loads the nodes that are connected to
// the "Category" edge. The optional arguments used to configure the query builder of the edge.
func (iq *InquiryQuery) WithCategory(opts ...func(*CategoryQuery)) *InquiryQuery {
	query := &CategoryQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withCategory = query
	return iq
}

//  WithOfficer tells the query-builder to eager-loads the nodes that are connected to
// the "Officer" edge. The optional arguments used to configure the query builder of the edge.
func (iq *InquiryQuery) WithOfficer(opts ...func(*OfficerQuery)) *InquiryQuery {
	query := &OfficerQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withOfficer = query
	return iq
}

//  WithProduct tells the query-builder to eager-loads the nodes that are connected to
// the "Product" edge. The optional arguments used to configure the query builder of the edge.
func (iq *InquiryQuery) WithProduct(opts ...func(*ProductQuery)) *InquiryQuery {
	query := &ProductQuery{config: iq.config}
	for _, opt := range opts {
		opt(query)
	}
	iq.withProduct = query
	return iq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		InquiryInguiryMessages string `json:"Inquiry_inguiry_messages,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Inquiry.Query().
//		GroupBy(inquiry.FieldInquiryInguiryMessages).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (iq *InquiryQuery) GroupBy(field string, fields ...string) *InquiryGroupBy {
	group := &InquiryGroupBy{config: iq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		InquiryInguiryMessages string `json:"Inquiry_inguiry_messages,omitempty"`
//	}
//
//	client.Inquiry.Query().
//		Select(inquiry.FieldInquiryInguiryMessages).
//		Scan(ctx, &v)
//
func (iq *InquiryQuery) Select(field string, fields ...string) *InquirySelect {
	selector := &InquirySelect{config: iq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return iq.sqlQuery(), nil
	}
	return selector
}

func (iq *InquiryQuery) prepareQuery(ctx context.Context) error {
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *InquiryQuery) sqlAll(ctx context.Context) ([]*Inquiry, error) {
	var (
		nodes       = []*Inquiry{}
		withFKs     = iq.withFKs
		_spec       = iq.querySpec()
		loadedTypes = [4]bool{
			iq.withMember != nil,
			iq.withCategory != nil,
			iq.withOfficer != nil,
			iq.withProduct != nil,
		}
	)
	if iq.withMember != nil || iq.withCategory != nil || iq.withOfficer != nil || iq.withProduct != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, inquiry.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Inquiry{config: iq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := iq.withMember; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Inquiry)
		for i := range nodes {
			if fk := nodes[i].member_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(member.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "member_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Member = n
			}
		}
	}

	if query := iq.withCategory; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Inquiry)
		for i := range nodes {
			if fk := nodes[i].category_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(category.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "category_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Category = n
			}
		}
	}

	if query := iq.withOfficer; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Inquiry)
		for i := range nodes {
			if fk := nodes[i].officer_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(officer.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "officer_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Officer = n
			}
		}
	}

	if query := iq.withProduct; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Inquiry)
		for i := range nodes {
			if fk := nodes[i].product_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(product.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "product_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Product = n
			}
		}
	}

	return nodes, nil
}

func (iq *InquiryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InquiryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := iq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (iq *InquiryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   inquiry.Table,
			Columns: inquiry.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: inquiry.FieldID,
			},
		},
		From:   iq.sql,
		Unique: true,
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *InquiryQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(inquiry.Table)
	selector := builder.Select(t1.Columns(inquiry.Columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(inquiry.Columns...)...)
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InquiryGroupBy is the builder for group-by Inquiry entities.
type InquiryGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InquiryGroupBy) Aggregate(fns ...AggregateFunc) *InquiryGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the group-by query and scan the result into the given value.
func (igb *InquiryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := igb.path(ctx)
	if err != nil {
		return err
	}
	igb.sql = query
	return igb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (igb *InquiryGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := igb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (igb *InquiryGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InquiryGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (igb *InquiryGroupBy) StringsX(ctx context.Context) []string {
	v, err := igb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (igb *InquiryGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = igb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{inquiry.Label}
	default:
		err = fmt.Errorf("ent: InquiryGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (igb *InquiryGroupBy) StringX(ctx context.Context) string {
	v, err := igb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (igb *InquiryGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InquiryGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (igb *InquiryGroupBy) IntsX(ctx context.Context) []int {
	v, err := igb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (igb *InquiryGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = igb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{inquiry.Label}
	default:
		err = fmt.Errorf("ent: InquiryGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (igb *InquiryGroupBy) IntX(ctx context.Context) int {
	v, err := igb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (igb *InquiryGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InquiryGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (igb *InquiryGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := igb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (igb *InquiryGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = igb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{inquiry.Label}
	default:
		err = fmt.Errorf("ent: InquiryGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (igb *InquiryGroupBy) Float64X(ctx context.Context) float64 {
	v, err := igb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (igb *InquiryGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(igb.fields) > 1 {
		return nil, errors.New("ent: InquiryGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := igb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (igb *InquiryGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := igb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (igb *InquiryGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = igb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{inquiry.Label}
	default:
		err = fmt.Errorf("ent: InquiryGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (igb *InquiryGroupBy) BoolX(ctx context.Context) bool {
	v, err := igb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (igb *InquiryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := igb.sqlQuery().Query()
	if err := igb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (igb *InquiryGroupBy) sqlQuery() *sql.Selector {
	selector := igb.sql
	columns := make([]string, 0, len(igb.fields)+len(igb.fns))
	columns = append(columns, igb.fields...)
	for _, fn := range igb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(igb.fields...)
}

// InquirySelect is the builder for select fields of Inquiry entities.
type InquirySelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (is *InquirySelect) Scan(ctx context.Context, v interface{}) error {
	query, err := is.path(ctx)
	if err != nil {
		return err
	}
	is.sql = query
	return is.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (is *InquirySelect) ScanX(ctx context.Context, v interface{}) {
	if err := is.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (is *InquirySelect) Strings(ctx context.Context) ([]string, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InquirySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (is *InquirySelect) StringsX(ctx context.Context) []string {
	v, err := is.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (is *InquirySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = is.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{inquiry.Label}
	default:
		err = fmt.Errorf("ent: InquirySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (is *InquirySelect) StringX(ctx context.Context) string {
	v, err := is.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (is *InquirySelect) Ints(ctx context.Context) ([]int, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InquirySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (is *InquirySelect) IntsX(ctx context.Context) []int {
	v, err := is.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (is *InquirySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = is.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{inquiry.Label}
	default:
		err = fmt.Errorf("ent: InquirySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (is *InquirySelect) IntX(ctx context.Context) int {
	v, err := is.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (is *InquirySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InquirySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (is *InquirySelect) Float64sX(ctx context.Context) []float64 {
	v, err := is.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (is *InquirySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = is.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{inquiry.Label}
	default:
		err = fmt.Errorf("ent: InquirySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (is *InquirySelect) Float64X(ctx context.Context) float64 {
	v, err := is.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (is *InquirySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(is.fields) > 1 {
		return nil, errors.New("ent: InquirySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := is.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (is *InquirySelect) BoolsX(ctx context.Context) []bool {
	v, err := is.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (is *InquirySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = is.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{inquiry.Label}
	default:
		err = fmt.Errorf("ent: InquirySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (is *InquirySelect) BoolX(ctx context.Context) bool {
	v, err := is.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (is *InquirySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := is.sqlQuery().Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (is *InquirySelect) sqlQuery() sql.Querier {
	selector := is.sql
	selector.Select(selector.Columns(is.fields...)...)
	return selector
}
