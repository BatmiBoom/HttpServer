package routes

import "net/http"

type Route struct {
	Route   string
	Handler http.HandlerFunc
}

func GetRoutes() [3]Route {
	template_path := http.Dir("./template")
	root_route := Route{
		Route:   "/",
		Handler: http.FileServer(template_path).ServeHTTP,
	}

	assets_route := Route{
		Route:   "peto/",
		Handler: http.FileServer(template_path).ServeHTTP,
	}

	readiness_route := Route{
		Route: "healthz/",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			headers := w.Header()
			headers.Add("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte("OK"))
			w.WriteHeader(http.StatusOK)
		},
	}

	routes := [3]Route{
		root_route,
		assets_route,
		readiness_route,
	}

	return routes
}
