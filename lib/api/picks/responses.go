package picks

import (
	"github.com/dwalker109/record-club-api/lib/model"
	"github.com/go-chi/render"
	"net/http"
)

type PickResponse struct {
	*model.Pick
	Test string
}

func (pr *PickResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewPickResponse(Pick *model.Pick) *PickResponse {
	return &PickResponse{
		Pick,
		"123",
	}
}

func NewPickListResponse(picks *[]model.Pick) []render.Renderer {
	list := make([]render.Renderer, 0)
	for _, pick := range *picks {
		list = append(list, NewPickResponse(&pick))
	}
	return list
}
