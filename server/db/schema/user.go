package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.String("name").NotEmpty().Unique(),
	}
}

// Edges of the Node.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bestFriend", User.Type).
			Unique().
			Required().
			From("friends"),
		edge.From("adminGroups", UserGroup.Type).
			Annotations(
				entgql.MapsTo("adminGroups"),
				entgql.RelayConnection(),
			).
			Ref("admin"),
	}
}

// Annotations of the User
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
	}
}
