package todoAPI

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sep1ol/new-stack/client/components"
	"github.com/sep1ol/new-stack/utils"
)

func CreateRoutes(app *fiber.App, todos []utils.Todo) {
	app.Post("/todo/add", func(ctx *fiber.Ctx) error {
		task := ctx.FormValue("task")
		if task == "" {
			return ctx.Status(fiber.StatusBadRequest).SendString("Task is required")
		}

		todo := utils.Todo{
			ID:        len(todos) + 1,
			Task:      task,
			Completed: false,
		}

		todos = append(todos, todo)

		fmt.Println(">> Todo added.")
		fmt.Println("ID:", todo.ID)

		return components.TodosList(ctx, todos)
	})

	app.Delete("/todo/remove/:id", func(ctx *fiber.Ctx) error {
		id, _ := strconv.Atoi(ctx.Params("id"))
		fmt.Println(id)
		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)

				fmt.Println(">> Todo deleted.")
				fmt.Println("ID:", todos[i].ID)

				return components.TodosList(ctx, todos)
			}
		}

		return ctx.Status(fiber.StatusNotFound).SendString("Todo not found")
	})
}
