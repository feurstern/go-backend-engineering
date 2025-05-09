package controller

import (
	"go-fiber-react/pkg/database"
	"go-fiber-react/pkg/model"

	"github.com/gofiber/fiber/v2"
)

func UserRoleList(c *fiber.Ctx) error {
	context := fiber.Map{
		"message": "User role list",
	}

	db := database.DBConnection

	var user_role *model.UserRoles

	db.Find(&user_role)

	c.Status(200)

	context["user_roles"] = user_role

	return c.JSON(context)

}
