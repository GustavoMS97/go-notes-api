package note

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// ListNotes godoc
// @Summary List notes for the logged-in user
// @Tags Notes
// @Produce json
// @Security BearerAuth
// @Param search query string false "Filter by title (case insensitive)"
// @Success 200 {array} internal_note.Note
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /notes [get]
func (nc *NoteController) GetNotes(c *fiber.Ctx) error {
	log.Println("[GetNotes] Fetching user notes...")
	userID := c.Locals("user_id").(string)
	search := c.Query("search", "")

	notes, err := nc.service.GetNotesByUser(userID, search)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch notes",
		})
	}

	return c.JSON(notes)
}
