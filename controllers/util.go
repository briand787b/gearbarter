package controllers

import (
	"github.com/briand787b/gearbarter/auth"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

// GetRouter is the only exported function in this package.
// It allows the main package to call a single function and
// then pass that to listen and serve
func GetRouter() *mux.Router {
	r := mux.NewRouter()

	// not authenticated yet
	registerAuths(r)

	authenticatedRouter := mux.NewRouter()
	r.PathPrefix("/").Handler(negroni.New(
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("../assets/")),
		negroni.NewRecovery(),
		negroni.HandlerFunc(auth.GetJWTMiddleware().HandlerWithNext),
		negroni.Wrap(authenticatedRouter),
	))

	// home := homeBase.PathPrefix("/").Subrouter()
	//home.HandleFunc("/", homeRoute) // root path

	registerUsers(authenticatedRouter)

	return r
}
