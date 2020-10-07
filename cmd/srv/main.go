package main

import (
	"github.com/dwalker109/record-club-api/lib/api"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	http.HandleFunc("/v1/picks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			api.PicksIndex(w, r)
		case http.MethodPost:
			api.PicksPost(w, r)
		}
	})
	http.HandleFunc("/v1/picks/", api.PicksGet)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}