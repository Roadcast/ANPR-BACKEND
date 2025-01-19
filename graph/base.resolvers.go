package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"fmt"
	"go-ent-project/graph/model"
	"go-ent-project/internal/ent"

	"entgo.io/contrib/entgql"
)

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context, where *ent.UserWhereInput, after *string, first *int32, before *string, last *int32, orderBy *ent.UserOrder) (*model.UserConnection, error) {
	var afterCursor, beforeCursor *entgql.Cursor[int]
	if after != nil {
		cursor, err := DecodeCursor(*after)
		if err != nil {
			return nil, fmt.Errorf("invalid 'after' cursor: %v", err)
		}
		afterCursor = &entgql.Cursor[int]{Value: cursor}
	}

	if before != nil {
		cursor, err := DecodeCursor(*before)
		if err != nil {
			return nil, fmt.Errorf("invalid 'before' cursor: %v", err)
		}
		beforeCursor = &entgql.Cursor[int]{Value: cursor}
	}

	var firstInt, lastInt *int
	if first != nil {
		val := int(*first)
		firstInt = &val
	}
	if last != nil {
		val := int(*last)
		lastInt = &val
	}
	query := r.Client.User.Query()

	// Apply `where` filters if provided
	if where != nil {
		var err error
		query, err = where.Filter(query)
		if err != nil {
			return nil, fmt.Errorf("failed to apply filters: %v", err)
		}
	}

	paginate, err := query.Paginate(ctx, afterCursor, firstInt, beforeCursor, lastInt, ent.WithUserOrder(orderBy))
	if err != nil {
		return nil, err
	}
	var data []*ent.User
	for _, node := range paginate.Edges {
		data = append(data, node.Node)
	}

	startCursor := convertCursorToString(paginate.PageInfo.StartCursor)
	endCursor := convertCursorToString(paginate.PageInfo.EndCursor)

	return &model.UserConnection{
		PageInfo: &model.PageInfo{ // PageInfo struct
			HasNextPage:     paginate.PageInfo.HasNextPage,
			HasPreviousPage: paginate.PageInfo.HasPreviousPage,
			StartCursor:     startCursor,
			EndCursor:       endCursor,
		},
		Edges: data,
	}, nil
}

// GetPoliceStations is the resolver for the getPoliceStations field.
func (r *queryResolver) GetPoliceStations(ctx context.Context, where *ent.PoliceStationWhereInput, after *string, first *int32, before *string, last *int32, orderBy *ent.PoliceStationOrder) (*model.PoliceStationConnection, error) {
	var afterCursor, beforeCursor *entgql.Cursor[int]
	if after != nil {
		cursor, err := DecodeCursor(*after)
		if err != nil {
			return nil, fmt.Errorf("invalid 'after' cursor: %v", err)
		}
		afterCursor = &entgql.Cursor[int]{Value: cursor}
	}

	if before != nil {
		cursor, err := DecodeCursor(*before)
		if err != nil {
			return nil, fmt.Errorf("invalid 'before' cursor: %v", err)
		}
		beforeCursor = &entgql.Cursor[int]{Value: cursor}
	}

	var firstInt, lastInt *int
	if first != nil {
		val := int(*first)
		firstInt = &val
	}
	if last != nil {
		val := int(*last)
		lastInt = &val
	}
	query := r.Client.PoliceStation.Query()

	// Apply `where` filters if provided
	if where != nil {
		var err error
		query, err = where.Filter(query)
		if err != nil {
			return nil, fmt.Errorf("failed to apply filters: %v", err)
		}
	}

	paginate, err := query.Paginate(ctx, afterCursor, firstInt, beforeCursor, lastInt, ent.WithPoliceStationOrder(orderBy))
	if err != nil {
		return nil, err
	}
	var data []*ent.PoliceStation
	for _, node := range paginate.Edges {
		data = append(data, node.Node)
	}

	startCursor := convertCursorToString(paginate.PageInfo.StartCursor)
	endCursor := convertCursorToString(paginate.PageInfo.EndCursor)

	return &model.PoliceStationConnection{
		PageInfo: &model.PageInfo{ // PageInfo struct
			HasNextPage:     paginate.PageInfo.HasNextPage,
			HasPreviousPage: paginate.PageInfo.HasPreviousPage,
			StartCursor:     startCursor,
			EndCursor:       endCursor,
		},
		Edges: data,
	}, nil
}

// GetCameras is the resolver for the getCameras field.
func (r *queryResolver) GetCameras(ctx context.Context, where *ent.CameraWhereInput, after *string, first *int32, before *string, last *int32, orderBy *ent.CameraOrder) (*model.CameraConnection, error) {
	panic(fmt.Errorf("not implemented: GetCameras - getCameras"))
}
