package route

import "github.com/go-chi/chi/v5"

type Route interface {
	Apply(router chi.Router)
}
