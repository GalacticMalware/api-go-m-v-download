package middleware

import (
	"log/slog"
	"net/http"
)

// Recover captura panicos y responde con 500 sin tumbar el servidor.
func Recover(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					logger.Error("panico recuperado",
						"error", rec,
						"path", r.URL.Path,
					)
					http.Error(w, `{"success":false,"error":"Error interno del servidor"}`, http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
