package api

import (
	"encoding/json"
	"net/http"
	"phone-number-manager/validation"
	"phone-number-manager/errors"
	"go.uber.org/zap"
	"phone-number-manager/logging"
)

// PhoneNumber example model
type PhoneNumber struct {
	Number string `json:"number" validate:"required,e164"`
	Label  string `json:"label" validate:"required"`
}

// CreatePhoneNumberHandler demonstrates validation, error, logging, tracing
func CreatePhoneNumberHandler(w http.ResponseWriter, r *http.Request) {
	var req PhoneNumber
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logging.L().Error("Invalid JSON", zap.Error(err))
		error.WriteAPIError(w, errors.New(400, "Invalid JSON", err.Error()))
		return
	}
	if err := validation.ValidateStruct(&req); err != nil {
		logging.L().Error("Validation failed", zap.Error(err))
		fieldErrors := errors.FormatValidationError(err)
		resp := struct {
			Error       errors.APIError     `json:"error"`
			FieldErrors []errors.FieldError `json:"fields,omitempty"`
		}{
			Error:       errors.New(422, "Validation failed", "One or more fields are invalid"),
			FieldErrors: fieldErrors,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// Success response stub
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(req)
}