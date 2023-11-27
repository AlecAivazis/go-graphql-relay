package api

import (
	"projectname/api/resolvers"
	"projectname/api/runtime"
	"projectname/db"

	"github.com/99designs/gqlgen/graphql"
)

// NewSchema creates a new instance of the schema
func NewSchema(client *db.Client) graphql.ExecutableSchema {
	return runtime.NewExecutableSchema(runtime.Config{
		Resolvers:  &resolver{},
		Directives: runtime.DirectiveRoot{},
	})
}

// the struct that provides access to the custom resolver for each type
type resolver struct{ client *db.Client }

func (r *resolver) Query() runtime.QueryResolver {
	return &resolvers.Query{Client: r.client}
}
