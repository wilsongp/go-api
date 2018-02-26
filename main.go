package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/wilsongp/go-api/graphql"
	"github.com/wilsongp/go-api/routing"

	"github.com/gorilla/handlers"
)

var addr = "localhost:8080"

func main() {
	routes := append(graphql.Routes)
	router := routing.NewRouter(routes)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	server := &http.Server{
		Addr:         addr,
		Handler:      handlers.CORS(allowedOrigins, allowedMethods)(router),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Launching server at %s\n\n", addr)

	// launch server
	log.Fatal(server.ListenAndServe())
}
