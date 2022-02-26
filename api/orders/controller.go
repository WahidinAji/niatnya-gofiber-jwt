package orders

import (
	"github.com/gofiber/fiber/v2"
)

func (d *OrderDeps) GetAll(ctx *fiber.Ctx) error {
	data, err := FindAll(ctx.Context(), d.DB)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"data":   data,
		"status": "success",
	})
}

func (d *OrderDeps) GetById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	data, err := d.ServiceById(ctx.Context(), int64(id))
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"data":   data,
		"status": "success",
	})
}
