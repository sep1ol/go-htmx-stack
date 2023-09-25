package handlers

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sep1ol/new-stack/client/components"
	"github.com/sep1ol/new-stack/services"
)

func RegisterClientHandlers(app *fiber.App, db *sql.DB) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		todos, err := services.Todos(db).GetTodos()
		if err != nil {
			fmt.Println(">> Error fetching todos.")
			fmt.Println(err)
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		fmt.Println(">> Fetching successful.")
		fmt.Println(todos)
		fmt.Println(">> Rendering App...")

		return components.HomePage(
			components.HomePageProps{
				Todos:  todos,
				Layout: "layout",
				Ctx:    ctx,
			},
		)
	})
}
