// Code generated by ent, DO NOT EDIT.

package db

import (
	"fmt"
	"projectname/db/user"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges            UserEdges `json:"edges"`
	user_best_friend *int
	selectValues     sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Friends holds the value of the friends edge.
	Friends []*User `json:"friends,omitempty"`
	// BestFriend holds the value of the bestFriend edge.
	BestFriend *User `json:"bestFriend,omitempty"`
	// AdminGroups holds the value of the adminGroups edge.
	AdminGroups []*UserGroup `json:"adminGroups,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int

	namedFriends     map[string][]*User
	namedAdminGroups map[string][]*UserGroup
}

// FriendsOrErr returns the Friends value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FriendsOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Friends, nil
	}
	return nil, &NotLoadedError{edge: "friends"}
}

// BestFriendOrErr returns the BestFriend value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) BestFriendOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.BestFriend == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.BestFriend, nil
	}
	return nil, &NotLoadedError{edge: "bestFriend"}
}

// AdminGroupsOrErr returns the AdminGroups value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AdminGroupsOrErr() ([]*UserGroup, error) {
	if e.loadedTypes[2] {
		return e.AdminGroups, nil
	}
	return nil, &NotLoadedError{edge: "adminGroups"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldName:
			values[i] = new(sql.NullString)
		case user.ForeignKeys[0]: // user_best_friend
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_best_friend", value)
			} else if value.Valid {
				u.user_best_friend = new(int)
				*u.user_best_friend = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryFriends queries the "friends" edge of the User entity.
func (u *User) QueryFriends() *UserQuery {
	return NewUserClient(u.config).QueryFriends(u)
}

// QueryBestFriend queries the "bestFriend" edge of the User entity.
func (u *User) QueryBestFriend() *UserQuery {
	return NewUserClient(u.config).QueryBestFriend(u)
}

// QueryAdminGroups queries the "adminGroups" edge of the User entity.
func (u *User) QueryAdminGroups() *UserGroupQuery {
	return NewUserClient(u.config).QueryAdminGroups(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("db: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedFriends returns the Friends named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedFriends(name string) ([]*User, error) {
	if u.Edges.namedFriends == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedFriends[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedFriends(name string, edges ...*User) {
	if u.Edges.namedFriends == nil {
		u.Edges.namedFriends = make(map[string][]*User)
	}
	if len(edges) == 0 {
		u.Edges.namedFriends[name] = []*User{}
	} else {
		u.Edges.namedFriends[name] = append(u.Edges.namedFriends[name], edges...)
	}
}

// NamedAdminGroups returns the AdminGroups named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedAdminGroups(name string) ([]*UserGroup, error) {
	if u.Edges.namedAdminGroups == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedAdminGroups[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedAdminGroups(name string, edges ...*UserGroup) {
	if u.Edges.namedAdminGroups == nil {
		u.Edges.namedAdminGroups = make(map[string][]*UserGroup)
	}
	if len(edges) == 0 {
		u.Edges.namedAdminGroups[name] = []*UserGroup{}
	} else {
		u.Edges.namedAdminGroups[name] = append(u.Edges.namedAdminGroups[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User
