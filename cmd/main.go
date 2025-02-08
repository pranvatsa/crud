package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"crud/config"
	"crud/internal/database"
	"crud/internal/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func main() {
	storageMode := getStorageMode()
	fmt.Println("🚀 Running in", storageMode, "mode")

	if storageMode == "mongo" {
		initMongoDB()
	} else {
		fmt.Println("📁 Using JSON storage (users.json)")
		database.InitJSONStorage("users.json")
	}

	router := gin.Default()
	routes.RegisterUserRoutes(router)

	// Start server
	port := getPort()
	fmt.Println("✅ Server running on port:", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// getStorageMode returns "json" (default) or "mongo" based on config or environment variable
func getStorageMode() string {
	mode := os.Getenv("STORAGE_MODE")
	if mode == "mongo" {
		return "mongo"
	}
	return "json" // Default to JSON mode
}

// getPort returns the port from environment variable or default (8080)
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func initMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(config.MongoURI)
	var err error
	mongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("❌ Failed to connect to MongoDB:", err)
	}

	// Ping MongoDB
	if err = mongoClient.Ping(ctx, nil); err != nil {
		log.Fatal("❌ MongoDB not responding:", err)
	}

	fmt.Println("✅ Connected to MongoDB!")
	database.InitMongoDB() // ✅ Ensure MongoDBName exists in config.go
}
