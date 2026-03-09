package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Data    any         `json:"data,omitempty"`
	Error   *ErrorData  `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ErrorData struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]any         `json:"details,omitempty"`
}

func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		Success: statusCode >= 200 && statusCode < 300,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

func Success(w http.ResponseWriter, data any, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := Response{
		Success: true,
		Data:    data,
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}

func Created(w http.ResponseWriter, data any, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := Response{
		Success: true,
		Data:    data,
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}

func Error(w http.ResponseWriter, statusCode int, code, message string, details map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		Success: false,
		Error: &ErrorData{
			Code:    code,
			Message: message,
			Details: details,
		},
	}

	json.NewEncoder(w).Encode(response)
}

func BadRequest(w http.ResponseWriter, message string, details map[string]any) {
	Error(w, http.StatusBadRequest, "BAD_REQUEST", message, details)
}

func NotFound(w http.ResponseWriter, message string) {
	Error(w, http.StatusNotFound, "NOT_FOUND", message, nil)
}

func InternalServerError(w http.ResponseWriter, message string) {
	Error(w, http.StatusInternalServerError, "INTERNAL_ERROR", message, nil)
}

func Unauthorized(w http.ResponseWriter, message string) {
	Error(w, http.StatusUnauthorized, "UNAUTHORIZED", message, nil)
}

func Forbidden(w http.ResponseWriter, message string) {
	Error(w, http.StatusForbidden, "FORBIDDEN", message, nil)
}
