// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// EdgeBestFriend holds the string denoting the bestfriend edge name in mutations.
	EdgeBestFriend = "bestFriend"
	// EdgeAdminGroups holds the string denoting the admingroups edge name in mutations.
	EdgeAdminGroups = "adminGroups"
	// Table holds the table name of the user in the database.
	Table = "users"
	// FriendsTable is the table that holds the friends relation/edge.
	FriendsTable = "users"
	// FriendsColumn is the table column denoting the friends relation/edge.
	FriendsColumn = "user_best_friend"
	// BestFriendTable is the table that holds the bestFriend relation/edge.
	BestFriendTable = "users"
	// BestFriendColumn is the table column denoting the bestFriend relation/edge.
	BestFriendColumn = "user_best_friend"
	// AdminGroupsTable is the table that holds the adminGroups relation/edge.
	AdminGroupsTable = "user_groups"
	// AdminGroupsInverseTable is the table name for the UserGroup entity.
	// It exists in this package in order to avoid circular dependency with the "usergroup" package.
	AdminGroupsInverseTable = "user_groups"
	// AdminGroupsColumn is the table column denoting the adminGroups relation/edge.
	AdminGroupsColumn = "user_group_admin"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_best_friend",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByFriendsCount orders the results by friends count.
func ByFriendsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFriendsStep(), opts...)
	}
}

// ByFriends orders the results by friends terms.
func ByFriends(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFriendsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBestFriendField orders the results by bestFriend field.
func ByBestFriendField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBestFriendStep(), sql.OrderByField(field, opts...))
	}
}

// ByAdminGroupsCount orders the results by adminGroups count.
func ByAdminGroupsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAdminGroupsStep(), opts...)
	}
}

// ByAdminGroups orders the results by adminGroups terms.
func ByAdminGroups(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAdminGroupsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newFriendsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, FriendsTable, FriendsColumn),
	)
}
func newBestFriendStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, BestFriendTable, BestFriendColumn),
	)
}
func newAdminGroupsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AdminGroupsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, AdminGroupsTable, AdminGroupsColumn),
	)
}
