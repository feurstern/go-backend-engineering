package controller

import (
	"go-fiber-react/pkg/database"
	"go-fiber-react/pkg/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func TodoList(c *fiber.Ctx) error {
	db := database.DBConnection

	var todo []model.TodoList

	if err := db.Where("deleted_at IS NULL").Find(&todo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     "Faield to retrieve the data ",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    todo,
	})
}

func DeleteTodoList(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": true,
			"msg":     "Invalid id!",
		})
	}
	db := database.DBConnection
	var todo []model.TodoList

	if err := db.Model(&model.TodoList{}).Where("id = ?", id).Delete(&todo).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"succes": false,
			"msg":    "Failed to delete the to do list",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"msg":     "The todo list has been deleted succesfully!",
	})

}
