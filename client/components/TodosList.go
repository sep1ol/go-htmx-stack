package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sep1ol/new-stack/utils"
)

func TodosList(Ctx *fiber.Ctx, Todos []utils.Todo) error {
	return Ctx.Render("html/todos-list", fiber.Map{
		"Todos": Todos,
	})
}
