package main

import (
	"log"
	"tasklybe/internal/task"
	"tasklybe/pkg/db"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error Loading Env")
	}

	app := fiber.New() //initialize web server
	log.Println("Conecting to Database...")
	db.Connect()

	log.Println("Migrating Table...")
	if err := db.DB.AutoMigrate(&task.Task{}); err != nil {
			log.Fatal("Failed to Migrate table", err)
		}

		task.RegisterTaskRoute(app)


	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}