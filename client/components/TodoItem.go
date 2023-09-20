package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sep1ol/new-stack/utils"
)

func TodoItem(ctx *fiber.Ctx, todo utils.Todo) error {
	return ctx.Render("html/todo-item", fiber.Map{
		"Task":      todo.Task,
		"Completed": todo.Completed,
	})
}
