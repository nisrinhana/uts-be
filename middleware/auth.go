package middleware

import (
	"strings"
	"tugas4go/utils"


	"github.com/gofiber/fiber/v2"
)

// Middleware untuk memerlukan login
func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token akses diperlukan"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Format token tidak valid"})
		}

		claims, err := utils.ValidateToken(parts[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid atau expired", "detail": err.Error()})
		}

		// set user info ke context
		c.Locals("user_id", claims.UserID)
		c.Locals("username", claims.Username)
		c.Locals("role", claims.Role)
		c.Locals("user", claims)

		return c.Next()
	}
}

// AdminOnly memastikan role
func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		r := c.Locals("role")
		role, ok := r.(string)
		if !ok || role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Akses ditolak. Hanya admin yang diizinkan"})
		}
		return c.Next()
	}
}
