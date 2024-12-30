// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"timeline/backend/ent/event"
	"timeline/backend/ent/tag"
	"timeline/backend/ent/timeline"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EventCreate is the builder for creating a Event entity.
type EventCreate struct {
	config
	mutation *EventMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ec *EventCreate) SetCreatedAt(t time.Time) *EventCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EventCreate) SetNillableCreatedAt(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// SetDate sets the "date" field.
func (ec *EventCreate) SetDate(t time.Time) *EventCreate {
	ec.mutation.SetDate(t)
	return ec
}

// SetType sets the "type" field.
func (ec *EventCreate) SetType(e event.Type) *EventCreate {
	ec.mutation.SetType(e)
	return ec
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ec *EventCreate) SetNillableType(e *event.Type) *EventCreate {
	if e != nil {
		ec.SetType(*e)
	}
	return ec
}

// SetTime sets the "time" field.
func (ec *EventCreate) SetTime(s string) *EventCreate {
	ec.mutation.SetTime(s)
	return ec
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (ec *EventCreate) SetNillableTime(s *string) *EventCreate {
	if s != nil {
		ec.SetTime(*s)
	}
	return ec
}

// SetShowTime sets the "showTime" field.
func (ec *EventCreate) SetShowTime(b bool) *EventCreate {
	ec.mutation.SetShowTime(b)
	return ec
}

// SetNillableShowTime sets the "showTime" field if the given value is not nil.
func (ec *EventCreate) SetNillableShowTime(b *bool) *EventCreate {
	if b != nil {
		ec.SetShowTime(*b)
	}
	return ec
}

// SetTitle sets the "title" field.
func (ec *EventCreate) SetTitle(s string) *EventCreate {
	ec.mutation.SetTitle(s)
	return ec
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ec *EventCreate) SetNillableTitle(s *string) *EventCreate {
	if s != nil {
		ec.SetTitle(*s)
	}
	return ec
}

// SetDescription sets the "description" field.
func (ec *EventCreate) SetDescription(s string) *EventCreate {
	ec.mutation.SetDescription(s)
	return ec
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ec *EventCreate) SetNillableDescription(s *string) *EventCreate {
	if s != nil {
		ec.SetDescription(*s)
	}
	return ec
}

// SetURL sets the "url" field.
func (ec *EventCreate) SetURL(s string) *EventCreate {
	ec.mutation.SetURL(s)
	return ec
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (ec *EventCreate) SetNillableURL(s *string) *EventCreate {
	if s != nil {
		ec.SetURL(*s)
	}
	return ec
}

// SetPreviewlyImageID sets the "previewly_image_id" field.
func (ec *EventCreate) SetPreviewlyImageID(i int) *EventCreate {
	ec.mutation.SetPreviewlyImageID(i)
	return ec
}

// SetNillablePreviewlyImageID sets the "previewly_image_id" field if the given value is not nil.
func (ec *EventCreate) SetNillablePreviewlyImageID(i *int) *EventCreate {
	if i != nil {
		ec.SetPreviewlyImageID(*i)
	}
	return ec
}

// SetTimelineID sets the "timeline" edge to the Timeline entity by ID.
func (ec *EventCreate) SetTimelineID(id int) *EventCreate {
	ec.mutation.SetTimelineID(id)
	return ec
}

// SetNillableTimelineID sets the "timeline" edge to the Timeline entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillableTimelineID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetTimelineID(*id)
	}
	return ec
}

// SetTimeline sets the "timeline" edge to the Timeline entity.
func (ec *EventCreate) SetTimeline(t *Timeline) *EventCreate {
	return ec.SetTimelineID(t.ID)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (ec *EventCreate) AddTagIDs(ids ...int) *EventCreate {
	ec.mutation.AddTagIDs(ids...)
	return ec
}

// AddTags adds the "tags" edges to the Tag entity.
func (ec *EventCreate) AddTags(t ...*Tag) *EventCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ec.AddTagIDs(ids...)
}

// Mutation returns the EventMutation object of the builder.
func (ec *EventCreate) Mutation() *EventMutation {
	return ec.mutation
}

// Save creates the Event in the database.
func (ec *EventCreate) Save(ctx context.Context) (*Event, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EventCreate) SaveX(ctx context.Context) *Event {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EventCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EventCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EventCreate) defaults() {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		v := event.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.GetType(); !ok {
		v := event.DefaultType
		ec.mutation.SetType(v)
	}
	if _, ok := ec.mutation.ShowTime(); !ok {
		v := event.DefaultShowTime
		ec.mutation.SetShowTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EventCreate) check() error {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Event.created_at"`)}
	}
	if _, ok := ec.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "Event.date"`)}
	}
	if _, ok := ec.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Event.type"`)}
	}
	if v, ok := ec.mutation.GetType(); ok {
		if err := event.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Event.type": %w`, err)}
		}
	}
	if _, ok := ec.mutation.ShowTime(); !ok {
		return &ValidationError{Name: "showTime", err: errors.New(`ent: missing required field "Event.showTime"`)}
	}
	return nil
}

