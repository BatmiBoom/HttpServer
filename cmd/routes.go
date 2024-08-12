package routes

import "net/http"

type Route struct {
	Route   string
	Handler http.Handler
}

func GetRoutes() []Route {
	routes := make([]Route, 1)

	template_path := http.Dir("../template")
	root_route := Route{
		Route:   "/",
		Handler: http.FileServer(template_path),
	}

	routes = append(routes, root_route)
	return routes
}
