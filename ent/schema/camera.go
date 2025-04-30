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

// Camera holds the schema definition for the Camera entity.
type Camera struct {
	ent.Schema
}

// Mixin adds the base mixin to the schema.
func (Camera) Mixin() []ent.Mixin {
	return []ent.Mixin{
		base.BMixin{},
	}
}

// Fields of the Camera.
func (Camera) Fields() []ent.Field {
	return []ent.Field{
		// Name of the camera
		field.String("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
		// Model of the camera
		field.String("model").
			NotEmpty().
			Annotations(),

		// IMEI number of the camera
		field.String("imei").
			Unique().
			NotEmpty().
			Annotations(),

		field.String("location"). // Store as WKB (Bytes)
						SchemaType(map[string]string{
				"postgres": "GEOMETRY(Point, 4326)",
			}).
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),

		field.Bool("active").
			Default(true).
			Annotations(),

		field.String("address").Optional(),

		field.Bool("is_working").Default(true).Annotations(),
		field.String("district").Default("N/A").Nillable().Annotations(),
		field.Time("last_ping"),
		field.UUID("police_station_id", uuid.UUID{}).Nillable().
			Optional().Annotations(entgql.Type("ID"), entgql.Skip(entgql.SkipWhereInput)),
	}
}

// Indexes defines the indexes of the User entity.
func (Camera) Indexes() []ent.Index {
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

func (Camera) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("police_station", PoliceStation.Type).
			Ref("camera").Field("police_station_id").Unique().
			Annotations(entsql.OnDelete(entsql.SetNull)),
	}
}
