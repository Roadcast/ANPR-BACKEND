package middleware

import (
	"context"
	"fmt"
	"go-ent-project/internal/ent"
	"go-ent-project/utils/auth"
	"go-ent-project/utils/constant"
	"net/http"
)

// AuthMiddleware validates the token and attaches the user to the context.
func AuthMiddleware(client *ent.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				// If no token is provided, continue without authentication
				next.ServeHTTP(w, r)
				return
			}

			token, _ := auth.ExtractTokenFromHeader(authHeader)
			ctx := context.WithValue(r.Context(), constant.BypassPrivacyKey, true)
			claims, err := auth.ValidateTokenWithDB(ctx, token, client)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid token: %v", err), http.StatusUnauthorized)
				return
			}

			// Fetch user details from the database
			authUser, err := client.User.Get(ctx, claims.UserID)
			if err != nil {
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			}

			fmt.Printf("User: %v\n", authUser)
			// Attach user to the context

			ctx = context.WithValue(ctx, constant.UserCtxKey, authUser)
			ctx = context.WithValue(ctx, constant.TokenContextKey, token)
			// Proceed with the next handler
			a := ctx.Value(constant.UserCtxKey)
			fmt.Printf("User: %v\n", a)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserFromContext(ctx context.Context) (*ent.User, error) {
	user, ok := ctx.Value(constant.UserCtxKey).(*ent.User)
	if !ok || user == nil {
		return nil, fmt.Errorf("user not found in context")
	}
	return user, nil
}

// ExtractTokenFromContext retrieves the token from the request context.
func ExtractTokenFromContext(ctx context.Context) (string, error) {
	token, ok := ctx.Value(constant.TokenContextKey).(string)
	if !ok || token == "" {
		return "", fmt.Errorf("token not found in context")
	}
	return token, nil
}

func isPublicOperation(operation string) bool {
	publicOperations := map[string]bool{
		"XXX":   true,
		"Login": true,
	}

	return publicOperations[operation]
}
