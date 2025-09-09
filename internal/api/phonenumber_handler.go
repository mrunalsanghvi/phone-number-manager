package api

import (
	"encoding/json"
	"net/http"
	"phone-number-manager/internal/errors"
	logger "phone-number-manager/internal/logging"
	"phone-number-manager/internal/models"
	"phone-number-manager/internal/service"
	"phone-number-manager/internal/validation"

	"go.uber.org/zap"
)

type PhoneBookHandler struct {
	PhoneBookService *service.PhoneBookService
}

func NewHandler(phoneBookService *service.PhoneBookService) *PhoneBookHandler {
	return &PhoneBookHandler{
		PhoneBookService: phoneBookService,
	}
}

// CreatePhoneNumberHandler demonstrates validation, error, logging, tracing
func (h *PhoneBookHandler) CreatePhoneNumberHandler(w http.ResponseWriter, r *http.Request) {
	req := models.PhoneBook{
		PhoneNumber: r.URL.Query().Get("phoneNumber"),
	}
	logger.Log.Info("Received phone number:", zap.String("phoneNumber", req.PhoneNumber))
	if err := validation.ValidateStruct(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		errors.WriteAPIError(w, errors.New(http.StatusBadRequest, "Invalid request", err.Error()))
		return
	}
	
	req.PhoneNumber = addPlusIfNotAlready(req.PhoneNumber)

	num, err := validation.ValidateE164Phone(req.PhoneNumber)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		errors.WriteAPIError(w, errors.New(http.StatusBadRequest, "Invalid phone number", err.Error()))
		return
	}
	if err := h.PhoneBookService.CreatePhoneBookEntry(r.Context(), &req, num); err != nil {
		w.Header().Set("Content-Type", "application/json")
		errors.WriteAPIError(w, errors.New(http.StatusInternalServerError, "Failed to create phone book entry", err.Error()))
		return
	}
	// Success response stub
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(req)
}

func addPlusIfNotAlready(phoneNumber string) string {
	if len(phoneNumber) > 0 && phoneNumber[0] != '+' {
		return "+" + phoneNumber
	}
	return phoneNumber
}