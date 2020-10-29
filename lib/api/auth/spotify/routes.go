package spotify

import "github.com/go-chi/chi"

func Register(r *chi.Mux) {
	r.Get("/tokens/spotify-redirect", HandleGetAuthRedirectURL)
	r.Get("/tokens/spotify-cb", HandleAuthCallback)
	//r.Post("/jwt-decode", HandleDecodeJWT)
}
