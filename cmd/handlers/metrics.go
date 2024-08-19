package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func (s *Server) Metrics(w http.ResponseWriter, r *http.Request) {
	log.Print(s.FileServerHits)
	headers := w.Header()
	headers.Add("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("<html><body><h1>Welcome, Chirpy Admin</h1><p>Chirpy has been visited %d times!</p></body></html>", s.FileServerHits)))
}
