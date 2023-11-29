// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package runtime

import (
	"bytes"
	"context"
	"errors"
	"projectname/db"
	"sync/atomic"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		schema:     cfg.Schema,
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Schema     *ast.Schema
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	PageInfo struct {
		EndCursor       func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		StartCursor     func(childComplexity int) int
	}

	Query struct {
		Node       func(childComplexity int, id string) int
		Nodes      func(childComplexity int, ids []string) int
		UserGroups func(childComplexity int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *db.UserGroupWhereInput) int
		Users      func(childComplexity int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *db.UserWhereInput) int
	}

	User struct {
		AdminGroups func(childComplexity int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *db.UserGroupWhereInput) int
		BestFriend  func(childComplexity int) int
		Friends     func(childComplexity int) int
		ID          func(childComplexity int) int
		Name        func(childComplexity int) int
	}

	UserConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	UserEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	UserGroup struct {
		Admin func(childComplexity int) int
		ID    func(childComplexity int) int
		Name  func(childComplexity int) int
	}

	UserGroupConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	UserGroupEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}
}

type executableSchema struct {
	schema     *ast.Schema
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	if e.schema != nil {
		return e.schema
	}
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e, 0, 0, nil}
	_ = ec
	switch typeName + "." + field {

	case "PageInfo.endCursor":
		if e.complexity.PageInfo.EndCursor == nil {
			break
		}

		return e.complexity.PageInfo.EndCursor(childComplexity), true

	case "PageInfo.hasNextPage":
		if e.complexity.PageInfo.HasNextPage == nil {
			break
		}

		return e.complexity.PageInfo.HasNextPage(childComplexity), true

	case "PageInfo.hasPreviousPage":
		if e.complexity.PageInfo.HasPreviousPage == nil {
			break
		}

		return e.complexity.PageInfo.HasPreviousPage(childComplexity), true

	case "PageInfo.startCursor":
		if e.complexity.PageInfo.StartCursor == nil {
			break
		}

		return e.complexity.PageInfo.StartCursor(childComplexity), true

	case "Query.node":
		if e.complexity.Query.Node == nil {
			break
		}

		args, err := ec.field_Query_node_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Node(childComplexity, args["id"].(string)), true

	case "Query.nodes":
		if e.complexity.Query.Nodes == nil {
			break
		}

		args, err := ec.field_Query_nodes_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Nodes(childComplexity, args["ids"].([]string)), true

	case "Query.userGroups":
		if e.complexity.Query.UserGroups == nil {
			break
		}

		args, err := ec.field_Query_userGroups_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.UserGroups(childComplexity, args["after"].(*entgql.Cursor[int]), args["first"].(*int), args["before"].(*entgql.Cursor[int]), args["last"].(*int), args["where"].(*db.UserGroupWhereInput)), true

	case "Query.users":
		if e.complexity.Query.Users == nil {
			break
		}

		args, err := ec.field_Query_users_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Users(childComplexity, args["after"].(*entgql.Cursor[int]), args["first"].(*int), args["before"].(*entgql.Cursor[int]), args["last"].(*int), args["where"].(*db.UserWhereInput)), true

	case "User.adminGroups":
		if e.complexity.User.AdminGroups == nil {
			break
		}

		args, err := ec.field_User_adminGroups_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.User.AdminGroups(childComplexity, args["after"].(*entgql.Cursor[int]), args["first"].(*int), args["before"].(*entgql.Cursor[int]), args["last"].(*int), args["where"].(*db.UserGroupWhereInput)), true

	case "User.bestfriend":
		if e.complexity.User.BestFriend == nil {
			break
		}

		return e.complexity.User.BestFriend(childComplexity), true

	case "User.friends":
		if e.complexity.User.Friends == nil {
			break
		}

		return e.complexity.User.Friends(childComplexity), true

	case "User.id":
		if e.complexity.User.ID == nil {
			break
		}

		return e.complexity.User.ID(childComplexity), true

	case "User.name":
		if e.complexity.User.Name == nil {
			break
		}

		return e.complexity.User.Name(childComplexity), true

	case "UserConnection.edges":
		if e.complexity.UserConnection.Edges == nil {
			break
		}

		return e.complexity.UserConnection.Edges(childComplexity), true

	case "UserConnection.pageInfo":
		if e.complexity.UserConnection.PageInfo == nil {
			break
		}

		return e.complexity.UserConnection.PageInfo(childComplexity), true

	case "UserConnection.totalCount":
		if e.complexity.UserConnection.TotalCount == nil {
			break
		}

		return e.complexity.UserConnection.TotalCount(childComplexity), true

	case "UserEdge.cursor":
		if e.complexity.UserEdge.Cursor == nil {
			break
		}

		return e.complexity.UserEdge.Cursor(childComplexity), true

	case "UserEdge.node":
		if e.complexity.UserEdge.Node == nil {
			break
		}

		return e.complexity.UserEdge.Node(childComplexity), true

	case "UserGroup.admin":
		if e.complexity.UserGroup.Admin == nil {
			break
		}

		return e.complexity.UserGroup.Admin(childComplexity), true

	case "UserGroup.id":
		if e.complexity.UserGroup.ID == nil {
			break
		}

		return e.complexity.UserGroup.ID(childComplexity), true

	case "UserGroup.name":
		if e.complexity.UserGroup.Name == nil {
			break
		}

		return e.complexity.UserGroup.Name(childComplexity), true

	case "UserGroupConnection.edges":
		if e.complexity.UserGroupConnection.Edges == nil {
			break
		}

		return e.complexity.UserGroupConnection.Edges(childComplexity), true

	case "UserGroupConnection.pageInfo":
		if e.complexity.UserGroupConnection.PageInfo == nil {
			break
		}

		return e.complexity.UserGroupConnection.PageInfo(childComplexity), true

	case "UserGroupConnection.totalCount":
		if e.complexity.UserGroupConnection.TotalCount == nil {
			break
		}

		return e.complexity.UserGroupConnection.TotalCount(childComplexity), true

	case "UserGroupEdge.cursor":
		if e.complexity.UserGroupEdge.Cursor == nil {
			break
		}

		return e.complexity.UserGroupEdge.Cursor(childComplexity), true

	case "UserGroupEdge.node":
		if e.complexity.UserGroupEdge.Node == nil {
			break
		}

		return e.complexity.UserGroupEdge.Node(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputUserGroupWhereInput,
		ec.unmarshalInputUserWhereInput,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			var response graphql.Response
			var data graphql.Marshaler
			if first {
				first = false
				ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
				data = ec._Query(ctx, rc.Operation.SelectionSet)
			} else {
				if atomic.LoadInt32(&ec.pendingDeferred) > 0 {
					result := <-ec.deferredResults
					atomic.AddInt32(&ec.pendingDeferred, -1)
					data = result.Result
					response.Path = result.Path
					response.Label = result.Label
					response.Errors = result.Errors
				} else {
					return nil
				}
			}
			var buf bytes.Buffer
			data.MarshalGQL(&buf)
			response.Data = buf.Bytes()
			if atomic.LoadInt32(&ec.deferred) > 0 {
				hasNext := atomic.LoadInt32(&ec.pendingDeferred) > 0
				response.HasNext = &hasNext
			}

			return &response
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
	deferred        int32
	pendingDeferred int32
	deferredResults chan graphql.DeferredResult
}

func (ec *executionContext) processDeferredGroup(dg graphql.DeferredGroup) {
	atomic.AddInt32(&ec.pendingDeferred, 1)
	go func() {
		ctx := graphql.WithFreshResponseContext(dg.Context)
		dg.FieldSet.Dispatch(ctx)
		ds := graphql.DeferredResult{
			Path:   dg.Path,
			Label:  dg.Label,
			Result: dg.FieldSet,
			Errors: graphql.GetErrors(ctx),
		}
		// null fields should bubble up
		if dg.FieldSet.Invalids > 0 {
			ds.Result = graphql.Null
		}
		ec.deferredResults <- ds
	}()
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(ec.Schema()), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(ec.Schema(), ec.Schema().Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../schema/generated_ent.graphql", Input: `directive @goField(forceResolver: Boolean, name: String) on FIELD_DEFINITION | INPUT_FIELD_DEFINITION
directive @goModel(model: String, models: [String!]) on OBJECT | INPUT_OBJECT | SCALAR | ENUM | INTERFACE | UNION
"""
Define a Relay Cursor type:
https://relay.dev/graphql/connections.htm#sec-Cursor
"""
scalar Cursor
"""
An object with an ID.
Follows the [Relay Global Object Identification Specification](https://relay.dev/graphql/objectidentification.htm)
"""
interface Node @goModel(model: "projectname/db.Noder") {
  """The id of the object."""
  id: ID!
}
"""Possible directions in which to order a list of items when provided an ` + "`" + `orderBy` + "`" + ` argument."""
enum OrderDirection {
  """Specifies an ascending order for a given ` + "`" + `orderBy` + "`" + ` argument."""
  ASC
  """Specifies a descending order for a given ` + "`" + `orderBy` + "`" + ` argument."""
  DESC
}
"""
Information about pagination in a connection.
https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
"""
type PageInfo {
  """When paginating forwards, are there more items?"""
  hasNextPage: Boolean!
  """When paginating backwards, are there more items?"""
  hasPreviousPage: Boolean!
  """When paginating backwards, the cursor to continue."""
  startCursor: Cursor
  """When paginating forwards, the cursor to continue."""
  endCursor: Cursor
}
type Query {
  """Fetches an object given its ID."""
  node(
    """ID of the object."""
    id: ID!
  ): Node
  """Lookup nodes by a list of IDs."""
  nodes(
    """The list of node IDs."""
    ids: [ID!]!
  ): [Node]!
  users(
    """Returns the elements in the list that come after the specified cursor."""
    after: Cursor

    """Returns the first _n_ elements from the list."""
    first: Int

    """Returns the elements in the list that come before the specified cursor."""
    before: Cursor

    """Returns the last _n_ elements from the list."""
    last: Int

    """Filtering options for Users returned from the connection."""
    where: UserWhereInput
  ): UserConnection!
  userGroups(
    """Returns the elements in the list that come after the specified cursor."""
    after: Cursor

    """Returns the first _n_ elements from the list."""
    first: Int

    """Returns the elements in the list that come before the specified cursor."""
    before: Cursor

    """Returns the last _n_ elements from the list."""
    last: Int

    """Filtering options for UserGroups returned from the connection."""
    where: UserGroupWhereInput
  ): UserGroupConnection!
}
type User implements Node {
  id: ID!
  name: String!
  friends: [User!]
  bestfriend: User! @goField(name: "BestFriend", forceResolver: false)
  adminGroups(
    """Returns the elements in the list that come after the specified cursor."""
    after: Cursor

    """Returns the first _n_ elements from the list."""
    first: Int

    """Returns the elements in the list that come before the specified cursor."""
    before: Cursor

    """Returns the last _n_ elements from the list."""
    last: Int

    """Filtering options for UserGroups returned from the connection."""
    where: UserGroupWhereInput
  ): UserGroupConnection!
}
"""A connection to a list of items."""
type UserConnection {
  """A list of edges."""
  edges: [UserEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  """Identifies the total count of items in the connection."""
  totalCount: Int!
}
"""An edge in a connection."""
type UserEdge {
  """The item at the end of the edge."""
  node: User
  """A cursor for use in pagination."""
  cursor: Cursor!
}
type UserGroup implements Node {
  id: ID!
  name: String!
  admin: User!
}
"""A connection to a list of items."""
type UserGroupConnection {
  """A list of edges."""
  edges: [UserGroupEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  """Identifies the total count of items in the connection."""
  totalCount: Int!
}
"""An edge in a connection."""
type UserGroupEdge {
  """The item at the end of the edge."""
  node: UserGroup
  """A cursor for use in pagination."""
  cursor: Cursor!
}
"""
UserGroupWhereInput is used for filtering UserGroup objects.
Input was generated by ent.
"""
input UserGroupWhereInput {
  not: UserGroupWhereInput
  and: [UserGroupWhereInput!]
  or: [UserGroupWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """name field predicates"""
  name: String
  nameNEQ: String
  nameIn: [String!]
  nameNotIn: [String!]
  nameGT: String
  nameGTE: String
  nameLT: String
  nameLTE: String
  nameContains: String
  nameHasPrefix: String
  nameHasSuffix: String
  nameEqualFold: String
  nameContainsFold: String
  """admin edge predicates"""
  hasAdmin: Boolean
  hasAdminWith: [UserWhereInput!]
}
"""
UserWhereInput is used for filtering User objects.
Input was generated by ent.
"""
input UserWhereInput {
  not: UserWhereInput
  and: [UserWhereInput!]
  or: [UserWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """name field predicates"""
  name: String
  nameNEQ: String
  nameIn: [String!]
  nameNotIn: [String!]
  nameGT: String
  nameGTE: String
  nameLT: String
  nameLTE: String
  nameContains: String
  nameHasPrefix: String
  nameHasSuffix: String
  nameEqualFold: String
  nameContainsFold: String
  """friends edge predicates"""
  hasFriends: Boolean
  hasFriendsWith: [UserWhereInput!]
  """bestFriend edge predicates"""
  hasBestFriend: Boolean
  hasBestFriendWith: [UserWhereInput!]
  """adminGroups edge predicates"""
  hasAdminGroups: Boolean
  hasAdminGroupsWith: [UserGroupWhereInput!]
}
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)