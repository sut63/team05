// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team05/ent/inquiry"
	"github.com/sut63/team05/ent/insurance"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/payback"
	"github.com/sut63/team05/ent/payment"
	"github.com/sut63/team05/ent/predicate"
	"github.com/sut63/team05/ent/recordinsurance"
)

// MemberQuery is the builder for querying Member entities.
type MemberQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Member
	// eager-loading edges.
	withMemberInsurance       *InsuranceQuery
	withMemberPayment         *PaymentQuery
	withMemberInquiry         *InquiryQuery
	withMemberPayback         *PaybackQuery
	withMemberRecordinsurance *RecordinsuranceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (mq *MemberQuery) Where(ps ...predicate.Member) *MemberQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit adds a limit step to the query.
func (mq *MemberQuery) Limit(limit int) *MemberQuery {
	mq.limit = &limit
	return mq
}

// Offset adds an offset step to the query.
func (mq *MemberQuery) Offset(offset int) *MemberQuery {
	mq.offset = &offset
	return mq
}

// Order adds an order step to the query.
func (mq *MemberQuery) Order(o ...OrderFunc) *MemberQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryMemberInsurance chains the current query on the member_insurance edge.
func (mq *MemberQuery) QueryMemberInsurance() *InsuranceQuery {
	query := &InsuranceQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, mq.sqlQuery()),
			sqlgraph.To(insurance.Table, insurance.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.MemberInsuranceTable, member.MemberInsuranceColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMemberPayment chains the current query on the member_payment edge.
func (mq *MemberQuery) QueryMemberPayment() *PaymentQuery {
	query := &PaymentQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, mq.sqlQuery()),
			sqlgraph.To(payment.Table, payment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.MemberPaymentTable, member.MemberPaymentColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMemberInquiry chains the current query on the member_inquiry edge.
func (mq *MemberQuery) QueryMemberInquiry() *InquiryQuery {
	query := &InquiryQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, mq.sqlQuery()),
			sqlgraph.To(inquiry.Table, inquiry.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.MemberInquiryTable, member.MemberInquiryColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMemberPayback chains the current query on the member_payback edge.
func (mq *MemberQuery) QueryMemberPayback() *PaybackQuery {
	query := &PaybackQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, mq.sqlQuery()),
			sqlgraph.To(payback.Table, payback.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.MemberPaybackTable, member.MemberPaybackColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMemberRecordinsurance chains the current query on the member_recordinsurance edge.
func (mq *MemberQuery) QueryMemberRecordinsurance() *RecordinsuranceQuery {
	query := &RecordinsuranceQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(member.Table, member.FieldID, mq.sqlQuery()),
			sqlgraph.To(recordinsurance.Table, recordinsurance.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, member.MemberRecordinsuranceTable, member.MemberRecordinsuranceColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Member entity in the query. Returns *NotFoundError when no member was found.
func (mq *MemberQuery) First(ctx context.Context) (*Member, error) {
	ms, err := mq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ms) == 0 {
		return nil, &NotFoundError{member.Label}
	}
	return ms[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MemberQuery) FirstX(ctx context.Context) *Member {
	m, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return m
}

// FirstID returns the first Member id in the query. Returns *NotFoundError when no id was found.
func (mq *MemberQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{member.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (mq *MemberQuery) FirstXID(ctx context.Context) int {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Member entity in the query, returns an error if not exactly one entity was returned.
func (mq *MemberQuery) Only(ctx context.Context) (*Member, error) {
	ms, err := mq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ms) {
	case 1:
		return ms[0], nil
	case 0:
		return nil, &NotFoundError{member.Label}
	default:
		return nil, &NotSingularError{member.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MemberQuery) OnlyX(ctx context.Context) *Member {
	m, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// OnlyID returns the only Member id in the query, returns an error if not exactly one id was returned.
func (mq *MemberQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = mq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{member.Label}
	default:
		err = &NotSingularError{member.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MemberQuery) OnlyIDX(ctx context.Context) int {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Members.
func (mq *MemberQuery) All(ctx context.Context) ([]*Member, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mq *MemberQuery) AllX(ctx context.Context) []*Member {
	ms, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ms
}

// IDs executes the query and returns a list of Member ids.
func (mq *MemberQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := mq.Select(member.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MemberQuery) IDsX(ctx context.Context) []int {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MemberQuery) Count(ctx context.Context) (int, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MemberQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MemberQuery) Exist(ctx context.Context) (bool, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MemberQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MemberQuery) Clone() *MemberQuery {
	return &MemberQuery{
		config:     mq.config,
		limit:      mq.limit,
		offset:     mq.offset,
		order:      append([]OrderFunc{}, mq.order...),
		unique:     append([]string{}, mq.unique...),
		predicates: append([]predicate.Member{}, mq.predicates...),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

//  WithMemberInsurance tells the query-builder to eager-loads the nodes that are connected to
// the "member_insurance" edge. The optional arguments used to configure the query builder of the edge.
func (mq *MemberQuery) WithMemberInsurance(opts ...func(*InsuranceQuery)) *MemberQuery {
	query := &InsuranceQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withMemberInsurance = query
	return mq
}

//  WithMemberPayment tells the query-builder to eager-loads the nodes that are connected to
// the "member_payment" edge. The optional arguments used to configure the query builder of the edge.
func (mq *MemberQuery) WithMemberPayment(opts ...func(*PaymentQuery)) *MemberQuery {
	query := &PaymentQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withMemberPayment = query
	return mq
}

//  WithMemberInquiry tells the query-builder to eager-loads the nodes that are connected to
// the "member_inquiry" edge. The optional arguments used to configure the query builder of the edge.
func (mq *MemberQuery) WithMemberInquiry(opts ...func(*InquiryQuery)) *MemberQuery {
	query := &InquiryQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withMemberInquiry = query
	return mq
}

//  WithMemberPayback tells the query-builder to eager-loads the nodes that are connected to
// the "member_payback" edge. The optional arguments used to configure the query builder of the edge.
func (mq *MemberQuery) WithMemberPayback(opts ...func(*PaybackQuery)) *MemberQuery {
	query := &PaybackQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withMemberPayback = query
	return mq
}

//  WithMemberRecordinsurance tells the query-builder to eager-loads the nodes that are connected to
// the "member_recordinsurance" edge. The optional arguments used to configure the query builder of the edge.
func (mq *MemberQuery) WithMemberRecordinsurance(opts ...func(*RecordinsuranceQuery)) *MemberQuery {
	query := &RecordinsuranceQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withMemberRecordinsurance = query
	return mq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MemberEmail string `json:"member_email,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Member.Query().
//		GroupBy(member.FieldMemberEmail).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mq *MemberQuery) GroupBy(field string, fields ...string) *MemberGroupBy {
	group := &MemberGroupBy{config: mq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		MemberEmail string `json:"member_email,omitempty"`
//	}
//
//	client.Member.Query().
//		Select(member.FieldMemberEmail).
//		Scan(ctx, &v)
//
func (mq *MemberQuery) Select(field string, fields ...string) *MemberSelect {
	selector := &MemberSelect{config: mq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(), nil
	}
	return selector
}

func (mq *MemberQuery) prepareQuery(ctx context.Context) error {
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MemberQuery) sqlAll(ctx context.Context) ([]*Member, error) {
	var (
		nodes       = []*Member{}
		_spec       = mq.querySpec()
		loadedTypes = [5]bool{
			mq.withMemberInsurance != nil,
			mq.withMemberPayment != nil,
			mq.withMemberInquiry != nil,
			mq.withMemberPayback != nil,
			mq.withMemberRecordinsurance != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Member{config: mq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
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
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := mq.withMemberInsurance; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Member)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Insurance(func(s *sql.Selector) {
			s.Where(sql.InValues(member.MemberInsuranceColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.member_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "member_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "member_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.MemberInsurance = append(node.Edges.MemberInsurance, n)
		}
	}

	if query := mq.withMemberPayment; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Member)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Payment(func(s *sql.Selector) {
			s.Where(sql.InValues(member.MemberPaymentColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.member_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "member_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "member_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.MemberPayment = append(node.Edges.MemberPayment, n)
		}
	}

	if query := mq.withMemberInquiry; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Member)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Inquiry(func(s *sql.Selector) {
			s.Where(sql.InValues(member.MemberInquiryColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.member_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "member_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "member_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.MemberInquiry = append(node.Edges.MemberInquiry, n)
		}
	}

	if query := mq.withMemberPayback; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Member)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Payback(func(s *sql.Selector) {
			s.Where(sql.InValues(member.MemberPaybackColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.member_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "member_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "member_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.MemberPayback = append(node.Edges.MemberPayback, n)
		}
	}

	if query := mq.withMemberRecordinsurance; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Member)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Recordinsurance(func(s *sql.Selector) {
			s.Where(sql.InValues(member.MemberRecordinsuranceColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.member_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "member_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "member_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.MemberRecordinsurance = append(node.Edges.MemberRecordinsurance, n)
		}
	}

	return nodes, nil
}

func (mq *MemberQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MemberQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (mq *MemberQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   member.Table,
			Columns: member.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: member.FieldID,
			},
		},
		From:   mq.sql,
		Unique: true,
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MemberQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(member.Table)
	selector := builder.Select(t1.Columns(member.Columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(member.Columns...)...)
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MemberGroupBy is the builder for group-by Member entities.
type MemberGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MemberGroupBy) Aggregate(fns ...AggregateFunc) *MemberGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the group-by query and scan the result into the given value.
func (mgb *MemberGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mgb.path(ctx)
	if err != nil {
		return err
	}
	mgb.sql = query
	return mgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mgb *MemberGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (mgb *MemberGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MemberGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mgb *MemberGroupBy) StringsX(ctx context.Context) []string {
	v, err := mgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (mgb *MemberGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{member.Label}
	default:
		err = fmt.Errorf("ent: MemberGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mgb *MemberGroupBy) StringX(ctx context.Context) string {
	v, err := mgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (mgb *MemberGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MemberGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mgb *MemberGroupBy) IntsX(ctx context.Context) []int {
	v, err := mgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (mgb *MemberGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{member.Label}
	default:
		err = fmt.Errorf("ent: MemberGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mgb *MemberGroupBy) IntX(ctx context.Context) int {
	v, err := mgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (mgb *MemberGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MemberGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mgb *MemberGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (mgb *MemberGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{member.Label}
	default:
		err = fmt.Errorf("ent: MemberGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mgb *MemberGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (mgb *MemberGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mgb.fields) > 1 {
		return nil, errors.New("ent: MemberGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mgb *MemberGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (mgb *MemberGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{member.Label}
	default:
		err = fmt.Errorf("ent: MemberGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mgb *MemberGroupBy) BoolX(ctx context.Context) bool {
	v, err := mgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mgb *MemberGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mgb.sqlQuery().Query()
	if err := mgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mgb *MemberGroupBy) sqlQuery() *sql.Selector {
	selector := mgb.sql
	columns := make([]string, 0, len(mgb.fields)+len(mgb.fns))
	columns = append(columns, mgb.fields...)
	for _, fn := range mgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(mgb.fields...)
}

// MemberSelect is the builder for select fields of Member entities.
type MemberSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ms *MemberSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ms.path(ctx)
	if err != nil {
		return err
	}
	ms.sql = query
	return ms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ms *MemberSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ms *MemberSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MemberSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ms *MemberSelect) StringsX(ctx context.Context) []string {
	v, err := ms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ms *MemberSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{member.Label}
	default:
		err = fmt.Errorf("ent: MemberSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ms *MemberSelect) StringX(ctx context.Context) string {
	v, err := ms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ms *MemberSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MemberSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ms *MemberSelect) IntsX(ctx context.Context) []int {
	v, err := ms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ms *MemberSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{member.Label}
	default:
		err = fmt.Errorf("ent: MemberSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ms *MemberSelect) IntX(ctx context.Context) int {
	v, err := ms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ms *MemberSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MemberSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ms *MemberSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ms *MemberSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{member.Label}
	default:
		err = fmt.Errorf("ent: MemberSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ms *MemberSelect) Float64X(ctx context.Context) float64 {
	v, err := ms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ms *MemberSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ms.fields) > 1 {
		return nil, errors.New("ent: MemberSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ms *MemberSelect) BoolsX(ctx context.Context) []bool {
	v, err := ms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ms *MemberSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{member.Label}
	default:
		err = fmt.Errorf("ent: MemberSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ms *MemberSelect) BoolX(ctx context.Context) bool {
	v, err := ms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ms *MemberSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ms.sqlQuery().Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ms *MemberSelect) sqlQuery() sql.Querier {
	selector := ms.sql
	selector.Select(selector.Columns(ms.fields...)...)
	return selector
}
