package picks

import (
	"github.com/go-chi/chi"
)

func Register(r *chi.Mux) {
	r.Route("/v1/picks", func(r chi.Router) {
		r.Get("/{ThemeID}", HandleIndexTheme)
		r.Get("/{ThemeID}/{OwnerID}", HandleIndexThemeAndOwner)
		r.Post("/{ThemeID}/{OwnerID}", HandlePostAndPut)
		r.Put("/{ThemeID}/{OwnerID}", HandlePostAndPut)
	})

}
