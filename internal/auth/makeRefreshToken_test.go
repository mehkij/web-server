package auth

import "testing"

func TestMakeRefreshToken(t *testing.T) {
	t.Run("MakeRefreshToken creates a valid token with length of 250 characters", func(t *testing.T) {
		token, err := MakeRefreshToken()
		if err != nil {
			t.Fatalf("MakeRefreshToken returned an error: %v", err)
		}

		if len(token) != 64 {
			t.Fatalf("Expected token of length 64, got token of length %v", len(token))
		}
	})
}
