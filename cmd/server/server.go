// Package server Logic for handling connections
package server

import (
	"log"
	"net/http"

	routes "github.com/BatmiBoom/http_server_go/cmd"
)

// Start the server on port 8080
func Start(routes [1]routes.Route) {
	mux := http.NewServeMux()

	setupRoutes(mux, routes)

	log.Print("Listening on port 8080...:")

	http.ListenAndServe(":8080", mux)
}

func setupRoutes(server *http.ServeMux, routes [1]routes.Route) {
	for _, r := range routes {
		server.Handle(r.Route, http.FileServer(http.Dir("../../template")))
	}
}
