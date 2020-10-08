package picks

import (
	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	s := r.PathPrefix("/v1/picks").Subrouter()
	s.HandleFunc("/", HandleIndex).Methods("GET")
	s.HandleFunc("/", HandlePost).Methods("POST")
	s.HandleFunc("/{pick_id}", HandleGet).Methods("GET")
}
