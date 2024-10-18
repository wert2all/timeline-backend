package app

import (
	"context"
	"net/http"
	"strconv"
	"timeline/backend/app/config"
	"timeline/backend/app/di"
	"timeline/backend/lib/utils"

	"github.com/go-chi/chi/v5"
	"github.com/golobby/container/v3"
)

type (
	closer func()

	Application interface {
		Start()
		Closer() closer
	}

	appImpl struct {
		router *chi.Mux
		listen config.ListenHost
		closer closer
	}
)

// Close implements App.
func (a appImpl) Closer() closer {
	return a.closer
}

// Start implements App.
func (a appImpl) Start() {
	utils.F("Error: %v", http.ListenAndServe(a.listen.Host+":"+strconv.Itoa(a.listen.Port), a.router))
}

func NewApplication() (Application, error) {
	di.InitDi(config.NewConfig(), context.Background())
	var application Application
	err := container.Resolve(&application)
	return application, err
}
