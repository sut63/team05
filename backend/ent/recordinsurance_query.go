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
	"github.com/sut63/team05/ent/amountpaid"
	"github.com/sut63/team05/ent/hospital"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/predicate"
	"github.com/sut63/team05/ent/product"
	"github.com/sut63/team05/ent/recordinsurance"
)

// RecordinsuranceQuery is the builder for querying Recordinsurance entities.
type RecordinsuranceQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Recordinsurance
	// eager-loading edges.
	withMember     *MemberQuery
	withHospital   *HospitalQuery
	withOfficer    *OfficerQuery
	withProduct    *ProductQuery
	withAmountpaid *AmountpaidQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (rq *RecordinsuranceQuery) Where(ps ...predicate.Recordinsurance) *RecordinsuranceQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit adds a limit step to the query.
func (rq *RecordinsuranceQuery) Limit(limit int) *RecordinsuranceQuery {
	rq.limit = &limit
	return rq
}

// Offset adds an offset step to the query.
func (rq *RecordinsuranceQuery) Offset(offset int) *RecordinsuranceQuery {
	rq.offset = &offset
	return rq
}

// Order adds an order step to the query.
func (rq *RecordinsuranceQuery) Order(o ...OrderFunc) *RecordinsuranceQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryMember chains the current query on the Member edge.
func (rq *RecordinsuranceQuery) QueryMember() *MemberQuery {
	query := &MemberQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(recordinsurance.Table, recordinsurance.FieldID, rq.sqlQuery()),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, recordinsurance.MemberTable, recordinsurance.MemberColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHospital chains the current query on the Hospital edge.
func (rq *RecordinsuranceQuery) QueryHospital() *HospitalQuery {
	query := &HospitalQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(recordinsurance.Table, recordinsurance.FieldID, rq.sqlQuery()),
			sqlgraph.To(hospital.Table, hospital.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, recordinsurance.HospitalTable, recordinsurance.HospitalColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOfficer chains the current query on the Officer edge.
func (rq *RecordinsuranceQuery) QueryOfficer() *OfficerQuery {
	query := &OfficerQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(recordinsurance.Table, recordinsurance.FieldID, rq.sqlQuery()),
			sqlgraph.To(officer.Table, officer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, recordinsurance.OfficerTable, recordinsurance.OfficerColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProduct chains the current query on the Product edge.
func (rq *RecordinsuranceQuery) QueryProduct() *ProductQuery {
	query := &ProductQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(recordinsurance.Table, recordinsurance.FieldID, rq.sqlQuery()),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, recordinsurance.ProductTable, recordinsurance.ProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAmountpaid chains the current query on the Amountpaid edge.
func (rq *RecordinsuranceQuery) QueryAmountpaid() *AmountpaidQuery {
	query := &AmountpaidQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(recordinsurance.Table, recordinsurance.FieldID, rq.sqlQuery()),
			sqlgraph.To(amountpaid.Table, amountpaid.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, recordinsurance.AmountpaidTable, recordinsurance.AmountpaidColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Recordinsurance entity in the query. Returns *NotFoundError when no recordinsurance was found.
func (rq *RecordinsuranceQuery) First(ctx context.Context) (*Recordinsurance, error) {
	rs, err := rq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(rs) == 0 {
		return nil, &NotFoundError{recordinsurance.Label}
	}
	return rs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RecordinsuranceQuery) FirstX(ctx context.Context) *Recordinsurance {
	r, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return r
}

// FirstID returns the first Recordinsurance id in the query. Returns *NotFoundError when no id was found.
func (rq *RecordinsuranceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{recordinsurance.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (rq *RecordinsuranceQuery) FirstXID(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Recordinsurance entity in the query, returns an error if not exactly one entity was returned.
func (rq *RecordinsuranceQuery) Only(ctx context.Context) (*Recordinsurance, error) {
	rs, err := rq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(rs) {
	case 1:
		return rs[0], nil
	case 0:
		return nil, &NotFoundError{recordinsurance.Label}
	default:
		return nil, &NotSingularError{recordinsurance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RecordinsuranceQuery) OnlyX(ctx context.Context) *Recordinsurance {
	r, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// OnlyID returns the only Recordinsurance id in the query, returns an error if not exactly one id was returned.
func (rq *RecordinsuranceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{recordinsurance.Label}
	default:
		err = &NotSingularError{recordinsurance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RecordinsuranceQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Recordinsurances.
func (rq *RecordinsuranceQuery) All(ctx context.Context) ([]*Recordinsurance, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rq *RecordinsuranceQuery) AllX(ctx context.Context) []*Recordinsurance {
	rs, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return rs
}

// IDs executes the query and returns a list of Recordinsurance ids.
func (rq *RecordinsuranceQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rq.Select(recordinsurance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RecordinsuranceQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RecordinsuranceQuery) Count(ctx context.Context) (int, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RecordinsuranceQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RecordinsuranceQuery) Exist(ctx context.Context) (bool, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RecordinsuranceQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RecordinsuranceQuery) Clone() *RecordinsuranceQuery {
	return &RecordinsuranceQuery{
		config:     rq.config,
		limit:      rq.limit,
		offset:     rq.offset,
		order:      append([]OrderFunc{}, rq.order...),
		unique:     append([]string{}, rq.unique...),
		predicates: append([]predicate.Recordinsurance{}, rq.predicates...),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

//  WithMember tells the query-builder to eager-loads the nodes that are connected to
// the "Member" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RecordinsuranceQuery) WithMember(opts ...func(*MemberQuery)) *RecordinsuranceQuery {
	query := &MemberQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withMember = query
	return rq
}

//  WithHospital tells the query-builder to eager-loads the nodes that are connected to
// the "Hospital" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RecordinsuranceQuery) WithHospital(opts ...func(*HospitalQuery)) *RecordinsuranceQuery {
	query := &HospitalQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withHospital = query
	return rq
}

//  WithOfficer tells the query-builder to eager-loads the nodes that are connected to
// the "Officer" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RecordinsuranceQuery) WithOfficer(opts ...func(*OfficerQuery)) *RecordinsuranceQuery {
	query := &OfficerQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withOfficer = query
	return rq
}

//  WithProduct tells the query-builder to eager-loads the nodes that are connected to
// the "Product" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RecordinsuranceQuery) WithProduct(opts ...func(*ProductQuery)) *RecordinsuranceQuery {
	query := &ProductQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withProduct = query
	return rq
}

//  WithAmountpaid tells the query-builder to eager-loads the nodes that are connected to
// the "Amountpaid" edge. The optional arguments used to configure the query builder of the edge.
func (rq *RecordinsuranceQuery) WithAmountpaid(opts ...func(*AmountpaidQuery)) *RecordinsuranceQuery {
	query := &AmountpaidQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withAmountpaid = query
	return rq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		NumberOfDaysOfTreat int `json:"number_of_days_of_treat,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Recordinsurance.Query().
//		GroupBy(recordinsurance.FieldNumberOfDaysOfTreat).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rq *RecordinsuranceQuery) GroupBy(field string, fields ...string) *RecordinsuranceGroupBy {
	group := &RecordinsuranceGroupBy{config: rq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		NumberOfDaysOfTreat int `json:"number_of_days_of_treat,omitempty"`
//	}
//
//	client.Recordinsurance.Query().
//		Select(recordinsurance.FieldNumberOfDaysOfTreat).
//		Scan(ctx, &v)
//
func (rq *RecordinsuranceQuery) Select(field string, fields ...string) *RecordinsuranceSelect {
	selector := &RecordinsuranceSelect{config: rq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(), nil
	}
	return selector
}

func (rq *RecordinsuranceQuery) prepareQuery(ctx context.Context) error {
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RecordinsuranceQuery) sqlAll(ctx context.Context) ([]*Recordinsurance, error) {
	var (
		nodes       = []*Recordinsurance{}
		withFKs     = rq.withFKs
		_spec       = rq.querySpec()
		loadedTypes = [5]bool{
			rq.withMember != nil,
			rq.withHospital != nil,
			rq.withOfficer != nil,
			rq.withProduct != nil,
			rq.withAmountpaid != nil,
		}
	)
	if rq.withMember != nil || rq.withHospital != nil || rq.withOfficer != nil || rq.withProduct != nil || rq.withAmountpaid != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, recordinsurance.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Recordinsurance{config: rq.config}
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
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rq.withMember; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Recordinsurance)
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

	if query := rq.withHospital; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Recordinsurance)
		for i := range nodes {
			if fk := nodes[i].hospital_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(hospital.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "hospital_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Hospital = n
			}
		}
	}

	if query := rq.withOfficer; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Recordinsurance)
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

	if query := rq.withProduct; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Recordinsurance)
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

	if query := rq.withAmountpaid; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Recordinsurance)
		for i := range nodes {
			if fk := nodes[i].amountpaid_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(amountpaid.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "amountpaid_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Amountpaid = n
			}
		}
	}

	return nodes, nil
}

func (rq *RecordinsuranceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RecordinsuranceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (rq *RecordinsuranceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recordinsurance.Table,
			Columns: recordinsurance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: recordinsurance.FieldID,
			},
		},
		From:   rq.sql,
		Unique: true,
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RecordinsuranceQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(recordinsurance.Table)
	selector := builder.Select(t1.Columns(recordinsurance.Columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(recordinsurance.Columns...)...)
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RecordinsuranceGroupBy is the builder for group-by Recordinsurance entities.
type RecordinsuranceGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RecordinsuranceGroupBy) Aggregate(fns ...AggregateFunc) *RecordinsuranceGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the group-by query and scan the result into the given value.
func (rgb *RecordinsuranceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rgb.path(ctx)
	if err != nil {
		return err
	}
	rgb.sql = query
	return rgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rgb *RecordinsuranceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (rgb *RecordinsuranceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RecordinsuranceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rgb *RecordinsuranceGroupBy) StringsX(ctx context.Context) []string {
	v, err := rgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (rgb *RecordinsuranceGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{recordinsurance.Label}
	default:
		err = fmt.Errorf("ent: RecordinsuranceGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rgb *RecordinsuranceGroupBy) StringX(ctx context.Context) string {
	v, err := rgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (rgb *RecordinsuranceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RecordinsuranceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rgb *RecordinsuranceGroupBy) IntsX(ctx context.Context) []int {
	v, err := rgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (rgb *RecordinsuranceGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{recordinsurance.Label}
	default:
		err = fmt.Errorf("ent: RecordinsuranceGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rgb *RecordinsuranceGroupBy) IntX(ctx context.Context) int {
	v, err := rgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (rgb *RecordinsuranceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RecordinsuranceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rgb *RecordinsuranceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (rgb *RecordinsuranceGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{recordinsurance.Label}
	default:
		err = fmt.Errorf("ent: RecordinsuranceGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rgb *RecordinsuranceGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (rgb *RecordinsuranceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RecordinsuranceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rgb *RecordinsuranceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (rgb *RecordinsuranceGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{recordinsurance.Label}
	default:
		err = fmt.Errorf("ent: RecordinsuranceGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rgb *RecordinsuranceGroupBy) BoolX(ctx context.Context) bool {
	v, err := rgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rgb *RecordinsuranceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rgb.sqlQuery().Query()
	if err := rgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rgb *RecordinsuranceGroupBy) sqlQuery() *sql.Selector {
	selector := rgb.sql
	columns := make([]string, 0, len(rgb.fields)+len(rgb.fns))
	columns = append(columns, rgb.fields...)
	for _, fn := range rgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(rgb.fields...)
}

// RecordinsuranceSelect is the builder for select fields of Recordinsurance entities.
type RecordinsuranceSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (rs *RecordinsuranceSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := rs.path(ctx)
	if err != nil {
		return err
	}
	rs.sql = query
	return rs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rs *RecordinsuranceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (rs *RecordinsuranceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RecordinsuranceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rs *RecordinsuranceSelect) StringsX(ctx context.Context) []string {
	v, err := rs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (rs *RecordinsuranceSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{recordinsurance.Label}
	default:
		err = fmt.Errorf("ent: RecordinsuranceSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rs *RecordinsuranceSelect) StringX(ctx context.Context) string {
	v, err := rs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (rs *RecordinsuranceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RecordinsuranceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rs *RecordinsuranceSelect) IntsX(ctx context.Context) []int {
	v, err := rs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (rs *RecordinsuranceSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{recordinsurance.Label}
	default:
		err = fmt.Errorf("ent: RecordinsuranceSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rs *RecordinsuranceSelect) IntX(ctx context.Context) int {
	v, err := rs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (rs *RecordinsuranceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RecordinsuranceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rs *RecordinsuranceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (rs *RecordinsuranceSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{recordinsurance.Label}
	default:
		err = fmt.Errorf("ent: RecordinsuranceSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rs *RecordinsuranceSelect) Float64X(ctx context.Context) float64 {
	v, err := rs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (rs *RecordinsuranceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RecordinsuranceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rs *RecordinsuranceSelect) BoolsX(ctx context.Context) []bool {
	v, err := rs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (rs *RecordinsuranceSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{recordinsurance.Label}
	default:
		err = fmt.Errorf("ent: RecordinsuranceSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rs *RecordinsuranceSelect) BoolX(ctx context.Context) bool {
	v, err := rs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rs *RecordinsuranceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rs.sqlQuery().Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rs *RecordinsuranceSelect) sqlQuery() sql.Querier {
	selector := rs.sql
	selector.Select(selector.Columns(rs.fields...)...)
	return selector
}
