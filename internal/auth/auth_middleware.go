package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization header"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Authorization header format"})
		}

		tokenStr := parts[1]
		claims, err := ParseAndValidateJWT(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
		}

		if refreshClaim, ok := claims["refresh"]; ok && refreshClaim == true {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Refresh tokens are not allowed"})
		}

		userID := claims["user_id"].(string)
		c.Locals("user_id", userID)

		return c.Next()
	}
}
