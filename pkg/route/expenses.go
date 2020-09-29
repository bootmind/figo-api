package route

import (
	"github.com/bootmind/figo/pkg/handler"
	"github.com/gofiber/fiber/v2"
)

// Expenses route
func Expenses(app *fiber.App) {
	var h handler.ExpenseHandler
	r := app.Group("/expenses")
	r.Get("/", h.Index)
}
