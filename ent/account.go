// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"timeline/backend/ent/account"
	"timeline/backend/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// PreviewlyToken holds the value of the "previewly_token" field.
	PreviewlyToken string `json:"previewly_token,omitempty"`
	// AvatarID holds the value of the "avatar_id" field.
	AvatarID *int `json:"avatar_id,omitempty"`
	// About holds the value of the "about" field.
	About *string `json:"about,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges        AccountEdges `json:"edges"`
	user_account *int
	selectValues sql.SelectValues
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// Timeline holds the value of the timeline edge.
	Timeline []*Timeline `json:"timeline,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TimelineOrErr returns the Timeline value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) TimelineOrErr() ([]*Timeline, error) {
	if e.loadedTypes[0] {
		return e.Timeline, nil
	}
	return nil, &NotLoadedError{edge: "timeline"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldID, account.FieldAvatarID:
			values[i] = new(sql.NullInt64)
		case account.FieldName, account.FieldPreviewlyToken, account.FieldAbout:
			values[i] = new(sql.NullString)
		case account.ForeignKeys[0]: // user_account
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case account.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				a.Name = value.String
			}
		case account.FieldPreviewlyToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field previewly_token", values[i])
			} else if value.Valid {
				a.PreviewlyToken = value.String
			}
		case account.FieldAvatarID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_id", values[i])
			} else if value.Valid {
				a.AvatarID = new(int)
				*a.AvatarID = int(value.Int64)
			}
		case account.FieldAbout:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field about", values[i])
			} else if value.Valid {
				a.About = new(string)
				*a.About = value.String
			}
		case account.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_account", value)
			} else if value.Valid {
				a.user_account = new(int)
				*a.user_account = int(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Account.
// This includes values selected through modifiers, order, etc.
func (a *Account) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryTimeline queries the "timeline" edge of the Account entity.
func (a *Account) QueryTimeline() *TimelineQuery {
	return NewAccountClient(a.config).QueryTimeline(a)
}

// QueryUser queries the "user" edge of the Account entity.
func (a *Account) QueryUser() *UserQuery {
	return NewAccountClient(a.config).QueryUser(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return NewAccountClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("name=")
	builder.WriteString(a.Name)
	builder.WriteString(", ")
	builder.WriteString("previewly_token=")
	builder.WriteString(a.PreviewlyToken)
	builder.WriteString(", ")
	if v := a.AvatarID; v != nil {
		builder.WriteString("avatar_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := a.About; v != nil {
		builder.WriteString("about=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account
