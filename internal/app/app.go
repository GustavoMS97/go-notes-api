package app

import (
	"github.com/GustavoMS97/go-notes-api/internal/di"
	"github.com/GustavoMS97/go-notes-api/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func InitApp() *fiber.App {
	app := fiber.New()

	// basic health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Go + Fiber!")
	})

	// setup dependencies
	container := di.NewContainer()

	// register routes
	routes.Register(app, container)

	return app
}
