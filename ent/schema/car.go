package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		// Make of the car
		field.String("make").
			NotEmpty().
			Annotations(
				// Expose in GraphQL
			),

		// Model of the car
		field.String("model").
			NotEmpty().
			Annotations(),

		// Year of manufacture
		field.Int("year").
			Positive().
			Annotations(),

		// Registration number of the car
		field.String("registration").
			Unique().
			NotEmpty().
			Annotations(),

		// Color of the car
		field.String("color").
			NotEmpty().
			Annotations(),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		// Owner edge: A car belongs to a single user

	}
}
