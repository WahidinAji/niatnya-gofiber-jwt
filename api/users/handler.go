package users

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

func Handler(db *pgx.Conn, val *validator.Validate) *fiber.App {
	app := fiber.New()
	deps := &UserDeps{
		DB:        db,
		Validator: val,
	}
	app.Post("/login", deps.LoginUser)
	return app
}
