// Package server Logic for handling connections
package server

import (
	"log"
	"net/http"

	routes "github.com/BatmiBoom/http_server_go/cmd"
)

// Start the server on port 8080
func Start(routes routes.Routes) {
	const port = "8080"

	mux := http.NewServeMux()
	setupRoutes(mux, routes)

	http_server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Listening on port : %s", port)
	log.Fatal(http_server.ListenAndServe())
}

func setupRoutes(server *http.ServeMux, routes routes.Routes) {
	for _, route := range routes {
		server.HandleFunc(route.Methods+" "+route.Route, route.Handler)
	}
}
