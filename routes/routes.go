package routes

import (
	"sully/todo-app/controller"
	"sully/todo-app/database"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	cont := controller.NewController(database.DB)
	api := app.Group("/api/v1")
	setupTodoRoutes(cont, api.Group("/todo-items"))
}

func setupTodoRoutes(cont *controller.Controller, router fiber.Router) {

	router.Get("/", cont.GetAllTodoItems)
	router.Get("/:id", cont.GetTodoItem)

	router.Post("/", cont.AddTodoItem)

	router.Put("/:id", cont.CompleteTodoItem)
	router.Delete("/:id", cont.DeleteTodoItem)
}
