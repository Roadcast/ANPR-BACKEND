package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
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
		field.Time("created_at").
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"), // Adds ordering support
			),
		field.Time("updated_at").
			Annotations(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).
			Ref("users").
			Unique().
			Annotations(
				entgql.Bind(), // Binds this edge to GraphQL
			),
	}
}
