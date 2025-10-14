package service

import (
	"tugas4go/app/model"
	"tugas4go/app/repository"

	"github.com/gofiber/fiber/v2"
)

// GET semua
func GetAllPekerjaan(c *fiber.Ctx) error {
	data, err := repository.GetAllPekerjaan()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "data": data})
}

// GET by ID
func GetPekerjaanByID(c *fiber.Ctx) error {
	id := c.Params("id")

	pekerjaan, err := repository.GetPekerjaanByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"success": true, "data": pekerjaan})
}

// POST
func CreatePekerjaan(c *fiber.Ctx) error {
	var input model.PekerjaanAlumni
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := repository.CreatePekerjaan(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Pekerjaan berhasil ditambahkan",
		"
