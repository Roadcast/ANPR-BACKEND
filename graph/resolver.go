package graph

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkb"
	"go-ent-project/internal/ent"
)

type Resolver struct {
	Client *ent.Client
	SQL    *sql.DB
}

// Convert WKB (Bytes) to WKT (String)
func decodeWKBToWKT(wkbBytes []byte) (string, error) {
	if len(wkbBytes) == 0 {
		return "", fmt.Errorf("empty WKB data")
	}

	// Detect if data is in HEX format (ASCII-encoded)
	if isHexEncoded(wkbBytes) {
		decodedBytes, err := hex.DecodeString(string(wkbBytes))
		if err != nil {
			return "", fmt.Errorf("failed to decode HEX to Binary: %v", err)
		}
		wkbBytes = decodedBytes
	}

	// Decode WKB to Geometry
	point, err := wkb.Unmarshal(wkbBytes)
	if err != nil {
		return "", fmt.Errorf("failed to decode WKB: %v", err)
	}

	// Ensure it's a Point Geometry
	gp, ok := point.(*geom.Point)
	if !ok {
		return "", fmt.Errorf("invalid geometry type: %T", point)
	}

	// Convert to WKT format
	return fmt.Sprintf("POINT(%f %f)", gp.X(), gp.Y()), nil
}

// Helper function: Detects if data is HEX encoded
func isHexEncoded(data []byte) bool {
	for _, b := range data {
		if !(b >= 48 && b <= 57) && !(b >= 65 && b <= 70) && !(b >= 97 && b <= 102) { // 0-9, A-F, a-f
			return false
		}
	}
	return len(data)%2 == 0 // Must be an even number of hex chars
}
