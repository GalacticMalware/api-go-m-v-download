// Package handler contiene los manejadores HTTP.
package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// APIResponse es la estructura estandar de respuesta de la API.
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// writeJSON escribe una respuesta JSON con el codigo de estado dado.
func writeJSON(w http.ResponseWriter, status int, payload APIResponse) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		slog.Error("error al escribir respuesta JSON", "error", err)
	}
}

// respondError es un helper para respuestas de error.
func respondError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, APIResponse{
		Success: false,
		Error:   msg,
	})
}

// respondSuccess es un helper para respuestas exitosas.
func respondSuccess(w http.ResponseWriter, status int, msg string, data interface{}) {
	writeJSON(w, status, APIResponse{
		Success: true,
		Message: msg,
		Data:    data,
	})
}
