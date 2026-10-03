// Package domain contiene las entidades y reglas de negocio puras.
// No debe depender de ninguna capa externa (ni HTTP, ni infraestructura).
package domain

// Tipo de descarga soportada.
type DownloadType string

const (
	DownloadTypeVideo DownloadType = "video"
	DownloadTypeMusic DownloadType = "music"
)

// Formatos validos de video.
const (
	VideoFormat1080p = "1080p"
	VideoFormat720p  = "720p"
)

// Formatos validos de musica.
const (
	MusicFormatMP3 = "mp3"
	MusicFormatMP4 = "mp4"
)

// VideoDownloadRequest representa una solicitud de descarga de video.
type VideoDownloadRequest struct {
	URL    string `json:"url"`
	Format string `json:"format"`
}

// MusicDownloadRequest representa una solicitud de descarga de musica.
type MusicDownloadRequest struct {
	URL    string `json:"url"`
	Format string `json:"format"`
}

// DownloadResult contiene el resultado de una descarga.
type DownloadResult struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	Format   string `json:"format"`
	Type     string `json:"type"`
}

// IsValidVideoFormat valida que el formato de video sea 720p o 1080p.
func IsValidVideoFormat(format string) bool {
	return format == VideoFormat1080p || format == VideoFormat720p
}

// IsValidMusicFormat valida que el formato de musica sea mp3 o mp4.
func IsValidMusicFormat(format string) bool {
	return format == MusicFormatMP3 || format == MusicFormatMP4
}
