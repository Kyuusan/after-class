package user

import (
	"tasklybe/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App) {
	user := app.Group("/user")
	user.Post("/register", HandleRegister)
	user.Post("/login", HandleLogin)
	user.Get("/me", middleware.Auth(), HandleMe)
}