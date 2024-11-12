package schema

import (
	enumvalues "timeline/backend/lib/enum-values"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Settings holds the schema definition for the Settings entity.
type Settings struct {
	ent.Schema
}

// Fields of the Settings.
func (Settings) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").GoType(enumvalues.SettingsType("")).Default(string(enumvalues.Account)),
		field.Int("entity_id").Positive(),
		field.String("key").NotEmpty(),
		field.String("value").Default(""),
	}
}

func (Settings) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type", "entity_id", "key").Unique(),
	}
}

// Edges of the Settings.
func (Settings) Edges() []ent.Edge {
	return nil
}
