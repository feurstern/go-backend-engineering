package controller

import (
	"go-fiber-react/pkg/database"
	"go-fiber-react/pkg/model"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func HashedPasswordConversion(p string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)

	if err != nil {
		return ""
	}

	return string(hashedPassword)
}

func Registration(c *fiber.Ctx) error {

	payload := model.RegisterPayload{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,

			"message": "Invalid Request"})
	}

	hashedPassword := HashedPasswordConversion(payload.Password)
	db := database.DBConnection

	user := model.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashedPassword,
		RoleId:   1,
	}

	if err := db.Create(user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to register",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "The user  has been registered succesfully",
	})

}

func UserRegistration(c *fiber.Ctx) error {
	payload := model.RegisterPayload{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid  Request",
		})
	}

	hashedPassword := HashedPasswordConversion(payload.Password)

	user := model.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashedPassword,
		RoleId:   1,
	}

	db := database.DBConnection

	if err := db.Create(user).Error; err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"Message": "Failed to register the data!",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "The data has been submitted successfully!",
		"data":    user,
	})
}
