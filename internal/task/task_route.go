package task

import "github.com/gofiber/fiber/v2"

func RegisterTaskRoute(app *fiber.App) {
	task := app.Group("/task")
	task.Get("/", HandleGetTask)
}
