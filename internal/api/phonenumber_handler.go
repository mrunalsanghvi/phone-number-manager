package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"phone-number-manager/internal/errors"
	"phone-number-manager/internal/models"
	"phone-number-manager/internal/service"
	"phone-number-manager/internal/validation"
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
	fmt.Println("Received phone number:", req.PhoneNumber)
	if err := validation.ValidateStruct(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		errors.WriteAPIError(w, errors.New(http.StatusBadRequest, "Invalid request", err.Error()))
		return
	}
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
