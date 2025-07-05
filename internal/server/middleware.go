package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/httplog/v3"
	"github.com/google/uuid"
)

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
