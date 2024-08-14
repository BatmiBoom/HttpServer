// Package server Logic for handling connections
package server

import (
	"log"
	"net/http"

	routes "github.com/BatmiBoom/http_server_go/cmd"
)

// Start the server on port 8080
func Start(routes [2]routes.Route) {
	const port = "8080"

	mux := http.NewServeMux()
	setupRoutes(mux, routes)

	http_server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", "./template/", port)
	log.Fatal(http_server.ListenAndServe())
}

func setupRoutes(server *http.ServeMux, routes [2]routes.Route) {
	for _, route := range routes {
		log.Print(route)
		server.Handle(route.Route, route.Handler)
	}
}
