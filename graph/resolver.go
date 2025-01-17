package graph

import (
	"encoding/base64"
	"fmt"
	"go-ent-project/internal/ent"
	"strconv"
)

// Resolver struct for dependency injection
type Resolver struct {
	Client *ent.Client
}

func DecodeCursor(cursor string) (int, error) {
	decoded, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return 0, fmt.Errorf("failed to decode cursor: %v", err)
	}

	value, err := strconv.Atoi(string(decoded))
	if err != nil {
		return 0, fmt.Errorf("invalid cursor value: %v", err)
	}

	return value, nil
}
