package service

import (
	"strconv"                 
	"tugas4go/app/model"
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


func GetAlumniByID(c *fiber.Ctx) error {
	id := c.Params("id")

	alumni, err := repository.GetAlumniByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Alumni tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"success": true, "data": alumni})
}


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

func GetAlumniWithPagination(c *fiber.Ctx) error {
    page, _ := strconv.Atoi(c.Query("page", "1"))
    limit, _ := strconv.Atoi(c.Query("limit", "10"))
    sortBy := c.Query("sortBy", "id")
    order := c.Query("order", "asc")
    search := c.Query("search", "")

    offset := (page - 1) * limit
    if order != "desc" {
        order = "asc"
    }

    data, err := repository.GetAlumniWithPagination(search, sortBy, order, limit, offset)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    total, _ := repository.CountAlumni(search)

    return c.JSON(model.AlumniResponse{
        Data: data,
        Meta: model.MetaInfo{
            Page:   page,
            Limit:  limit,
            Total:  total,
            Pages:  (total + limit - 1) / limit,
            SortBy: sortBy,
            Order:  order,
            Search: search,
        },
    })
}


// // //
func  AlumniStatusKerjaLebih1Tahun(c *fiber.Ctx) error {
	data, err := repository.GetAlumniStatusKerjaLebih1Tahun()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

