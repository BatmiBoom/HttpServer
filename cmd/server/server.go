// Package server Logic for handling connections
package server

import (
	"log"
	"net/http"

	routes "github.com/BatmiBoom/http_server_go/cmd"
)

// Start the server on port 8080
func Start(routes []routes.Route) {

	mux := http.NewServeMux()

	setupRoutes(mux, routes)

	log.Print("Listening on port 8080...:")

	http.ListenAndServe(":8080", mux)
}

func setupRoutes(server *http.ServeMux, routes []routes.Route) {
	for i := 0; i < 1; i++ {
		log.Print(http.Dir("."))
		server.Handle("/", http.FileServer(http.Dir("/template")))
	}
}
