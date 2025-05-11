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
