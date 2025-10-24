package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

func ConnectMongo() {
	fmt.Println("=== [DEBUG] Mulai koneksi MongoDB ===")

	uri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGODB_NAME")

	// Debug print untuk memastikan env terbaca
	fmt.Println("=== [DEBUG] MONGODB_URI:", uri)
	fmt.Println("=== [DEBUG] MONGODB_NAME:", dbName)

	if uri == "" {
		uri = "mongodb://localhost:27017"
		fmt.Println("=== [DEBUG] URI kosong, pakai default:", uri)
	}

	clientOpts := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Fatal("❌ Gagal konek Mongo:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ Ping Mongo gagal:", err)
	}

	if dbName == "" {
		log.Fatal("❌ [DEBUG] database name cannot be empty — Pastikan kamu punya file .env dan variabel MONGODB_NAME diisi!")
	}

	MongoDB = client.Database(dbName)
	fmt.Println("✅ MongoDB connected:", dbName)
	fmt.Println("=== [DEBUG] MongoDB connection sukses ===")
}
