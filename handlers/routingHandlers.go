package handlers

import (
	"database/sql"
	"fmt"
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/sep1ol/new-stack/components"
	"github.com/sep1ol/new-stack/services"
)

var bootstrap *template.Template

func UseRoutingHandlers(app *fiber.App, db *sql.DB) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		todos, err := services.Todos(db).GetTodos()
		if err != nil {
			fmt.Println(">> Error fetching todos.")
			fmt.Println(err)
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		return components.HomePage(
			components.HomePageProps{
				Todos: todos,
				Ctx:   ctx,
			},
		)
	})
}
