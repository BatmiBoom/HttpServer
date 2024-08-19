package routes

import (
	"net/http"
	"strings"

	"github.com/BatmiBoom/http_server_go/cmd/handlers"
)

const (
	GET    = "GET"
	POST   = "POST"
	PATCH  = "PATCH"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func selectMethods(methods ...string) string {
	return strings.Join(methods[:], " ")
}

type Route struct {
	Route   string
	Handler http.HandlerFunc
	Methods string
}

type Routes = [5]Route

func GetRoutes() Routes {
	server := handlers.Server{
		FileServerHits: 0,
	}

	template_path := http.Dir("./template")
	root_route := Route{
		Route:   "/app/*",
		Handler: server.CountRequest(http.StripPrefix("/app", http.FileServer(template_path)).ServeHTTP),
		Methods: selectMethods(GET),
	}

	readiness_route := Route{
		Route:   "/api/healthz/",
		Handler: server.CountRequest(server.Healthz),
		Methods: selectMethods(GET),
	}

	reset_metrics_route := Route{
		Route:   "/api/reset/",
		Handler: server.ResetMetrics,
		Methods: selectMethods(GET),
	}

	validate_chirp := Route{
		Route:   "/api/validate_chirp",
		Handler: server.ValidateChirp,
		Methods: selectMethods(POST),
	}

	metrics_route := Route{
		Route:   "/admin/metrics/",
		Handler: server.Metrics,
		Methods: selectMethods(GET),
	}

	routes := Routes{
		root_route,
		readiness_route,
		metrics_route,
		reset_metrics_route,
		validate_chirp,
	}

	return routes
}
