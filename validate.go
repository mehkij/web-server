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

	type returnVals struct {
		Error       string `json:"error"`
		CleanedBody string `json:"cleaned_body"`
	}

	if len(params.Body) > 140 {
		respBody := returnVals{
			Error: "Chirp is too long",
		}

		data, err := json.Marshal(respBody)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(data)
		return
	}

	// If length OK, censor disallowed words
	profanity := []string{
		"kerfuffle",
		"sharbert",
		"fornax",
	}

	split := strings.Split(params.Body, " ")

	for i, word := range split {
		if Contains(profanity, strings.ToLower(word)) {
			split[i] = strings.Repeat("*", len(word))
		}
	}

	respBody := returnVals{
		CleanedBody: strings.Join(split, " "),
	}

	data, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}
