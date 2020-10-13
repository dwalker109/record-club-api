package pick

import (
	"github.com/google/uuid"
)

type Entity struct {
	PickID  uuid.UUID `bson:"_id,omitempty"`
	ThemeID uuid.UUID `bson:"theme,omitempty"`
	OwnerID uuid.UUID `bson:"owner,omitempty"`
	Title   string    `bson:"title,omitempty"`
	Artist  string    `json:"artist,omitempty"`
}

func (ent Entity) ToDTO() *DTO {
	return &DTO{
		PickID:  ent.PickID,
		ThemeID: ent.ThemeID,
		OwnerID: ent.OwnerID,
		Title:   ent.Title,
		Artist:  ent.Artist,
	}
}

type DTO struct {
	PickID  uuid.UUID `json:"pick_id"`
	ThemeID uuid.UUID `json:"theme_id"`
	OwnerID uuid.UUID `json:"owner_id"`
	Title   string    `json:"title"`
	Artist  string    `json:"artist"`
}

func (dto DTO) ToEntity() *Entity {
	return &Entity{
		PickID:  dto.PickID,
		ThemeID: dto.ThemeID,
		OwnerID: dto.OwnerID,
		Title:   dto.Title,
		Artist:  dto.Artist,
	}
}
