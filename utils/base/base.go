package base

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

type BMixin struct {
	mixin.Schema
}

func (BMixin) Fields() []ent.Field {
	return []ent.Field{
		// Primary key
		field.Int("id").
			Unique().
			Immutable().
			Annotations(entgql.OrderField("ID")),

		// CreatedAt timestamp
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(),

		// UpdatedAt timestamp
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(),
	}
}

//func (BMixin) Annotations() []schema.Annotation {
//	return []schema.Annotation{
//
//		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
//	}
//}
