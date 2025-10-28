package mongo

import (
	"strconv"
	modelMongo "tugas4go/app/model/mongo"
	repoMongo "tugas4go/app/repository/mongo"

	"github.com/gofiber/fiber/v2"
)

// GET ALL
func GetAllPekerjaanMongo(c *fiber.Ctx) error {
	data, err := repoMongo.GetAllPekerjaanMongo()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

// GET BY ID
func GetPekerjaanByIDMongo(c *fiber.Ctx) error {
	id := c.Params("id")
	data, err := repoMongo.GetPekerjaanByIDMongo(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Data tidak ditemukan"})
	}
	return c.JSON(data)
}

// GET BY ALUMNI_ID
func GetPekerjaanByAlumniIDMongo(c *fiber.Ctx) error {
	alumniID, _ := strconv.Atoi(c.Params("alumni_id"))
	data, err := repoMongo.GetPekerjaanByAlumniIDMongo(alumniID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(data)
}

// CREATE
func CreatePekerjaanMongo(c *fiber.Ctx) error {
	var p modelMongo.PekerjaanAlumniMongo
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := repoMongo.CreatePekerjaanMongo(p)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "Pekerjaan berhasil ditambahkan"})
}

// UPDATE
func UpdatePekerjaanMongo(c *fiber.Ctx) error {
	id := c.Params("id")
	var p modelMongo.PekerjaanAlumniMongo
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := repoMongo.UpdatePekerjaanMongo(id, p)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Pekerjaan berhasil diperbarui"})
}

// DELETE
func DeletePekerjaanMongo(c *fiber.Ctx) error {
	id := c.Params("id")
	err := repoMongo.DeletePekerjaanMongo(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Pekerjaan berhasil dihapus"})
}
