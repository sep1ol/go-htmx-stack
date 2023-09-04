package main

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/sep1ol/new-stack/utils"
)

func main() {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	configureRoutes(app)

	app.Listen(":3000")
}

func configureRoutes(app *fiber.App) {
	todos := []utils.Todo{
		{ID: 1, Task: "Todo 1", Completed: false},
		{ID: 2, Task: "Todo 2", Completed: false},
		{ID: 3, Task: "Todo 3", Completed: false},
	}

	// App
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Todos": todos,
		})
	})

	// API
	app.Post("/todo/add", func(c *fiber.Ctx) error {
		var todo utils.Todo
		json.Unmarshal(c.Body(), &todo)
		todos = append(todos, todo)
		return c.SendStatus(fiber.StatusCreated)
	})

	app.Delete("/todo/remove/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)

				return c.SendStatus(fiber.StatusOK)
			}
		}
		return c.Status(fiber.StatusNotFound).SendString("Todo not found")
	})
}
