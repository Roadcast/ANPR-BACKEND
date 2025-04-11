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
	hook "go-ent-project/ent/hooks"
	"go-ent-project/ent/privacy"
	"go-ent-project/utils/base"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		base.BMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"), // Adds ordering
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.String("email").
			Unique().
			NotEmpty().
			Annotations(
				entgql.OrderField("EMAIL"),
			),
		field.String("password").
			Sensitive().
			NotEmpty().
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.String("phone").
			Optional().
			Annotations(
			// Expose this field in GraphQL
			),
		field.Bool("active").Default(true).Annotations(),
		field.UUID("role_id", uuid.UUID{}).Annotations(entgql.Type("ID"), entgql.Skip(entgql.SkipWhereInput)),
		field.UUID("police_station_id", uuid.UUID{}).Optional().Annotations(entgql.Type("ID"), entgql.Skip(entgql.SkipWhereInput)),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).
			Ref("users").Field("role_id").Unique().
			Required().
			Annotations(),
		edge.From("police_station", PoliceStation.Type).
			Ref("users").Unique().Field("police_station_id").
			Unique().
			Annotations(),
	}
}

// Indexes defines the indexes of the User entity.
func (User) Indexes() []ent.Index {
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

func (User) Hooks() []ent.Hook {
	return hook.MakeHooks()
}

func (User) Policy() ent.Policy {
	return privacy.UserPolicy
}
