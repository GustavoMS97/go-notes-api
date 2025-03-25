package routes

import (
	"github.com/GustavoMS97/go-notes-api/internal/di"
	"github.com/GustavoMS97/go-notes-api/internal/user"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, c *di.Container) {
	api := app.Group("/api")

	// /api/users
	user.RegisterUserRoutes(api.Group("/users"), c.UserController)
}
