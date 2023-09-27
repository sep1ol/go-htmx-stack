package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sep1ol/new-stack/pkg/structs"
)

// Home page
type HomePageProps struct {
	Todos []structs.Todo
	Ctx   *fiber.Ctx
}

func HomePage(props HomePageProps) error {
	return props.Ctx.Render("home-page", fiber.Map{
		"Todos": props.Todos,
	}, "layout")
}
