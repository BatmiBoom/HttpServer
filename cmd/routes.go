package routes

import "net/http"

type Route struct {
	Route   string
	Handler http.Handler
}

func GetRoutes() [1]Route {
	template_path := http.Dir("../template")
	root_route := Route{
		Route:   "/",
		Handler: http.FileServer(template_path),
	}

	routes := [1]Route{
		root_route,
	}

	return routes
}
