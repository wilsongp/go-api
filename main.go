package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"mongoapi/shortcut"
)

func main() {
	router := shortcut.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// launch server
	log.Fatal(http.ListenAndServe(":9000",
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
