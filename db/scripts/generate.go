package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereInputs(true),
		entgql.WithConfigPath("./gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./api/schema/generated_ent.graphql"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.Extensions(
			ex,
		),
	}

	if err := entc.Generate("./db/schema", &gen.Config{
		Target:  "./db",
		Package: "projectname/db",
		Features: []gen.Feature{
			gen.FeatureSnapshot,
			gen.FeatureUpsert,
			gen.FeatureIntercept,
		},
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
