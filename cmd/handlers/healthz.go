package handlers

import (
	"net/http"
)

func (s *Server) Healthz(w http.ResponseWriter, r *http.Request) {
		headers := w.Header()
		headers.Add("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("OK"))
}
