package handler

import (
	"net/http"
	"todolist/service"

	"github.com/gofiber/fiber/v2"
)

type todoHandler struct {
	todoService service.TodoListService
}

func NewGroupHandler(todoService service.TodoListService) *todoHandler {
	return &todoHandler{todoService}
}

func (h *todoHandler) GetTodos(c *fiber.Ctx) error {
	id, ok := c.Locals("id").(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "user_id not found"})
	}
	todoLists, err := h.todoService.GetByUserID(int(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"data": todoLists})
}
