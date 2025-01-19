package middleware

import (
	"context"
	"fmt"
	"go-ent-project/internal/ent"
	"go-ent-project/utils/auth"
	"net/http"
)

type contextKey string

const userCtxKey = contextKey("user")
const TokenContextKey = contextKey("token")

// AuthMiddleware validates the token and attaches the user to the context.
func AuthMiddleware(client *ent.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header missing", http.StatusUnauthorized)
				return
			}

			token, _ := auth.ExtractTokenFromHeader(authHeader)
			claims, err := auth.ValidateTokenWithDB(r.Context(), token, client)
			if err != nil {
				http.Error(w, fmt.Sprintf("Invalid token: %v", err), http.StatusUnauthorized)
				return
			}

			// Fetch user details from the database
			authUser, err := client.User.Get(r.Context(), claims.UserID)
			if err != nil {
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			}

			// Attach user to the context
			ctx := context.WithValue(r.Context(), userCtxKey, authUser)
			ctx = context.WithValue(r.Context(), TokenContextKey, token)
			// Proceed with the next handler
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserFromContext(ctx context.Context) (*ent.User, error) {
	user, ok := ctx.Value(contextKey("user")).(*ent.User)
	if !ok || user == nil {
		return nil, fmt.Errorf("user not found in context")
	}
	return user, nil
}

// ExtractTokenFromContext retrieves the token from the request context.
func ExtractTokenFromContext(ctx context.Context) (string, error) {
	token, ok := ctx.Value(TokenContextKey).(string)
	if !ok || token == "" {
		return "", fmt.Errorf("token not found in context")
	}
	return token, nil
}
