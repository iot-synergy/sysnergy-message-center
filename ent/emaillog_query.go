// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/synergy-message-center/ent/emaillog"
	"github.com/iot-synergy/synergy-message-center/ent/predicate"
)

// EmailLogQuery is the builder for querying EmailLog entities.
type EmailLogQuery struct {
	config
	ctx        *QueryContext
	order      []emaillog.OrderOption
	inters     []Interceptor
	predicates []predicate.EmailLog
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EmailLogQuery builder.
func (elq *EmailLogQuery) Where(ps ...predicate.EmailLog) *EmailLogQuery {
	elq.predicates = append(elq.predicates, ps...)
	return elq
}

// Limit the number of records to be returned by this query.
func (elq *EmailLogQuery) Limit(limit int) *EmailLogQuery {
	elq.ctx.Limit = &limit
	return elq
}

// Offset to start from.
func (elq *EmailLogQuery) Offset(offset int) *EmailLogQuery {
	elq.ctx.Offset = &offset
	return elq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (elq *EmailLogQuery) Unique(unique bool) *EmailLogQuery {
	elq.ctx.Unique = &unique
	return elq
}

// Order specifies how the records should be ordered.
func (elq *EmailLogQuery) Order(o ...emaillog.OrderOption) *EmailLogQuery {
	elq.order = append(elq.order, o...)
	return elq
}

// First returns the first EmailLog entity from the query.
// Returns a *NotFoundError when no EmailLog was found.
func (elq *EmailLogQuery) First(ctx context.Context) (*EmailLog, error) {
	nodes, err := elq.Limit(1).All(setContextOp(ctx, elq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{emaillog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (elq *EmailLogQuery) FirstX(ctx context.Context) *EmailLog {
	node, err := elq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EmailLog ID from the query.
// Returns a *NotFoundError when no EmailLog ID was found.
func (elq *EmailLogQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = elq.Limit(1).IDs(setContextOp(ctx, elq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{emaillog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (elq *EmailLogQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := elq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EmailLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EmailLog entity is found.
// Returns a *NotFoundError when no EmailLog entities are found.
func (elq *EmailLogQuery) Only(ctx context.Context) (*EmailLog, error) {
	nodes, err := elq.Limit(2).All(setContextOp(ctx, elq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{emaillog.Label}
	default:
		return nil, &NotSingularError{emaillog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (elq *EmailLogQuery) OnlyX(ctx context.Context) *EmailLog {
	node, err := elq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EmailLog ID in the query.
// Returns a *NotSingularError when more than one EmailLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (elq *EmailLogQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = elq.Limit(2).IDs(setContextOp(ctx, elq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{emaillog.Label}
	default:
		err = &NotSingularError{emaillog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (elq *EmailLogQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := elq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EmailLogs.
func (elq *EmailLogQuery) All(ctx context.Context) ([]*EmailLog, error) {
	ctx = setContextOp(ctx, elq.ctx, "All")
	if err := elq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EmailLog, *EmailLogQuery]()
	return withInterceptors[[]*EmailLog](ctx, elq, qr, elq.inters)
}

// AllX is like All, but panics if an error occurs.
func (elq *EmailLogQuery) AllX(ctx context.Context) []*EmailLog {
	nodes, err := elq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EmailLog IDs.
func (elq *EmailLogQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if elq.ctx.Unique == nil && elq.path != nil {
		elq.Unique(true)
	}
	ctx = setContextOp(ctx, elq.ctx, "IDs")
	if err = elq.Select(emaillog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (elq *EmailLogQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := elq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (elq *EmailLogQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, elq.ctx, "Count")
	if err := elq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, elq, querierCount[*EmailLogQuery](), elq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (elq *EmailLogQuery) CountX(ctx context.Context) int {
	count, err := elq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (elq *EmailLogQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, elq.ctx, "Exist")
	switch _, err := elq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (elq *EmailLogQuery) ExistX(ctx context.Context) bool {
	exist, err := elq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EmailLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (elq *EmailLogQuery) Clone() *EmailLogQuery {
	if elq == nil {
		return nil
	}
	return &EmailLogQuery{
		config:     elq.config,
		ctx:        elq.ctx.Clone(),
		order:      append([]emaillog.OrderOption{}, elq.order...),
		inters:     append([]Interceptor{}, elq.inters...),
		predicates: append([]predicate.EmailLog{}, elq.predicates...),
		// clone intermediate query.
		sql:  elq.sql.Clone(),
		path: elq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EmailLog.Query().
//		GroupBy(emaillog.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (elq *EmailLogQuery) GroupBy(field string, fields ...string) *EmailLogGroupBy {
	elq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EmailLogGroupBy{build: elq}
	grbuild.flds = &elq.ctx.Fields
	grbuild.label = emaillog.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.EmailLog.Query().
//		Select(emaillog.FieldCreatedAt).
//		Scan(ctx, &v)
func (elq *EmailLogQuery) Select(fields ...string) *EmailLogSelect {
	elq.ctx.Fields = append(elq.ctx.Fields, fields...)
	sbuild := &EmailLogSelect{EmailLogQuery: elq}
	sbuild.label = emaillog.Label
	sbuild.flds, sbuild.scan = &elq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EmailLogSelect configured with the given aggregations.
func (elq *EmailLogQuery) Aggregate(fns ...AggregateFunc) *EmailLogSelect {
	return elq.Select().Aggregate(fns...)
}

func (elq *EmailLogQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range elq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, elq); err != nil {
				return err
			}
		}
	}
	for _, f := range elq.ctx.Fields {
		if !emaillog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if elq.path != nil {
		prev, err := elq.path(ctx)
		if err != nil {
			return err
		}
		elq.sql = prev
	}
	return nil
}

func (elq *EmailLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EmailLog, error) {
	var (
		nodes = []*EmailLog{}
		_spec = elq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EmailLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EmailLog{config: elq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, elq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (elq *EmailLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := elq.querySpec()
	_spec.Node.Columns = elq.ctx.Fields
	if len(elq.ctx.Fields) > 0 {
		_spec.Unique = elq.ctx.Unique != nil && *elq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, elq.driver, _spec)
}

func (elq *EmailLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(emaillog.Table, emaillog.Columns, sqlgraph.NewFieldSpec(emaillog.FieldID, field.TypeUUID))
	_spec.From = elq.sql
	if unique := elq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if elq.path != nil {
		_spec.Unique = true
	}
	if fields := elq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emaillog.FieldID)
		for i := range fields {
			if fields[i] != emaillog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := elq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := elq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := elq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := elq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (elq *EmailLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(elq.driver.Dialect())
	t1 := builder.Table(emaillog.Table)
	columns := elq.ctx.Fields
	if len(columns) == 0 {
		columns = emaillog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if elq.sql != nil {
		selector = elq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if elq.ctx.Unique != nil && *elq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range elq.predicates {
		p(selector)
	}
	for _, p := range elq.order {
		p(selector)
	}
	if offset := elq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := elq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EmailLogGroupBy is the group-by builder for EmailLog entities.
type EmailLogGroupBy struct {
	selector
	build *EmailLogQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (elgb *EmailLogGroupBy) Aggregate(fns ...AggregateFunc) *EmailLogGroupBy {
	elgb.fns = append(elgb.fns, fns...)
	return elgb
}

// Scan applies the selector query and scans the result into the given value.
func (elgb *EmailLogGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, elgb.build.ctx, "GroupBy")
	if err := elgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmailLogQuery, *EmailLogGroupBy](ctx, elgb.build, elgb, elgb.build.inters, v)
}

func (elgb *EmailLogGroupBy) sqlScan(ctx context.Context, root *EmailLogQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(elgb.fns))
	for _, fn := range elgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*elgb.flds)+len(elgb.fns))
		for _, f := range *elgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*elgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := elgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EmailLogSelect is the builder for selecting fields of EmailLog entities.
type EmailLogSelect struct {
	*EmailLogQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (els *EmailLogSelect) Aggregate(fns ...AggregateFunc) *EmailLogSelect {
	els.fns = append(els.fns, fns...)
	return els
}

// Scan applies the selector query and scans the result into the given value.
func (els *EmailLogSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, els.ctx, "Select")
	if err := els.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmailLogQuery, *EmailLogSelect](ctx, els.EmailLogQuery, els, els.inters, v)
}

func (els *EmailLogSelect) sqlScan(ctx context.Context, root *EmailLogQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(els.fns))
	for _, fn := range els.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*els.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := els.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
