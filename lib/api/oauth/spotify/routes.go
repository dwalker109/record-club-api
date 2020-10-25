package spotify

import "github.com/go-chi/chi"

func Register(r *chi.Mux) {
	r.Get("/oauth/spotify-redirect", HandleGetAuthRedirectURL)
	r.Get("/oauth/spotify-cb", HandleAuthCallback)
	r.Post("/jwt-decode", HandleDecodeJWT)
}
