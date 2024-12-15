package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) resetHitsHandler(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, 403, "Forbidden")
		return
	}

	err := cfg.queries.ResetUsers(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error resetting users: %s", err))
		return
	}

	cfg.fileserverHits.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hits reset to 0"))
}
