package todoClient

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sep1ol/new-stack/client/components"
	"github.com/sep1ol/new-stack/utils"
)

func CreateRoutes(app *fiber.App, todos []utils.Todo) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		log.Print("Rendering App")
		return components.App(
			components.AppProps{
				Todos:  todos,
				Layout: "html/layout",
				Ctx:    ctx,
			},
		)
	})

	app.Get("/render-todos", func(ctx *fiber.Ctx) error {
		return components.TodosList(ctx, todos)
	})
}
