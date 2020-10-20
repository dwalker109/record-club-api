package spotify

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dwalker109/record-club-api/lib/api"
	"github.com/go-chi/render"
	"github.com/zmb3/spotify"
	"io/ioutil"
	"net/http"
	"time"
)

var callbackURL string = "http://localhost:8080/oauth-sp-cb"
var stateTMP = "123"
var auth = spotify.NewAuthenticator(callbackURL, spotify.ScopeUserReadPrivate)

var HMACKey = []byte("theendisthebeginningistheend")

func HandleGetAuthRedirectUrl(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, auth.AuthURL(stateTMP))
}

type RCClaims struct {
	UserID string `json:"sub"`
	jwt.StandardClaims
}

func HandleAuthCallback(w http.ResponseWriter, r *http.Request) {
	token, err := auth.Token(stateTMP, r)
	if err != nil {
		render.Respond(w, r, api.NewErrorResponse(http.StatusNotFound, errors.New("couldn't get token")))
		return
	}

	client := auth.NewClient(token)
	user, err := client.CurrentUser()
	if err != nil {
		render.Respond(w, r, api.NewErrorResponse(http.StatusUnauthorized, err))
		return
	}

	claims := RCClaims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "dwrc",
		},
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := tok.SignedString(HMACKey)
	if err != nil {
		render.Respond(w, r, api.NewErrorResponse(http.StatusUnauthorized, err))
		return
	}

	render.Respond(w, r, ss)
}

func HandleDecodeJWT(w http.ResponseWriter, r *http.Request) {
	sb, _ := ioutil.ReadAll(r.Body)
	ss := string(sb)

	tok, err := jwt.ParseWithClaims(ss, &RCClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			msg := fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New(msg)
		}

		return HMACKey, nil
	})
	if err != nil {
		render.Respond(w, r, api.NewErrorResponse(http.StatusBadRequest, err))
		return
	}

	if claims, ok := tok.Claims.(*RCClaims); ok && tok.Valid {
		render.Respond(w, r, claims)
	} else {
		render.Respond(w, r, api.NewErrorResponse(http.StatusBadRequest, err))
	}
}
