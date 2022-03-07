package users

import (
	"github.com/gofiber/fiber/v2"
)

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
