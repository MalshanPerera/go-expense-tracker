package handlers

import (
	"net/http"
)

func HandleFunc(handlerFunc func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerFunc(w, r)
	}
}
