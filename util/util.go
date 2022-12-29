package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	secret := []byte("secret")

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	// Validate token claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return nil, fmt.Errorf("token expired")
		}
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
