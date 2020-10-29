package tokens

import (
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

type RefreshTokensMemCache map[string]*RefreshToken

func (t *RefreshTokensMemCache) IssueRefreshToken(id primitive.ObjectID) string {
	key := uniuri.NewLen(uniuri.UUIDLen)
	(*t)[key] = &RefreshToken{
		UserID:    id,
		ExpiresAt: time.Now().Add(MaxAge).Unix(),
	}

	return key
}

func (t *RefreshTokensMemCache) Clear(key string) {
	delete(*t, key)
}
