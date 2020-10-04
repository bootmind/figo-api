package handler

import (
	"github.com/bootmind/figo/pkg/entity"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ExpenseHandler type
type ExpenseHandler struct {
	DB *gorm.DB
}

// Index to list all expenses
func (h ExpenseHandler) Index(ctx *fiber.Ctx) error {
	expenses := []entity.Expense{}
	return ctx.JSON(fiber.Map{"data": expenses})
}
