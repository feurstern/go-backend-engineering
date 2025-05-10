package controller

import (
	"go-fiber-react/pkg/database"
	"go-fiber-react/pkg/model"

	"github.com/gofiber/fiber/v2"
)

func TodoList(c *fiber.Ctx) error {

	ctx := fiber.Map{
		"success": false,
	}

	db := database.DBConnection
	var todo *model.TodoParent

	db.Find(&todo).Where("deleted_at is NULL")

	c.Status(200)

	ctx["todo"] = todo

	return c.JSON(ctx)

}
