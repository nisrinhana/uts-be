package route

import (
	"tugas4go/app/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Root → tampilkan data alumni
	app.Get("/", service.GetAllAlumni)

	// Group API
	api := app.Group("/api")

	// Alumni endpoints
	alumni := api.Group("/alumni")
	alumni.Get("/", service.GetAllAlumni)

	// Pekerjaan endpoints
	pekerjaan := api.Group("/pekerjaan")
	pekerjaan.Get("/", service.GetAllPekerjaan)
}
