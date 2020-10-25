package spotify

import (
	"errors"
	"fmt"
	"github.com/dchest/uniuri"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dwalker109/record-club-api/lib/api"
	"github.com/go-chi/render"
	"github.com/zmb3/spotify"
)

var callbackURL string = "http://localhost:3000/oauth/spotify-cb"
var stateCookieName = "state"
var auth = spotify.NewAuthenticator(callbackURL, spotify.ScopeUserReadPrivate)
var HMACKey = []byte(uniuri.NewLen(256))

func HandleGetAuthRedirectURL(w http.ResponseWriter, r *http.Request) {
	stateCookieValue := uniuri.New()
	url := auth.AuthURL(stateCookieValue)

	http.SetCookie(w, &http.Cookie{
		Name:  stateCookieName,
		Value: stateCookieValue,
		Path:  "",
		//Secure:     false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	render.Respond(w, r, map[string]string{"url": url})
}

type RCClaims struct {
	UserID string `json:"sub"`
	jwt.StandardClaims
}

func HandleAuthCallback(w http.ResponseWriter, r *http.Request) {
	state, _ := r.Cookie(stateCookieName)
	http.SetCookie(w, &http.Cookie{
		Name:   stateCookieName,
		Value:  "",
		Path:   "",
		MaxAge: -1,
		//Secure:     false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	token, err := auth.Token(state.Value, r)
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

	render.Respond(w, r, map[string]string{"jwt": ss})
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
