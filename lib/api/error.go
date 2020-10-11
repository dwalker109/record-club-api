package api

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrorResponse struct {
	HttpStatusCode int    `json:"-"`
	Err            error  `json:"-"`
	Msg            string `json:"msg"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HttpStatusCode)
	e.Msg = e.Err.Error()
	return nil
}

type AppErrString string

func NewErrorResponse(status int, err error) render.Renderer {
	return &ErrorResponse{status, err, "unknown error"}
}
