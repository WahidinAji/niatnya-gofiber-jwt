package users

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte("secret"),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})

	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
	}
}

func (d *UserDeps) LoginUser(c *fiber.Ctx) error {
	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// errs := d.ValidateStruct(input)
	// if errs != nil {
	// 	return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
	// 		"status":  "error",
	// 		"message": errs,
	// 	})
	// }

	user, err := d.LoginServiceUser(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"data":   user,
		"status": "success",
	})
}
