// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-ent-project/internal/ent/predicate"
	"go-ent-project/internal/ent/vehicledata"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// VehicleDataQuery is the builder for querying VehicleData entities.
type VehicleDataQuery struct {
	config
	ctx        *QueryContext
	order      []vehicledata.OrderOption
	inters     []Interceptor
	predicates []predicate.VehicleData
	loadTotal  []func(context.Context, []*VehicleData) error
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VehicleDataQuery builder.
func (vdq *VehicleDataQuery) Where(ps ...predicate.VehicleData) *VehicleDataQuery {
	vdq.predicates = append(vdq.predicates, ps...)
	return vdq
}

// Limit the number of records to be returned by this query.
func (vdq *VehicleDataQuery) Limit(limit int) *VehicleDataQuery {
	vdq.ctx.Limit = &limit
	return vdq
}

// Offset to start from.
func (vdq *VehicleDataQuery) Offset(offset int) *VehicleDataQuery {
	vdq.ctx.Offset = &offset
	return vdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vdq *VehicleDataQuery) Unique(unique bool) *VehicleDataQuery {
	vdq.ctx.Unique = &unique
	return vdq
}

// Order specifies how the records should be ordered.
func (vdq *VehicleDataQuery) Order(o ...vehicledata.OrderOption) *VehicleDataQuery {
	vdq.order = append(vdq.order, o...)
	return vdq
}

// First returns the first VehicleData entity from the query.
// Returns a *NotFoundError when no VehicleData was found.
func (vdq *VehicleDataQuery) First(ctx context.Context) (*VehicleData, error) {
	nodes, err := vdq.Limit(1).All(setContextOp(ctx, vdq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{vehicledata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vdq *VehicleDataQuery) FirstX(ctx context.Context) *VehicleData {
	node, err := vdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VehicleData ID from the query.
// Returns a *NotFoundError when no VehicleData ID was found.
func (vdq *VehicleDataQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vdq.Limit(1).IDs(setContextOp(ctx, vdq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{vehicledata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vdq *VehicleDataQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := vdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VehicleData entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VehicleData entity is found.
// Returns a *NotFoundError when no VehicleData entities are found.
func (vdq *VehicleDataQuery) Only(ctx context.Context) (*VehicleData, error) {
	nodes, err := vdq.Limit(2).All(setContextOp(ctx, vdq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{vehicledata.Label}
	default:
		return nil, &NotSingularError{vehicledata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vdq *VehicleDataQuery) OnlyX(ctx context.Context) *VehicleData {
	node, err := vdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VehicleData ID in the query.
// Returns a *NotSingularError when more than one VehicleData ID is found.
// Returns a *NotFoundError when no entities are found.
func (vdq *VehicleDataQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vdq.Limit(2).IDs(setContextOp(ctx, vdq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{vehicledata.Label}
	default:
		err = &NotSingularError{vehicledata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vdq *VehicleDataQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := vdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VehicleDataSlice.
func (vdq *VehicleDataQuery) All(ctx context.Context) ([]*VehicleData, error) {
	ctx = setContextOp(ctx, vdq.ctx, ent.OpQueryAll)
	if err := vdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VehicleData, *VehicleDataQuery]()
	return withInterceptors[[]*VehicleData](ctx, vdq, qr, vdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vdq *VehicleDataQuery) AllX(ctx context.Context) []*VehicleData {
	nodes, err := vdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VehicleData IDs.
func (vdq *VehicleDataQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if vdq.ctx.Unique == nil && vdq.path != nil {
		vdq.Unique(true)
	}
	ctx = setContextOp(ctx, vdq.ctx, ent.OpQueryIDs)
	if err = vdq.Select(vehicledata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vdq *VehicleDataQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := vdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vdq *VehicleDataQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vdq.ctx, ent.OpQueryCount)
	if err := vdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vdq, querierCount[*VehicleDataQuery](), vdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vdq *VehicleDataQuery) CountX(ctx context.Context) int {
	count, err := vdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vdq *VehicleDataQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vdq.ctx, ent.OpQueryExist)
	switch _, err := vdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vdq *VehicleDataQuery) ExistX(ctx context.Context) bool {
	exist, err := vdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VehicleDataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vdq *VehicleDataQuery) Clone() *VehicleDataQuery {
	if vdq == nil {
		return nil
	}
	return &VehicleDataQuery{
		config:     vdq.config,
		ctx:        vdq.ctx.Clone(),
		order:      append([]vehicledata.OrderOption{}, vdq.order...),
		inters:     append([]Interceptor{}, vdq.inters...),
		predicates: append([]predicate.VehicleData{}, vdq.predicates...),
		// clone intermediate query.
		sql:  vdq.sql.Clone(),
		path: vdq.path,
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
//	client.VehicleData.Query().
//		GroupBy(vehicledata.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vdq *VehicleDataQuery) GroupBy(field string, fields ...string) *VehicleDataGroupBy {
	vdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VehicleDataGroupBy{build: vdq}
	grbuild.flds = &vdq.ctx.Fields
	grbuild.label = vehicledata.Label
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
//	client.VehicleData.Query().
//		Select(vehicledata.FieldCreatedAt).
//		Scan(ctx, &v)
func (vdq *VehicleDataQuery) Select(fields ...string) *VehicleDataSelect {
	vdq.ctx.Fields = append(vdq.ctx.Fields, fields...)
	sbuild := &VehicleDataSelect{VehicleDataQuery: vdq}
	sbuild.label = vehicledata.Label
	sbuild.flds, sbuild.scan = &vdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VehicleDataSelect configured with the given aggregations.
func (vdq *VehicleDataQuery) Aggregate(fns ...AggregateFunc) *VehicleDataSelect {
	return vdq.Select().Aggregate(fns...)
}

func (vdq *VehicleDataQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vdq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vdq); err != nil {
				return err
			}
		}
	}
	for _, f := range vdq.ctx.Fields {
		if !vehicledata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vdq.path != nil {
		prev, err := vdq.path(ctx)
		if err != nil {
			return err
		}
		vdq.sql = prev
	}
	return nil
}

func (vdq *VehicleDataQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VehicleData, error) {
	var (
		nodes = []*VehicleData{}
		_spec = vdq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VehicleData).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VehicleData{config: vdq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(vdq.modifiers) > 0 {
		_spec.Modifiers = vdq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range vdq.loadTotal {
		if err := vdq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vdq *VehicleDataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vdq.querySpec()
	if len(vdq.modifiers) > 0 {
		_spec.Modifiers = vdq.modifiers
	}
	_spec.Node.Columns = vdq.ctx.Fields
	if len(vdq.ctx.Fields) > 0 {
		_spec.Unique = vdq.ctx.Unique != nil && *vdq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vdq.driver, _spec)
}

func (vdq *VehicleDataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(vehicledata.Table, vehicledata.Columns, sqlgraph.NewFieldSpec(vehicledata.FieldID, field.TypeUUID))
	_spec.From = vdq.sql
	if unique := vdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vdq.path != nil {
		_spec.Unique = true
	}
	if fields := vdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vehicledata.FieldID)
		for i := range fields {
			if fields[i] != vehicledata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vdq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vdq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vdq *VehicleDataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vdq.driver.Dialect())
	t1 := builder.Table(vehicledata.Table)
	columns := vdq.ctx.Fields
	if len(columns) == 0 {
		columns = vehicledata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vdq.sql != nil {
		selector = vdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vdq.ctx.Unique != nil && *vdq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range vdq.modifiers {
		m(selector)
	}
	for _, p := range vdq.predicates {
		p(selector)
	}
	for _, p := range vdq.order {
		p(selector)
	}
	if offset := vdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (vdq *VehicleDataQuery) ForUpdate(opts ...sql.LockOption) *VehicleDataQuery {
	if vdq.driver.Dialect() == dialect.Postgres {
		vdq.Unique(false)
	}
	vdq.modifiers = append(vdq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return vdq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (vdq *VehicleDataQuery) ForShare(opts ...sql.LockOption) *VehicleDataQuery {
	if vdq.driver.Dialect() == dialect.Postgres {
		vdq.Unique(false)
	}
	vdq.modifiers = append(vdq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return vdq
}

// VehicleDataGroupBy is the group-by builder for VehicleData entities.
type VehicleDataGroupBy struct {
	selector
	build *VehicleDataQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vdgb *VehicleDataGroupBy) Aggregate(fns ...AggregateFunc) *VehicleDataGroupBy {
	vdgb.fns = append(vdgb.fns, fns...)
	return vdgb
}

// Scan applies the selector query and scans the result into the given value.
func (vdgb *VehicleDataGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vdgb.build.ctx, ent.OpQueryGroupBy)
	if err := vdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VehicleDataQuery, *VehicleDataGroupBy](ctx, vdgb.build, vdgb, vdgb.build.inters, v)
}

func (vdgb *VehicleDataGroupBy) sqlScan(ctx context.Context, root *VehicleDataQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vdgb.fns))
	for _, fn := range vdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vdgb.flds)+len(vdgb.fns))
		for _, f := range *vdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VehicleDataSelect is the builder for selecting fields of VehicleData entities.
type VehicleDataSelect struct {
	*VehicleDataQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vds *VehicleDataSelect) Aggregate(fns ...AggregateFunc) *VehicleDataSelect {
	vds.fns = append(vds.fns, fns...)
	return vds
}

// Scan applies the selector query and scans the result into the given value.
func (vds *VehicleDataSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vds.ctx, ent.OpQuerySelect)
	if err := vds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VehicleDataQuery, *VehicleDataSelect](ctx, vds.VehicleDataQuery, vds, vds.inters, v)
}

func (vds *VehicleDataSelect) sqlScan(ctx context.Context, root *VehicleDataQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vds.fns))
	for _, fn := range vds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
