// Package main es el punto de entrada de la aplicacion.
// Se encarga unicamente de inicializar dependencias y arrancar el servidor.
package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"backend-go-download-music-video/internal/adapter/ytdlp"
	"backend-go-download-music-video/internal/config"
	"backend-go-download-music-video/internal/handler"
	"backend-go-download-music-video/internal/middleware"
	"backend-go-download-music-video/internal/router"
	"backend-go-download-music-video/internal/usecase"
	"backend-go-download-music-video/pkg/logger"
)

func main() {
	// 1. Cargar configuracion
	cfg := config.Load()

	// Convertir DownloadDir a ruta absoluta
	if abs, err := filepath.Abs(cfg.DownloadDir); err == nil {
		cfg.DownloadDir = abs
	}

	if err := os.MkdirAll(cfg.DownloadDir, 0o755); err != nil {
		panic("no se pudo crear la carpeta de descargas: " + err.Error())
	}

	// 2. Inicializar logger centralizado
	log := logger.New(cfg.LogLevel, cfg.LogFile)
	log.Info("iniciando aplicacion", "env", cfg.Env, "port", cfg.Port, "log_level", cfg.LogLevel, "log_file", cfg.LogFile, "download_dir", cfg.DownloadDir, "yt-dlp_path", cfg.YTDLPPath, "ffmpeg_path", cfg.FfmpegPath)

	// 3. Inicializar adaptador (infraestructura)
	ytdlpAdapter := ytdlp.NewAdapter(cfg.DownloadDir, cfg.YTDLPPath, cfg.FfmpegPath, log)

	// 4. Inicializar casos de uso con sus dependencias
	videoUC := usecase.NewVideoUseCase(ytdlpAdapter, log)
	musicUC := usecase.NewMusicUseCase(ytdlpAdapter, log)

	// 5. Inicializar handlers
	videoHandler := handler.NewVideoHandler(videoUC, log)
	musicHandler := handler.NewMusicHandler(musicUC, log)

	// 6. Configurar router
	mux := router.NewRouter(videoHandler, musicHandler)

	// 7. Envolver con middleware
	handlerChain := middleware.Logging(log)(middleware.Recover(log)(mux))

	// 8. Servidor HTTP
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handlerChain,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 10 * time.Minute, // descargas pueden tardar
		IdleTimeout:  60 * time.Second,
	}

	// 9. Arranque en goroutine para permitir shutdown graceful
	go func() {
		log.Info("servidor escuchando", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("error en el servidor", "error", err)
			os.Exit(1)
		}
	}()

	// 10. Esperar señal de apagado
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("apagando servidor...")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("error en shutdown", "error", err)
	}
	log.Info("servidor detenido correctamente")
}
