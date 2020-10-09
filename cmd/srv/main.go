package main

import (
	"github.com/dwalker109/record-club-api/lib/api/picks"
	"github.com/dwalker109/record-club-api/lib/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db.Activate(db.SQLiteConn)

	r := mux.NewRouter()
	picks.Register(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
