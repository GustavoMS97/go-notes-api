package note

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// DeleteNote godoc
// @Summary Delete a note by ID
// @Tags Notes
// @Produce json
// @Security BearerAuth
// @Param id path string true "Note ID"
// @Success 204 "No Content"
// @Failure 401 {object} docs.ErrorResponse
// @Failure 404 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /notes/{id} [delete]
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
