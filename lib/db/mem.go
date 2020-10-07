package db

import (
	"errors"
	"github.com/google/uuid"
)

type DB struct {
	Picks map[uuid.UUID]Pick
}

type Pick struct {
	PickID  uuid.UUID `json:"pick_id"`
	ThemeID uuid.UUID `json:"theme_id"`
	OwnerID uuid.UUID `json:"owner_id"`
	Title   string    `json:"title"`
	Artist  string    `json:"artist"`
}

var Database = &DB {
	Picks: make(map[uuid.UUID]Pick),
}

func (db *DB) AddPick(p Pick) {
	db.Picks[p.PickID] = p
}

func (db *DB) GetPickByPickID(id string) (*Pick, error) {
	uuid, err := uuid.Parse(id); if err != nil {
		return nil, err
	}
	p, ok := db.Picks[uuid]; if !ok {
		return nil, errors.New("pick not found")
	}
	return &p, nil
}

