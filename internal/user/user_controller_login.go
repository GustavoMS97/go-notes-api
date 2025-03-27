package user

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/GustavoMS97/go-notes-api/internal/auth"
	"github.com/GustavoMS97/go-notes-api/internal/internal_error"
)

// Login godoc
// @Summary Authenticate user
// @Description Log in with email and password and receive access and refresh tokens
// @Tags Users
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} docs.LoginResponse
// @Failure 400 {object} docs.ErrorResponse
// @Failure 401 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /users/login [post]
func (uc *UserController) Login(c *fiber.Ctx) error {
	log.Println("[Login] Authenticating user...")
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
