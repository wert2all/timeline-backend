package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Timeline holds the schema definition for the Timeline entity.
type Timeline struct {
	ent.Schema
}

// Fields of the Timeline.
func (Timeline) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Timeline.
func (Timeline) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).Ref("timeline").Unique(),
		edge.To("event", Event.Type),
	}
}
