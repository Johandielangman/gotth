// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~
//      /\_/\
//     ( o.o )
//      > ^ <
//
// Author: Johan Hanekom
// Date: July 2025
//
// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~

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
	// ====> (1) DEFINE A NEW ABOUT PAGE
	c := About()

	// ====> (2) RENDER AND WRITE
	err := views.Render(w, r, views.Layout(c, "About"))
	views.HandleErr(err, w)
}
