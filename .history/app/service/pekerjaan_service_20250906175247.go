package service

import (
	"tugas4go/app/model"
	"tugas4go/app/repository"

	"github.com/gofiber/fiber/v2"
)


// GET semua pekerjaan
func GetAllPekerjaan(c *fiber.Ctx) error {
	data, err := repository.GetAllPekerjaan()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "data": data})
}

// GET pekerjaan by ID
func GetPekerjaanByID(c *fiber.Ctx) error {
	id := c.Params("id")

	pekerjaan, err := repository.GetPekerjaanByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"success": true, "data": pekerjaan})
}

// GET pekerjaan by alumni_id
func GetPekerjaanByAlumniID(c *fiber.Ctx) error {
	alumniID := c.Params("alumni_id")

	data, err := repository.GetPekerjaanByAlumniID(alumniID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if len(data) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan untuk alumni ini tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"success": true, "data": data})
}

// POST tambah pekerjaan
func CreatePekerjaan(c *fiber.Ctx) error {
	var input model.PekerjaanAlumni
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Pastikan AlumniID valid sebelum insert (opsional, bisa tambahkan cek di sini)

	err := repository.CreatePekerjaan(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Pekerjaan berhasil ditambahkan",
		"data":    input,
	})
}

// PUT update pekerjaan
func UpdatePekerjaan(c *fiber.Ctx) error {
	id := c.Params("id")
	var input model.PekerjaanAlumni
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := repository.UpdatePekerjaan(id, input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Pekerjaan berhasil diupdate",
	})
}

// DELETE pekerjaan
func DeletePekerjaan(c *fiber.Ctx) error {
	id := c.Params("id")

	err := repository.DeletePekerjaan(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Pekerjaan berhasil dihapus",
	})
}
