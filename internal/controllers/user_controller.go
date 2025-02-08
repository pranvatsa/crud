package controllers

import (
	"net/http"

	"crud/config"
	"crud/internal/database"
	"crud/internal/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := fetchUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, found, err := fetchUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if config.StorageMode == "mongo" {
		insertedID, err := database.CreateMongoUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		user.ID = insertedID // Assign MongoDB-generated ObjectID
	} else {
		if err := database.CreateJSONUser(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updatedUser models.User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if config.StorageMode == "mongo" {
		if err := database.UpdateMongoUser(id, updatedUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := database.UpdateJSONUser(id, updatedUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if config.StorageMode == "mongo" {
		if err := database.DeleteMongoUser(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := database.DeleteJSONUser(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func fetchUsers() ([]models.User, error) {
	if config.StorageMode == "mongo" {
		return database.GetMongoUsers()
	}
	return database.GetJSONUsers()
}

func fetchUserByID(id string) (models.User, bool, error) {
	if config.StorageMode == "mongo" {
		user, err := database.GetMongoUserByID(id)
		if err != nil {
			return models.User{}, false, err
		}
		return user, true, nil
	}
	user, found := database.GetJSONUserByID(id)
	return user, found, nil
}
