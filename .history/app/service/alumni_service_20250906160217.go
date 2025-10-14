package service

import (
	"tugas4go/app/model"
	"tugas4go/app/repository"

	"github.com/gofiber/fiber/v2"
)

// GET semua
func GetAllAlumni(c *fiber.Ctx) error {
	data, err := repository.GetAllAlumni()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "data": data})
}

// GET by ID
func GetAlumniByID(c *fiber.Ctx) error {
	id := c.Params("id")

	alumni, err := repository.GetAlumniByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Alumni tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"success": true, "data": alumni})
}

// POST tambah alumni
func CreateAlumni(c *fiber.Ctx) error {
	var input model.Alumni
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := repository.CreateAlumni(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Alumni berhasil ditambahkan",
		"data":    input,
	})
}

// PUT update alumni
func UpdateAlumni(c *fiber.Ctx) error {
	id := c.Params("id")
	var input model.Alumni
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := repository.UpdateAlumni(id, input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Alumni berhasil diupdate",
	})
}

// DELETE hapus alumni
func DeleteAlumni(c *fiber.Ctx) error {
	id := c.Params("id")

	err := repository.DeleteAlumni(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Alumni berhasil dihapus",
	})
}
