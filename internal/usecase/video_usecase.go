// Package usecase contiene la logica de negocio orquestada.
package usecase

import (
	"context"
	"log/slog"
	"net/url"
	"strings"

	"backend-go-download-music-video/internal/domain"
	"backend-go-download-music-video/internal/ports"
)

// VideoUseCase maneja la logica de descarga de videos.
type VideoUseCase struct {
	downloader ports.Downloader
	logger     *slog.Logger
}

// NewVideoUseCase crea una nueva instancia del caso de uso de video.
func NewVideoUseCase(downloader ports.Downloader, logger *slog.Logger) *VideoUseCase {
	return &VideoUseCase{
		downloader: downloader,
		logger:     logger,
	}
}

// Download valida y ejecuta la descarga de un video.
func (uc *VideoUseCase) Download(ctx context.Context, req domain.VideoDownloadRequest) (*domain.DownloadResult, error) {
	// Validaciones de campos requeridos
	if strings.TrimSpace(req.URL) == "" {
		return nil, domain.ErrURLRequired
	}
	if strings.TrimSpace(req.Format) == "" {
		return nil, domain.ErrFormatRequired
	}

	// Validar URL
	if !isValidURL(req.URL) {
		return nil, domain.ErrInvalidURL
	}

	// Validar formato permitido
	if !domain.IsValidVideoFormat(req.Format) {
		return nil, domain.ErrInvalidFormat
	}

	uc.logger.Info("iniciando descarga de video", "url", req.URL, "format", req.Format)

	result, err := uc.downloader.DownloadVideo(ctx, req)
	if err != nil {
		uc.logger.Error("error al descargar video", "error", err, "url", req.URL)
		return nil, err
	}

	return result, nil
}

// isValidURL verifica que la cadena sea una URL valida con esquema http/https.
func isValidURL(raw string) bool {
	raw = strings.TrimSpace(raw)
	// Permitir URLs sin esquema agregandolo
	if !strings.HasPrefix(raw, "http://") && !strings.HasPrefix(raw, "https://") {
		raw = "https://" + raw
	}
	u, err := url.ParseRequestURI(raw)
	if err != nil {
		return false
	}
	return u.Scheme == "http" || u.Scheme == "https"
}
