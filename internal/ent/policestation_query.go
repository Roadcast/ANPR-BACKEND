// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"go-ent-project/internal/ent/camera"
	"go-ent-project/internal/ent/policestation"
	"go-ent-project/internal/ent/predicate"
	"go-ent-project/internal/ent/user"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PoliceStationQuery is the builder for querying PoliceStation entities.
type PoliceStationQuery struct {
	config
	ctx                    *QueryContext
	order                  []policestation.OrderOption
	inters                 []Interceptor
	predicates             []predicate.PoliceStation
	withUsers              *UserQuery
	withCamera             *CameraQuery
	withParent             *PoliceStationQuery
	withChildStations      *PoliceStationQuery
	loadTotal              []func(context.Context, []*PoliceStation) error
	modifiers              []func(*sql.Selector)
	withNamedUsers         map[string]*UserQuery
	withNamedCamera        map[string]*CameraQuery
	withNamedChildStations map[string]*PoliceStationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PoliceStationQuery builder.
func (psq *PoliceStationQuery) Where(ps ...predicate.PoliceStation) *PoliceStationQuery {
	psq.predicates = append(psq.predicates, ps...)
	return psq
}

// Limit the number of records to be returned by this query.
func (psq *PoliceStationQuery) Limit(limit int) *PoliceStationQuery {
	psq.ctx.Limit = &limit
	return psq
}

// Offset to start from.
func (psq *PoliceStationQuery) Offset(offset int) *PoliceStationQuery {
	psq.ctx.Offset = &offset
	return psq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (psq *PoliceStationQuery) Unique(unique bool) *PoliceStationQuery {
	psq.ctx.Unique = &unique
	return psq
}

// Order specifies how the records should be ordered.
func (psq *PoliceStationQuery) Order(o ...policestation.OrderOption) *PoliceStationQuery {
	psq.order = append(psq.order, o...)
	return psq
}

// QueryUsers chains the current query on the "users" edge.
func (psq *PoliceStationQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: psq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(policestation.Table, policestation.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, policestation.UsersTable, policestation.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(psq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCamera chains the current query on the "camera" edge.
func (psq *PoliceStationQuery) QueryCamera() *CameraQuery {
	query := (&CameraClient{config: psq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(policestation.Table, policestation.FieldID, selector),
			sqlgraph.To(camera.Table, camera.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, policestation.CameraTable, policestation.CameraColumn),
		)
		fromU = sqlgraph.SetNeighbors(psq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParent chains the current query on the "parent" edge.
func (psq *PoliceStationQuery) QueryParent() *PoliceStationQuery {
	query := (&PoliceStationClient{config: psq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(policestation.Table, policestation.FieldID, selector),
			sqlgraph.To(policestation.Table, policestation.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, policestation.ParentTable, policestation.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(psq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildStations chains the current query on the "child_stations" edge.
func (psq *PoliceStationQuery) QueryChildStations() *PoliceStationQuery {
	query := (&PoliceStationClient{config: psq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := psq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := psq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(policestation.Table, policestation.FieldID, selector),
			sqlgraph.To(policestation.Table, policestation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, policestation.ChildStationsTable, policestation.ChildStationsColumn),
		)
		fromU = sqlgraph.SetNeighbors(psq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PoliceStation entity from the query.
// Returns a *NotFoundError when no PoliceStation was found.
func (psq *PoliceStationQuery) First(ctx context.Context) (*PoliceStation, error) {
	nodes, err := psq.Limit(1).All(setContextOp(ctx, psq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{policestation.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (psq *PoliceStationQuery) FirstX(ctx context.Context) *PoliceStation {
	node, err := psq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PoliceStation ID from the query.
// Returns a *NotFoundError when no PoliceStation ID was found.
func (psq *PoliceStationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = psq.Limit(1).IDs(setContextOp(ctx, psq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{policestation.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (psq *PoliceStationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := psq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PoliceStation entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PoliceStation entity is found.
// Returns a *NotFoundError when no PoliceStation entities are found.
func (psq *PoliceStationQuery) Only(ctx context.Context) (*PoliceStation, error) {
	nodes, err := psq.Limit(2).All(setContextOp(ctx, psq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{policestation.Label}
	default:
		return nil, &NotSingularError{policestation.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (psq *PoliceStationQuery) OnlyX(ctx context.Context) *PoliceStation {
	node, err := psq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PoliceStation ID in the query.
// Returns a *NotSingularError when more than one PoliceStation ID is found.
// Returns a *NotFoundError when no entities are found.
func (psq *PoliceStationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = psq.Limit(2).IDs(setContextOp(ctx, psq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{policestation.Label}
	default:
		err = &NotSingularError{policestation.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (psq *PoliceStationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := psq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PoliceStations.
func (psq *PoliceStationQuery) All(ctx context.Context) ([]*PoliceStation, error) {
	ctx = setContextOp(ctx, psq.ctx, ent.OpQueryAll)
	if err := psq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PoliceStation, *PoliceStationQuery]()
	return withInterceptors[[]*PoliceStation](ctx, psq, qr, psq.inters)
}

// AllX is like All, but panics if an error occurs.
func (psq *PoliceStationQuery) AllX(ctx context.Context) []*PoliceStation {
	nodes, err := psq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PoliceStation IDs.
func (psq *PoliceStationQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if psq.ctx.Unique == nil && psq.path != nil {
		psq.Unique(true)
	}
	ctx = setContextOp(ctx, psq.ctx, ent.OpQueryIDs)
	if err = psq.Select(policestation.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (psq *PoliceStationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := psq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (psq *PoliceStationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, psq.ctx, ent.OpQueryCount)
	if err := psq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, psq, querierCount[*PoliceStationQuery](), psq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (psq *PoliceStationQuery) CountX(ctx context.Context) int {
	count, err := psq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (psq *PoliceStationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, psq.ctx, ent.OpQueryExist)
	switch _, err := psq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (psq *PoliceStationQuery) ExistX(ctx context.Context) bool {
	exist, err := psq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PoliceStationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (psq *PoliceStationQuery) Clone() *PoliceStationQuery {
	if psq == nil {
		return nil
	}
	return &PoliceStationQuery{
		config:            psq.config,
		ctx:               psq.ctx.Clone(),
		order:             append([]policestation.OrderOption{}, psq.order...),
		inters:            append([]Interceptor{}, psq.inters...),
		predicates:        append([]predicate.PoliceStation{}, psq.predicates...),
		withUsers:         psq.withUsers.Clone(),
		withCamera:        psq.withCamera.Clone(),
		withParent:        psq.withParent.Clone(),
		withChildStations: psq.withChildStations.Clone(),
		// clone intermediate query.
		sql:  psq.sql.Clone(),
		path: psq.path,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (psq *PoliceStationQuery) WithUsers(opts ...func(*UserQuery)) *PoliceStationQuery {
	query := (&UserClient{config: psq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	psq.withUsers = query
	return psq
}

// WithCamera tells the query-builder to eager-load the nodes that are connected to
// the "camera" edge. The optional arguments are used to configure the query builder of the edge.
func (psq *PoliceStationQuery) WithCamera(opts ...func(*CameraQuery)) *PoliceStationQuery {
	query := (&CameraClient{config: psq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	psq.withCamera = query
	return psq
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (psq *PoliceStationQuery) WithParent(opts ...func(*PoliceStationQuery)) *PoliceStationQuery {
	query := (&PoliceStationClient{config: psq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	psq.withParent = query
	return psq
}

// WithChildStations tells the query-builder to eager-load the nodes that are connected to
// the "child_stations" edge. The optional arguments are used to configure the query builder of the edge.
func (psq *PoliceStationQuery) WithChildStations(opts ...func(*PoliceStationQuery)) *PoliceStationQuery {
	query := (&PoliceStationClient{config: psq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	psq.withChildStations = query
	return psq
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
//	client.PoliceStation.Query().
//		GroupBy(policestation.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (psq *PoliceStationQuery) GroupBy(field string, fields ...string) *PoliceStationGroupBy {
	psq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PoliceStationGroupBy{build: psq}
	grbuild.flds = &psq.ctx.Fields
	grbuild.label = policestation.Label
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
//	client.PoliceStation.Query().
//		Select(policestation.FieldCreatedAt).
//		Scan(ctx, &v)
func (psq *PoliceStationQuery) Select(fields ...string) *PoliceStationSelect {
	psq.ctx.Fields = append(psq.ctx.Fields, fields...)
	sbuild := &PoliceStationSelect{PoliceStationQuery: psq}
	sbuild.label = policestation.Label
	sbuild.flds, sbuild.scan = &psq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PoliceStationSelect configured with the given aggregations.
func (psq *PoliceStationQuery) Aggregate(fns ...AggregateFunc) *PoliceStationSelect {
	return psq.Select().Aggregate(fns...)
}

func (psq *PoliceStationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range psq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, psq); err != nil {
				return err
			}
		}
	}
	for _, f := range psq.ctx.Fields {
		if !policestation.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if psq.path != nil {
		prev, err := psq.path(ctx)
		if err != nil {
			return err
		}
		psq.sql = prev
	}
	return nil
}

func (psq *PoliceStationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PoliceStation, error) {
	var (
		nodes       = []*PoliceStation{}
		_spec       = psq.querySpec()
		loadedTypes = [4]bool{
			psq.withUsers != nil,
			psq.withCamera != nil,
			psq.withParent != nil,
			psq.withChildStations != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PoliceStation).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PoliceStation{config: psq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(psq.modifiers) > 0 {
		_spec.Modifiers = psq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, psq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := psq.withUsers; query != nil {
		if err := psq.loadUsers(ctx, query, nodes,
			func(n *PoliceStation) { n.Edges.Users = []*User{} },
			func(n *PoliceStation, e *User) {
				n.Edges.Users = append(n.Edges.Users, e)
				if !e.Edges.loadedTypes[1] {
					e.Edges.PoliceStation = n
				}
			}); err != nil {
			return nil, err
		}
	}
	if query := psq.withCamera; query != nil {
		if err := psq.loadCamera(ctx, query, nodes,
			func(n *PoliceStation) { n.Edges.Camera = []*Camera{} },
			func(n *PoliceStation, e *Camera) {
				n.Edges.Camera = append(n.Edges.Camera, e)
				if !e.Edges.loadedTypes[0] {
					e.Edges.PoliceStation = n
				}
			}); err != nil {
			return nil, err
		}
	}
	if query := psq.withParent; query != nil {
		if err := psq.loadParent(ctx, query, nodes, nil,
			func(n *PoliceStation, e *PoliceStation) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := psq.withChildStations; query != nil {
		if err := psq.loadChildStations(ctx, query, nodes,
			func(n *PoliceStation) { n.Edges.ChildStations = []*PoliceStation{} },
			func(n *PoliceStation, e *PoliceStation) {
				n.Edges.ChildStations = append(n.Edges.ChildStations, e)
				if !e.Edges.loadedTypes[2] {
					e.Edges.Parent = n
				}
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range psq.withNamedUsers {
		if err := psq.loadUsers(ctx, query, nodes,
			func(n *PoliceStation) { n.appendNamedUsers(name) },
			func(n *PoliceStation, e *User) {
				n.appendNamedUsers(name, e)
				if !e.Edges.loadedTypes[1] {
					e.Edges.PoliceStation = n
				}
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range psq.withNamedCamera {
		if err := psq.loadCamera(ctx, query, nodes,
			func(n *PoliceStation) { n.appendNamedCamera(name) },
			func(n *PoliceStation, e *Camera) {
				n.appendNamedCamera(name, e)
				if !e.Edges.loadedTypes[0] {
					e.Edges.PoliceStation = n
				}
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range psq.withNamedChildStations {
		if err := psq.loadChildStations(ctx, query, nodes,
			func(n *PoliceStation) { n.appendNamedChildStations(name) },
			func(n *PoliceStation, e *PoliceStation) {
				n.appendNamedChildStations(name, e)
				if !e.Edges.loadedTypes[2] {
					e.Edges.Parent = n
				}
			}); err != nil {
			return nil, err
		}
	}
	for i := range psq.loadTotal {
		if err := psq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (psq *PoliceStationQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*PoliceStation, init func(*PoliceStation), assign func(*PoliceStation, *User)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*PoliceStation)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(user.FieldPoliceStationID)
	}
	query.Where(predicate.User(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(policestation.UsersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PoliceStationID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "police_station_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (psq *PoliceStationQuery) loadCamera(ctx context.Context, query *CameraQuery, nodes []*PoliceStation, init func(*PoliceStation), assign func(*PoliceStation, *Camera)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*PoliceStation)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(camera.FieldPoliceStationID)
	}
	query.Where(predicate.Camera(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(policestation.CameraColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PoliceStationID
		if fk == nil {
			return fmt.Errorf(`foreign-key "police_station_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "police_station_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (psq *PoliceStationQuery) loadParent(ctx context.Context, query *PoliceStationQuery, nodes []*PoliceStation, init func(*PoliceStation), assign func(*PoliceStation, *PoliceStation)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*PoliceStation)
	for i := range nodes {
		if nodes[i].ParentStationID == nil {
			continue
		}
		fk := *nodes[i].ParentStationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(policestation.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "parent_station_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (psq *PoliceStationQuery) loadChildStations(ctx context.Context, query *PoliceStationQuery, nodes []*PoliceStation, init func(*PoliceStation), assign func(*PoliceStation, *PoliceStation)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*PoliceStation)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(policestation.FieldParentStationID)
	}
	query.Where(predicate.PoliceStation(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(policestation.ChildStationsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ParentStationID
		if fk == nil {
			return fmt.Errorf(`foreign-key "parent_station_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "parent_station_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (psq *PoliceStationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := psq.querySpec()
	if len(psq.modifiers) > 0 {
		_spec.Modifiers = psq.modifiers
	}
	_spec.Node.Columns = psq.ctx.Fields
	if len(psq.ctx.Fields) > 0 {
		_spec.Unique = psq.ctx.Unique != nil && *psq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, psq.driver, _spec)
}

func (psq *PoliceStationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(policestation.Table, policestation.Columns, sqlgraph.NewFieldSpec(policestation.FieldID, field.TypeUUID))
	_spec.From = psq.sql
	if unique := psq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if psq.path != nil {
		_spec.Unique = true
	}
	if fields := psq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, policestation.FieldID)
		for i := range fields {
			if fields[i] != policestation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if psq.withParent != nil {
			_spec.Node.AddColumnOnce(policestation.FieldParentStationID)
		}
	}
	if ps := psq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := psq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := psq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := psq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (psq *PoliceStationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(psq.driver.Dialect())
	t1 := builder.Table(policestation.Table)
	columns := psq.ctx.Fields
	if len(columns) == 0 {
		columns = policestation.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if psq.sql != nil {
		selector = psq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if psq.ctx.Unique != nil && *psq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range psq.modifiers {
		m(selector)
	}
	for _, p := range psq.predicates {
		p(selector)
	}
	for _, p := range psq.order {
		p(selector)
	}
	if offset := psq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := psq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (psq *PoliceStationQuery) ForUpdate(opts ...sql.LockOption) *PoliceStationQuery {
	if psq.driver.Dialect() == dialect.Postgres {
		psq.Unique(false)
	}
	psq.modifiers = append(psq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return psq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (psq *PoliceStationQuery) ForShare(opts ...sql.LockOption) *PoliceStationQuery {
	if psq.driver.Dialect() == dialect.Postgres {
		psq.Unique(false)
	}
	psq.modifiers = append(psq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return psq
}

// WithNamedUsers tells the query-builder to eager-load the nodes that are connected to the "users"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (psq *PoliceStationQuery) WithNamedUsers(name string, opts ...func(*UserQuery)) *PoliceStationQuery {
	query := (&UserClient{config: psq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if psq.withNamedUsers == nil {
		psq.withNamedUsers = make(map[string]*UserQuery)
	}
	psq.withNamedUsers[name] = query
	return psq
}

// WithNamedCamera tells the query-builder to eager-load the nodes that are connected to the "camera"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (psq *PoliceStationQuery) WithNamedCamera(name string, opts ...func(*CameraQuery)) *PoliceStationQuery {
	query := (&CameraClient{config: psq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if psq.withNamedCamera == nil {
		psq.withNamedCamera = make(map[string]*CameraQuery)
	}
	psq.withNamedCamera[name] = query
	return psq
}

// WithNamedChildStations tells the query-builder to eager-load the nodes that are connected to the "child_stations"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (psq *PoliceStationQuery) WithNamedChildStations(name string, opts ...func(*PoliceStationQuery)) *PoliceStationQuery {
	query := (&PoliceStationClient{config: psq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if psq.withNamedChildStations == nil {
		psq.withNamedChildStations = make(map[string]*PoliceStationQuery)
	}
	psq.withNamedChildStations[name] = query
	return psq
}

// PoliceStationGroupBy is the group-by builder for PoliceStation entities.
type PoliceStationGroupBy struct {
	selector
	build *PoliceStationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (psgb *PoliceStationGroupBy) Aggregate(fns ...AggregateFunc) *PoliceStationGroupBy {
	psgb.fns = append(psgb.fns, fns...)
	return psgb
}

// Scan applies the selector query and scans the result into the given value.
func (psgb *PoliceStationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, psgb.build.ctx, ent.OpQueryGroupBy)
	if err := psgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PoliceStationQuery, *PoliceStationGroupBy](ctx, psgb.build, psgb, psgb.build.inters, v)
}

func (psgb *PoliceStationGroupBy) sqlScan(ctx context.Context, root *PoliceStationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(psgb.fns))
	for _, fn := range psgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*psgb.flds)+len(psgb.fns))
		for _, f := range *psgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*psgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := psgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PoliceStationSelect is the builder for selecting fields of PoliceStation entities.
type PoliceStationSelect struct {
	*PoliceStationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pss *PoliceStationSelect) Aggregate(fns ...AggregateFunc) *PoliceStationSelect {
	pss.fns = append(pss.fns, fns...)
	return pss
}

// Scan applies the selector query and scans the result into the given value.
func (pss *PoliceStationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pss.ctx, ent.OpQuerySelect)
	if err := pss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PoliceStationQuery, *PoliceStationSelect](ctx, pss.PoliceStationQuery, pss, pss.inters, v)
}

func (pss *PoliceStationSelect) sqlScan(ctx context.Context, root *PoliceStationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pss.fns))
	for _, fn := range pss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
