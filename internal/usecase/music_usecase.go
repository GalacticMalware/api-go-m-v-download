package usecase

import (
	"context"
	"log/slog"
	"strings"

	"backend-download-youtube/internal/domain"
	"backend-download-youtube/internal/ports"
)

// MusicUseCase maneja la logica de descarga de musica.
type MusicUseCase struct {
	downloader ports.Downloader
	logger     *slog.Logger
}

// NewMusicUseCase crea una nueva instancia del caso de uso de musica.
func NewMusicUseCase(downloader ports.Downloader, logger *slog.Logger) *MusicUseCase {
	return &MusicUseCase{
		downloader: downloader,
		logger:     logger,
	}
}

// Download valida y ejecuta la descarga de musica.
func (uc *MusicUseCase) Download(ctx context.Context, req domain.MusicDownloadRequest) (*domain.DownloadResult, error) {
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
	if !domain.IsValidMusicFormat(req.Format) {
		return nil, domain.ErrInvalidFormat
	}

	uc.logger.Info("iniciando descarga de musica", "url", req.URL, "format", req.Format)

	result, err := uc.downloader.DownloadMusic(ctx, req)
	if err != nil {
		uc.logger.Error("error al descargar musica", "error", err, "url", req.URL)
		return nil, err
	}

	return result, nil
}
