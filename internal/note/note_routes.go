package note

import (
	"github.com/GustavoMS97/go-notes-api/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func RegisterNoteRoutes(router fiber.Router, controller *NoteController) {
	router.Post("/", auth.JWTMiddleware(), controller.CreateNote)
	router.Get("/", auth.JWTMiddleware(), controller.GetNotes)
	router.Put("/:id", auth.JWTMiddleware(), controller.UpdateNote)
	router.Delete("/:id", auth.JWTMiddleware(), controller.DeleteNote)
}
