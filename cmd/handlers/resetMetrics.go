package handlers

import (
	"net/http"
)

func (s *Server) ResetMetrics(w http.ResponseWriter, r *http.Request) {
	s.FileServerHits = 0
	headers := w.Header()
	headers.Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
