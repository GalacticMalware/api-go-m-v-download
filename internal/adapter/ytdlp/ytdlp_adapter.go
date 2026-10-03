// Package ytdlp implementa el puerto Downloader usando el binario yt-dlp
// como herramienta externa invocada via os/exec.
package ytdlp

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os/exec"
	"path/filepath"
	"strings"

	"backend-go-download-music-video/internal/domain"
)

// Adapter es la implementacion concreta del puerto Downloader.
type Adapter struct {
	downloadDir string
	ytdlpPath   string
	ffmpegPath  string
	logger      *slog.Logger
}

// NewAdapter crea una nueva instancia del adaptador.
func NewAdapter(downloadDir, ytdlpPath, ffmpegPath string, logger *slog.Logger) *Adapter {
	return &Adapter{
		downloadDir: downloadDir,
		ytdlpPath:   ytdlpPath,
		ffmpegPath:  ffmpegPath,
		logger:      logger,
	}
}

// DownloadVideo descarga un video con la resolucion solicitada.
func (a *Adapter) DownloadVideo(ctx context.Context, req domain.VideoDownloadRequest) (*domain.DownloadResult, error) {
	// Mapear formato de dominio a filtro de yt-dlp
	var height string
	switch req.Format {
	case domain.VideoFormat1080p:
		height = "1080"
	case domain.VideoFormat720p:
		height = "720"
	default:
		return nil, domain.ErrInvalidFormat
	}

	outputTemplate := filepath.Join(a.downloadDir, "%(title)s.%(ext)s")

	// Filtro: mejor video+audio hasta la altura indicada, merge a mp4
	formatSelector := fmt.Sprintf("bestvideo[height<=%s]+bestaudio/best[height<=%s]", height, height)

	args := []string{
		"-f", formatSelector,
		"--merge-output-format", "mp4",
		"-o", outputTemplate,
		"--no-playlist",
		"--print", "after_move:filepath",
		req.URL,
	}

	if a.ffmpegPath != "" {
		args = append(args, "--ffmpeg-location", a.ffmpegPath)
	}

	args = append(args, req.URL)

	filePath, err := a.run(ctx, args)
	if err != nil {
		return nil, err
	}

	return &domain.DownloadResult{
		FileName: filepath.Base(filePath),
		FilePath: filePath,
		Format:   req.Format,
		Type:     string(domain.DownloadTypeVideo),
	}, nil
}

// DownloadMusic descarga el audio del video en formato mp3 o mp4.
func (a *Adapter) DownloadMusic(ctx context.Context, req domain.MusicDownloadRequest) (*domain.DownloadResult, error) {
	outputTemplate := filepath.Join(a.downloadDir, "%(title)s.%(ext)s")

	args := []string{
		"--no-playlist",
		"-o", outputTemplate,
		"--print", "after_move:filepath",
	}

	// Añadir ruta de ffmpeg (necesario para conversion de audio)
	if a.ffmpegPath != "" {
		args = append(args, "--ffmpeg-location", a.ffmpegPath)
	}

	switch req.Format {
	case domain.MusicFormatMP3:
		args = append(args, "-x", "--audio-format", "mp3", "--audio-quality", "0")
	case domain.MusicFormatMP4:
		args = append(
			args,
			"-f",
			"bestaudio[ext=m4a]/bestaudio",
			"--extract-audio",
			"--audio-format",
			"m4a")
	default:
		return nil, domain.ErrInvalidFormat
	}

	args = append(args, req.URL)

	filePath, err := a.run(ctx, args)
	if err != nil {
		return nil, err
	}

	return &domain.DownloadResult{
		FileName: filepath.Base(filePath),
		FilePath: filePath,
		Format:   req.Format,
		Type:     string(domain.DownloadTypeMusic),
	}, nil
}

// run ejecuta yt-dlp con los argumentos dados y devuelve la ruta del archivo.
func (a *Adapter) run(ctx context.Context, args []string) (string, error) {
	a.logger.Debug("ejecutando yt-dlp", "args", args)

	cmd := exec.CommandContext(ctx, a.ytdlpPath, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		a.logger.Error("yt-dlp fallo",
			"error", err,
			"stderr", stderr.String(),
			"args", args,
		)
		return "", fmt.Errorf("%w: %s", domain.ErrDownloadFailed, strings.TrimSpace(stderr.String()))
	}

	// IMPORTANTE: yt-dlp imprime varias lineas (una por formato descargado).
	// Tomamos SOLO la ULTIMA linea no vacia, que es la ruta final del archivo mergeado.
	raw := strings.TrimSpace(stdout.String())
	if raw == "" {
		a.logger.Error("yt-dlp no devolvio ruta",
			"stdout", stdout.String(),
			"stderr", stderr.String(),
		)
		return "", fmt.Errorf("%w: no se obtuvo la ruta del archivo", domain.ErrDownloadFailed)
	}

	lines := strings.Split(raw, "\n")
	var filePath string
	for i := len(lines) - 1; i >= 0; i-- {
		line := strings.TrimSpace(lines[i])
		if line != "" {
			filePath = line
			break
		}
	}

	if filePath == "" {
		return "", fmt.Errorf("%w: no se obtuvo la ruta del archivo", domain.ErrDownloadFailed)
	}

	a.logger.Info("descarga completada", "file", filePath)
	return filePath, nil
}