func (ec *EventCreate) sqlSave(ctx context.Context) (*Event, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EventCreate) createSpec() (*Event, *sqlgraph.CreateSpec) {
	var (
		_node = &Event{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(event.Table, sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ec.conflict
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(event.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.Date(); ok {
		_spec.SetField(event.FieldDate, field.TypeTime, value)
		_node.Date = value
	}
	if value, ok := ec.mutation.GetType(); ok {
		_spec.SetField(event.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := ec.mutation.Time(); ok {
		_spec.SetField(event.FieldTime, field.TypeString, value)
		_node.Time = value
	}
	if value, ok := ec.mutation.ShowTime(); ok {
		_spec.SetField(event.FieldShowTime, field.TypeBool, value)
		_node.ShowTime = value
	}
	if value, ok := ec.mutation.Title(); ok {
		_spec.SetField(event.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := ec.mutation.Description(); ok {
		_spec.SetField(event.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ec.mutation.URL(); ok {
		_spec.SetField(event.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := ec.mutation.PreviewlyImageID(); ok {
		_spec.SetField(event.FieldPreviewlyImageID, field.TypeInt, value)
		_node.PreviewlyImageID = &value
	}
	if nodes := ec.mutation.TimelineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.TimelineTable,
			Columns: []string{event.TimelineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timeline.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.timeline_event = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   event.TagsTable,
			Columns: event.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Event.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EventUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ec *EventCreate) OnConflict(opts ...sql.ConflictOption) *EventUpsertOne {
	ec.conflict = opts
	return &EventUpsertOne{
		create: ec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ec *EventCreate) OnConflictColumns(columns ...string) *EventUpsertOne {
	ec.conflict = append(ec.conflict, sql.ConflictColumns(columns...))
	return &EventUpsertOne{
		create: ec,
	}
}

type (
	// EventUpsertOne is the builder for "upsert"-ing
	//  one Event node.
	EventUpsertOne struct {
		create *EventCreate
	}

	// EventUpsert is the "OnConflict" setter.
	EventUpsert struct {
		*sql.UpdateSet
	}
)

// SetDate sets the "date" field.
func (u *EventUpsert) SetDate(v time.Time) *EventUpsert {
	u.Set(event.FieldDate, v)
	return u
}

// UpdateDate sets the "date" field to the value that was provided on create.
func (u *EventUpsert) UpdateDate() *EventUpsert {
	u.SetExcluded(event.FieldDate)
	return u
}

// SetType sets the "type" field.
func (u *EventUpsert) SetType(v event.Type) *EventUpsert {
	u.Set(event.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *EventUpsert) UpdateType() *EventUpsert {
	u.SetExcluded(event.FieldType)
	return u
}

// SetTime sets the "time" field.
func (u *EventUpsert) SetTime(v string) *EventUpsert {
	u.Set(event.FieldTime, v)
	return u
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *EventUpsert) UpdateTime() *EventUpsert {
	u.SetExcluded(event.FieldTime)
	return u
}

// ClearTime clears the value of the "time" field.
func (u *EventUpsert) ClearTime() *EventUpsert {
	u.SetNull(event.FieldTime)
	return u
}

// SetShowTime sets the "showTime" field.
func (u *EventUpsert) SetShowTime(v bool) *EventUpsert {
	u.Set(event.FieldShowTime, v)
	return u
}

// UpdateShowTime sets the "showTime" field to the value that was provided on create.
func (u *EventUpsert) UpdateShowTime() *EventUpsert {
	u.SetExcluded(event.FieldShowTime)
	return u
}

// SetTitle sets the "title" field.
func (u *EventUpsert) SetTitle(v string) *EventUpsert {
	u.Set(event.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *EventUpsert) UpdateTitle() *EventUpsert {
	u.SetExcluded(event.FieldTitle)
	return u
}

// ClearTitle clears the value of the "title" field.
func (u *EventUpsert) ClearTitle() *EventUpsert {
	u.SetNull(event.FieldTitle)
	return u
}

// SetDescription sets the "description" field.
func (u *EventUpsert) SetDescription(v string) *EventUpsert {
	u.Set(event.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *EventUpsert) UpdateDescription() *EventUpsert {
	u.SetExcluded(event.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *EventUpsert) ClearDescription() *EventUpsert {
	u.SetNull(event.FieldDescription)
	return u
}

// SetURL sets the "url" field.
func (u *EventUpsert) SetURL(v string) *EventUpsert {
	u.Set(event.FieldURL, v)
	return u
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *EventUpsert) UpdateURL() *EventUpsert {
	u.SetExcluded(event.FieldURL)
	return u
}

// ClearURL clears the value of the "url" field.
func (u *EventUpsert) ClearURL() *EventUpsert {
	u.SetNull(event.FieldURL)
	return u
}

// SetPreviewlyImageID sets the "previewly_image_id" field.
func (u *EventUpsert) SetPreviewlyImageID(v int) *EventUpsert {
	u.Set(event.FieldPreviewlyImageID, v)
	return u
}

// UpdatePreviewlyImageID sets the "previewly_image_id" field to the value that was provided on create.
func (u *EventUpsert) UpdatePreviewlyImageID() *EventUpsert {
	u.SetExcluded(event.FieldPreviewlyImageID)
	return u
}

// AddPreviewlyImageID adds v to the "previewly_image_id" field.
func (u *EventUpsert) AddPreviewlyImageID(v int) *EventUpsert {
	u.Add(event.FieldPreviewlyImageID, v)
	return u
}

// ClearPreviewlyImageID clears the value of the "previewly_image_id" field.
func (u *EventUpsert) ClearPreviewlyImageID() *EventUpsert {
	u.SetNull(event.FieldPreviewlyImageID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *EventUpsertOne) UpdateNewValues() *EventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(event.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Event.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *EventUpsertOne) Ignore() *EventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EventUpsertOne) DoNothing() *EventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EventCreate.OnConflict
// documentation for more info.
func (u *EventUpsertOne) Update(set func(*EventUpsert)) *EventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EventUpsert{UpdateSet: update})
	}))
	return u
}

// SetDate sets the "date" field.
func (u *EventUpsertOne) SetDate(v time.Time) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetDate(v)
	})
}

// UpdateDate sets the "date" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateDate() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateDate()
	})
}

// SetType sets the "type" field.
func (u *EventUpsertOne) SetType(v event.Type) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateType() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateType()
	})
}

// SetTime sets the "time" field.
func (u *EventUpsertOne) SetTime(v string) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetTime(v)
	})
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateTime() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateTime()
	})
}

