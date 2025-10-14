package main

import (
	"tugas4go/config"
	"tugas4go/database"
	"tugas4go/middleware"
	"tugas4go/route"
	"tugas4go/utils"

	"fmt"
	"log"
)

func main() {
	config.LoadEnv()

	database.ConnectDB()
	defer database.DB.Close()

	logger := middleware.InitLogger()
	logger.Println("Server starting...")

	app := config.NewFiberApp()
	app.Use(middleware.CorsConfig())

	hash, _ := utils.HashPassword("123456")
	fmt.Println(hash)

	route.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

