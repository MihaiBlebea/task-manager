package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
)

var (
	ErrInvalidUserID = errors.New("Could not cast user_id claim as int")
	ErrNoSecretKey   = errors.New("Please provide a secret key")
	ErrNoBearerToken = errors.New("Expected auth token is missing")
)

func GetUserIDFromRequest(r *http.Request) (int, error) {
	token, err := extractBearerToken(r)
	if err != nil {
		return 0, err
	}

	return extractUserID(token)
}

func extractUserID(tkn string) (userID int, err error) {
	secretKey := os.Getenv("AUTH_SECRET")
	if secretKey == "" {
		return 0, ErrNoSecretKey
	}

	token, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
		if err != nil {
			return 0, err
		}

		return userID, nil
	}

	return 0, err
}

func extractBearerToken(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return "", ErrNoBearerToken
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	if token == "" {
		return "", ErrNoBearerToken
	}

	return token, nil
}
