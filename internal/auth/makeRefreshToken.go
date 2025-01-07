package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func MakeRefreshToken() (string, error) {
	byte := make([]byte, 32)

	_, err := rand.Read(byte)
	if err != nil {
		return "", fmt.Errorf("error creating refresh token: %w", err)
	}

	hex := hex.EncodeToString(byte)

	return hex, nil
}
