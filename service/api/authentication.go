package api

import (
	"net/http"
	"strconv"
	"strings"
)

func getToken(authorization string) uint64 {
	tokens := strings.Split(authorization, " ")
	if len(tokens) != 2 || tokens[0] != "Bearer" {
		return http.StatusUnauthorized
	}

	token := strings.TrimSpace(tokens[1])
	if token == "" {
		return http.StatusForbidden
	}

	authID, err := strconv.ParseUint(token, 10, 64)
	if err != nil {
		return http.StatusInternalServerError
	}

	return authID
}

func checkAuthorization(authorization string, id uint64) int {
	tokens := strings.Split(authorization, " ")
	if len(tokens) != 2 || tokens[0] != "Bearer" {
		return http.StatusUnauthorized
	}

	token := strings.TrimSpace(tokens[1])
	if token == "" {
		return http.StatusForbidden
	}

	authID, err := strconv.ParseUint(token, 10, 64)
	if err != nil {
		return http.StatusInternalServerError
	}
	if id != authID {
		return http.StatusUnauthorized
	}

	return 0
}
