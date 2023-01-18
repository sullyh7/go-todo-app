package main

import (
	"sully/todo-app/database"
	"sully/todo-app/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {
	database.Initialise()
}

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":8080")
}
