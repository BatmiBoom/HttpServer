// Package server Logic for handling connections
package server

import (
	"log"
	"net/http"
)

// Start the server on port 8080
func Start() {

	mux := http.NewServeMux()

	log.Print("Listening on port 8080...:")

	http.ListenAndServe(":8080", mux)
}
