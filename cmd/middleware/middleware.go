package middleware

import "net/http"

type ApiConfig struct {
	FileServerHits int
}

func (cfg *ApiConfig) Metrics(next http.HandlerFunc) http.HandlerFunc {
	cfg.FileServerHits += 1

	return next.ServeHTTP
}
