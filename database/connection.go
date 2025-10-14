package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Gagal koneksi DB:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("❌ Gagal ping DB:", err)
	}

	fmt.Println("✅ Berhasil konek ke PostgreSQL")
}
