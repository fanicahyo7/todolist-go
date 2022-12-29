package util

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// JWTAuthMiddleware checks if the request has a valid JWT token in the authorization header
func JWTAuthMiddleware(c *fiber.Ctx) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{"Status": "error", "message": "Unauthorized"})
		return
	}

	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{"Status": "error", "message": "Unauthorized"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.Locals("userID", claims["userID"])
		c.Next()
	} else {
		c.Status(http.StatusUnauthorized).JSON(fiber.Map{"Status": "error", "message": "Unauthorized"})
		return
	}
}
