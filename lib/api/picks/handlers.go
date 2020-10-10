package picks

import (
	"encoding/json"
	"github.com/dwalker109/record-club-api/lib/model"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	p, err := GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "PickID")
	p, err := GetOne(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	p := model.Pick{}
	if err := dec.Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if err := AddOne(&p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(p)
}
