package graph

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"encoding/base64"
	"entgo.io/contrib/entgql"
	"fmt"
	"go-ent-project/internal/ent"
	"strconv"
)

type Resolver struct {
	Client *ent.Client
}

func DecodeCursor(cursor string) (int, error) {
	// Decode Base64 string
	decoded, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return 0, fmt.Errorf("failed to decode cursor: %v", err)
	}

	// Convert the decoded string to an integer
	value, err := strconv.Atoi(string(decoded))

	if err != nil {
		return 0, fmt.Errorf("invalid cursor value: %v", err)
	}

	return value, nil
}
func convertCursorToString[T any](cursor *entgql.Cursor[T]) *string {
	str := fmt.Sprintf("%v", cursor.ID)
	// Encode the value as Base64
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	return &encoded
}
