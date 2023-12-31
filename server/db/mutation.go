// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"projectname/db/predicate"
	"projectname/db/user"
	"projectname/db/usergroup"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser      = "User"
	TypeUserGroup = "UserGroup"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op                 Op
	typ                string
	id                 *int
	name               *string
	clearedFields      map[string]struct{}
	friends            map[int]struct{}
	removedfriends     map[int]struct{}
	clearedfriends     bool
	bestFriend         *int
	clearedbestFriend  bool
	adminGroups        map[int]struct{}
	removedadminGroups map[int]struct{}
	clearedadminGroups bool
	done               bool
	oldValue           func(context.Context) (*User, error)
	predicates         []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("db: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// AddFriendIDs adds the "friends" edge to the User entity by ids.
func (m *UserMutation) AddFriendIDs(ids ...int) {
	if m.friends == nil {
		m.friends = make(map[int]struct{})
	}
	for i := range ids {
		m.friends[ids[i]] = struct{}{}
	}
}

// ClearFriends clears the "friends" edge to the User entity.
func (m *UserMutation) ClearFriends() {
	m.clearedfriends = true
}

// FriendsCleared reports if the "friends" edge to the User entity was cleared.
func (m *UserMutation) FriendsCleared() bool {
	return m.clearedfriends
}

// RemoveFriendIDs removes the "friends" edge to the User entity by IDs.
func (m *UserMutation) RemoveFriendIDs(ids ...int) {
	if m.removedfriends == nil {
		m.removedfriends = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.friends, ids[i])
		m.removedfriends[ids[i]] = struct{}{}
	}
}

// RemovedFriends returns the removed IDs of the "friends" edge to the User entity.
func (m *UserMutation) RemovedFriendsIDs() (ids []int) {
	for id := range m.removedfriends {
		ids = append(ids, id)
	}
	return
}

// FriendsIDs returns the "friends" edge IDs in the mutation.
func (m *UserMutation) FriendsIDs() (ids []int) {
	for id := range m.friends {
		ids = append(ids, id)
	}
	return
}

// ResetFriends resets all changes to the "friends" edge.
func (m *UserMutation) ResetFriends() {
	m.friends = nil
	m.clearedfriends = false
	m.removedfriends = nil
}

// SetBestFriendID sets the "bestFriend" edge to the User entity by id.
func (m *UserMutation) SetBestFriendID(id int) {
	m.bestFriend = &id
}

// ClearBestFriend clears the "bestFriend" edge to the User entity.
func (m *UserMutation) ClearBestFriend() {
	m.clearedbestFriend = true
}

// BestFriendCleared reports if the "bestFriend" edge to the User entity was cleared.
func (m *UserMutation) BestFriendCleared() bool {
	return m.clearedbestFriend
}

// BestFriendID returns the "bestFriend" edge ID in the mutation.
func (m *UserMutation) BestFriendID() (id int, exists bool) {
	if m.bestFriend != nil {
		return *m.bestFriend, true
	}
	return
}

// BestFriendIDs returns the "bestFriend" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// BestFriendID instead. It exists only for internal usage by the builders.
func (m *UserMutation) BestFriendIDs() (ids []int) {
	if id := m.bestFriend; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetBestFriend resets all changes to the "bestFriend" edge.
func (m *UserMutation) ResetBestFriend() {
	m.bestFriend = nil
	m.clearedbestFriend = false
}

// AddAdminGroupIDs adds the "adminGroups" edge to the UserGroup entity by ids.
func (m *UserMutation) AddAdminGroupIDs(ids ...int) {
	if m.adminGroups == nil {
		m.adminGroups = make(map[int]struct{})
	}
	for i := range ids {
		m.adminGroups[ids[i]] = struct{}{}
	}
}

// ClearAdminGroups clears the "adminGroups" edge to the UserGroup entity.
func (m *UserMutation) ClearAdminGroups() {
	m.clearedadminGroups = true
}

// AdminGroupsCleared reports if the "adminGroups" edge to the UserGroup entity was cleared.
func (m *UserMutation) AdminGroupsCleared() bool {
	return m.clearedadminGroups
}

// RemoveAdminGroupIDs removes the "adminGroups" edge to the UserGroup entity by IDs.
func (m *UserMutation) RemoveAdminGroupIDs(ids ...int) {
	if m.removedadminGroups == nil {
		m.removedadminGroups = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.adminGroups, ids[i])
		m.removedadminGroups[ids[i]] = struct{}{}
	}
}

// RemovedAdminGroups returns the removed IDs of the "adminGroups" edge to the UserGroup entity.
func (m *UserMutation) RemovedAdminGroupsIDs() (ids []int) {
	for id := range m.removedadminGroups {
		ids = append(ids, id)
	}
	return
}

// AdminGroupsIDs returns the "adminGroups" edge IDs in the mutation.
func (m *UserMutation) AdminGroupsIDs() (ids []int) {
	for id := range m.adminGroups {
		ids = append(ids, id)
	}
	return
}

