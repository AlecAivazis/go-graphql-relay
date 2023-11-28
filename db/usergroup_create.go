// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"projectname/db/user"
	"projectname/db/usergroup"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserGroupCreate is the builder for creating a UserGroup entity.
type UserGroupCreate struct {
	config
	mutation *UserGroupMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (ugc *UserGroupCreate) SetName(s string) *UserGroupCreate {
	ugc.mutation.SetName(s)
	return ugc
}

// SetAdminID sets the "admin" edge to the User entity by ID.
func (ugc *UserGroupCreate) SetAdminID(id int) *UserGroupCreate {
	ugc.mutation.SetAdminID(id)
	return ugc
}

// SetAdmin sets the "admin" edge to the User entity.
func (ugc *UserGroupCreate) SetAdmin(u *User) *UserGroupCreate {
	return ugc.SetAdminID(u.ID)
}

// Mutation returns the UserGroupMutation object of the builder.
func (ugc *UserGroupCreate) Mutation() *UserGroupMutation {
	return ugc.mutation
}

// Save creates the UserGroup in the database.
func (ugc *UserGroupCreate) Save(ctx context.Context) (*UserGroup, error) {
	return withHooks(ctx, ugc.sqlSave, ugc.mutation, ugc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ugc *UserGroupCreate) SaveX(ctx context.Context) *UserGroup {
	v, err := ugc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ugc *UserGroupCreate) Exec(ctx context.Context) error {
	_, err := ugc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ugc *UserGroupCreate) ExecX(ctx context.Context) {
	if err := ugc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ugc *UserGroupCreate) check() error {
	if _, ok := ugc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`db: missing required field "UserGroup.name"`)}
	}
	if v, ok := ugc.mutation.Name(); ok {
		if err := usergroup.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "UserGroup.name": %w`, err)}
		}
	}
	if _, ok := ugc.mutation.AdminID(); !ok {
		return &ValidationError{Name: "admin", err: errors.New(`db: missing required edge "UserGroup.admin"`)}
	}
	return nil
}

func (ugc *UserGroupCreate) sqlSave(ctx context.Context) (*UserGroup, error) {
	if err := ugc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ugc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ugc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ugc.mutation.id = &_node.ID
	ugc.mutation.done = true
	return _node, nil
}

func (ugc *UserGroupCreate) createSpec() (*UserGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &UserGroup{config: ugc.config}
		_spec = sqlgraph.NewCreateSpec(usergroup.Table, sqlgraph.NewFieldSpec(usergroup.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ugc.conflict
	if value, ok := ugc.mutation.Name(); ok {
		_spec.SetField(usergroup.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := ugc.mutation.AdminIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usergroup.AdminTable,
			Columns: []string{usergroup.AdminColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_group_admin = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserGroup.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserGroupUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ugc *UserGroupCreate) OnConflict(opts ...sql.ConflictOption) *UserGroupUpsertOne {
	ugc.conflict = opts
	return &UserGroupUpsertOne{
		create: ugc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserGroup.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ugc *UserGroupCreate) OnConflictColumns(columns ...string) *UserGroupUpsertOne {
	ugc.conflict = append(ugc.conflict, sql.ConflictColumns(columns...))
	return &UserGroupUpsertOne{
		create: ugc,
	}
}

type (
	// UserGroupUpsertOne is the builder for "upsert"-ing
	//  one UserGroup node.
	UserGroupUpsertOne struct {
		create *UserGroupCreate
	}

	// UserGroupUpsert is the "OnConflict" setter.
	UserGroupUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *UserGroupUpsert) SetName(v string) *UserGroupUpsert {
	u.Set(usergroup.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserGroupUpsert) UpdateName() *UserGroupUpsert {
	u.SetExcluded(usergroup.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.UserGroup.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserGroupUpsertOne) UpdateNewValues() *UserGroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserGroup.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserGroupUpsertOne) Ignore() *UserGroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserGroupUpsertOne) DoNothing() *UserGroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserGroupCreate.OnConflict
// documentation for more info.
func (u *UserGroupUpsertOne) Update(set func(*UserGroupUpsert)) *UserGroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserGroupUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *UserGroupUpsertOne) SetName(v string) *UserGroupUpsertOne {
	return u.Update(func(s *UserGroupUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserGroupUpsertOne) UpdateName() *UserGroupUpsertOne {
	return u.Update(func(s *UserGroupUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *UserGroupUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for UserGroupCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserGroupUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserGroupUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserGroupUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserGroupCreateBulk is the builder for creating many UserGroup entities in bulk.
type UserGroupCreateBulk struct {
	config
	err      error
	builders []*UserGroupCreate
	conflict []sql.ConflictOption
}

// Save creates the UserGroup entities in the database.
func (ugcb *UserGroupCreateBulk) Save(ctx context.Context) ([]*UserGroup, error) {
	if ugcb.err != nil {
		return nil, ugcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ugcb.builders))
	nodes := make([]*UserGroup, len(ugcb.builders))
	mutators := make([]Mutator, len(ugcb.builders))
	for i := range ugcb.builders {
		func(i int, root context.Context) {
			builder := ugcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserGroupMutation)
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
					_, err = mutators[i+1].Mutate(root, ugcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ugcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ugcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ugcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ugcb *UserGroupCreateBulk) SaveX(ctx context.Context) []*UserGroup {
	v, err := ugcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ugcb *UserGroupCreateBulk) Exec(ctx context.Context) error {
	_, err := ugcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ugcb *UserGroupCreateBulk) ExecX(ctx context.Context) {
	if err := ugcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserGroup.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserGroupUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ugcb *UserGroupCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserGroupUpsertBulk {
	ugcb.conflict = opts
	return &UserGroupUpsertBulk{
		create: ugcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserGroup.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ugcb *UserGroupCreateBulk) OnConflictColumns(columns ...string) *UserGroupUpsertBulk {
	ugcb.conflict = append(ugcb.conflict, sql.ConflictColumns(columns...))
	return &UserGroupUpsertBulk{
		create: ugcb,
	}
}

// UserGroupUpsertBulk is the builder for "upsert"-ing
// a bulk of UserGroup nodes.
type UserGroupUpsertBulk struct {
	create *UserGroupCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserGroup.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserGroupUpsertBulk) UpdateNewValues() *UserGroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserGroup.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserGroupUpsertBulk) Ignore() *UserGroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserGroupUpsertBulk) DoNothing() *UserGroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserGroupCreateBulk.OnConflict
// documentation for more info.
func (u *UserGroupUpsertBulk) Update(set func(*UserGroupUpsert)) *UserGroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserGroupUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *UserGroupUpsertBulk) SetName(v string) *UserGroupUpsertBulk {
	return u.Update(func(s *UserGroupUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserGroupUpsertBulk) UpdateName() *UserGroupUpsertBulk {
	return u.Update(func(s *UserGroupUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *UserGroupUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("db: OnConflict was set for builder %d. Set it on the UserGroupCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for UserGroupCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserGroupUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
