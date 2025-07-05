package views

import (
	"net/http"

	"github.com/a-h/templ"
)

func Render(w http.ResponseWriter, r *http.Request, comp templ.Component) error {
	return comp.Render(r.Context(), w)
}

func HandleErr(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
