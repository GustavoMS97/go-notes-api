package user

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/GustavoMS97/go-notes-api/internal/internal_error"
)

// CreateUser godoc
// @Summary Register a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "User data"
// @Success 201 {object} User
// @Failure 400 {object} docs.ErrorResponse
// @Failure 422 {object} docs.ErrorResponse
// @Failure 500 {object} docs.ErrorResponse
// @Router /users [post]
func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	log.Println("[CreateUser] Creating user...")
	var body CreateUserRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": internal_error.FormatValidationError(err),
		})
	}

	user, err := uc.service.CreateUser(body.Name, body.Email, body.Password)
	if err != nil {
		if err.Error() == "email already registered" {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "Email already in use.",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user.",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
