package svc

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type accessTokenService struct{}

func (a accessTokenService) makeAccessToken(id primitive.ObjectID, signingKey []byte) (string, error) {
	claims := jwt.StandardClaims{
		Subject:   id.String(),
		Issuer:    "rc",
		ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := tok.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}
