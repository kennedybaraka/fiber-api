package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kennedybaraka/fiber-api/pkg/controllers"
)

// users handler for routes
var (
	controller controllers.UserController = controllers.NewUserController()
)

func Users(router fiber.Router) {
	api := router.Group("/api/user")

	api.Post("/register", controller.RegisterUser)
	api.Post("/login", controller.LoginUser)
	api.Get("/all", controller.AllUsers)

}
