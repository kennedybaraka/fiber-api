package services

import (
	"fmt"

	"github.com/kennedybaraka/fiber-api/database"
	"github.com/kennedybaraka/fiber-api/pkg/models"
)

type services struct {
}

func NewUserServices() *services {
	return &services{}
}

type UserServices interface {
	InsertOne() (string, error)
	FindByEmail() (string, error)
	FindById() (string, error)
	UpdateOne() (string, error)
	DeleteOne() error
}

// insert user into database
func InsertOne(user models.User) (models.User, error) {
	result := database.DB.Create(&user)
	if result.RowsAffected != 1 {
		fmt.Println("User not created")
		return models.User{}, result.Error
	}
	return user, nil
}

// find user by email
func FindByEmail(email string) (models.User, int64) {
	var user models.User

	result := database.DB.First(&user, "email = ?", email)

	if result.RowsAffected != 1 {
		return models.User{}, result.RowsAffected
	}
	return user, result.RowsAffected
}

// find user by id
func FindById(id string) (string, error) {
	return id, nil
}

// update user data
func UpdateOne(id string) (string, error) {
	return id, nil
}

// delete user data
func DeleteOne(id string) error {
	return nil
}
