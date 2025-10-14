package route

import (
	"tugas4go/app/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Root
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Selamat datang di API Alumni 🚀",
			"routes": fiber.Map{
				"GET /api/alumni":    "Ambil semua data alumni",
				"GET /api/pekerjaan": "Ambil semua data pekerjaan alumni",
			},
		})
	})

	// Group API
	api := app.Group("/api")

	alumni := api.Group("/alumni")
	alumni.Get("/", service.GetAllAlumni)

	pekerjaan := api.Group("/pekerjaan")
	pekerjaan.Get("/", service.GetAllPekerjaan)
}