// ClearTime clears the value of the "time" field.
func (u *EventUpsertOne) ClearTime() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.ClearTime()
	})
}

// SetShowTime sets the "showTime" field.
func (u *EventUpsertOne) SetShowTime(v bool) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetShowTime(v)
	})
}

// UpdateShowTime sets the "showTime" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateShowTime() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateShowTime()
	})
}

// SetTitle sets the "title" field.
func (u *EventUpsertOne) SetTitle(v string) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateTitle() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateTitle()
	})
}

// ClearTitle clears the value of the "title" field.
func (u *EventUpsertOne) ClearTitle() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.ClearTitle()
	})
}

// SetDescription sets the "description" field.
func (u *EventUpsertOne) SetDescription(v string) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateDescription() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *EventUpsertOne) ClearDescription() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.ClearDescription()
	})
}

// SetURL sets the "url" field.
func (u *EventUpsertOne) SetURL(v string) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateURL() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateURL()
	})
}

// ClearURL clears the value of the "url" field.
func (u *EventUpsertOne) ClearURL() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.ClearURL()
	})
}

// SetPreviewlyImageID sets the "previewly_image_id" field.
func (u *EventUpsertOne) SetPreviewlyImageID(v int) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetPreviewlyImageID(v)
	})
}

// AddPreviewlyImageID adds v to the "previewly_image_id" field.
func (u *EventUpsertOne) AddPreviewlyImageID(v int) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.AddPreviewlyImageID(v)
	})
}

