package picks

import (
	"github.com/dwalker109/record-club-api/lib/api"
	"github.com/dwalker109/record-club-api/lib/domain/pick"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	ent, err := pick.GetAll()
	if err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	render.Respond(w, r, NewListResponse(ent))
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "PickID")
	ent, err := pick.GetOne(id)
	if err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusNotFound, err))
		return
	}

	render.Respond(w, r, NewResponse(ent))
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	dto := &pick.DTO{}
	if err := render.DecodeJSON(r.Body, dto); err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusBadRequest, err))
		return
	}

	ent := dto.ToEntity()
	if err := pick.AddOne(ent); err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	render.Render(w, r, NewResponse(ent))
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "PickID")
	if err := pick.DeleteOne(id); err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusBadRequest, err))
		return
	}

	render.Render(w, r, NewResponse(&pick.Entity{}))
}
