package routes

import (
	"github.com/gofiber/fiber/v2"

	"go-todo-app/controllers"
)

func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
}
