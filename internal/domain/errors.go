package domain

import "errors"

// Errores de dominio reutilizables.
var (
	ErrInvalidFormat  = errors.New("Formato incorrecto")
	ErrURLRequired    = errors.New("El campo 'url' es requerido")
	ErrFormatRequired = errors.New("El campo 'format' es requerido")
	ErrInvalidURL     = errors.New("La URL proporcionada no es valida")
	ErrDownloadFailed = errors.New("Fallo la descarga del recurso")
)
