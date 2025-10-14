package service

import (
	"tugas4go/app/repository"

	"github.com/gofiber/fiber/v2"
)

func GetAllPekerjaan(c *fiber.Ctx) error {
	data, err := repository.GetAllPekerjaan()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "data": data})
}
