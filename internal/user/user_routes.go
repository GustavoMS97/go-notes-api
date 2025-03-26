package user

import (
	"github.com/gofiber/fiber/v2"

	"github.com/GustavoMS97/go-notes-api/internal/auth"
)

func RegisterUserRoutes(router fiber.Router, controller *UserController) {
	router.Post("/", controller.CreateUser)
	router.Post("/login", controller.Login)
	router.Post("/refresh", controller.RefreshToken)
	router.Get("/me", auth.JWTMiddleware(), controller.Me)
}
