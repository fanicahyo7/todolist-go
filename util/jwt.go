package util

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JWTAuthMiddleware(secret string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Get token from request header
		tokenString := c.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Initialize a new instance of `jwt.Parser`
		parser := jwt.Parser{ValidMethods: []string{jwt.SigningMethodHS256.Name}}

		// Parse the token
		token, err := parser.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		// If token is invalid or expired
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// set user_id to context locals
			c.Locals("id", claims["id"])
			return nil
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
	}
}
