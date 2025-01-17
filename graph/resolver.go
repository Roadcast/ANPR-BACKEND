package graph

import (
	"encoding/base64"
	"entgo.io/contrib/entgql"
	"fmt"
	"go-ent-project/internal/ent"
	"strconv"
)

// Resolver struct for dependency injection
type Resolver struct {
	Client *ent.Client
}

func DecodeCursor(cursor string) (int, error) {
	// Decode Base64 string
	decoded, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return 0, fmt.Errorf("failed to decode cursor: %v", err)
	}

	// Print the decoded value as a string
	fmt.Printf("Decoded Cursor: %s\n", string(decoded))

	// Convert the decoded string to an integer
	value, err := strconv.Atoi(string(decoded))

	if err != nil {
		return 0, fmt.Errorf("invalid cursor value: %v", err)
	}

	return value, nil
}

func convertCursorToString[T any](cursor *entgql.Cursor[T]) *string {
	println(cursor.ID, cursor.Value)
	if cursor == nil {
		return nil
	}
	str := fmt.Sprintf("%v", cursor.ID)
	// Encode the value as Base64
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	return &encoded
}
