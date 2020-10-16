package handler

import (
	"fmt"
	"strconv"

	"github.com/bootmind/figo/pkg/entity"
	"github.com/bootmind/figo/pkg/validation"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ExpenseHandler type
type ExpenseHandler struct {
	DB *gorm.DB
}

// Index to list all expenses
func (h ExpenseHandler) Index(ctx *fiber.Ctx) error {
	var expenses []entity.Expense
	h.DB.Find(&expenses)
	return ctx.JSON(fiber.Map{"data": expenses})
}

// Show an expense
func (h ExpenseHandler) Show(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your expense"}})
	}

	var expenseDB entity.Expense
	expenseDB.ID = uint(id)
	h.DB.First(&expenseDB)

	return ctx.JSON(fiber.Map{"data": expenseDB})
}

// Store a new expense
func (h ExpenseHandler) Store(ctx *fiber.Ctx) error {
	expense := new(entity.Expense)

	if err := ctx.BodyParser(expense); err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your request"}})
	}

	if validate := validation.CreateOrUpdate(*expense); len(validate) > 0 {
		return ctx.Status(422).JSON(fiber.Map{"errors": validate})
	}

	h.DB.Create(&expense)
	return ctx.JSON(fiber.Map{"message": "Expense successfully registered"})
}

// Update an expense
func (h ExpenseHandler) Update(ctx *fiber.Ctx) error {
	expense := new(entity.Expense)

	if err := ctx.BodyParser(expense); err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your request"}})
	}

	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your expense"}})
	}

	if validate := validation.CreateOrUpdate(*expense); len(validate) > 0 {
		return ctx.Status(422).JSON(fiber.Map{"errors": validate})
	}

	var expenseDB entity.Expense
	h.DB.First(&expenseDB, id)
	h.DB.Model(&expenseDB).Updates(map[string]interface{}{
		"title": expense.Title,
		"total": expense.Total,
	})

	return ctx.JSON(fiber.Map{"message": "Expense successfully updated"})
}

// Destroy an expense
func (h ExpenseHandler) Destroy(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your expense"}})
	}

	var expenseDB entity.Expense
	expenseDB.ID = uint(id)
	h.DB.Delete(&expenseDB)

	return ctx.JSON(fiber.Map{"message": "Expense successfully removed"})
}

// Upload an attachment
func (h ExpenseHandler) Upload(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)

	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able to process your expense"}})
	}

	file, err := ctx.FormFile("attachment")

	if err != nil {
		return ctx.Status(422).JSON(fiber.Map{"errors": [1]string{"We were not able upload your attachment"}})
	}

	ctx.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))

	var expenseDB entity.Expense
	h.DB.First(&expenseDB, id)
	h.DB.Model(&expenseDB).Update("attachment", file.Filename)

	return ctx.JSON(fiber.Map{"message": "Attachment uploaded successfully"})
}
