package main

import (
	"log"
	"net/http"
	"timeline/backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
)

const defaultPort = "3000"

func main() {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	http.Handle("/graphql", srv)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
