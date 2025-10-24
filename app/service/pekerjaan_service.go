package service

import (
	"strconv"                 
	"tugas4go/app/model"
	"tugas4go/app/repository"
	 "database/sql"

	"github.com/gofiber/fiber/v2"
)

func GetAllPekerjaan(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	isAdmin := role == "admin"

	data, err := repository.GetAllPekerjaan(isAdmin)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "data": data})
}

func GetPekerjaanByID(c *fiber.Ctx) error {
    id := c.Params("id")

    pekerjaan, err := repository.GetPekerjaanByID(id)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": "Data pekerjaan tidak ditemukan",
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(pekerjaan)
}



func GetPekerjaanByAlumniID(c *fiber.Ctx) error {
	alumniID := c.Params("alumni_id")
	role := c.Locals("role").(string)
	isAdmin := role == "admin"

	data, err := repository.GetPekerjaanByAlumniID(alumniID, isAdmin)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if len(data) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan untuk alumni ini tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"success": true, "data": data})
}

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
		"data":    input,
	})
}

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

func GetPekerjaanWithPagination(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortBy := c.Query("sortBy", "id")
	order := c.Query("order", "asc")
	search := c.Query("search", "")

	offset := (page - 1) * limit
	if order != "desc" {
		order = "asc"
	}

	data, err := repository.GetPekerjaanWithPagination(search, sortBy, order, limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	total, _ := repository.CountPekerjaan(search)

	return c.JSON(model.PekerjaanResponse{
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

//SoftDelete 
func SoftDeletePekerjaan(c *fiber.Ctx) error {
	id := c.Params("id")
	role := c.Locals("role").(string)
	userID := c.Locals("user_id").(int)

	isAdmin := role == "admin"

	err := repository.SoftDeletePekerjaan(id, isAdmin, userID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Data berhasil dihapus (soft delete)",
	})
}


// Menampilkan (Trash)
func GetTrashedPekerjaan(c *fiber.Ctx) error {
	data, err := repository.GetTrashedPekerjaan()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Data pekerjaan yang berada di trash",
		"data":    data,
	})
}

// Restore data dari Trash
func RestorePekerjaan(c *fiber.Ctx) error {
	id := c.Params("id")

	err := repository.RestorePekerjaan(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Data pekerjaan berhasil dikembalikan dari trash",
	})
}

// Hard delete 
func HardDeletePekerjaan(c *fiber.Ctx) error {
	id := c.Params("id")

	err := repository.HardDeletePekerjaan(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Data berhasil dihapus permanen",
	})
}

