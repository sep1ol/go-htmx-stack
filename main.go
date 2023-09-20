package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	todoAPI "github.com/sep1ol/new-stack/API"
	todoClient "github.com/sep1ol/new-stack/client"
	"github.com/sep1ol/new-stack/utils"
)

func main() {
	engine := html.New("./client", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	configureRoutes(app)

	app.Listen(":3000")
}

func configureRoutes(app *fiber.App) {
	todos := []utils.Todo{
		{ID: 1, Task: "Todo 1", Completed: true},
		{ID: 2, Task: "Todo 2", Completed: false},
		{ID: 3, Task: "Todo 3", Completed: false},
	}

	todoClient.CreateRoutes(app, todos)
	todoAPI.CreateRoutes(app, todos)
}
