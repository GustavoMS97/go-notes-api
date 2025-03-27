package user

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/GustavoMS97/go-notes-api/internal/auth"
)

// RefreshToken godoc
// @Summary Refresh access token
// @Description Use a valid refresh token to generate a new access and refresh token
// @Tags Users
// @Accept json
// @Produce json
// @Param request body RefreshRequest true "Refresh token"
// @Success 201 {object} docs.LoginResponse
// @Failure 400 {object} docs.ErrorResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /users/refresh [post]
func (uc *UserController) RefreshToken(c *fiber.Ctx) error {
	log.Println("[RefreshToken] Refreshing user auth token...")
	var body RefreshRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	userID, err := auth.ParseAndValidateRefreshToken(body.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	newAccessToken, err := auth.GenerateJWT(userID, false)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate access token",
		})
	}
	newRefreshToken, err := auth.GenerateJWT(userID, true)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate refresh token",
		})
	}

	return c.JSON(fiber.Map{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}
