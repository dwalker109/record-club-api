package themes

import (
	"github.com/go-chi/chi"
)

func Register(r *chi.Mux) {
	r.Route("/v1/themes", func(r chi.Router) {
		r.Get("/", HandleIndex)
		r.Post("/", HandlePost)
		r.Get("/{ThemeID}", HandleGet)
	})

}
