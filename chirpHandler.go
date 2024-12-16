package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mehkij/web-server/internal/database"
)

type Chirp struct {
	ID        uuid.UUID     `json:"id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Body      string        `json:"body"`
	UserID    uuid.NullUUID `json:"user_id"`
}

func (cfg *apiConfig) createChirpHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body   string        `json:"body"`
		UserID uuid.NullUUID `json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding params: %s", err)
		w.WriteHeader(500)
		return
	}

	validatedChirp, code := validateChirp(params.Body)
	if code == 400 {
		respondWithError(w, code, "Chirp too long")
		return
	}

	params.Body = validatedChirp

	chirp, err := cfg.queries.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   params.Body,
		UserID: params.UserID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating chirp: %s", err))
		return
	}

	respondWithJSON(w, 201, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}

func (cfg *apiConfig) getChirpsHandler(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.queries.GetChirps(r.Context())
	if err != nil {
		log.Printf("Error decoding params: %s", err)
		w.WriteHeader(500)
		return
	}

	var payload []Chirp

	for _, chirp := range chirps {
		payload = append(payload, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}

	respondWithJSON(w, 200, payload)
}

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

func validateChirp(body string) (string, int) {
	if len(body) > 140 {
		return "", 400
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
		CleanedBody: filterMessage(profanity, body),
	}

	return respBody.CleanedBody, 200
}
