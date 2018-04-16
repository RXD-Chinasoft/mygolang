package handlers

import (
	"github.com/gorilla/mux"
	"github.com/RXD-Chinasoft/mygolang/version"
)

// Router register necessary routes and returns an instance of a router.
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", home(version.BuildTime, version.Commit, version.Release)).Methods("GET")
	return r
} 