package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kennedybaraka/fiber-api/database"
	"github.com/kennedybaraka/fiber-api/pkg/models"
	"github.com/kennedybaraka/fiber-api/pkg/services"
	"github.com/kennedybaraka/fiber-api/pkg/utilities"
)

type controller struct {
}

func NewUserController() *controller {
	return &controller{}
}

type UserController interface {
	RegisterUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	ResetUserPassword(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	AllUsers(c *fiber.Ctx) error
}

// REGISTER USER
func (*controller) RegisterUser(c *fiber.Ctx) error {
	// get data from body parser
	var user models.User

	if e := c.BodyParser(&user); e != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "credentials",
				"message": "bad formatting of input data",
			},
		})
	}
	// validate user input
	err := utilities.UserInputValidation(user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "credentials",
				"message": err.Error(),
			},
		})
	}

	// check if users email exists
	_, exists := services.FindByEmail(user.Email)

	if exists == 1 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "credentials",
				"message": "the email already exists",
			},
		})
	}

	// hash user password
	hash, err := utilities.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "credentials",
				"message": err.Error(),
			},
		})
	}

	user.Password = string(hash)

	// add user to database
	res, err := services.InsertOne(user)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "credentials",
				"message": err.Error(),
			},
		})
	}

	// return response

	return c.JSON(fiber.Map{
		"message": "Thank you for registering!",
		"doc":     res,
	})
}

// LOGIN USER
func (*controller) LoginUser(c *fiber.Ctx) error {
	// get data from body parser

	var user models.User

	if e := c.BodyParser(&user); e != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "credentials",
				"message": "bad formatting of input data",
			},
		})
	}
	// validate user input
	err := utilities.UserInputValidation(user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "credentials",
				"message": err.Error(),
			},
		})
	}

	// check if users email exists
	res, exists := services.FindByEmail(user.Email)

	if exists != 1 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "credentials",
				"message": "the email does not exists",
			},
		})
	}

	// compare hashed password
	err = utilities.VerifyHash(user.Password, res.Password)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "authentication",
				"message": "passwords do not match",
			},
		})
	}

	// generate access and refresh tokens
	access_token, err := utilities.SignAccessToken(3600*time.Second, res.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "authentication",
				"message": err.Error(),
			},
		})

	}
	refresh_token, err := utilities.SignAccessToken(36000*time.Second, res.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "authentication",
				"message": err.Error(),
			},
		})

	}

	// set cookies

	// return response to user

	return c.JSON(fiber.Map{
		"message": "Thank you for logging in!",
		"doc": fiber.Map{
			"refresh_token": refresh_token,
			"access_token":  access_token,
			"res":           res,
		},
	})
}

// LOGIN USER
func (*controller) AllUsers(c *fiber.Ctx) error {
	// get data from body parser

	var users []models.User
	// Get all records
	result := database.DB.Find(&users)
	// SELECT * FROM users;

	if result.RowsAffected == 0 {
		return c.JSON(fiber.Map{
			"message": "There are no users yet!",
		})
	}
	// returns found records count, equals `len(users)`
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "server",
				"message": result.Error.Error(),
			},
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success!",
		"doc": fiber.Map{
			"users": users,
		},
	})
}

// RESET USER PASSWORD
func (*controller) ResetUserPassword(c *fiber.Ctx) error {
	// get data from body parser (email)

	// validate user input

	// check if users email exists

	// generate new password and hash it

	// save new password to database

	// send a password reset email

	// return response to user

	return nil
}

// UPDATE USER
func (*controller) UpdateUser(c *fiber.Ctx) error {
	// get id from params

	// get data from body parser

	// validate user input

	// hash new password if it exists for update

	// update database

	// return response to user

	return nil
}

// DELETE USER
func (*controller) DeleteUser(c *fiber.Ctx) error {
	// get id from params

	// delete from database

	// return response to user

	return nil
}
