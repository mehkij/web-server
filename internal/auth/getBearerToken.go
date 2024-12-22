package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	token := headers.Get("Authorization")
	if token == "" {
		return "", fmt.Errorf("unauthorized: missing token")
	}

	split := strings.Split(token, " ")
	if len(split) != 2 || split[0] != "Bearer" {
		return "", fmt.Errorf("unauthorized: malformed token")
	}

	return split[1], nil
}
