package routes

import (
	"crud/internal/controllers"
	"crud/internal/database"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, db database.Database) {
	userController := controllers.NewUserController(db)

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", userController.GetUsers)
		userRoutes.GET("/:id", userController.GetUserByID)
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}
}
