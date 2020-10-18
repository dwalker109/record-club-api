package theme

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	ThemeID     primitive.ObjectID `bson:"_id,omitempty"`
	OwnerID     primitive.ObjectID `bson:"owner,omitempty"`
	Theme       string             `bson:"theme,omitempty"`
	Description string             `bson:"description,omitempty"`
	Quantity    int8               `bson:"quantity,omitempty"`
}

func (ent Entity) ToDTO() *DTO {
	return &DTO{
		ThemeID:     ent.ThemeID,
		OwnerID:     ent.OwnerID,
		Theme:       ent.Theme,
		Description: ent.Description,
		Quantity:    ent.Quantity,
	}
}

type DTO struct {
	ThemeID     primitive.ObjectID `json:"theme_id,omitempty"`
	OwnerID     primitive.ObjectID `json:"owner_id,omitempty"`
	Theme       string             `json:"theme,omitempty"`
	Description string             `json:"Description,omitempty"`
	Quantity    int8               `json:"quantity,omitempty"`
}

func (dto DTO) ToEntity() *Entity {
	return &Entity{
		ThemeID:     dto.ThemeID,
		OwnerID:     dto.OwnerID,
		Theme:       dto.Theme,
		Description: dto.Description,
		Quantity:    dto.Quantity,
	}
}
