package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"backend-download-youtube/internal/domain"
	"backend-download-youtube/internal/usecase"
)

// MusicHandler maneja las peticiones HTTP relacionadas con musica.
type MusicHandler struct {
	useCase *usecase.MusicUseCase
	logger  *slog.Logger
}

// NewMusicHandler crea una nueva instancia del handler de musica.
func NewMusicHandler(uc *usecase.MusicUseCase, logger *slog.Logger) *MusicHandler {
	return &MusicHandler{useCase: uc, logger: logger}
}

// Download maneja POST /api/v1/download/music
func (h *MusicHandler) Download(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Método no permitido")
		return
	}

	var req domain.MusicDownloadRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		h.logger.Warn("JSON invalido", "error", err)
		respondError(w, http.StatusBadRequest, "JSON invalido")
		return
	}

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
func (h *MusicHandler) handleError(w http.ResponseWriter, err error) {
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
