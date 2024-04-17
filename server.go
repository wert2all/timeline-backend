package main

import (
	"log"
	"net/http"
	"timeline/backend/auth"
	"timeline/backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi"
)

const defaultPort = "8000"

func main() {
	router := chi.NewRouter()
	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/graphql", srv)
	http.Handle("/graphql", srv)
	log.Fatal(http.ListenAndServe(":"+defaultPort, router))
}
