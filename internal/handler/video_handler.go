package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"backend-go-download-music-video/internal/domain"
	"backend-go-download-music-video/internal/usecase"
)

// VideoHandler maneja las peticiones HTTP relacionadas con video.
type VideoHandler struct {
	useCase *usecase.VideoUseCase
	logger  *slog.Logger
}

// NewVideoHandler crea una nueva instancia del handler de video.
func NewVideoHandler(uc *usecase.VideoUseCase, logger *slog.Logger) *VideoHandler {
	return &VideoHandler{useCase: uc, logger: logger}
}

// Download maneja POST /api/v1/download/videos
func (h *VideoHandler) Download(w http.ResponseWriter, r *http.Request) {
	// Solo POST permitido
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Método no permitido")
		return
	}

	var req domain.VideoDownloadRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		h.logger.Warn("JSON invalido", "error", err)
		respondError(w, http.StatusBadRequest, "JSON invalido")
		return
	}

	// Contexto con timeout para la descarga
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Minute)
	defer cancel()

	result, err := h.useCase.Download(ctx, req)
	if err != nil {
		h.handleError(w, err)
		return
	}

	respondSuccess(w, http.StatusOK, "Descarga completada", result)
}

// handleError mapea errores de dominio a codigos HTTP.
func (h *VideoHandler) handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrInvalidFormat),
		errors.Is(err, domain.ErrURLRequired),
		errors.Is(err, domain.ErrFormatRequired),
		errors.Is(err, domain.ErrInvalidURL):
		respondError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, domain.ErrDownloadFailed):
		respondError(w, http.StatusInternalServerError, err.Error())
	default:
		h.logger.Error("error inesperado", "error", err)
		respondError(w, http.StatusInternalServerError, "Error interno del servidor")
	}
}
