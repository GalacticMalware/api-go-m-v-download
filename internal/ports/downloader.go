// Package ports define las interfaces (contratos) que la infraestructura
// debe implementar. Es el "puerto" en la arquitectura hexagonal.
package ports

import (
	"context"

	"backend-go-download-music-video/internal/domain"
)

// Downloader define el contrato para cualquier motor de descarga.
// Cualquier implementacion (yt-dlp, otra API, etc.) debe cumplirlo.
type Downloader interface {
	// DownloadVideo descarga un video en el formato indicado.
	DownloadVideo(ctx context.Context, req domain.VideoDownloadRequest) (*domain.DownloadResult, error)

	// DownloadMusic descarga musica (audio) en el formato indicado.
	DownloadMusic(ctx context.Context, req domain.MusicDownloadRequest) (*domain.DownloadResult, error)
}
