// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~
//      /\_/\
//     ( o.o )
//      > ^ <
//
// Author: Johan Hanekom
// Date: July 2025
//
// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~

package server

import (
	"context"
	"fmt"
	"gotth/internal/nonce"
	"log/slog"
	"net/http"

	"github.com/go-chi/httplog/v3"
	"github.com/google/uuid"
)

// Middleware that will add a trace ID to the request context.
func TraceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := uuid.New().String()
		ctx := r.Context()

		httplog.SetAttrs(ctx,
			slog.String("trace_id", traceID),
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// See: https://github.com/TomDoesTech/GOTTH/blob/main/internal/middleware/middleware.go
func CSPMiddleware(next http.Handler) http.Handler {
	// Static hash for HTMX CSS - this is public and doesn't change
	const htmxCSSHash = "sha256-pgn1TCGZX6O77zDvy0oTODMOxemn0oj0LeCnQTRj7Kg="

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new Nonces struct for every request when here.
		// move to outside the handler to use the same nonces in all responses
		nonceSet := nonce.Nonces{
			Htmx:            nonce.GenerateRandomString(16),
			ResponseTargets: nonce.GenerateRandomString(16),
			Tw:              nonce.GenerateRandomString(16),
		}

		// set nonces in context
		ctx := context.WithValue(r.Context(), nonce.NonceKey, nonceSet)
		// insert the nonces into the content security policy header
		cspHeader := fmt.Sprintf("default-src 'self'; script-src 'nonce-%s' 'nonce-%s' ; style-src 'nonce-%s' '%s' https://fonts.googleapis.com; font-src https://fonts.gstatic.com;",
			nonceSet.Htmx,
			nonceSet.ResponseTargets,
			nonceSet.Tw,
			htmxCSSHash)
		w.Header().Set("Content-Security-Policy", cspHeader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TextHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
