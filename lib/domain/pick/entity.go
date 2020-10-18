package pick

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entity struct {
	ThemeID primitive.ObjectID `bson:"theme"`
	OwnerID primitive.ObjectID `bson:"owner"`
	Tracks  []EntityTrack      `bson:"tracks,omitempty"`
}

type EntityTrack struct {
	ID      string         `bson:"id"`
	Title   string         `bson:"title,omitempty"`
	Artist  string         `bson:"artist,omitempty"`
	Reason  string         `bson:"reason,omitempty"`
	Reviews []EntityReview `bson:"reviews,omitempty"`
}

type EntityReview struct {
	OwnerID primitive.ObjectID `bson:"owner_id"`
	Review  string             `bson:"review,omitempty"`
}

func (ent Entity) ToDTO() *DTO {
	dt := make([]DTOTrack, 0, len(ent.Tracks))

	for _, et := range ent.Tracks {
		dr := make([]DTOReview, 0, len(et.Reviews))

		for _, er := range et.Reviews {
			dr = append(dr, DTOReview{
				er.OwnerID,
				er.Review,
			})
		}

		dt = append(dt, DTOTrack{
			et.ID,
			et.Title,
			et.Artist,
			et.Reason,
			dr,
		})
	}

	return &DTO{
		ent.ThemeID,
		ent.OwnerID,
		dt,
	}
}

type DTO struct {
	ThemeID primitive.ObjectID `json:"theme"`
	OwnerID primitive.ObjectID `json:"owner"`
	Tracks  []DTOTrack         `json:"tracks,omitempty"`
}

type DTOTrack struct {
	ID      string      `json:"id"`
	Title   string      `json:"title,omitempty"`
	Artist  string      `json:"artist,omitempty"`
	Reason  string      `json:reason,omitempty`
	Reviews []DTOReview `json:"reviews,omitempty"`
}

type DTOReview struct {
	OwnerID primitive.ObjectID `json:"owner_id"`
	Review  string             `json:"review,omitempty"`
}

func (dto DTO) ToEntity() *Entity {
	et := make([]EntityTrack, 0, len(dto.Tracks))

	for _, dt := range dto.Tracks {
		er := make([]EntityReview, 0, len(dt.Reviews))

		for _, dr := range dt.Reviews {
			er = append(er, EntityReview{
				dr.OwnerID,
				dr.Review,
			})
		}

		et = append(et, EntityTrack{
			dt.ID,
			dt.Title,
			dt.Artist,
			dt.Reason,
			er,
		})
	}

	return &Entity{
		dto.ThemeID,
		dto.OwnerID,
		et,
	}
}
