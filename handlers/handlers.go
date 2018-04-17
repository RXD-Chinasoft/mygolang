package handlers

import (
	"github.com/gorilla/mux"
	"time"
	"log"
	"github.com/RXD-Chinasoft/mygolang/version"
	"sync/atomic"
)

// Router register necessary routes and returns an instance of a router.
func Router() *mux.Router {
	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		log.Printf("Readyz probe is negative by default...")
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Printf("Readyz probe is positive.")
	}()
	r := mux.NewRouter()
	r.HandleFunc("/home", home(version.BuildTime, version.Commit, version.Release)).Methods("GET")
	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/readyz", readyz(isReady))
	return r
} 