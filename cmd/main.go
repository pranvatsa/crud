package main

import (
	"log"
	"os"

	"crud/internal/database"
	"crud/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.ConnectDB(); err != nil {
		log.Fatalf("❌ Failed to connect to MongoDB: %v", err)
		os.Exit(1) // Exit safely instead of panic
	}

	router := gin.Default()
	routes.RegisterRoutes(router)

	log.Println("🚀 Server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
