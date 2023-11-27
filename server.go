package main

import (
	"log"
	"net/http"
	"os"
	"projectname/api"
	"projectname/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	// connect to the database
	dbClient := db.NewX()

	// setup the routes
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handler.NewDefaultServer(api.NewSchema(dbClient)))

	// figure out what port to run on
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// start the server
	log.Printf("Server is now running on http://localhost:%s/ 🚀", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
