package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/GustavoMS97/go-notes-api/internal/auth"
	"github.com/GustavoMS97/go-notes-api/internal/internal_error"
)

func (uc *UserController) Login(c *fiber.Ctx) error {
	var body LoginRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": internal_error.FormatValidationError(err),
		})
	}

	user, err := uc.service.GetByEmail(body.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	if !auth.CheckPasswordHash(body.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	accessToken, err := auth.GenerateJWT(user.ObjectID.Hex(), false)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate access token",
		})
	}

	refreshToken, err := auth.GenerateJWT(user.ObjectID.Hex(), true)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate refresh token",
		})
	}

	return c.JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
