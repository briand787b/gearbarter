package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/briand787b/gearbarter/models"
	"github.com/gorilla/mux"
	"net/http"
)

func registerUsers(r *mux.Router) {
	// create subrouter in cases where all paths in the controller
	// should have a common behavior
	users := r.PathPrefix("/users").Subrouter()

	users.HandleFunc("/", HandleUserCreate).Methods("POST")
}

// HandleUserCreate handles the creation of a user
func HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	// this is going to simply create a user out of nothing
	// to validate that the backend is working properly

	u := &models.User{
		Username: "brian3",
	}

	if err := u.Create(r.Context()); err != nil {
		http.Error(w, fmt.Sprintf("error: %s", err), http.StatusInternalServerError)
	}

	ret, err := json.Marshal(u)
	if err != nil {
		http.Error(w, fmt.Sprintf("could not marshal json: %s", err), 500)
	}

	w.Write(ret)
}
