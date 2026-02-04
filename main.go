package main

import (
	"log"

	"belajar-fiber/database"
	"belajar-fiber/handlers"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Initialize Database
	database.ConnectDB()

	// Initialize Fiber
	app := fiber.New()

	// Routes
	app.Get("/users", handlers.GetAllUsers)
	app.Get("/users/:id", handlers.GetUserByID)
	app.Post("/users", handlers.CreateUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.DeleteUser)

	// Start Server
	log.Fatal(app.Listen(":3000"))
}
