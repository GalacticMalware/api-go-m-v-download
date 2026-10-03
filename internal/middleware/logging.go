// Package middleware contiene middlewares HTTP reutilizables.
package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// responseWriter envuelve http.ResponseWriter para capturar el status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Logging registra cada peticion HTTP entrante.
func Logging(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

			next.ServeHTTP(rw, r)

			logger.Info("peticion HTTP",
				"method", r.Method,
				"path", r.URL.Path,
				"status", rw.statusCode,
				"remote", r.RemoteAddr,
				"duration", time.Since(start).String(),
			)
		})
	}
}
