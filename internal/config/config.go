// Package config gestiona la configuracion de la aplicacion
// cargada desde variables de entorno.
package config

import (
	"os"
)

// Config contiene toda la configuracion de la aplicacion.
type Config struct {
	Env         string // development | production
	Port        string // puerto del servidor
	LogLevel    string // debug | info | warn | error
	LogFile     string // ruta del archivo de logs
	DownloadDir string // carpeta donde se guardan las descargas
	YTDLPPath   string // ruta al binario yt-dlp
	FfmpegPath  string // ruta al binario ffmpeg
}

// Load lee las variables de entorno y devuelve una Config con valores por defecto.
func Load() *Config {
	return &Config{
		Env:         getEnv("APP_ENV", "development"),
		Port:        getEnv("APP_PORT", "8080"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		LogFile:     getEnv("LOG_FILE", "logs/app.log"),
		DownloadDir: getEnv("DOWNLOAD_DIR", "downloads"),
		YTDLPPath:   getEnv("YTDLP_PATH", "C:\\Users\\Daniel Alejandro\\AppData\\Local\\Microsoft\\WinGet\\Packages\\yt-dlp.yt-dlp_Microsoft.Winget.Source_8wekyb3d8bbwe\\yt-dlp.exe"),
		FfmpegPath:  getEnv("FFMPEG_PATH", "C:\\Users\\Daniel Alejandro\\AppData\\Local\\Microsoft\\WinGet\\Packages\\yt-dlp.FFmpeg_Microsoft.Winget.Source_8wekyb3d8bbwe\\ffmpeg-N-126374-g089a48eb36-win64-gpl\\bin\\ffmpeg.exe"),
	}
}

// getEnv devuelve el valor de una variable de entorno o un valor por defecto.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
