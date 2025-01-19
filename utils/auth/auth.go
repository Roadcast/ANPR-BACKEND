package auth

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	accessTokenSecret  = []byte("your-access-token-secret")
	refreshTokenSecret = []byte("your-refresh-token-secret")
	tokenExpiry        = time.Minute * 15   // Access token expiration
	refreshExpiry      = time.Hour * 24 * 7 // Refresh token expiration
)

// Claims defines custom claims structure.
type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	Hash   string `json:"hash"`
	jwt.RegisteredClaims
}

// GenerateMD5 generates an MD5 hash for a given string.
func GenerateMD5(value string) string {
	hash := md5.Sum([]byte(value))
	return hex.EncodeToString(hash[:])
}

// GenerateToken generates an access token and refresh token with a combined MD5 hash.
func GenerateToken(userID int, role, password string, active bool) (string, string, error) {
	now := time.Now()

	// Generate MD5 hash of concatenated fields
	hashInput := fmt.Sprintf("%s|%s|%t", role, password, active)
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
