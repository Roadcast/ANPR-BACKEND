package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	hook "go-ent-project/ent/hooks"
	"go-ent-project/utils/base"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Mixin adds the base mixin to the schema.
func (Car) Mixin() []ent.Mixin {
	return []ent.Mixin{
		base.BMixin{},
	}
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		// Make of the car
		field.String("make").
			Optional().
			Annotations(
			// Expose in GraphQL
			),

		// Model of the car
		field.String("model").
			Optional().
			Annotations(),

		// Year of manufacture
		field.Int("year").
			Optional().
			Annotations(),

		// Registration number of the car
		field.String("registration").
			Unique().
			NotEmpty().
			Annotations(),

		// Color of the car
		field.String("color").
			Optional().
			Annotations(),

		field.UUID("police_station_id", uuid.UUID{}).Nillable().
			Optional().Annotations(entgql.Type("ID"), entgql.Skip(entgql.SkipWhereInput)),

		field.Time("stolen_date").Optional().Annotations(),
		field.String("fir_number").Optional().Annotations(),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("police_station", PoliceStation.Type).
			Ref("car").Field("police_station_id").Unique().
			Annotations(),
	}
}

func (Car) Hooks() []ent.Hook {
	return hook.CarHooks()
}
