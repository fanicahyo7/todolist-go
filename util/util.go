package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// func ValidateToken(tokenString string) (jwt.MapClaims, error) {
// 	secret := []byte(os.Getenv("JWT_SECRET"))

// 	// Parse token
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Validate signing algorithm
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return secret, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Validate token claims
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		if claims.VerifyExpiresAt(time.Now().Unix(), true) {
// 			return nil, fmt.Errorf("token expired")
// 		}
// 		return claims, nil
// 	}
// 	return nil, fmt.Errorf("invalid token")
// }

func CreateJWT(id int) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(secret)
}

// HashPassword hashes the given password and returns the hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash checks if the given password and hash match
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
