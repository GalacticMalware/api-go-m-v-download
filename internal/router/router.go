// Package router configura las rutas HTTP usando net/http nativo (Go 1.22+).
package router

import (
	"net/http"

	"backend-go-download-music-video/internal/handler"
)

// NewRouter construye el enrutador con todas las rutas de la API.
func NewRouter(
	videoHandler *handler.VideoHandler,
	musicHandler *handler.MusicHandler,
) *http.ServeMux {
	mux := http.NewServeMux()

	// POST /api/v1/download/videos
	mux.HandleFunc("POST /api/v1/download/videos", videoHandler.Download)

	// POST /api/v1/download/music
	mux.HandleFunc("POST /api/v1/download/music", musicHandler.Download)

	// Endpoint de salud
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	return mux
}
