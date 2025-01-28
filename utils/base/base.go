package base

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"time"
)

type BMixin struct {
	mixin.Schema
}

func (BMixin) Fields() []ent.Field {
	return []ent.Field{
		// Primary key
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique().
			Immutable().
			Annotations(entgql.Type("ID")),

		// CreatedAt timestamp
		field.Time("created_at").
			Default(time.Now).
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone", // PostgreSQL type
			}).
			Immutable().
			Annotations(entgql.OrderField("CreatedAt")),

		// UpdatedAt timestamp
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone", // PostgreSQL type
			}).
			Annotations(entgql.OrderField("UpdatedAt")),
	}
}

func (BMixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entgql.MultiOrder(),
	}
}
