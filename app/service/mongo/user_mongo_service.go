package mongo

import (
	"github.com/gofiber/fiber/v2"
	"tugas4go/app/repository/mongo"
	"tugas4go/utils"
)

// LoginUser: login tanpa token di request, langsung dapat JWT
func LoginUser(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Body tidak valid"})
	}

	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email dan password harus diisi"})
	}

	user, err := mongo.FindUserByEmail(req.Email)
	if err != nil || user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User tidak ditemukan"})
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Password salah"})
	}

	// Generate JWT token
	token, err := utils.GenerateTokenMongo(*user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal membuat token"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Login berhasil",
		"token":   token,
		"user":    user,
	})
}
