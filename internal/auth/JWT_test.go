package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeAndValidateJWT(t *testing.T) {
	tokenSecret := "supersecretkey"
	userID := uuid.New()
	expiresIn := time.Minute * 15

	t.Run("MakeJWT creates a valid token", func(t *testing.T) {
		token, err := MakeJWT(userID, tokenSecret, expiresIn)
		if err != nil {
			t.Fatalf("MakeJWT returned an error: %v", err)
		}
		if token == "" {
			t.Fatalf("MakeJWT returned an empty token")
		}
	})

	t.Run("ValidateJWT returns the correct user ID for a valid token", func(t *testing.T) {
		token, err := MakeJWT(userID, tokenSecret, expiresIn)
		if err != nil {
			t.Fatalf("MakeJWT returned an error: %v", err)
		}

		returnedUserID, err := ValidateJWT(token, tokenSecret)
		if err != nil {
			t.Fatalf("ValidateJWT returned an error: %v", err)
		}

		if returnedUserID != userID {
			t.Fatalf("Expected userID %v, got %v", userID, returnedUserID)
		}
	})

	t.Run("ValidateJWT returns an error for an invalid token", func(t *testing.T) {
		invalidToken := "invalid.token.string"

		_, err := ValidateJWT(invalidToken, tokenSecret)
		if err == nil {
			t.Fatalf("ValidateJWT did not return an error for an invalid token")
		}
	})

	t.Run("ValidateJWT returns an error for an expired token", func(t *testing.T) {
		expiredToken, err := MakeJWT(userID, tokenSecret, -time.Minute)
		if err != nil {
			t.Fatalf("MakeJWT returned an error: %v", err)
		}

		_, err = ValidateJWT(expiredToken, tokenSecret)
		if err == nil {
			t.Fatalf("ValidateJWT did not return an error for an expired token")
		}
	})

	t.Run("ValidateJWT returns an error for a token signed with a different secret", func(t *testing.T) {
		token, err := MakeJWT(userID, tokenSecret, expiresIn)
		if err != nil {
			t.Fatalf("MakeJWT returned an error: %v", err)
		}

		_, err = ValidateJWT(token, "wrongsecretkey")
		if err == nil {
			t.Fatalf("ValidateJWT did not return an error for a token signed with a different secret key")
		}
	})
}
