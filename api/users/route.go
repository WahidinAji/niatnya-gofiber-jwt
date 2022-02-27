package users

import "github.com/gofiber/fiber/v2"

func (d *UserDeps) UserRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/users", d.LoginUser)
}
