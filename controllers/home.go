package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func registerHome(r *mux.Router) {
	r.HandleFunc("/", handleHome).Methods("GET")
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home")
}
