package app

import (
	"net/http"
	"strconv"
	"timeline/backend/app/config"
	"timeline/backend/lib/utils"

	"github.com/go-chi/chi/v5"
)

type (
	closer func()

	Application interface {
		Start()
		Closer() closer
	}

	simpleApplication struct {
		router  *chi.Mux
		listen  config.ListenHost
		onClose closer
	}
)

// Close implements App.
func (a simpleApplication) Closer() closer {
	return a.onClose
}

// Start implements App.
func (a simpleApplication) Start() {
	utils.F("Error: %v", http.ListenAndServe(a.listen.Host+":"+strconv.Itoa(a.listen.Port), a.router))
}

func NewApplication(router *chi.Mux, listen config.ListenHost, onClose closer) Application {
	return simpleApplication{
		router:  router,
		listen:  listen,
		onClose: onClose,
	}
}
