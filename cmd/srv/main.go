package main

import (
	"github.com/dwalker109/record-club-api/lib/api/picks"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	picks.Register(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
