package controllers

import (
	"fmt"
	"github.com/briand787b/gearbarter/auth"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func registerAuths(r *mux.Router) {
	authBase := mux.NewRouter()
	r.PathPrefix("/auth").Handler(negroni.New(
		negroni.NewLogger(),
		negroni.NewRecovery(),
		negroni.Wrap(authBase),
	))

	auth := authBase.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", login)
}

func login(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GenerateJWT()
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot get token: %s", err), 500)
		return
	}

	w.Header().Set("Authorization", fmt.Sprintf("Bearer: %s", token))
}
