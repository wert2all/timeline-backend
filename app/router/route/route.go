package route

import "github.com/go-chi/chi"

type Route interface {
	Apply(router chi.Router)
}
