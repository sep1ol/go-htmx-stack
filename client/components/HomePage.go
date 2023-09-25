package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sep1ol/new-stack/pkg/structs"
)

type HomePageProps struct {
	Todos  []structs.Todo
	Layout string
	Ctx    *fiber.Ctx
}

func HomePage(props HomePageProps) error {
	return props.Ctx.Render("home-page", fiber.Map{
		"Todos": props.Todos,
	}, props.Layout)
}
