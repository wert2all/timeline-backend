// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"timeline/backend/ent/account"
	"timeline/backend/ent/timeline"
	"timeline/backend/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (ac *AccountCreate) SetName(s string) *AccountCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetAvatar sets the "avatar" field.
func (ac *AccountCreate) SetAvatar(s string) *AccountCreate {
	ac.mutation.SetAvatar(s)
	return ac
}

// AddTimelineIDs adds the "timeline" edge to the Timeline entity by IDs.
func (ac *AccountCreate) AddTimelineIDs(ids ...int) *AccountCreate {
	ac.mutation.AddTimelineIDs(ids...)
	return ac
}

// AddTimeline adds the "timeline" edges to the Timeline entity.
func (ac *AccountCreate) AddTimeline(t ...*Timeline) *AccountCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddTimelineIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ac *AccountCreate) SetUserID(id int) *AccountCreate {
	ac.mutation.SetUserID(id)
	return ac
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ac *AccountCreate) SetNillableUserID(id *int) *AccountCreate {
	if id != nil {
		ac = ac.SetUserID(*id)
	}
	return ac
}

// SetUser sets the "user" edge to the User entity.
func (ac *AccountCreate) SetUser(u *User) *AccountCreate {
	return ac.SetUserID(u.ID)
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AccountCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AccountCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Account.name"`)}
	}
	if _, ok := ac.mutation.Avatar(); !ok {
		return &ValidationError{Name: "avatar", err: errors.New(`ent: missing required field "Account.avatar"`)}
	}
	return nil
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		_node = &Account{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(account.Table, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ac.conflict
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(account.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Avatar(); ok {
		_spec.SetField(account.FieldAvatar, field.TypeString, value)
		_node.Avatar = &value
	}
	if nodes := ac.mutation.TimelineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.TimelineTable,
			Columns: []string{account.TimelineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(timeline.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   account.UserTable,
			Columns: []string{account.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_account = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Account.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccountUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ac *AccountCreate) OnConflict(opts ...sql.ConflictOption) *AccountUpsertOne {
	ac.conflict = opts
	return &AccountUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AccountCreate) OnConflictColumns(columns ...string) *AccountUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AccountUpsertOne{
		create: ac,
	}
}

type (
	// AccountUpsertOne is the builder for "upsert"-ing
	//  one Account node.
	AccountUpsertOne struct {
		create *AccountCreate
	}

	// AccountUpsert is the "OnConflict" setter.
	AccountUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *AccountUpsert) SetName(v string) *AccountUpsert {
	u.Set(account.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AccountUpsert) UpdateName() *AccountUpsert {
	u.SetExcluded(account.FieldName)
	return u
}

// SetAvatar sets the "avatar" field.
func (u *AccountUpsert) SetAvatar(v string) *AccountUpsert {
	u.Set(account.FieldAvatar, v)
	return u
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *AccountUpsert) UpdateAvatar() *AccountUpsert {
	u.SetExcluded(account.FieldAvatar)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AccountUpsertOne) UpdateNewValues() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AccountUpsertOne) Ignore() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AccountUpsertOne) DoNothing() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AccountCreate.OnConflict
// documentation for more info.
func (u *AccountUpsertOne) Update(set func(*AccountUpsert)) *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AccountUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *AccountUpsertOne) SetName(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateName() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateName()
	})
}

// SetAvatar sets the "avatar" field.
func (u *AccountUpsertOne) SetAvatar(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetAvatar(v)
	})
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateAvatar() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateAvatar()
	})
}

// Exec executes the query.
func (u *AccountUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AccountCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AccountUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AccountUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AccountUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AccountCreateBulk is the builder for creating many Account entities in bulk.
type AccountCreateBulk struct {
	config
	err      error
	builders []*AccountCreate
	conflict []sql.ConflictOption
}

// Save creates the Account entities in the database.
func (acb *AccountCreateBulk) Save(ctx context.Context) ([]*Account, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Account, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccountCreateBulk) SaveX(ctx context.Context) []*Account {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AccountCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AccountCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Account.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccountUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (acb *AccountCreateBulk) OnConflict(opts ...sql.ConflictOption) *AccountUpsertBulk {
	acb.conflict = opts
	return &AccountUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AccountCreateBulk) OnConflictColumns(columns ...string) *AccountUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AccountUpsertBulk{
		create: acb,
	}
}

// AccountUpsertBulk is the builder for "upsert"-ing
// a bulk of Account nodes.
type AccountUpsertBulk struct {
	create *AccountCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AccountUpsertBulk) UpdateNewValues() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AccountUpsertBulk) Ignore() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AccountUpsertBulk) DoNothing() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AccountCreateBulk.OnConflict
// documentation for more info.
func (u *AccountUpsertBulk) Update(set func(*AccountUpsert)) *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AccountUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *AccountUpsertBulk) SetName(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateName() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateName()
	})
}

// SetAvatar sets the "avatar" field.
func (u *AccountUpsertBulk) SetAvatar(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetAvatar(v)
	})
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateAvatar() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateAvatar()
	})
}

// Exec executes the query.
func (u *AccountUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AccountCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AccountCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AccountUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
