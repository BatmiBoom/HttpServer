package handlers

import (
	"net/http"
)

func (s *Server) CountRequest(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.FileServerHits++
		next.ServeHTTP(w, r)
	})
}
