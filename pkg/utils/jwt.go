package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

func NewToken(secret string, id uint64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   "auth",
		"exp":   time.Now().Add(time.Hour).Unix(),
		"id":    id,
		"email": email,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
