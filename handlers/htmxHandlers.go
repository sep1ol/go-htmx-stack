package handlers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sep1ol/new-stack/components"
	"github.com/sep1ol/new-stack/pkg/structs"
	"github.com/sep1ol/new-stack/services"
)

func UseHTMXHandlers(app *fiber.App, db *sql.DB) {
	app.Post("/api/todos/add", func(ctx *fiber.Ctx) error {
		fmt.Println(">> [POST]: Create Todo")

		task := ctx.FormValue("task")
		if task == "" {
			fmt.Println(">> Missing: 'task'")
			return ctx.Status(fiber.StatusBadRequest).SendString("Task is required")
		}

		addedTodo, err := services.Todos(db).AddTodo(
			structs.AddTodo{
				Task:      task,
				Completed: false,
			},
		)
		if err != nil {
			fmt.Println(err)
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		fmt.Println(">> Todo added.")

		return components.TodoItem(
			components.TodoItemProps{
				Ctx:  ctx,
				Todo: *addedTodo,
			})
	})

	app.Delete("/api/todos/remove/:id", func(ctx *fiber.Ctx) error {
		fmt.Println(">> [DELETE]: Todo by ID")
		id, _ := strconv.Atoi(ctx.Params("id"))

		err := services.Todos(db).DeleteTodo(id)
		if err != nil {
			fmt.Println(">> Error deleting todo.")
			fmt.Println(fmt.Sprintf("[id= %d]", id))
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		return nil
	})
}
