package route

import (
	"github.com/bootmind/figo/pkg/handler"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Expenses route
func Expenses(app *fiber.App, db *gorm.DB) {
	h := &handler.ExpenseHandler{
		DB: db,
	}
	r := app.Group("/expenses")
	r.Get("/", h.Index)
}
