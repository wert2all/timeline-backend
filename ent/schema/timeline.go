package schema

import (
	"entgo.io/ent"
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
	return nil
}
