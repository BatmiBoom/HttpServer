package routes

import (
	"net/http"

	"github.com/BatmiBoom/http_server_go/cmd/handlers"
	"github.com/BatmiBoom/http_server_go/cmd/middleware"
)

type Route struct {
	Route   string
	Handler http.HandlerFunc
}

func GetRoutes() [4]Route {
	apiConfig := middleware.ApiConfig{
		FileServerHits: 0,
	}

	template_path := http.Dir("./template")
	root_route := Route{
		Route:   "/app/*",
		Handler: apiConfig.Metrics(http.StripPrefix("/app", http.FileServer(template_path)).ServeHTTP),
	}

	readiness_route := Route{
		Route:   "/healthz/",
		Handler: apiConfig.Metrics(handlers.Healthz()),
	}

	metrics_route := Route{
		Route:   "/metrics/",
		Handler: handlers.Metrics(apiConfig),
	}

	reset_metrics_route := Route{
		Route:   "/reset/",
		Handler: handlers.ResetMetrics(apiConfig),
	}

	routes := [4]Route{
		root_route,
		readiness_route,
		metrics_route,
		reset_metrics_route,
	}

	return routes
}
