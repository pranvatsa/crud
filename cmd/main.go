package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"crud/config"
	"crud/internal/database"
	"crud/internal/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	config.LoadConfig()
	fmt.Println("🚀 Running in", config.StorageMode, "mode")

	var db database.Database
	if config.StorageMode == "mongo" {
		db = initMongoDB()
	} else {
		fmt.Println("📁 Using JSON storage (users.json)")
		db = database.NewJSONStorage("users.json")
	}

	router := gin.Default()
	routes.RegisterUserRoutes(router, db)

	port := getPort()
	fmt.Println("✅ Server running on port:", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func getPort() string {
	port := config.GetEnv("PORT", "8080")
	return port
}

func initMongoDB() database.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(config.MongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("❌ Failed to connect to MongoDB:", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("❌ MongoDB not responding:", err)
	}

	fmt.Println("✅ Connected to MongoDB!")
	return database.NewMongoDB(client, config.MongoDBName)
}
