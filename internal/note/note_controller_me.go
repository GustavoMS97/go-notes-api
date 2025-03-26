package note

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

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
