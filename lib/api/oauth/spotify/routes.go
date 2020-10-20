package spotify

import "github.com/go-chi/chi"

func Register(r *chi.Mux) {
	r.Get("/oauth-sp-init", HandleGetAuthRedirectUrl)
	r.Get("/oauth-sp-cb", HandleAuthCallback)
	r.Post("/jwt-decode", HandleDecodeJWT)
}
