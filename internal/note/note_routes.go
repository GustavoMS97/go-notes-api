package note

import (
	"github.com/GustavoMS97/go-notes-api/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func RegisterNoteRoutes(router fiber.Router, controller *NoteController) {
	router.Post("/", auth.JWTMiddleware(), controller.CreateNote)
}
