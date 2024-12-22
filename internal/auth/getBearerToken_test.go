package auth

import (
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestGetBearerToken(t *testing.T) {
	userID := uuid.New()
	tokenSecret := "supersecretkey"
	expiresIn := time.Minute * 15

	expectedToken, err := MakeJWT(userID, tokenSecret, expiresIn)
	if err != nil {
		t.Fatalf("MakeJWT returned an error: %v", err)
	}

	t.Run("GetBearerToken returns the token when Authorization header is valid", func(t *testing.T) {
		headers := http.Header{}

		headers.Set("Authorization", "Bearer "+expectedToken)

		token, err := GetBearerToken(headers)
		if err != nil {
			t.Fatalf("GetBearerToken returned an error: %v", err)
		}

		if token != expectedToken {
			t.Fatalf("Expected token %q, got token %q", expectedToken, token)
		}
	})

	t.Run("GetBearerToken returns an error when Authorization header is malformed", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "Bearer")

		_, err := GetBearerToken(headers)
		if err == nil {
			t.Fatalf("Expected an error when Authorization header is malformed, got nil")
		}

		if !strings.Contains(err.Error(), "unauthorized: malformed token") {
			t.Fatalf("Expected error message to contain 'unauthorized: malformed token', got %q", err)
		}
	})

	t.Run("GetBearerToken returns an error when Authorization header is not Bearer", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "Basic somecredentials")

		_, err := GetBearerToken(headers)
		if err == nil {
			t.Fatalf("Expected an error when Authorization header is not Bearer, got nil")
		}
	})
}
