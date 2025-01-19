package auth

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"go-ent-project/internal/ent"
	redisDB "go-ent-project/utils/redis"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

var (
	accessTokenSecret  = []byte("your-access-token-secret")
	refreshTokenSecret = []byte("your-refresh-token-secret")
	tokenExpiry        = time.Hour * 2      // Access token expiration
	refreshExpiry      = time.Hour * 24 * 7 // Refresh token expiration
)

// Claims defines custom claims structure.
type Claims struct {
	UserID int    `json:"user_id"`
	Role   int    `json:"role"`
	Hash   string `json:"hash"`
	jwt.RegisteredClaims
}

type contextKey string

// GenerateMD5 generates an MD5 hash for a given string.
func GenerateMD5(value string) string {
	hash := md5.Sum([]byte(value))
	return hex.EncodeToString(hash[:])
}

// GenerateToken generates an access token and refresh token with a combined MD5 hash.
func GenerateToken(userID int, role int, password string, active bool) (string, string, error) {
	now := time.Now()

	// Generate MD5 hash of concatenated fields
	hashInput := fmt.Sprintf("%d|%s|%t", role, password, active)
	hash := GenerateMD5(hashInput)

	// Access token claims
	accessClaims := Claims{
		UserID: userID,
		Role:   role,
		Hash:   hash,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(tokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(accessTokenSecret)
	if err != nil {
		return "", "", err
	}

	// Refresh token claims
	refreshClaims := Claims{
		UserID: userID,
		Role:   role,
		Hash:   hash,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(refreshExpiry)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(refreshTokenSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func ValidateTokenWithDB(ctx context.Context, tokenString string, db *ent.Client) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return accessTokenSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Fetch user data from the database
	user, err := db.User.Get(ctx, claims.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	// Compute hash of current database values
	hashInput := fmt.Sprintf("%d|%s|%t", user.RoleID, user.Password, user.Active)
	expectedHash := GenerateMD5(hashInput)

	// Compare hashes
	if claims.Hash != expectedHash {
		return nil, errors.New("token hash mismatch, data may have been tampered with")
	}

	return claims, nil
}

func ValidatePassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

// AddToBlacklist adds a token to the blacklist.
func AddToBlacklist(token string, expiry time.Duration) error {
	return redisDB.Client.SetEX(redisDB.Ctx, token, "blacklisted", expiry).Err()
}

// IsTokenBlacklisted checks if a token is in the blacklist.
func IsTokenBlacklisted(token string) (bool, error) {
	result, err := redisDB.Client.Get(redisDB.Ctx, token).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil // Token is not blacklisted
	} else if err != nil {
		return false, err
	}
	return result == "blacklisted", nil
}

// ExtractTokenFromHeader extracts the token from the "Authorization" header.
func ExtractTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", errors.New("authorization header is empty")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("authorization header format must be 'Bearer <token>'")
	}

	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", errors.New("token is empty in authorization header")
	}

	return token, nil
}
