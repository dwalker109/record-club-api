package svc

import (
	"fmt"
	"github.com/dchest/uniuri"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	MaxAge = time.Hour * 24 * 365
)

type RefreshToken struct {
	UserID    primitive.ObjectID
	ExpiresAt int64
}

type refreshTokenMemCache map[string]*RefreshToken

func (t *refreshTokenMemCache) issueRefreshToken(id primitive.ObjectID) string {
	key := uniuri.NewLen(uniuri.UUIDLen)
	(*t)[key] = &RefreshToken{
		UserID:    id,
		ExpiresAt: time.Now().Add(MaxAge).Unix(),
	}

	return key
}

func (t *refreshTokenMemCache) pullRefreshToken(key string) (*RefreshToken, error) {
	rtok, ok := (*t)[key]
	if ok == false {
		return &RefreshToken{}, fmt.Errorf("could not retrieve refresh token")
	}
	delete(*t, key)
	return rtok, nil
}

func (t *refreshTokenMemCache) Clear(key string) {
	delete(*t, key)
}
