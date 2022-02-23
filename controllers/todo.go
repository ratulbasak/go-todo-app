package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id	int	`json:"id"`
	Title	string	`json:"title"`
	Completed	bool `json:"completed"`
}

var todos = []*Todo{
	{
		Id: 1,
		Title: "aaa",
		Completed: false,
	},
	{
		Id: 2,
		Title: "bbb",
		Completed: false,
	},
	{
		Id: 3,
		Title: "ccc",
		Completed: true,
	},
}

func GetTodos(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": todos,
		},
	})
}
