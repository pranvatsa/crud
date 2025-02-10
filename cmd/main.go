package main

import (
	"context"
	"flag"
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
	https := flag.Bool("s", false, "Run server with HTTPS")
	certFile := flag.String("cert", "", "Path to the SSL certificate file")
	keyFile := flag.String("key", "", "Path to the SSL key file")
	flag.Parse()

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
	if *https {
		if *certFile == "" || *keyFile == "" {
			log.Fatal("SSL certificate and key files must be provided for HTTPS")
		}
		if err := router.RunTLS(":"+port, *certFile, *keyFile); err != nil {
			log.Fatal("Failed to start HTTPS server:", err)
		}
	} else {
		if err := router.Run(":" + port); err != nil {
			log.Fatal("Failed to start server:", err)
		}
	}
}

func getPort() string {
	return config.GetEnv("PORT", "8080")
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
