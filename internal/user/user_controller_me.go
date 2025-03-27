package user

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Me godoc
// @Summary Get logged in user
// @Description Return the currently authenticated user
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} user.User
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /users/me [get]
func (uc *UserController) Me(c *fiber.Ctx) error {
	log.Println("[Me] Fetching user own info...")
	userID := c.Locals("user_id").(string)

	user, err := uc.service.FindByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}
