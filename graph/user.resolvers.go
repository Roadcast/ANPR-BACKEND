package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"fmt"
	"go-ent-project/internal/ent"
	"go-ent-project/internal/ent/user"
	"strconv"
)

// UpdateUserPassword is the resolver for the update_user_password field.
func (r *mutationResolver) UpdateUserPassword(ctx context.Context, userID string, password string) (*ent.User, error) {
	// convert userID to int
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user id")
	}
	authUser, err := r.Client.User.Query().Where(user.IDEQ(userIDInt)).Only(ctx)
	fmt.Printf("authUser: %v\n", authUser)
	if err != nil {
		return nil, fmt.Errorf("invalid email")
	}
	user1, err := authUser.Update().SetPassword(password).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update password")
	}
	fmt.Printf("user1: %v\n", user1)
	return user1, nil
}
