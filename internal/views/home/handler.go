package home

import (
	"gotth/internal/views"
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := Home()
	err := views.Render(w, r, views.Layout(c, "Home"))
	views.HandleErr(err, w)
}
