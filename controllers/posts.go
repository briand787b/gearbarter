package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func registerPosts(r *mux.Router) {
	r.HandleFunc("/posts", getPosts).Methods("GET")
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "getting posts")
}
