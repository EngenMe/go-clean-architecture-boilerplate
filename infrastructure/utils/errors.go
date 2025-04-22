package utils

import (
	"errors"
	"net/http"
)

// Common error types
var (
	ErrNotFound     = errors.New("resource not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrBadRequest   = errors.New("bad request")
	ErrConflict     = errors.New("resource already exists")
)

// APIError represents an API error response
type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// ErrorToStatusCode maps error types to HTTP status codes
func ErrorToStatusCode(err error) int {
	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrUnauthorized):
		return http.StatusUnauthorized
	case errors.Is(err, ErrBadRequest):
		return http.StatusBadRequest
	case errors.Is(err, ErrConflict):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

// NewAPIError creates a new API error
func NewAPIError(status int, message string) APIError {
	return APIError{
		Status:  status,
		Message: message,
	}
}
