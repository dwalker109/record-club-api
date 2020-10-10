package main

import (
	"github.com/dwalker109/record-club-api/lib/api/picks"
	"github.com/dwalker109/record-club-api/lib/db"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	db.Activate(db.SQLiteConn)

	r := chi.NewRouter()
	picks.Register(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
