package handlers

import (
	"fmt"
	"net/http"

	"github.com/BatmiBoom/http_server_go/cmd/middleware"
)

func Metrics(cfg middleware.ApiConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		headers := w.Header()
		headers.Add("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.FileServerHits)))
		w.WriteHeader(http.StatusOK)
	}
}
