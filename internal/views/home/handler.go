package home

import (
	"gotth/internal/components"
	"gotth/internal/views"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type HomeHandler struct {
	logger *slog.Logger
}

func NewHomeHandler(logger *slog.Logger) *HomeHandler {
	return &HomeHandler{
		logger: logger,
	}
}

func (h *HomeHandler) ServeGetHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling the home")
	c := Home(0)
	err := views.Render(w, r, views.Layout(c, "Home"))
	views.HandleErr(err, w)
}

func (h *HomeHandler) Count(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Counting a click")

	currentStr := chi.URLParam(r, "count")

	cnt, err := strconv.Atoi(currentStr)
	if err != nil {
		h.logger.Error("Failed to parse count from URL", "error", err, "count", currentStr)
		cnt = 0
	}
	cnt++

	h.logger.Info("New count", "count", cnt)

	components.CounterWithButton(cnt).Render(r.Context(), w)
}
