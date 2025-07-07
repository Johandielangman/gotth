// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~
//      /\_/\
//     ( o.o )
//      > ^ <
//
// Author: Johan Hanekom
// Date: July 2025
//
// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~

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

	// ====> (1) DEFINE A NEW HOME PAGE WITH A DEFAULT COUNTER VALUE
	c := Home(0)

	// ====> (2) RENDER THE COMPONENT AND WRITE THE RESPONSE
	err := views.Render(w, r, views.Layout(c, "Home"))
	views.HandleErr(err, w)
}

func (h *HomeHandler) Count(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Counting a click")

	// ====> (1) GET THE CURRENT COUNT
	currentStr := chi.URLParam(r, "count")

	// ====> (2) CONVERT TO INT AND ADVANCE THE COUNT
	cnt, err := strconv.Atoi(currentStr)
	if err != nil {
		h.logger.Error("Failed to parse count from URL", "error", err, "count", currentStr)
		cnt = 0
	}
	cnt++
	h.logger.Info("New count", "count", cnt)

	// ====> (3) RENDER A NEW COMPONENT BASED ON THE NEW COUNT
	err = views.Render(w, r, components.CounterWithButton(cnt))
	views.HandleErr(err, w)
}
