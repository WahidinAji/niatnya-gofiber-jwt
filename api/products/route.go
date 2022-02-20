package products

import (
	"github.com/gofiber/fiber/v2"
)

//create funtion for routing
func (d *ProductDeps) ProductRouter(app *fiber.App) {

	api := app.Group("/api")
	api.Get("/products", d.GetAll)
	// app.Get("/products/:id", Get)
	// app.Post("/products", Post)
	// app.Put("/products/:id", Put)
	// app.Delete("/products/:id", Delete)
}
