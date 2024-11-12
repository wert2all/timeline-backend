// Code generated by ent, DO NOT EDIT.

package settings

import (
	"fmt"
	enumvalues "timeline/backend/lib/enum-values"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the settings type in the database.
	Label = "settings"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldEntityID holds the string denoting the entity_id field in the database.
	FieldEntityID = "entity_id"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// Table holds the table name of the settings in the database.
	Table = "settings"
)

// Columns holds all SQL columns for settings fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldEntityID,
	FieldKey,
	FieldValue,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// EntityIDValidator is a validator for the "entity_id" field. It is called by the builders before save.
	EntityIDValidator func(int) error
	// KeyValidator is a validator for the "key" field. It is called by the builders before save.
	KeyValidator func(string) error
	// DefaultValue holds the default value on creation for the "value" field.
	DefaultValue string
)

const DefaultType enumvalues.SettingsType = "ACCOUNT"

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type enumvalues.SettingsType) error {
	switch _type {
	case "ACCOUNT":
		return nil
	default:
		return fmt.Errorf("settings: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the Settings queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByEntityID orders the results by the entity_id field.
func ByEntityID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntityID, opts...).ToFunc()
}

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}