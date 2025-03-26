package note

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/GustavoMS97/go-notes-api/internal/internal_error"
)

func (nc *NoteController) UpdateNote(c *fiber.Ctx) error {
	log.Println("[UpdateNote] Updating Note...")
	noteID := c.Params("id")
	userID := c.Locals("user_id").(string)

	var body UpdateNoteRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": internal_error.FormatValidationError(err),
		})
	}

	note, err := nc.service.UpdateNote(noteID, userID, body.Title, body.Content)
	if err != nil {
		if err.Error() == "note not found or not owned by user" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(note)
}
