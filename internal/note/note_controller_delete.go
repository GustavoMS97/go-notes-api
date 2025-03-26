package note

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (nc *NoteController) DeleteNote(c *fiber.Ctx) error {
	log.Println("[DeleteNote] Deleting Note...")
	noteID := c.Params("id")
	userID := c.Locals("user_id").(string)

	err := nc.service.DeleteNote(noteID, userID)
	if err != nil {
		if err.Error() == "note not found or not owned by user" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete note"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
