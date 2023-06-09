package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cast"
	"time"
)

func NewToken(secret string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   "auth",
		"exp":   time.Now().Add(time.Hour).Unix(),
		"email": email,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(secret string, tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString[6:], func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := claims["exp"].(float64)
		expTime := time.Unix(cast.ToInt64(exp), 0)
		if claims["exp"].(float64) != 0.0 && !expTime.After(time.Now()) {
			return map[string]interface{}{}, errors.New("token is invalid")
		}
		return claims, nil
	} else {
		return map[string]interface{}{}, err
	}
}

func GetClaimByTokenAndName(secret string, tokenString string, name string) (string, error) {
	claims, err := ParseToken(secret, tokenString)
	if err != nil {
		return "", err
	}

	if claims[name] != "" {
		return claims[name].(string), nil
	}
	return "", nil
}
