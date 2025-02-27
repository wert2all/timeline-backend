// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"timeline/backend/ent/predicate"
	"timeline/backend/ent/settings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SettingsDelete is the builder for deleting a Settings entity.
type SettingsDelete struct {
	config
	hooks    []Hook
	mutation *SettingsMutation
}

// Where appends a list predicates to the SettingsDelete builder.
func (sd *SettingsDelete) Where(ps ...predicate.Settings) *SettingsDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *SettingsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *SettingsDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *SettingsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(settings.Table, sqlgraph.NewFieldSpec(settings.FieldID, field.TypeInt))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// SettingsDeleteOne is the builder for deleting a single Settings entity.
type SettingsDeleteOne struct {
	sd *SettingsDelete
}

// Where appends a list predicates to the SettingsDelete builder.
func (sdo *SettingsDeleteOne) Where(ps ...predicate.Settings) *SettingsDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *SettingsDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{settings.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *SettingsDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}
