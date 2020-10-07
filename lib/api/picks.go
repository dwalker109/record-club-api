package api

import (
	"encoding/json"
	"github.com/dwalker109/record-club-api/lib/db"
	"net/http"
	"path"
)

func PicksIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db.Database.Picks)
}

func PicksGet(w http.ResponseWriter, r *http.Request) {
	p, err := db.Database.GetPickByPickID(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func PicksPost(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	p := db.Pick{}
	err := dec.Decode(&p); if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db.Database.AddPick(p)
	w.WriteHeader(http.StatusAccepted)
}
