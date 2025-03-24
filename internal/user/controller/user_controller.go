package controller

import (
	"github.com/GustavoMS97/go-notes-api/internal/user/entity"
	"github.com/GustavoMS97/go-notes-api/internal/user/service"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := c.service.GetUser(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(user)
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	var user entity.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Dados inválidos"})
	}
	c.service.CreateUser(user)
	return ctx.JSON(user)
}
