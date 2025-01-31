package schema

import (
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkb"
)

// Point is a custom struct for storing lat/lon
type Point struct {
	geom.Point
}

// Implement database serialization for Point
func (p *Point) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("invalid type for Point: %T", value)
	}

	// Decode WKB
	point, err := wkb.Unmarshal(bytes)
	if err != nil {
		return fmt.Errorf("failed to decode WKB: %v", err)
	}

	// Convert to go-geom format
	gp, ok := point.(*geom.Point)
	if !ok {
		return fmt.Errorf("invalid geometry type: %T", point)
	}

	p.Point = *gp
	return nil
}

// Convert Point to WKB for storage
func (p Point) Value() (driver.Value, error) {
	// Ensure it's a valid point
	if p.Point.Empty() {
		return nil, nil // Store as NULL if empty
	}

	// Marshal to WKB format
	wkbData, err := wkb.Marshal(&p.Point, binary.BigEndian)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal WKB: %v", err)
	}

	return wkbData, nil
}
