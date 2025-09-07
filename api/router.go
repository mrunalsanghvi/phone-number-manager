package api

import (
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/phone-numbers", CreatePhoneNumberHandler)
	return mux
}