// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"timeline/backend/ent/settings"
	enumvalues "timeline/backend/lib/enum-values"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Settings is the model entity for the Settings schema.
type Settings struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type enumvalues.SettingsType `json:"type,omitempty"`
	// EntityID holds the value of the "entity_id" field.
	EntityID int `json:"entity_id,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// Value holds the value of the "value" field.
	Value        string `json:"value,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Settings) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case settings.FieldID, settings.FieldEntityID:
			values[i] = new(sql.NullInt64)
		case settings.FieldType, settings.FieldKey, settings.FieldValue:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Settings fields.
func (s *Settings) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case settings.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case settings.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				s.Type = enumvalues.SettingsType(value.String)
			}
		case settings.FieldEntityID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field entity_id", values[i])
			} else if value.Valid {
				s.EntityID = int(value.Int64)
			}
		case settings.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				s.Key = value.String
			}
		case settings.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				s.Value = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Settings.
// This includes values selected through modifiers, order, etc.
func (s *Settings) GetValue(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Settings.
// Note that you need to call Settings.Unwrap() before calling this method if this Settings
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Settings) Update() *SettingsUpdateOne {
	return NewSettingsClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Settings entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Settings) Unwrap() *Settings {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Settings is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Settings) String() string {
	var builder strings.Builder
	builder.WriteString("Settings(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", s.Type))
	builder.WriteString(", ")
	builder.WriteString("entity_id=")
	builder.WriteString(fmt.Sprintf("%v", s.EntityID))
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(s.Key)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(s.Value)
	builder.WriteByte(')')
	return builder.String()
}

// SettingsSlice is a parsable slice of Settings.
type SettingsSlice []*Settings
