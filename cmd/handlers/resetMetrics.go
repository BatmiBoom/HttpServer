package handlers

import (
	"net/http"

	"github.com/BatmiBoom/http_server_go/cmd/middleware"
)

func ResetMetrics(cfg middleware.ApiConfig) func(http.ResponseWriter, *http.Request) {
	cfg.FileServerHits = 0

	return func(w http.ResponseWriter, r *http.Request) {
		headers := w.Header()
		headers.Add("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("OK"))
		w.WriteHeader(http.StatusOK)
	}
}
