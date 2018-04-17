package main

import (
	"net/http"
	"log"
	"github.com/RXD-Chinasoft/mygolang/handlers"
	"github.com/RXD-Chinasoft/mygolang/version"
	"os"
	"os/signal"
	"syscall"
	"context"
)

func main() {
	log.Printf(
		"Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)
	port := os.Getenv("PORT")
	router := handlers.Router()

	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server {
		Addr: ":"+port,
		Handler: router,
	}

	shutdown := make(chan struct{}, 1)
	go func(){
		err := srv.ListenAndServe()
		if (err != nil)  {
			shutdown <- struct{}{}
			log.Printf("%v", err)
		}
	}()

	log.Print("The service is ready to listen and serve.")
	// log.Fatal(http.ListenAndServe(":"+port, router))
	select {
	case killSignal := <- interupt:
		switch killSignal {
		case os.Interrupt:
			log.Print("Got SIGINT...")
		case syscall.SIGTERM:
			log.Print("Got SIGTERM...")
		}
	case <- shutdown:
		log.Printf("Got an error...")
	}
	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Done")
}