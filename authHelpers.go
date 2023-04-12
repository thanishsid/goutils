package goutils

import (
	"errors"
	"net/http"
	"strings"
)

func GetAuthHeader(r *http.Request) string {
	var authorizationHeader string

	for _, key := range []string{"authorization", "x-authorization"} {
		authorizationHeader = r.Header.Get(key)
		if authorizationHeader != "" {
			break
		}
	}

	return authorizationHeader
}

func GetBearerToken(r *http.Request) (string, error) {
	tokenString := GetAuthHeader(r)

	tokenParts := strings.Split(tokenString, " ")

	if len(tokenParts) < 2 {
		return "", errors.New("invalid bearer token")
	}

	return tokenParts[1], nil
}
