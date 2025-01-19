package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"go-ent-project/utils/base"
)

// PoliceStation holds the schema definition for the PoliceStation entity.
type PoliceStation struct {
	ent.Schema
}

// Mixin adds the base mixin to the schema.
func (PoliceStation) Mixin() []ent.Mixin {
	return []ent.Mixin{
		base.BMixin{},
	}
}

// Fields of the PoliceStation.
func (PoliceStation) Fields() []ent.Field {
	return []ent.Field{
		// Name of the police station
		field.String("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"), // Adds ordering support
				// Expose in GraphQL
			),

		// Location stored as geometry(Point) in PostgreSQL
		field.JSON("location", map[string]interface{}{}).
			Optional().
			Annotations().
			SchemaType(map[string]string{
				"postgres": "geometry(Point)",
			}),

		// Code for the police station
		field.String("code").
			Unique().
			NotEmpty().
			Annotations(),

		// Unique identifier for the police station
		field.String("identifier").
			Unique().
			NotEmpty().
			Annotations(),
	}
}

// Edges of the PoliceStation.
func (PoliceStation) Edges() []ent.Edge {
	return []ent.Edge{
		// Users assigned to this police station
		edge.To("users", User.Type),

		// Parent police station
		edge.From("parent_station", PoliceStation.Type).
			Ref("child_stations").
			Unique(),

		// Child police stations
		edge.To("child_stations", PoliceStation.Type),
	}
}
