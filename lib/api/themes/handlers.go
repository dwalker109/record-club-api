package themes

import (
	"github.com/dwalker109/record-club-api/lib/api"
	"github.com/dwalker109/record-club-api/lib/domain/theme"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	ent, err := theme.GetAll()
	if err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	render.Respond(w, r, NewListResponse(http.StatusOK, ent))
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
	themeOID, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "ThemeID"))
	ent, err := theme.GetOne(themeOID)
	if err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusNotFound, err))
		return
	}

	render.Respond(w, r, NewApiResponse(http.StatusOK, ent))
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	dto := &theme.DTO{}
	if err := render.DecodeJSON(r.Body, dto); err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusBadRequest, err))
		return
	}

	ent := dto.ToEntity()
	ent.OwnerID = EnsureOID(ent.OwnerID)
	ent.ThemeID = EnsureOID(ent.ThemeID)

	if err := theme.AddOne(ent); err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusBadRequest, err))
		return
	}

	render.Render(w, r, NewApiResponse(http.StatusCreated, ent))
}

func EnsureOID(oid primitive.ObjectID) primitive.ObjectID {
	if oid == primitive.NilObjectID {
		return primitive.NewObjectID()
	}
	return oid
}
