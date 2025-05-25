package controller

import (
	"go-fiber-react/pkg/database"
	"go-fiber-react/pkg/model"

	"github.com/gofiber/fiber/v2"
)

func Registration(c *fiber.Ctx) error {

	payload := model.RegisterPayload{}

	var users model.User
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid Request"})
	}

	db := database.DBConnection

	if err := db.Create(&users{
		users.Username: payload.Username,
		users.Email:    payload.Email,
		users.Password: payload.Password,
	}); err != nil {

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    payload,
	})

}
