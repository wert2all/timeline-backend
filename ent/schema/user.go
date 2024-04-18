package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").NotEmpty().Unique(),
		field.String("name").NotEmpty(),
		field.String("email").NotEmpty(),
		field.String("avatar").Nillable(),
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Bool("active").Default(true),
		field.Bool("admin").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
