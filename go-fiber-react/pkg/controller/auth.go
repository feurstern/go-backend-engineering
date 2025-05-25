package controller

import (
	"go-fiber-react/pkg/model"

	"github.com/gofiber/fiber/v2"
)

func Registration(c *fiber.Ctx) error {

	payload := model.RegisterPayload{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Request",
		})
	}

	// db := database.DBConnection

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    payload,
	})

}
