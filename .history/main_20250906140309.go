package main

import (
	"tugas4go/config"
	"tugas4go/database"
	"tugas4go/middleware"
	"tugas4go/route"
	"log"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to DB
	database.ConnectDB()
	defer database.DB.Close()

	// Init logger
	logger := config.InitLogger()
	logger.Println("Server starting...")

	// Fiber app
	app := config.NewFiberApp()
	app.Use(middleware.CorsConfig())

	// Routes
	route.SetupRoutes(app)

	// Run server
	log.Fatal(app.Listen(":3000"))
}
