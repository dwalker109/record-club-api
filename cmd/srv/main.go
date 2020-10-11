package main

import (
	"github.com/dwalker109/record-club-api/lib/api/picks"
	"github.com/dwalker109/record-club-api/lib/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func main() {
	db.Activate(db.SQLiteConn)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	render.SetContentType(render.ContentTypeJSON)
	picks.Register(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
