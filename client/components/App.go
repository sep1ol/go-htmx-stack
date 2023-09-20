package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sep1ol/new-stack/utils"
)

type AppProps struct {
	Todos  []utils.Todo
	Layout string
	Ctx    *fiber.Ctx
}

func App(props AppProps) error {
	return props.Ctx.Render("html/home-page", fiber.Map{
		"Todos": props.Todos,
	}, props.Layout)
}
