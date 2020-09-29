package handler

import (
	"time"

	"github.com/bootmind/figo/pkg/entity"
	"github.com/gofiber/fiber/v2"
)

// ExpenseHandler type
type ExpenseHandler struct {
}

// Index to list all expenses
func (h ExpenseHandler) Index(ctx *fiber.Ctx) error {
	expenses := []entity.Expense{
		{
			ID:         "1",
			Title:      "Lunch at MyFood",
			Total:      14.95,
			Attachment: "photo.jpg",
			CreatedAt:  time.Now(),
		},
	}
	return ctx.JSON(fiber.Map{"data": expenses})
}
