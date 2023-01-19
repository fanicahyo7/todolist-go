package handler

import (
	"todolist/model"
	"todolist/service"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	service service.UserService
}

func NewuserHandler(service service.UserService) *userHandler {
	return &userHandler{service: service}
}

func (h *userHandler) RegisterUser(c *fiber.Ctx) error {
	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	data, token, err := h.service.RegisterUser(user, user.Password)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"message": "register successfully", "data": data, "token": token})
}

func (h *userHandler) LoginUser(c *fiber.Ctx) error {
	var user model.UserLoginInput
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	data, token, err := h.service.LoginUser(user.UsernameOrEmail, user.Password)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"message": "register successfully", "data": data, "token": token})
}
