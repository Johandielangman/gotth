package home

import (
	"gotth/internal/components"
	"gotth/internal/views"
	"log/slog"
	"net/http"
	"strconv"
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
	h.logger.Info("Counting a click - DEBUG START")

	// Parse form first
	err := r.ParseForm()
	if err != nil {
		h.logger.Error("Failed to parse form", "error", err)
	}

	current := r.FormValue("current")
	h.logger.Info("Received current value", "current", current, "formData", r.Form)

	cnt, err := strconv.Atoi(current)
	if err != nil {
		h.logger.Error("Failed to parse current value", "error", err, "current", current)
		cnt = 0
	}
	cnt++

	h.logger.Info("New count", "count", cnt)

	// Return updated Counter component
	components.CounterWithButton(cnt).Render(r.Context(), w)
}
