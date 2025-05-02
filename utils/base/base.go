package base

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
			SchemaType(map[string]string{
				"postgres": "uuid",
			}).
			Annotations(entgql.Type("ID"),
				entgql.Skip(entgql.SkipWhereInput),
				entsql.Default("uuid_generate_v4()")),

		// CreatedAt timestamp
		field.Time("created_at").
			Default(time.Now).
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone", // PostgreSQL type
			}).
			Immutable().
			Annotations(entgql.OrderField("CreatedAt"),
				entsql.Annotation{
					Default: "CURRENT_TIMESTAMP", // this becomes DEFAULT CURRENT_TIMESTAMP
				}),

		// UpdatedAt timestamp
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				"postgres": "timestamp with time zone", // PostgreSQL type
			}).
			Annotations(entgql.OrderField("UpdatedAt"),
				entgql.Skip(entgql.SkipWhereInput),
				entsql.Annotation{
					Default: "CURRENT_TIMESTAMP", // this becomes DEFAULT CURRENT_TIMESTAMP
				}),
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

func (BMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
	}
}
