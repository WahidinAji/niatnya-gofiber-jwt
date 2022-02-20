package products

import "github.com/gofiber/fiber/v2"

func (d *ProductDeps) GetAll(ctx *fiber.Ctx) error {
	rows, err := d.FindAll(ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(rows)
}