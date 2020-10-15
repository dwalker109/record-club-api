package themes

import (
	"github.com/dwalker109/record-club-api/lib/domain/theme"
	"github.com/go-chi/render"
	"net/http"
)

type ApiResponse struct {
	HttpStatusCode int `json:"-"`
	*theme.DTO
}

func (a *ApiResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, a.HttpStatusCode)
	return nil
}

func NewApiResponse(status int, ent *theme.Entity) *ApiResponse {
	return &ApiResponse{status, ent.ToDTO()}
}

func NewListResponse(status int, ents *[]theme.Entity) []render.Renderer {
	list := make([]render.Renderer, 0, len(*ents))
	for _, ent := range *ents {
		list = append(list, NewApiResponse(status, &ent))
	}
	return list
}
