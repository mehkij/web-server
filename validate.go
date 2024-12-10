package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func filterMessage(filter []string, msg string) string {
	split := strings.Split(msg, " ")

	for i, word := range split {
		if Contains(filter, strings.ToLower(word)) {
			split[i] = strings.Repeat("*", 4)
		}
	}

	return strings.Join(split, " ")
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding params: %s", err)
		w.WriteHeader(500)
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}

	type cleaned struct {
		CleanedBody string `json:"cleaned_body"`
	}

	// If length OK, censor disallowed words
	profanity := []string{
		"kerfuffle",
		"sharbert",
		"fornax",
	}

	respBody := cleaned{
		CleanedBody: filterMessage(profanity, params.Body),
	}

	respondWithJSON(w, 200, respBody)
}
