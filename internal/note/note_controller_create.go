package note

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/GustavoMS97/go-notes-api/internal/internal_error"
)

func (nc *NoteController) CreateNote(c *fiber.Ctx) error {
	var body CreateNoteRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": internal_error.FormatValidationError(err),
		})
	}

	userID := c.Locals("user_id").(string)

	note, err := nc.service.CreateNote(body.Title, body.Content, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create note"})
	}

	return c.Status(fiber.StatusCreated).JSON(note)
}
