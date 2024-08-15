package handlers

import (
	"log"
	"fmt"
	"net/http"
)

func (s *Server) Metrics(w http.ResponseWriter, r *http.Request) {
		log.Print(s.FileServerHits)
		headers := w.Header()
		headers.Add("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(fmt.Sprintf("Hits: %d", s.FileServerHits)))
}
