// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "date", Type: field.TypeTime},
		{Name: "time", Type: field.TypeString},
		{Name: "show_time", Type: field.TypeBool, Default: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
	}
	// TimelinesColumns holds the columns for the "timelines" table.
	TimelinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// TimelinesTable holds the schema information for the "timelines" table.
	TimelinesTable = &schema.Table{
		Name:       "timelines",
		Columns:    TimelinesColumns,
		PrimaryKey: []*schema.Column{TimelinesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "admin", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EventsTable,
		TimelinesTable,
		UsersTable,
	}
)

func init() {
}
