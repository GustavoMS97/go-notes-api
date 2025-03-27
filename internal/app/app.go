package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	_ "github.com/GustavoMS97/go-notes-api/docs"
	"github.com/GustavoMS97/go-notes-api/internal/di"
	"github.com/GustavoMS97/go-notes-api/internal/routes"
)

func InitApp() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())

	// basic health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Go + Fiber!")
	})
	// swagger
	app.Get("/docs/*", swagger.HandlerDefault)

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
