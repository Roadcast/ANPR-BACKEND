package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
				entgql.OrderField("NAME"), // Adds ordering support
				// Expose this field in GraphQL
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
			// Exposed for creating or updating a user
			),
		field.String("phone").
			Optional().
			Annotations(
			// Expose this field in GraphQL
			),
		field.Bool("active").Default(true).Annotations(),
		field.Int("role_id").Annotations(),
		field.Int("police_station_id").Annotations(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).
			Ref("users").
			Required().
			Annotations(),
		edge.From("police_station", PoliceStation.Type).
			Ref("users").
			Required().
			Annotations(),
	}
}

func (User) Hooks() []ent.Hook {
	return hook.MakeHooks()
}

func (User) Policy() ent.Policy {
	return privacy.UserPolicy
}
