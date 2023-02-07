package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/kennedybaraka/fiber-api/pkg/routes"
)

func init() {
	// load config, database, and other plugins
	// configs.InitializeAppConfig()
}
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	// main app entry point
	app := fiber.New()

	// health check route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "The server is up and running!",
		})
	})

	routes.Users(app)
	// 404 error handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fiber.Map{
				"type":    "api",
				"message": c.Path() + " does not exist",
			},
		})
	})

	// listening to server on a port (from env)
	app.Listen(getPort())

}
