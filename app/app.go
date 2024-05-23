package app

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type Application interface {
	Start()
}

type app struct {
	router chi.Router
}

func (a *app) Start() {
	log.Fatal(http.ListenAndServe(":8000", a.router))
}

func NewApplication(router chi.Router) Application {
	return &app{router: router}
}
