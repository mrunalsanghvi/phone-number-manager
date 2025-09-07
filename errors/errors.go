package errors

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

// WriteAPIError writes a standardized error response.
func WriteAPIError(w http.ResponseWriter, err APIError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(err)
}

// Helper to create APIError
func New(code int, msg, detail string) APIError {
	return APIError{
		Code:    code,
		Message: msg,
		Detail:  detail,
	}
}