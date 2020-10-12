package pick

import (
	"github.com/dwalker109/record-club-api/lib/db"
	"github.com/google/uuid"
)

func GetAll() (*[]Pick, error) {
	var p []Pick
	r := db.Conn.Limit(-1).Find(&p)
	return &p, r.Error
}

func GetOne(id string) (*Pick, error) {
	u, _ := uuid.Parse(id)
	var p Pick
	r := db.Conn.First(&p, "pick_id", u)
	return &p, r.Error
}

func AddOne(p *Pick) error {
	r := db.Conn.Create(p)
	return r.Error
}

//func DeleteOne(p *entity.Pick) {
//	db.Conn.Delete(p)
//}
