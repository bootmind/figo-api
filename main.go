package main

import (
	"github.com/bootmind/figo/pkg/connector"
	"github.com/bootmind/figo/pkg/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := connector.Connect()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Figo API",
		})
	})

	route.Expenses(app, db)
	app.Listen(":3000")
}
