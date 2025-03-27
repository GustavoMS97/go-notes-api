package note

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/GustavoMS97/go-notes-api/internal/internal_error"
)

// CreateNote godoc
// @Summary Create a new note
// @Tags Notes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateNoteRequest true "Note data"
// @Success 201 {object} Note
// @Failure 400 {object} docs.ErrorResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /notes [post]
func (nc *NoteController) CreateNote(c *fiber.Ctx) error {
	log.Println("[CreateNote] Creating Note...")
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
