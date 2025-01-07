package main

import (
	"log"
	"net/http"
	"time"

	"github.com/mehkij/web-server/internal/auth"
)

type AccessToken struct {
	Token string `json:"token"`
}

func (cfg *apiConfig) refreshHandler(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		log.Printf("Error getting refresh token: %s", err)
		w.WriteHeader(500)
		return
	}

	user, err := cfg.queries.GetUserFromRefreshToken(r.Context(), token)
	if err != nil {
		respondWithError(w, 401, "Refresh token not found")
		return
	}

	accessToken, err := auth.MakeJWT(user.UserID.UUID, cfg.jwtSecret, time.Duration(time.Hour))
	if err != nil {
		respondWithError(w, 400, "Error creating token")
	}

	respondWithJSON(w, 200, AccessToken{
		Token: accessToken,
	})
}
