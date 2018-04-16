package main

import (
	"net/http"
	"log"
	"github.com/RXD-Chinasoft/mygolang/handlers"
	"os"
)

func main() {
	log.Print("Starting the service....")
	port := os.Getenv("PORT")
	router := handlers.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, router))
}