package main

import (
    "fmt"
    "log"
    "tugas4go/config"
    "tugas4go/database"
    "tugas4go/middleware"
    "tugas4go/route"
    "tugas4go/utils"
    "tugas4go/app/repository/mongo" 
)
func main() {
    config.LoadEnv()

    // === Connect PostgreSQL ===
    database.ConnectDB()
    defer database.DB.Close()

    // === Connect MongoDB ===
    database.ConnectMongo()
    mongo.InitPekerjaanCollection() 

    defer func() {
        if database.MongoDB != nil {
            fmt.Println("✅ MongoDB connection closed")
        }
    }()

    logger := middleware.InitLogger()
    logger.Println("Server starting...")

    app := config.NewFiberApp()
    app.Use(middleware.CorsConfig())

    hash, _ := utils.HashPassword("123456")
    fmt.Println("Sample hash:", hash)

    // === Setup all routes ===
    route.SetupRoutes(app)

    log.Fatal(app.Listen(":3000"))
}
