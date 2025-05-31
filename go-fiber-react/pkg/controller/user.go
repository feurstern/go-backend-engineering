package controller

import (
	"go-fiber-react/pkg/database"
	"go-fiber-react/pkg/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]model.User, error) {
	var users []model.User
	err := db.Model(&model.User{}).Find(&users).Error
	return users, err
}

func Userlist(c *fiber.Ctx) error {

	db := database.DBConnection
	var users []model.User

	if err := db.Where("deleted_at IS NULL").Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to fetch the fuckjing user data",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"msg":     "user list",
		"success": true,
		"data":    users,
	})

}

func UpdateUser(c *fiber.Ctx) error {
	context := fiber.Map{
		"success": true,
		"message": "The data has been updated succesfully",
	}
	c.Status(200)
	return c.JSON(context)
}

// func UpdateUser(c *fiber.Ctx) error {

// }

func DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"msg":     "Invalid ID request",
		})
	}
	db := database.DBConnection
	var user *[]model.User

	if err := db.Model(&model.User{}).Where("id = ?", id).Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"msg":     "Failed to delete the data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"msg":     "The data has been deleted succesfully!",
	})

	// if err:= db.where("id = %d")
}

// func GetAllTodoList(db *gorm.DB) ([]model.TodoList, error) {
// 	var todo_list []model.TodoParent
// 	err := db.Model(&model.TodoParent{}).Preload("todo_list").Find(&todo_list).Error
// 	// return todo_list, err
// }
