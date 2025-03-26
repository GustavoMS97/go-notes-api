package app

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/GustavoMS97/go-notes-api/internal/di"
	"github.com/GustavoMS97/go-notes-api/internal/routes"
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

	log.Println("📦 Registered Routes:")
	for _, r := range app.GetRoutes() {
		log.Printf("👉 %s %s", r.Method, r.Path)
	}

	return app
}
