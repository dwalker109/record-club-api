package picks

import (
	"github.com/dwalker109/record-club-api/lib/api"
	"github.com/dwalker109/record-club-api/lib/model"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	p, err := GetAll()
	if err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	render.Respond(w, r, NewPickListResponse(p))
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "PickID")
	p, err := GetOne(id)
	if err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusNotFound, err))
		return
	}

	render.Respond(w, r, NewPickResponse(p))
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
	p := model.Pick{}
	if err := render.DecodeJSON(r.Body, &p); err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusBadRequest, err))
		return
	}

	if err := AddOne(&p); err != nil {
		render.Render(w, r, api.NewErrorResponse(http.StatusInternalServerError, err))
		return
	}

	render.Render(w, r, NewPickResponse(&p))
}
