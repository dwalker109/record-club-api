package picks

import (
	"fmt"
	"github.com/dwalker109/record-club-api/lib/api"
	"github.com/dwalker109/record-club-api/lib/domain/pick"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func HandleIndexTheme(w http.ResponseWriter, r *http.Request) {
	themeID, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "ThemeID"))
	ent, err := pick.GetThemePicks(themeID)

	if err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	render.Respond(w, r, NewListResponse(ent))
}
func HandleIndexThemeAndOwner(w http.ResponseWriter, r *http.Request) {
	themeID, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "ThemeID"))
	ownerID, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "OwnerID"))

	ent, err := pick.GetThemePicksForOwner(themeID, ownerID)
	if err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	render.Respond(w, r, NewListResponse(ent))
}

func HandlePostAndPut(w http.ResponseWriter, r *http.Request) {
	dto := &pick.DTO{}
	if err := render.DecodeJSON(r.Body, dto); err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusBadRequest, err))
		return
	}

	themeID, errT := primitive.ObjectIDFromHex(chi.URLParam(r, "ThemeID"))
	ownerID, errO := primitive.ObjectIDFromHex(chi.URLParam(r, "OwnerID"))
	if errT != nil || errO != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusBadRequest, fmt.Errorf("theme or owner invalid")))
		return
	}

	dto.ThemeID = themeID
	dto.OwnerID = ownerID

	ent := dto.ToEntity()
	if err := pick.AddOrUpdateThemePicksForOwner(ent); err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	render.Render(w, r, NewResponse(ent))
}

func EnsureOID(oid primitive.ObjectID) primitive.ObjectID {
	if oid == primitive.NilObjectID {
		return primitive.NewObjectID()
	}
	return oid
}
