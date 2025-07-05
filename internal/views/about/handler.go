package about

import (
	"gotth/internal/views"
	"net/http"
)

type AboutHandler struct{}

func NewAboutHandler() *AboutHandler {
	return &AboutHandler{}
}

func (h *AboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := About()
	err := views.Render(w, r, views.Layout(c, "About"))
	views.HandleErr(err, w)
}
