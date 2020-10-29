package main

import (
	"log"
	"net/http"

	"github.com/dwalker109/record-club-api/lib/api/auth/spotify"
	"github.com/dwalker109/record-club-api/lib/api/picks"
	"github.com/dwalker109/record-club-api/lib/api/themes"
	"github.com/dwalker109/record-club-api/lib/svc"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	defer svc.Ctr.Shutdown()
	runMigrations()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	//r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	render.SetContentType(render.ContentTypeJSON)

	picks.Register(r)
	themes.Register(r)
	spotify.Register(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func runMigrations() {
	//TODO! Mongo migrations, behind a flag
}
