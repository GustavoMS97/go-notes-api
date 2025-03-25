package user

import "github.com/gofiber/fiber/v2"

func RegisterUserRoutes(router fiber.Router, controller *UserController) {
	router.Post("/", controller.CreateUser)
}
