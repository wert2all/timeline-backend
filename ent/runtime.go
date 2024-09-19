// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
	"timeline/backend/ent/event"
	"timeline/backend/ent/schema"
	"timeline/backend/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	eventFields := schema.Event{}.Fields()
	_ = eventFields
	// eventDescCreatedAt is the schema descriptor for created_at field.
	eventDescCreatedAt := eventFields[0].Descriptor()
	// event.DefaultCreatedAt holds the default value on creation for the created_at field.
	event.DefaultCreatedAt = eventDescCreatedAt.Default.(func() time.Time)
	// eventDescShowTime is the schema descriptor for showTime field.
	eventDescShowTime := eventFields[4].Descriptor()
	// event.DefaultShowTime holds the default value on creation for the showTime field.
	event.DefaultShowTime = eventDescShowTime.Default.(bool)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUUID is the schema descriptor for uuid field.
	userDescUUID := userFields[0].Descriptor()
	// user.UUIDValidator is a validator for the "uuid" field. It is called by the builders before save.
	user.UUIDValidator = userDescUUID.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescActive is the schema descriptor for active field.
	userDescActive := userFields[6].Descriptor()
	// user.DefaultActive holds the default value on creation for the active field.
	user.DefaultActive = userDescActive.Default.(bool)
	// userDescAdmin is the schema descriptor for admin field.
	userDescAdmin := userFields[7].Descriptor()
	// user.DefaultAdmin holds the default value on creation for the admin field.
	user.DefaultAdmin = userDescAdmin.Default.(bool)
}
