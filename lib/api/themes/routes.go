package themes

import (
	"github.com/dwalker109/record-club-api/lib/svc"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func Register(r *chi.Mux) {
	r.Route("/v1/themes", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(svc.Ctr.JWTAuth()))
			r.Use(jwtauth.Authenticator)

			r.Get("/", HandleIndex)
			r.Post("/", HandlePost)
			r.Get("/{ThemeID}", HandleGet)
		})
	})

}
