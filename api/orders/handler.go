package orders

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

func Handler(db *pgx.Conn, val *validator.Validate) *fiber.App {
	app := fiber.New()
	deps := &OrderDeps{
		DB:        db,
		Validator: val,
	}

	//app.GET == anonimous func that returns a func(ctx *fiber.Ctx) error or u can
	//use deps.GetAll that returns the same response
	app.Get("/", func(c *fiber.Ctx) error {
		data, err := FindAll(c.Context(), deps.DB)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   data,
			"status": "success",
		})
	})

	app.Get("/:id", deps.GetById)
	app.Post("/", deps.CreateOrder)
	return app
}
