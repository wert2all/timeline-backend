package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("date"),
		field.Enum("type").Values("default", "selebrate").Default("default"),
		field.String("time").Nillable(),
		field.Bool("showTime").Default(true),
		field.String("title").Nillable(),
		field.String("description").Nillable(),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return nil
}
