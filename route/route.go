package route

import (
	"tugas4go/middleware"
	"tugas4go/app/service"

	"github.com/gofiber/fiber/v2"
)
func SetupRoutes(app *fiber.App) {
    app.Get("/", service.GetAllAlumni)

    api := app.Group("/api")

    // LOGIN (PUBLIC) 
    api.Post("/login", service.Login)

    // PROTECTED
    protected := api.Group("", middleware.AuthRequired())

 // Alumni routes
alumni := protected.Group("/alumni")
alumni.Get("/", service.GetAlumniWithPagination) 
alumni.Get("/:id", service.GetAlumniByID)
alumni.Get("/status-kerja-lebih-1-tahun", service.AlumniStatusKerjaLebih1Tahun)

alumni.Post("/", middleware.AdminOnly(), service.CreateAlumni)
alumni.Put("/:id", middleware.AdminOnly(), service.UpdateAlumni)
alumni.Delete("/:id", middleware.AdminOnly(), service.DeleteAlumni)

// Pekerjaan routes
pekerjaan := protected.Group("/pekerjaan")

// Soft delete, restore, dan trash 
pekerjaan.Delete("/softdelete/:id", service.SoftDeletePekerjaan)
pekerjaan.Delete("/harddelete/:id", middleware.AdminOnly(), service.HardDeletePekerjaan)
pekerjaan.Put("/restore/:id", service.RestorePekerjaan)
pekerjaan.Get("/trash", service.GetTrashedPekerjaan)

// utama
pekerjaan.Get("/", service.GetPekerjaanWithPagination)
pekerjaan.Get("/:id", service.GetPekerjaanByID)
pekerjaan.Get("/alumni/:alumni_id", middleware.AdminOnly(), service.GetPekerjaanByAlumniID)
pekerjaan.Post("/", middleware.AdminOnly(), service.CreatePekerjaan)
pekerjaan.Put("/:id", middleware.AdminOnly(), service.UpdatePekerjaan)
pekerjaan.Delete("/:id", middleware.AdminOnly(), service.DeletePekerjaan)

}
