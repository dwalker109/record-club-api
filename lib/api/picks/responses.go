package picks

import (
	"github.com/dwalker109/record-club-api/lib/domain/pick"
	"github.com/go-chi/render"
	"net/http"
)

type Response struct {
	*pick.DTO
}

func (pr *Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewResponse(p *pick.Pick) *Response {
	return &Response{p.ToDTO()}
}

func NewListResponse(pn *[]pick.Pick) []render.Renderer {
	list := make([]render.Renderer, 0)
	for _, p := range *pn {
		list = append(list, NewResponse(&p))
	}
	return list
}
