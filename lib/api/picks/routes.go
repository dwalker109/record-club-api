package picks

import (
	"github.com/go-chi/chi"
)

func Register(r *chi.Mux) {
	r.Route("/v1/picks", func(r chi.Router) {
		r.Get("/", HandleIndex)
		r.Post("/", HandlePost)
		r.Get("/{PickID}", HandleGet)
	})

}
