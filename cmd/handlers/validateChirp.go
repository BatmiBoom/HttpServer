package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"
)

func (s *Server) ValidateChirp(w http.ResponseWriter, r *http.Request) {
	type payload struct {
		Body string `json:"body"`
	}

	type resp struct {
		CreatedAt   time.Time `json:"created_at"`
		Error       string    `json:"error"`
		CleanedBody string    `json:"cleaned_body"`
	}

	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	p := payload{}

	err := decoder.Decode(&p)
	if err != nil {
		res := resp{
			CreatedAt:   time.Now(),
			Error:       "Something went wrong",
			CleanedBody: "",
		}

		dat, err := json.Marshal(res)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(400)
		w.Write(dat)
		return
	}

	if len(p.Body) > 140 {
		res := resp{
			CreatedAt:   time.Now(),
			Error:       "Chirp is too long",
			CleanedBody: "",
		}

		dat, err := json.Marshal(res)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(400)
		w.Write(dat)
		return
	}

	banned_words := []string{
		"kerfuffle",
		"sharbert",
		"fornax",
	}

	old_body := strings.Split(p.Body, " ")
	cleaned_body := make([]string, len(old_body))
	for _, word := range old_body {
		if slices.Contains(banned_words, strings.ToLower(word)) {
			cleaned_body = append(cleaned_body, "****")
		} else {
			cleaned_body = append(cleaned_body, word)
		}
	}

	log.Printf("Cleaned Body %v", cleaned_body)
	res := resp{
		CreatedAt:   time.Now(),
		Error:       "",
		CleanedBody: strings.TrimLeft(strings.Join(cleaned_body, " "), " "),
	}

	dat, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	w.Write(dat)
}
