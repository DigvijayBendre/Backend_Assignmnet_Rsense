package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HandlerRouting sets up API endpoint routing and starts the HTTP server.
func HandlerRouting() {
	// Create a new Gorilla router.
	r := mux.NewRouter()

	// Register the CreateEmployee handler function for the POST /employee endpoint.
	r.HandleFunc("/employee", CreateEmployee).Methods("POST")

	// Start the HTTP server on port 8181.
	log.Fatal(http.ListenAndServe(":8181", r))
}