// UpdatePreviewlyImageID sets the "previewly_image_id" field to the value that was provided on create.
func (u *EventUpsertOne) UpdatePreviewlyImageID() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdatePreviewlyImageID()
	})
}

// ClearPreviewlyImageID clears the value of the "previewly_image_id" field.
func (u *EventUpsertOne) ClearPreviewlyImageID() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.ClearPreviewlyImageID()
	})
}

// Exec executes the query.
func (u *EventUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EventCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EventUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EventUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EventUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EventCreateBulk is the builder for creating many Event entities in bulk.
type EventCreateBulk struct {
	config
	err      error
	builders []*EventCreate
	conflict []sql.ConflictOption
}

// Save creates the Event entities in the database.
func (ecb *EventCreateBulk) Save(ctx context.Context) ([]*Event, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Event, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EventCreateBulk) SaveX(ctx context.Context) []*Event {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EventCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EventCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Event.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EventUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ecb *EventCreateBulk) OnConflict(opts ...sql.ConflictOption) *EventUpsertBulk {
	ecb.conflict = opts
	return &EventUpsertBulk{
		create: ecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ecb *EventCreateBulk) OnConflictColumns(columns ...string) *EventUpsertBulk {
	ecb.conflict = append(ecb.conflict, sql.ConflictColumns(columns...))
	return &EventUpsertBulk{
		create: ecb,
	}
}

// EventUpsertBulk is the builder for "upsert"-ing
// a bulk of Event nodes.
type EventUpsertBulk struct {
	create *EventCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *EventUpsertBulk) UpdateNewValues() *EventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(event.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *EventUpsertBulk) Ignore() *EventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EventUpsertBulk) DoNothing() *EventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EventCreateBulk.OnConflict
// documentation for more info.
func (u *EventUpsertBulk) Update(set func(*EventUpsert)) *EventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EventUpsert{UpdateSet: update})
	}))
	return u
}

// SetDate sets the "date" field.
func (u *EventUpsertBulk) SetDate(v time.Time) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetDate(v)
	})
}

// UpdateDate sets the "date" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateDate() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateDate()
	})
}

// SetType sets the "type" field.
func (u *EventUpsertBulk) SetType(v event.Type) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateType() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateType()
	})
}

// SetTime sets the "time" field.
func (u *EventUpsertBulk) SetTime(v string) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetTime(v)
	})
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateTime() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateTime()
	})
}

// ClearTime clears the value of the "time" field.
func (u *EventUpsertBulk) ClearTime() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.ClearTime()
	})
}

// SetShowTime sets the "showTime" field.
func (u *EventUpsertBulk) SetShowTime(v bool) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetShowTime(v)
	})
}

// UpdateShowTime sets the "showTime" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateShowTime() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateShowTime()
	})
}

// SetTitle sets the "title" field.
func (u *EventUpsertBulk) SetTitle(v string) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateTitle() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateTitle()
	})
}

// ClearTitle clears the value of the "title" field.
func (u *EventUpsertBulk) ClearTitle() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.ClearTitle()
	})
}

// SetDescription sets the "description" field.
func (u *EventUpsertBulk) SetDescription(v string) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateDescription() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *EventUpsertBulk) ClearDescription() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.ClearDescription()
	})
}

// SetURL sets the "url" field.
func (u *EventUpsertBulk) SetURL(v string) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetURL(v)
	})
}

// UpdateURL sets the "url" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateURL() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateURL()
	})
}

// ClearURL clears the value of the "url" field.
func (u *EventUpsertBulk) ClearURL() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.ClearURL()
	})
}

// SetPreviewlyImageID sets the "previewly_image_id" field.
func (u *EventUpsertBulk) SetPreviewlyImageID(v int) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetPreviewlyImageID(v)
	})
}

// AddPreviewlyImageID adds v to the "previewly_image_id" field.
func (u *EventUpsertBulk) AddPreviewlyImageID(v int) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.AddPreviewlyImageID(v)
	})
}

// UpdatePreviewlyImageID sets the "previewly_image_id" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdatePreviewlyImageID() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdatePreviewlyImageID()
	})
}

// ClearPreviewlyImageID clears the value of the "previewly_image_id" field.
func (u *EventUpsertBulk) ClearPreviewlyImageID() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.ClearPreviewlyImageID()
	})
}

// Exec executes the query.
func (u *EventUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the EventCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EventCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EventUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
