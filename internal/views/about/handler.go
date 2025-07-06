package about

import (
	"gotth/internal/views"
	"log/slog"
	"net/http"
)

type AboutHandler struct {
	logger *slog.Logger
}

func NewAboutHandler(logger *slog.Logger) *AboutHandler {
	return &AboutHandler{
		logger: logger,
	}
}

func (h *AboutHandler) ServeGetHTTP(w http.ResponseWriter, r *http.Request) {
	c := About()
	err := views.Render(w, r, views.Layout(c, "About"))
	views.HandleErr(err, w)
}