// ResetAdminGroups resets all changes to the "adminGroups" edge.
func (m *UserMutation) ResetAdminGroups() {
	m.adminGroups = nil
	m.clearedadminGroups = false
	m.removedadminGroups = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 3)
	if m.friends != nil {
		edges = append(edges, user.EdgeFriends)
	}
	if m.bestFriend != nil {
		edges = append(edges, user.EdgeBestFriend)
	}
	if m.adminGroups != nil {
		edges = append(edges, user.EdgeAdminGroups)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeFriends:
		ids := make([]ent.Value, 0, len(m.friends))
		for id := range m.friends {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeBestFriend:
		if id := m.bestFriend; id != nil {
			return []ent.Value{*id}
		}
	case user.EdgeAdminGroups:
		ids := make([]ent.Value, 0, len(m.adminGroups))
		for id := range m.adminGroups {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 3)
	if m.removedfriends != nil {
		edges = append(edges, user.EdgeFriends)
	}
	if m.removedadminGroups != nil {
		edges = append(edges, user.EdgeAdminGroups)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeFriends:
		ids := make([]ent.Value, 0, len(m.removedfriends))
		for id := range m.removedfriends {
			ids = append(ids, id)
		}
		return ids
	case user.EdgeAdminGroups:
		ids := make([]ent.Value, 0, len(m.removedadminGroups))
		for id := range m.removedadminGroups {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 3)
	if m.clearedfriends {
		edges = append(edges, user.EdgeFriends)
	}
	if m.clearedbestFriend {
		edges = append(edges, user.EdgeBestFriend)
	}
	if m.clearedadminGroups {
		edges = append(edges, user.EdgeAdminGroups)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeFriends:
		return m.clearedfriends
	case user.EdgeBestFriend:
		return m.clearedbestFriend
	case user.EdgeAdminGroups:
		return m.clearedadminGroups
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	case user.EdgeBestFriend:
		m.ClearBestFriend()
		return nil
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeFriends:
		m.ResetFriends()
		return nil
	case user.EdgeBestFriend:
		m.ResetBestFriend()
		return nil
	case user.EdgeAdminGroups:
		m.ResetAdminGroups()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}

// UserGroupMutation represents an operation that mutates the UserGroup nodes in the graph.
type UserGroupMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	clearedFields map[string]struct{}
	admin         *int
	clearedadmin  bool
	done          bool
	oldValue      func(context.Context) (*UserGroup, error)
	predicates    []predicate.UserGroup
}

var _ ent.Mutation = (*UserGroupMutation)(nil)

// usergroupOption allows management of the mutation configuration using functional options.
type usergroupOption func(*UserGroupMutation)

// newUserGroupMutation creates new mutation for the UserGroup entity.
func newUserGroupMutation(c config, op Op, opts ...usergroupOption) *UserGroupMutation {
	m := &UserGroupMutation{
		config:        c,
		op:            op,
		typ:           TypeUserGroup,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserGroupID sets the ID field of the mutation.
func withUserGroupID(id int) usergroupOption {
	return func(m *UserGroupMutation) {
		var (
			err   error
			once  sync.Once
			value *UserGroup
		)
		m.oldValue = func(ctx context.Context) (*UserGroup, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().UserGroup.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUserGroup sets the old UserGroup of the mutation.
func withUserGroup(node *UserGroup) usergroupOption {
	return func(m *UserGroupMutation) {
		m.oldValue = func(context.Context) (*UserGroup, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserGroupMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserGroupMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("db: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserGroupMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserGroupMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().UserGroup.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *UserGroupMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserGroupMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the UserGroup entity.
// If the UserGroup object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserGroupMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserGroupMutation) ResetName() {
	m.name = nil
}

// SetAdminID sets the "admin" edge to the User entity by id.
func (m *UserGroupMutation) SetAdminID(id int) {
	m.admin = &id
}

// ClearAdmin clears the "admin" edge to the User entity.
func (m *UserGroupMutation) ClearAdmin() {
	m.clearedadmin = true
}

// AdminCleared reports if the "admin" edge to the User entity was cleared.
func (m *UserGroupMutation) AdminCleared() bool {
	return m.clearedadmin
}

// AdminID returns the "admin" edge ID in the mutation.
func (m *UserGroupMutation) AdminID() (id int, exists bool) {
	if m.admin != nil {
		return *m.admin, true
	}
	return
}

// AdminIDs returns the "admin" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// AdminID instead. It exists only for internal usage by the builders.
func (m *UserGroupMutation) AdminIDs() (ids []int) {
	if id := m.admin; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetAdmin resets all changes to the "admin" edge.
func (m *UserGroupMutation) ResetAdmin() {
	m.admin = nil
	m.clearedadmin = false
}

// Where appends a list predicates to the UserGroupMutation builder.
func (m *UserGroupMutation) Where(ps ...predicate.UserGroup) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserGroupMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserGroupMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.UserGroup, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserGroupMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserGroupMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (UserGroup).
func (m *UserGroupMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserGroupMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, usergroup.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserGroupMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case usergroup.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserGroupMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case usergroup.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown UserGroup field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserGroupMutation) SetField(name string, value ent.Value) error {
	switch name {
	case usergroup.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown UserGroup field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserGroupMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserGroupMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserGroupMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown UserGroup numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserGroupMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserGroupMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserGroupMutation) ClearField(name string) error {
	return fmt.Errorf("unknown UserGroup nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserGroupMutation) ResetField(name string) error {
	switch name {
	case usergroup.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown UserGroup field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserGroupMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.admin != nil {
		edges = append(edges, usergroup.EdgeAdmin)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserGroupMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case usergroup.EdgeAdmin:
		if id := m.admin; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserGroupMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserGroupMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserGroupMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedadmin {
		edges = append(edges, usergroup.EdgeAdmin)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserGroupMutation) EdgeCleared(name string) bool {
	switch name {
	case usergroup.EdgeAdmin:
		return m.clearedadmin
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserGroupMutation) ClearEdge(name string) error {
	switch name {
	case usergroup.EdgeAdmin:
		m.ClearAdmin()
		return nil
	}
	return fmt.Errorf("unknown UserGroup unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserGroupMutation) ResetEdge(name string) error {
	switch name {
	case usergroup.EdgeAdmin:
		m.ResetAdmin()
		return nil
	}
	return fmt.Errorf("unknown UserGroup edge %s", name)
}
