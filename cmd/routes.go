package routes

import "net/http"

type Route struct {
	Route   string
	Handler http.Handler
}

func GetRoutes() [2]Route {
	template_path := http.Dir("./template")
	root_route := Route{
		Route:   "/",
		Handler: http.FileServer(template_path),
	}

	assets_route := Route{
		Route:   "assets/",
		Handler: http.FileServer(template_path),
	}

	routes := [2]Route{
		root_route,
		assets_route,
	}

	return routes
}
