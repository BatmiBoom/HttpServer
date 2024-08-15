package routes

import (
	"net/http"

	"github.com/BatmiBoom/http_server_go/cmd/handlers"
)

type Route struct {
	Route   string
	Handler http.HandlerFunc
}

func GetRoutes() [4]Route {
	server := handlers.Server{
		FileServerHits: 0,
	}

	template_path := http.Dir("./template")
	root_route := Route{
		Route:   "/app/*",
		Handler: server.CountRequest(http.StripPrefix("/app", http.FileServer(template_path)).ServeHTTP),
	}

	readiness_route := Route{
		Route:   "/healthz/",
		Handler: server.CountRequest(server.Healthz),
	}

	metrics_route := Route{
		Route:   "/metrics/",
		Handler: server.Metrics,
	}

	reset_metrics_route := Route{
		Route:   "/reset/",
		Handler: server.ResetMetrics,
	}

	routes := [4]Route{
		root_route,
		readiness_route,
		metrics_route,
		reset_metrics_route,
	}

	return routes
}
