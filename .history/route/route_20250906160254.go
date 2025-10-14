// Pekerjaan endpoints
pekerjaan := api.Group("/pekerjaan")
pekerjaan.Get("/", service.GetAllPekerjaan)      // GET semua
pekerjaan.Get("/:id", service.GetPekerjaanByID)  // GET by ID
pekerjaan.Post("/", service.CreatePekerjaan)     // POST
pekerjaan.Put("/:id", service.UpdatePekerjaan)   // PUT update
pekerjaan.Delete("/:id", service.DeletePekerjaan) // DELETE hapus
