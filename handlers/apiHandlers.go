package handlers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sep1ol/new-stack/client/components"
	"github.com/sep1ol/new-stack/pkg/structs"
	"github.com/sep1ol/new-stack/services"
)

func RegisterAPIHandlers(app *fiber.App, db *sql.DB) {
	app.Post("/api/todos/add", func(ctx *fiber.Ctx) error {
		fmt.Println(">> [POST]: /api/todos/add")

		task := ctx.FormValue("task")
		if task == "" {
			fmt.Println(">> Missing: 'task'")
			return ctx.Status(fiber.StatusBadRequest).SendString("Task is required")
		}

		todos, err := services.Todos(db).AddTodo(
			structs.AddTodo{
				Task:      task,
				Completed: false,
			},
		)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(">> Todo added.")

		return components.TodosList(components.TodosListProps{
			Ctx:   ctx,
			Todos: todos,
		})
	})

	app.Delete("/api/todos/remove/:id", func(ctx *fiber.Ctx) error {
		fmt.Println(">> [DELETE]: /api/todos/remove/:id")
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
