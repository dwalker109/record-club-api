package pick

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pick struct {
	gorm.Model
	DTO
}

func (pick Pick) ToDTO() *DTO {
	return &DTO{
		PickID:  pick.PickID,
		ThemeID: pick.ThemeID,
		OwnerID: pick.OwnerID,
		Title:   pick.Title,
		Artist:  pick.Artist,
	}
}

type DTO struct {
	PickID  uuid.UUID `json:"pick_id"`
	ThemeID uuid.UUID `json:"theme_id"`
	OwnerID uuid.UUID `json:"owner_id"`
	Title   string    `json:"title"`
	Artist  string    `json:"artist"`
}

func (dto DTO) ToEntity() *Pick {
	return &Pick{
		DTO: dto,
	}
}
