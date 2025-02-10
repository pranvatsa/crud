package database

import "crud/internal/models"

type Database interface {
	GetUsers() ([]models.User, error)
	GetUserByID(id string) (models.User, error)
	CreateUser(user models.User) (string, error)
	UpdateUser(id string, updatedUser models.User) error
	DeleteUser(id string) error
}
