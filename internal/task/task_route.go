package task

import "github.com/gofiber/fiber/v2"

func RegisterTaskRoute(app *fiber.App) {
	task := app.Group("/task")
	task.Get("/", HandleGetTask)
	task.Post("/", HandleCreateTask)
	task.Put("/:id", HandleEditTask)
	task.Delete("/:id", HandleDeleteTask)
}
