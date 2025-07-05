package home

import (
	"gotth/internal/views"
	"log/slog"
	"net/http"
)

type HomeHandler struct {
	logger *slog.Logger
}

func NewHomeHandler(logger *slog.Logger) *HomeHandler {
	return &HomeHandler{
		logger: logger,
	}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling the home")
	c := Home()
	err := views.Render(w, r, views.Layout(c, "Home"))
	views.HandleErr(err, w)
}
