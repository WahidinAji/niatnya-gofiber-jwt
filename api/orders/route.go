package orders

import (
	"github.com/gofiber/fiber/v2"
)

func (d *OrderDeps) OrderRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/orders", d.GetAll)
	api.Get("/orders/:id", d.GetById)
	// router.Get("/:id", GetOne)
	// router.Post("/", Create)
	// router.Put("/:id", Update)
	// router.Delete("/:id", Delete)
}
