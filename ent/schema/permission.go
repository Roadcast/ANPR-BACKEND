package schema

import (
	_ "entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		// Name of the permission
		field.String("name").
			NotEmpty().
			Annotations(
			// Expose in GraphQL
			),

		// Read permission
		field.Bool("can_read").
			Default(false). // Defaults to false
			Annotations(
			// Expose in GraphQL
			),

		// Create permission
		field.Bool("can_create").
			Default(false).
			Annotations(),

		// Update permission
		field.Bool("can_update").
			Default(false).
			Annotations(),

		// Delete permission
		field.Bool("can_delete").
			Default(false).
			Annotations(),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return nil
}
