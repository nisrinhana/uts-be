package service

import (
	"tugas4go/app/model"
	"tugas4go/app/repository"
	"tugas4go/utils"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Request tidak valid"})
	}
	if req.Username == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username dan password harus diisi"})
	}

	user, err := repository.GetUserByUsername(req.Username)
	if err != nil {
		
		if err == sql.ErrNoRows {
			user, err = repository.GetUserByEmail(req.Username)
		}
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Username atau password salah"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error", "detail": err.Error()})
	}

	// cek password
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Username atau password salah"})
	}

	// buat token
	token, err := utils.GenerateTokenPostgres(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal generate token"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Login berhasil",
		"data": model.LoginResponse{
			User:  user,
			Token: token,
		},
	})
}
