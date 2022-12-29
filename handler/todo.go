package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"todolist/service"
	"todolist/util"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type todoHandler struct {
	todoService service.TodoListService
}

func NewGroupHandler(todoService service.TodoListService) *todoHandler {
	return &todoHandler{todoService}
}

func (h *todoHandler) GetTodos(c *fiber.Ctx) error {
	// Get JWT token from Authorization header
	authorizationHeader := c.Get("Authorization")
	if authorizationHeader == "" {
		return c.Status(http.StatusUnauthorized).JSON("Authorization header is required")

	}
	tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")
	if tokenString == "" {
		return c.Status(http.StatusUnauthorized).JSON("JWT token is required")

	}

	// Validate JWT token
	claims, err := util.ValidateToken(tokenString)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
	}

	// Save claims to context
	c.Locals("claims", claims)

	claims = c.Locals("claims").(jwt.MapClaims)
	userID, err := strconv.Atoi(fmt.Sprintf("%.0f", claims["ID"]))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})

	}
	todoLists, err := h.todoService.GetByUserID(userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})

	}
	return c.Status(200).JSON(fiber.Map{"Messgae": todoLists})
}
