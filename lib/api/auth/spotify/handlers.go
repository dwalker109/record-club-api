package spotify

import (
	"errors"
	"github.com/dchest/uniuri"
	"github.com/dwalker109/record-club-api/lib/domain/user"
	"github.com/dwalker109/record-club-api/lib/svc"
	"net/http"
	"time"

	"github.com/dwalker109/record-club-api/lib/api"
	"github.com/go-chi/render"
	"github.com/zmb3/spotify"
)

var callbackURL = "http://localhost:3000/oauth/spotify-cb"
var stateCookieName = "state"
var spAuthenticator = spotify.NewAuthenticator(callbackURL, spotify.ScopeUserReadPrivate)

func HandleGetAuthRedirectURL(w http.ResponseWriter, r *http.Request) {
	stateCookieValue := uniuri.New()
	url := spAuthenticator.AuthURL(stateCookieValue)

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

	spToken, err := spAuthenticator.Token(state.Value, r)
	if err != nil {
		render.Respond(w, r, api.NewErrorResponse(http.StatusNotFound, errors.New("couldn't get spToken")))
		return
	}

	spClient := spAuthenticator.NewClient(spToken)
	spUser, err := spClient.CurrentUser()
	if err != nil {
		render.Respond(w, r, api.NewErrorResponse(http.StatusUnauthorized, err))
		return
	}

	rcUser, err := user.GetUserFromSpotifyID(spUser.ID)
	if err != nil {
		render.Respond(w, r, api.NewErrorResponse(http.StatusInternalServerError, err))
	}

	if rcUser == nil {
		rcUser, err = user.AddUserFromSpotify(spUser)
		if err != nil {
			render.Respond(w, r, api.NewErrorResponse(http.StatusInternalServerError, err))
		}
	}

	atok, rtok, err := svc.Ctr.AuthService().MakeAuthTokens(rcUser.UserID)
	if err != nil {
		render.Respond(w, r, api.NewErrorResponse(http.StatusUnauthorized, err))
		return
	}

	svc.Ctr.SpotifyTokenStore().Store(rcUser.UserID, spToken)

	http.SetCookie(w, &http.Cookie{
		Name:   "refresh_token",
		Value:  rtok,
		Path:   "/tokens/refresh",
		MaxAge: int(svc.MaxAge / time.Second),
		//Secure:     false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	render.Respond(w, r, map[string]string{
		"token_type":   "bearer",
		"access_token": atok,
	})
}

func HandleRefreshToken(w http.ResponseWriter, r *http.Request) {
	key, _ := r.Cookie("refresh_token")

	atok, rtok, err := svc.Ctr.AuthService().RefreshAuthTokens(key.Value)
	if err != nil {
		render.Respond(w, r, api.NewErrorResponse(http.StatusUnauthorized, err))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "refresh_token",
		Value:  rtok,
		Path:   "/tokens/refresh",
		MaxAge: int(svc.MaxAge / time.Second),
		//Secure:     false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	render.Respond(w, r, map[string]string{
		"token_type":   "bearer",
		"access_token": atok,
	})
}

//func HandleDecodeJWT(w http.ResponseWriter, r *http.Request) {
//	sb, _ := ioutil.ReadAll(r.Body)
//	ss := string(sb)
//
//	tok, err := jwt.ParseWithClaims(ss, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			msg := fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])
//			return nil, errors.New(msg)
//		}
//
//		return HMACKey, nil
//	})
//	if err != nil {
//		render.Respond(w, r, api.NewErrorResponse(http.StatusBadRequest, err))
//		return
//	}
//
//	if claims, ok := tok.Claims.(*jwt.StandardClaims); ok && tok.Valid {
//		render.Respond(w, r, claims)
//	} else {
//		render.Respond(w, r, api.NewErrorResponse(http.StatusBadRequest, err))
//	}
//}
