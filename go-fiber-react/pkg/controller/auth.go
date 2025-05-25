package controller

import (
	"go-fiber-react/pkg/database"
	"go-fiber-react/pkg/model"
	"go-fiber-react/pkg/utils"

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

	var existingUser model.User

	if err := db.Take(&existingUser, "email = ?", payload.Email).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "The email is already exist",
		})
	}
	if err := db.Create(&user).Error; err != nil {
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

	var existingUser model.User
	if err := db.Take(&existingUser, "email = ?", payload.Email).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "The email is already taken",
		})
	}
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

func Login(c *fiber.Ctx) error {

	payload := model.LoginPayload{}

	var user model.User

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request",
		})
	}

	db := database.DBConnection

	if err := db.Where("email = ?", payload.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "failed",
		})
	}

	token, err := utils.GenerateJWT(user.ID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"token":   token,
		"data":    user,
	})
}

func Auth(c *fiber.Ctx) error {
	payload := model.LoginPayload{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "invalid request!",
		})
	}

	var user model.User
	db := database.DBConnection

	if err := db.Where("email = ?", payload.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"Message": "Invalid credentials",
		})
	}

	if err := db.Where("password = ?", bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid credentials",
		})
	}

	token, err := utils.GenerateJWT(user.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
		"token":   token,
	})
}
