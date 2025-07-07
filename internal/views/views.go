// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~
//      /\_/\
//     ( o.o )
//      > ^ <
//
// Author: Johan Hanekom
// Date: July 2025
//
// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~

// This file defines useful utilities used by the handlers

package views

import (
	"net/http"

	"github.com/a-h/templ"
)

// The Render function will render the template component to the provided http.ResponseWriter.
// It takes the request context to ensure that any context-specific data is available during rendering.
// It returns an error if the rendering fails.
func Render(w http.ResponseWriter, r *http.Request, comp templ.Component) error {
	return comp.Render(r.Context(), w)
}

// HandleErr is a utility function that checks if an error is not nil.
// If there is an error, it writes an HTTP error response with the error message and a status code of 500 (Internal Server Error).
func HandleErr(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
