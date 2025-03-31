// Code generated by ent, DO NOT EDIT.

package ent

import (
	"awesomeProject/ent/predicate"
	"awesomeProject/ent/subtask"
	"awesomeProject/ent/task"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubtaskQuery is the builder for querying Subtask entities.
type SubtaskQuery struct {
	config
	ctx        *QueryContext
	order      []subtask.OrderOption
	inters     []Interceptor
	predicates []predicate.Subtask
	withTask   *TaskQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SubtaskQuery builder.
func (sq *SubtaskQuery) Where(ps ...predicate.Subtask) *SubtaskQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SubtaskQuery) Limit(limit int) *SubtaskQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SubtaskQuery) Offset(offset int) *SubtaskQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SubtaskQuery) Unique(unique bool) *SubtaskQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SubtaskQuery) Order(o ...subtask.OrderOption) *SubtaskQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryTask chains the current query on the "task" edge.
func (sq *SubtaskQuery) QueryTask() *TaskQuery {
	query := (&TaskClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(subtask.Table, subtask.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, subtask.TaskTable, subtask.TaskColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Subtask entity from the query.
// Returns a *NotFoundError when no Subtask was found.
func (sq *SubtaskQuery) First(ctx context.Context) (*Subtask, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{subtask.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SubtaskQuery) FirstX(ctx context.Context) *Subtask {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Subtask ID from the query.
// Returns a *NotFoundError when no Subtask ID was found.
func (sq *SubtaskQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{subtask.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SubtaskQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Subtask entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Subtask entity is found.
// Returns a *NotFoundError when no Subtask entities are found.
func (sq *SubtaskQuery) Only(ctx context.Context) (*Subtask, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{subtask.Label}
	default:
		return nil, &NotSingularError{subtask.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SubtaskQuery) OnlyX(ctx context.Context) *Subtask {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Subtask ID in the query.
// Returns a *NotSingularError when more than one Subtask ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SubtaskQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{subtask.Label}
	default:
		err = &NotSingularError{subtask.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SubtaskQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Subtasks.
func (sq *SubtaskQuery) All(ctx context.Context) ([]*Subtask, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryAll)
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Subtask, *SubtaskQuery]()
	return withInterceptors[[]*Subtask](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SubtaskQuery) AllX(ctx context.Context) []*Subtask {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Subtask IDs.
func (sq *SubtaskQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryIDs)
	if err = sq.Select(subtask.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SubtaskQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SubtaskQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryCount)
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SubtaskQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SubtaskQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SubtaskQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, ent.OpQueryExist)
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SubtaskQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SubtaskQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SubtaskQuery) Clone() *SubtaskQuery {
	if sq == nil {
		return nil
	}
	return &SubtaskQuery{
		config:     sq.config,
		ctx:        sq.ctx.Clone(),
		order:      append([]subtask.OrderOption{}, sq.order...),
		inters:     append([]Interceptor{}, sq.inters...),
		predicates: append([]predicate.Subtask{}, sq.predicates...),
		withTask:   sq.withTask.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithTask tells the query-builder to eager-load the nodes that are connected to
// the "task" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SubtaskQuery) WithTask(opts ...func(*TaskQuery)) *SubtaskQuery {
	query := (&TaskClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withTask = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Subtask.Query().
//		GroupBy(subtask.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SubtaskQuery) GroupBy(field string, fields ...string) *SubtaskGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SubtaskGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = subtask.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Subtask.Query().
//		Select(subtask.FieldTitle).
//		Scan(ctx, &v)
func (sq *SubtaskQuery) Select(fields ...string) *SubtaskSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SubtaskSelect{SubtaskQuery: sq}
	sbuild.label = subtask.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SubtaskSelect configured with the given aggregations.
func (sq *SubtaskQuery) Aggregate(fns ...AggregateFunc) *SubtaskSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SubtaskQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !subtask.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SubtaskQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Subtask, error) {
	var (
		nodes       = []*Subtask{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withTask != nil,
		}
	)
	if sq.withTask != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, subtask.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Subtask).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Subtask{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withTask; query != nil {
		if err := sq.loadTask(ctx, query, nodes, nil,
			func(n *Subtask, e *Task) { n.Edges.Task = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SubtaskQuery) loadTask(ctx context.Context, query *TaskQuery, nodes []*Subtask, init func(*Subtask), assign func(*Subtask, *Task)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Subtask)
	for i := range nodes {
		if nodes[i].task_subtasks == nil {
			continue
		}
		fk := *nodes[i].task_subtasks
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(task.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "task_subtasks" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sq *SubtaskQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SubtaskQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(subtask.Table, subtask.Columns, sqlgraph.NewFieldSpec(subtask.FieldID, field.TypeInt))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subtask.FieldID)
		for i := range fields {
			if fields[i] != subtask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SubtaskQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(subtask.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = subtask.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SubtaskGroupBy is the group-by builder for Subtask entities.
type SubtaskGroupBy struct {
	selector
	build *SubtaskQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SubtaskGroupBy) Aggregate(fns ...AggregateFunc) *SubtaskGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SubtaskGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, ent.OpQueryGroupBy)
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubtaskQuery, *SubtaskGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SubtaskGroupBy) sqlScan(ctx context.Context, root *SubtaskQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SubtaskSelect is the builder for selecting fields of Subtask entities.
type SubtaskSelect struct {
	*SubtaskQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SubtaskSelect) Aggregate(fns ...AggregateFunc) *SubtaskSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SubtaskSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, ent.OpQuerySelect)
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubtaskQuery, *SubtaskSelect](ctx, ss.SubtaskQuery, ss, ss.inters, v)
}

func (ss *SubtaskSelect) sqlScan(ctx context.Context, root *SubtaskQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
