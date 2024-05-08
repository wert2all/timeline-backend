package schema

import (
	"entgo.io/ent/schema/edge"
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
		field.String("time").Optional(),
		field.Bool("showTime").Default(true),
		field.String("title").Optional(),
		field.String("description").Optional(),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("timeline", Timeline.Type).Ref("event").Unique(),
	}
}
