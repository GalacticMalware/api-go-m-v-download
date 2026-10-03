// Package logger provee un logger centralizado basado en log/slog (nativo de Go).
// Escribe simultaneamente a consola y a un archivo de log.
package logger

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

// New crea un logger que escribe a stdout y al archivo indicado.
func New(level, logFile string) *slog.Logger {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	// Asegurar que exista el directorio de logs
	if err := os.MkdirAll(filepath.Dir(logFile), 0o755); err != nil {
		// Si no se puede crear, solo consola
		logFile = ""
	}

	var writers []io.Writer
	writers = append(writers, os.Stdout)

	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err == nil {
			writers = append(writers, file)
		}
	}

	multi := io.MultiWriter(writers...)

	handler := slog.NewJSONHandler(multi, &slog.HandlerOptions{
		Level: logLevel,
	})

	return slog.New(handler)
}
