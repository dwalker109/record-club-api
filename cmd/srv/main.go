package main

import (
	"context"
	"github.com/dwalker109/record-club-api/lib/api/picks"
	"github.com/dwalker109/record-club-api/lib/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func main() {
	defer db.Conn.Disconnect(context.Background())
	runMigrations()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	render.SetContentType(render.ContentTypeJSON)
	picks.Register(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func runMigrations() {
	//TODO! Mongo migrations, behind a flag
}
