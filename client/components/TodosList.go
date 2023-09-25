package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sep1ol/new-stack/pkg/structs"
)

type TodosListProps struct {
	Todos []structs.Todo
	Ctx   *fiber.Ctx
}

func TodosList(props TodosListProps) error {
	return props.Ctx.Render("todos-list", fiber.Map{
		"Todos": props.Todos,
	})
}
