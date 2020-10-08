package picks

import (
	"encoding/json"
	"github.com/dwalker109/record-club-api/lib/db"
	"github.com/gorilla/mux"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db.Database.GetList())
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
	pickID := mux.Vars(r)["pick_id"]
	p, err := db.Database.GetPickByPickID(pickID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	p := db.Pick{}
	err := dec.Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	db.Database.AddPick(p)
	w.WriteHeader(http.StatusAccepted)
}
