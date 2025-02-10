package controllers

import (
	"net/http"

	"crud/internal/database"
	"crud/internal/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	db database.Database
}

func NewUserController(db database.Database) *UserController {
	return &UserController{db: db}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.db.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.db.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	id, err := uc.db.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.ID = id

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser models.User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := uc.db.UpdateUser(id, updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := uc.db.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
