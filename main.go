package main

import (
	"go-todo-app/config"
	"go-todo-app/routes"
	// "fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "/ endpoint",
		})
	})

	api := app.Group("/api")

	api.Get("", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "/api endpoint",
		})
	})

	routes.TodoRoute(api.Group("/todos"))
}

func main() {
	db, cancel, err := config.DatabaseInitConnection()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	db.Collection("books")


	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(),)

	setupRoutes(app)
	defer cancel()

	log.Fatal(app.Listen(":3000"))

}