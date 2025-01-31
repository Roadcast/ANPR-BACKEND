package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"fmt"
	"go-ent-project/internal/ent"
	"go-ent-project/utils/constant"

	"entgo.io/contrib/entgql"
	"github.com/google/uuid"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []uuid.UUID) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Cameras is the resolver for the cameras field.
func (r *queryResolver) Cameras(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy []*ent.CameraOrder, where *ent.CameraWhereInput) (*ent.CameraConnection, error) {
	fmt.Printf("where: %v\n", where)
	ctx = context.WithValue(ctx, constant.BypassPrivacyKey, true)
	paginate, err := r.Client.Camera.Query().Paginate(ctx, after, first, before, last, ent.WithCameraFilter(where.Filter), ent.WithCameraOrder(orderBy))
	fmt.Printf("paginate: %v\n", paginate)
	if err != nil {
		return nil, err
	}
	if paginate == nil {
		return nil, fmt.Errorf("no results found")

	}
	return paginate, nil
}

// Cars is the resolver for the cars field.
func (r *queryResolver) Cars(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy []*ent.CarOrder, where *ent.CarWhereInput) (*ent.CarConnection, error) {
	panic(fmt.Errorf("not implemented: Cars - cars"))
}

// Permissions is the resolver for the permissions field.
func (r *queryResolver) Permissions(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy []*ent.PermissionOrder, where *ent.PermissionWhereInput) (*ent.PermissionConnection, error) {
	panic(fmt.Errorf("not implemented: Permissions - permissions"))
}

// PoliceStations is the resolver for the policeStations field.
func (r *queryResolver) PoliceStations(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy []*ent.PoliceStationOrder, where *ent.PoliceStationWhereInput) (*ent.PoliceStationConnection, error) {
	fmt.Printf("where: %v\n", where)
	ctx = context.WithValue(ctx, constant.BypassPrivacyKey, true)
	paginate, err := r.Client.PoliceStation.Query().Paginate(ctx, after, first, before, last, ent.WithPoliceStationFilter(where.Filter), ent.WithPoliceStationOrder(orderBy))
	fmt.Printf("paginate: %v\n", paginate)
	if err != nil {
		return nil, err
	}
	if paginate == nil {
		return nil, fmt.Errorf("no results found")

	}
	return paginate, nil
}

// Roles is the resolver for the roles field.
func (r *queryResolver) Roles(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy []*ent.RoleOrder, where *ent.RoleWhereInput) (*ent.RoleConnection, error) {
	panic(fmt.Errorf("not implemented: Roles - roles"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy []*ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	fmt.Printf("where: %v\n", where)
	ctx = context.WithValue(ctx, constant.BypassPrivacyKey, true)
	paginate, err := r.Client.User.Query().Paginate(ctx, after, first, before, last, ent.WithUserFilter(where.Filter), ent.WithUserOrder(orderBy))
	fmt.Printf("paginate: %v\n", paginate)
	if err != nil {
		return nil, err
	}
	if paginate == nil {
		return nil, fmt.Errorf("no results found")

	}
	return paginate, nil
}

// VehicleDataSlice is the resolver for the vehicleDataSlice field.
func (r *queryResolver) VehicleDataSlice(ctx context.Context, after *entgql.Cursor[uuid.UUID], first *int, before *entgql.Cursor[uuid.UUID], last *int, orderBy []*ent.VehicleDataOrder, where *ent.VehicleDataWhereInput) (*ent.VehicleDataConnection, error) {
	fmt.Printf("where: %v\n", where)
	ctx = context.WithValue(ctx, constant.BypassPrivacyKey, true)
	paginate, err := r.Client.VehicleData.Query().Paginate(ctx, after, first, before, last, ent.WithVehicleDataFilter(where.Filter), ent.WithVehicleDataOrder(orderBy))
	fmt.Printf("paginate: %v\n", paginate)
	if err != nil {
		return nil, err
	}
	if paginate == nil {
		return nil, fmt.Errorf("no results found")

	}
	return paginate, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
