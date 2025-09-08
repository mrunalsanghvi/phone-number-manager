package api

import "github.com/gorilla/mux"

type Router struct {
	PhoneBookHandler *PhoneBookHandler
}

func NewRouter(phoneBookHandler *PhoneBookHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/v1/phone-numbers", phoneBookHandler.CreatePhoneNumberHandler).Methods("POST")
	return r
}
