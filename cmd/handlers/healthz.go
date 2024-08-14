package handlers

import "net/http"

func Healthz() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		headers := w.Header()
		headers.Add("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("OK"))
		w.WriteHeader(http.StatusOK)
	}
}
