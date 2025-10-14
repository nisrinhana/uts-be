package route

import (
	"tugas4go/app/service"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	alumni := api.Group("/alumni")
	alumni.Get("/", service.GetAllAlumni)

	pekerjaan := api.Group("/pekerjaan")
	pekerjaan.Get("/", service.GetAllPekerjaan)
}
