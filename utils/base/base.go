package mixin

import (
	"entgo.io/ent/schema"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// BaseMixin defines fields and behaviors common to all schemas.
type BaseMixin struct {
	ent.Mixin
}

// Fields adds `id`, `created_at`, and `updated_at` fields to the schema.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Unique().
			Immutable().
			Comment("Primary key"),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Comment("Timestamp when the entity was created"),

		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("Timestamp when the entity was last updated"),
	}
}

func (BaseMixin) Annotations() []schema.Annotation {
	fmt.Println("BaseMixin Annotations() called")
	return nil
}

func (BaseMixin) Edges() []ent.Edge {
	return nil
}

// Indexes returns no indexes by default.
func (BaseMixin) Indexes() []ent.Index {
	return nil
}

// Hooks explicitly returns nil, indicating no hooks are applied by the mixin.
func (BaseMixin) Hooks() []ent.Hook {
	return nil
}

// Interceptors explicitly returns nil, indicating no interceptors are applied by the mixin.
func (BaseMixin) Interceptors() []ent.Interceptor {
	return nil
}

// Policy explicitly returns nil, indicating no policies are applied by the mixin.
func (BaseMixin) Policy() ent.Policy {
	return nil
}
