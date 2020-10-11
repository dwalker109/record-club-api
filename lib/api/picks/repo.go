package picks

import (
	"github.com/dwalker109/record-club-api/lib/db"
	"github.com/dwalker109/record-club-api/lib/model"
	"github.com/google/uuid"
)

func GetAll() (*[]model.Pick, error) {
	var p []model.Pick
	r := db.Conn.Limit(-1).Find(&p)
	return &p, r.Error
}

func GetOne(id string) (*model.Pick, error) {
	u, _ := uuid.Parse(id)
	var p model.Pick
	r := db.Conn.First(&p, model.Pick{PickID: u})
	return &p, r.Error
}

func AddOne(p *model.Pick) error {
	r := db.Conn.Create(p)
	return r.Error
}

//func DeleteOne(p *model.Pick) {
//	db.Conn.Delete(p)
//}
