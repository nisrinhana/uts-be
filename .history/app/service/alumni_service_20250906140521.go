package service

import (
	"tugas4go/app/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllAlumni(c *fiber.Ctx) error {
	data, err := repository.GetAllAlumni()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "data": data})
}
