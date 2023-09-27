package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sep1ol/new-stack/pkg/structs"
)

// Todos List
type TodosListProps struct {
	Todos []structs.Todo
	Ctx   *fiber.Ctx
}

func TodosList(props TodosListProps) error {
	return props.Ctx.Render("todos-list", fiber.Map{
		"Todos": props.Todos,
	})
}

// Todo Item
type TodoItemProps struct {
	Todo structs.Todo
	Ctx  *fiber.Ctx
}

func TodoItem(props TodoItemProps) error {
	return props.Ctx.Render("todo-item", props.Todo)
}
