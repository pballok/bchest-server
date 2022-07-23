package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/pballok/bchest-server/graph"
	"github.com/pballok/bchest-server/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("BCHEST_PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
