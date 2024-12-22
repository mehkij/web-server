package auth

import "testing"

func TestHashPassword(t *testing.T) {
	t.Run("HashPassword generates a non-empty hash", func(t *testing.T) {
		password := "securepassword"

		hash, err := HashPassword(password)
		if err != nil {
			t.Fatalf("HashPassword returned an error: %v", err)
		}

		if hash == "" {
			t.Fatalf("HashPassword returned an empty hash")
		}
	})

	t.Run("HashPassword generates unique hashes for the same password", func(t *testing.T) {
		password := "securepassword"
		hash1, err := HashPassword(password)
		if err != nil {
			t.Fatalf("HashPassword returned an error: %v", err)
		}

		hash2, err := HashPassword(password)
		if err != nil {
			t.Fatalf("HashPassword returned an error: %v", err)
		}

		if hash1 == hash2 {
			t.Fatalf("Expected unique hashes for the same password, but got identical hashes")
		}
	})
}

func TestCheckPasswordHash(t *testing.T) {
	t.Run("CheckPasswordHash returns nil for a matching password and hash", func(t *testing.T) {
		password := "securepassword"

		hash, err := HashPassword(password)
		if err != nil {
			t.Fatalf("HashPassword returned an error: %v", err)
		}

		err = CheckPasswordHash(password, hash)
		if err != nil {
			t.Fatalf("CheckPasswordHash returned an error for a matching password and hash")
		}
	})

	t.Run("CheckPasswordHash returns an error for a non-matching password and hash", func(t *testing.T) {
		password := "securepassword"

		hash, err := HashPassword(password)
		if err != nil {
			t.Fatalf("HashPassword returned an error: %v", err)
		}

		err = CheckPasswordHash("wrongpassword", hash)
		if err == nil {
			t.Fatalf("CheckHashPassword did not return an error for a non-matching password and hash")
		}
	})
}
