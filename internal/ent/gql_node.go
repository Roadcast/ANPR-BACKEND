// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-ent-project/internal/ent/camera"
	"go-ent-project/internal/ent/car"
	"go-ent-project/internal/ent/event"
	"go-ent-project/internal/ent/permission"
	"go-ent-project/internal/ent/policestation"
	"go-ent-project/internal/ent/role"
	"go-ent-project/internal/ent/user"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"github.com/hashicorp/go-multierror"
)

// Noder wraps the basic Node method.
type Noder interface {
	IsNode()
}

var cameraImplementors = []string{"Camera", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Camera) IsNode() {}

var carImplementors = []string{"Car", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Car) IsNode() {}

var eventImplementors = []string{"Event", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Event) IsNode() {}

var permissionImplementors = []string{"Permission", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Permission) IsNode() {}

var policestationImplementors = []string{"PoliceStation", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*PoliceStation) IsNode() {}

var roleImplementors = []string{"Role", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*Role) IsNode() {}

var userImplementors = []string{"User", "Node"}

// IsNode implements the Node interface check for GQLGen.
func (*User) IsNode() {}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, uuid.UUID) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, uuid.UUID) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, uuid.UUID) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id uuid.UUID) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id uuid.UUID, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id uuid.UUID) (Noder, error) {
	switch table {
	case camera.Table:
		query := c.Camera.Query().
			Where(camera.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, cameraImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case car.Table:
		query := c.Car.Query().
			Where(car.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, carImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case event.Table:
		query := c.Event.Query().
			Where(event.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, eventImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case permission.Table:
		query := c.Permission.Query().
			Where(permission.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, permissionImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case policestation.Table:
		query := c.PoliceStation.Query().
			Where(policestation.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, policestationImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case role.Table:
		query := c.Role.Query().
			Where(role.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, roleImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	case user.Table:
		query := c.User.Query().
			Where(user.ID(id))
		if fc := graphql.GetFieldContext(ctx); fc != nil {
			if err := query.collectField(ctx, true, graphql.GetOperationContext(ctx), fc.Field, nil, userImplementors...); err != nil {
				return nil, err
			}
		}
		return query.Only(ctx)
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []uuid.UUID, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]uuid.UUID)
	id2idx := make(map[uuid.UUID][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []uuid.UUID) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[uuid.UUID][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case camera.Table:
		query := c.Camera.Query().
			Where(camera.IDIn(ids...))
		query, err := query.CollectFields(ctx, cameraImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case car.Table:
		query := c.Car.Query().
			Where(car.IDIn(ids...))
		query, err := query.CollectFields(ctx, carImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case event.Table:
		query := c.Event.Query().
			Where(event.IDIn(ids...))
		query, err := query.CollectFields(ctx, eventImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case permission.Table:
		query := c.Permission.Query().
			Where(permission.IDIn(ids...))
		query, err := query.CollectFields(ctx, permissionImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case policestation.Table:
		query := c.PoliceStation.Query().
			Where(policestation.IDIn(ids...))
		query, err := query.CollectFields(ctx, policestationImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case role.Table:
		query := c.Role.Query().
			Where(role.IDIn(ids...))
		query, err := query.CollectFields(ctx, roleImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		query := c.User.Query().
			Where(user.IDIn(ids...))
		query, err := query.CollectFields(ctx, userImplementors...)
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}
