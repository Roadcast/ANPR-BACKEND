package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
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

		field.String("location").Optional().Nillable(). // Store as WKB (Bytes)
								SchemaType(map[string]string{
				"postgres": "GEOMETRY(Point, 4326)",
			}).
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),

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
		field.UUID("parent_station_id", uuid.UUID{}).Optional().Nillable().Annotations(entgql.Type("ID")),
	}
}

// Edges of the PoliceStation.
func (PoliceStation) Edges() []ent.Edge {
	return []ent.Edge{
		// Users assigned to this police station
		edge.To("users", User.Type),
		edge.To("camera", Camera.Type),
		// Parent station (one-to-many relationship)
		edge.From("parent", PoliceStation.Type).
			Ref("child_stations").
			Unique().
			Field("parent_station_id").
			Annotations(entgql.Type("PoliceStation")),

		// Child stations (one-to-many relationship)
		edge.To("child_stations", PoliceStation.Type).
			Annotations(entgql.Type("[PoliceStation!]!")),
	}
}

// Indexes defines the indexes of the User entity.
func (PoliceStation) Indexes() []ent.Index {
	return []ent.Index{
		// Create a GIN index using the `pg_trgm` operator class
		index.Fields("name").
			Annotations(
				entsql.IndexTypes(map[string]string{
					dialect.MySQL:    "FULLTEXT",
					dialect.Postgres: "GIN",
				})),
	}
}
