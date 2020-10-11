package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pick struct {
	gorm.Model
	PickID  uuid.UUID `json:"pick_id"`
	ThemeID uuid.UUID `json:"theme_id"`
	OwnerID uuid.UUID `json:"owner_id"`
	Title   string    `json:"title"`
	Artist  string    `json:"artist"`
}
