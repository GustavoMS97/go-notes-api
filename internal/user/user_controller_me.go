package user

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (uc *UserController) Me(c *fiber.Ctx) error {
	log.Println("[Me] Fetching user own info...")
	userID := c.Locals("user_id").(string)

	user, err := uc.service.FindByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}
