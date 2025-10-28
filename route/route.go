package route

import (
	"tugas4go/middleware"
	"tugas4go/app/service"
	mongoservice "tugas4go/app/service/mongo"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", service.GetAllAlumni)

	api := app.Group("/api")

	// LOGIN (PUBLIC)
	api.Post("/login", service.Login)                  // Login Postgres
	api.Post("/mongo/login", mongoservice.LoginUser)   // Login Mongo ( tanpa token)

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

	// Pekerjaan routes (PostgreSQL)
	pekerjaan := protected.Group("/pekerjaan")
	pekerjaan.Get("/", service.GetPekerjaanWithPagination)
	pekerjaan.Get("/:id", service.GetPekerjaanByID)
	pekerjaan.Get("/alumni/:alumni_id", middleware.AdminOnly(), service.GetPekerjaanByAlumniID)
	pekerjaan.Post("/", middleware.AdminOnly(), service.CreatePekerjaan)
	pekerjaan.Put("/:id", middleware.AdminOnly(), service.UpdatePekerjaan)
	pekerjaan.Delete("/:id", middleware.AdminOnly(), service.DeletePekerjaan)

	// Pekerjaan routes (MongoDB)
	mongo := protected.Group("/pekerjaan-mongo")
	mongo.Get("/", mongoservice.GetAllPekerjaanMongo)
	mongo.Get("/:id", mongoservice.GetPekerjaanByIDMongo)
	mongo.Get("/alumni/:alumni_id", middleware.AdminOnly(), mongoservice.GetPekerjaanByAlumniIDMongo)
	mongo.Post("/", middleware.AdminOnly(), mongoservice.CreatePekerjaanMongo)
	mongo.Put("/:id", middleware.AdminOnly(), mongoservice.UpdatePekerjaanMongo)
	mongo.Delete("/:id", middleware.AdminOnly(), mongoservice.DeletePekerjaanMongo)

	// FILE ROUTES (MongoDB)
	files := protected.Group("/mongo/files")
	files.Post("/upload/foto", mongoservice.UploadFoto)
	files.Post("/upload/sertifikat", mongoservice.UploadSertifikat)
	files.Get("/", mongoservice.GetAllFilesMongo)
	files.Get("/:id", mongoservice.GetFileByIDMongo)
	files.Delete("/:id", middleware.AdminOnly(), mongoservice.DeleteFileByIDMongo)
}
