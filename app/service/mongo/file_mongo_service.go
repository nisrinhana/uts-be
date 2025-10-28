package mongo

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	modelMongo "tugas4go/app/model/mongo"
	repoMongo "tugas4go/app/repository/mongo"

	"github.com/gofiber/fiber/v2"
)

// ===============================
// UPLOAD FOTO
// ===============================
func UploadFoto(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	userID, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"error": "UserID invalid"})
	}
	username := c.Locals("username").(string)

	// Untuk admin bisa upload untuk user lain
	alumniID := userID
	if role == "admin" {
		if idStr := c.FormValue("alumni_id"); idStr != "" {
			alumniID = idStr
		}
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "File tidak ditemukan"})
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return c.Status(400).JSON(fiber.Map{"error": "Format file harus JPG, JPEG, atau PNG"})
	}
	if file.Size > 1*1024*1024 {
		return c.Status(400).JSON(fiber.Map{"error": "Ukuran file maksimal 1MB"})
	}

	uploadDir := "./uploads/foto"
	os.MkdirAll(uploadDir, os.ModePerm)
	savePath := filepath.Join(uploadDir, file.Filename)

	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan file"})
	}

	fileMeta := modelMongo.FileMongo{
		FileName:   file.Filename,
		FilePath:   savePath,
		FileType:   file.Header.Get("Content-Type"),
		FileSize:   file.Size,
		AlumniID:   alumniID,
		UploadedBy: username,
		UploadedAt: time.Now(),
	}
	_ = repoMongo.CreateFileMongo(fileMeta)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Foto berhasil diupload",
		"data":    fileMeta,
	})
}

// ======================================
// UPLOAD SERTIFIKAT (PDF ≤ 2MB)
// ======================================
func UploadSertifikat(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	userID, ok := c.Locals("user_id").(string)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"error": "UserID invalid"})
	}
	username := c.Locals("username").(string)

	alumniID := userID
	if role == "admin" {
		if idStr := c.FormValue("alumni_id"); idStr != "" {
			alumniID = idStr
		}
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "File tidak ditemukan"})
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".pdf" {
		return c.Status(400).JSON(fiber.Map{"error": "Format file harus PDF"})
	}
	if file.Size > 2*1024*1024 {
		return c.Status(400).JSON(fiber.Map{"error": "Ukuran file maksimal 2MB"})
	}

	uploadDir := "./uploads/sertifikat"
	os.MkdirAll(uploadDir, os.ModePerm)
	savePath := filepath.Join(uploadDir, file.Filename)

	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan file"})
	}

	fileMeta := modelMongo.FileMongo{
		FileName:   file.Filename,
		FilePath:   savePath,
		FileType:   file.Header.Get("Content-Type"),
		FileSize:   file.Size,
		AlumniID:   alumniID,
		UploadedBy: username,
		UploadedAt: time.Now(),
	}
	_ = repoMongo.CreateFileMongo(fileMeta)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Sertifikat berhasil diupload",
		"data":    fileMeta,
	})
}


// ======================================
// GET ALL FILES
// ======================================
func GetAllFilesMongo(c *fiber.Ctx) error {
	files, err := repoMongo.GetAllFilesMongo()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true, "data": files})
}

// GET FILE BY ID
func GetFileByIDMongo(c *fiber.Ctx) error {
    fileName := c.Params("id")

    // Ambil metadata file dari MongoDB
    file, err := repoMongo.GetFileByName(fileName)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "File tidak ditemukan"})
    }

    // Pastikan file path valid di filesystem
    if _, err := os.Stat(file.FilePath); os.IsNotExist(err) {
        return c.Status(404).JSON(fiber.Map{"error": "File tidak ditemukan"})
    }

    // Download file sesuai path di database
    return c.Download(file.FilePath)
}
func DeleteFileByIDMongo(c *fiber.Ctx) error {
    fileName := c.Params("id")

    // ambil file metadata dari MongoDB
    file, err := repoMongo.GetFileByName(fileName)
    if err != nil || file == nil {
        return c.Status(404).JSON(fiber.Map{"error": "File tidak ditemukan"})
    }

    // hapus file dari filesystem
    if err := os.Remove(file.FilePath); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus file"})
    }

    // hapus record MongoDB
    _ = repoMongo.DeleteFileRecord(fileName)

    return c.JSON(fiber.Map{
        "success": true,
        "message": "File berhasil dihapus",
    })
}

