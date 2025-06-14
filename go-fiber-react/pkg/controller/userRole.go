package controller

import (
	"go-fiber-react/pkg/database"
	"go-fiber-react/pkg/model"

	"github.com/gofiber/fiber/v2"
)

func UserRoleList(c *fiber.Ctx) error {

	db := database.DBConnection
	var user_role []model.UserRoles

	if err := db.Where("deleted_at IS NULL").Find(&user_role).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     "Error during retrieving the data",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"success": true,
		"data":    user_role,
	})

}

func CreateUserRole(c *fiber.Ctx) error {
	db := database.DBConnection

	payload := model.UserRolePayload{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "Invalid Request!"})
	}

	userRole := model.UserRoles{
		Name: payload.Name,
	}

	if err := db.Create(&userRole).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Internal server error"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    userRole,
	})
}
