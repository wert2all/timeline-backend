package app

import (
	"log"
	"net/http"
	"timeline/backend/app/router"
	"timeline/backend/di"

	"github.com/go-chi/chi"
)

type Application interface {
	Start()
	Close()
}

type app struct {
	router  chi.Router
	locator di.ServiceLocator
}

func (a *app) Close() {
	a.locator.Close()
}

func (a *app) Start() {
	log.Fatal(http.ListenAndServe(":8000", a.router))
}

func NewApplication(locator di.ServiceLocator) Application {
	return &app{router: router.NewRouterFactory(locator).Create(), locator: locator}
}
