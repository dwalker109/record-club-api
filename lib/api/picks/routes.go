package picks

import (
	"github.com/dwalker109/record-club-api/lib/svc"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func Register(r *chi.Mux) {
	r.Route("/v1/picks", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(svc.Ctr.GetJWTAuth()))
			r.Use(jwtauth.Authenticator)

			r.Get("/{ThemeID}", HandleIndexTheme)
			r.Get("/{ThemeID}/{OwnerID}", HandleIndexThemeAndOwner)
			r.Post("/{ThemeID}/{OwnerID}", HandlePostAndPut)
			r.Put("/{ThemeID}/{OwnerID}", HandlePostAndPut)
		})
	})
}
