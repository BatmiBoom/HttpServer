// Entry Point
package main

import (
	routes "github.com/BatmiBoom/http_server_go/cmd"
	"github.com/BatmiBoom/http_server_go/cmd/server"
)

func main() {
	server.Start(routes.GetRoutes())
}
