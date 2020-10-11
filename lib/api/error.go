package api

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrorResponse struct {
	HttpStatusCode int
	Err            error
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HttpStatusCode)
	return nil
}

type AppErrString string

func NewErrorResponse(status int, err error) render.Renderer {
	return &ErrorResponse{status, err}
}
