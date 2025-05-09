package controller

import (
	"go-fiber-react/pkg/database"
	"go-fiber-react/pkg/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]model.User, error) {
	var users []model.User
	err := db.Model(&model.User{}).Find(&users).Error
	return users, err
}

func Userlist(c *fiber.Ctx) error {
	context := fiber.Map{
		"success": true,
		"msg":     "User list",
	}

	db := database.DBConnection

	var users model.User

	db.Find(&users).Where("deleted_at IS NULL")

	context["users"] = users

	c.Status(200)
	return c.JSON(context)
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

// func DeleteUser(c *fiber.Ctx) error {

// }

// func GetAllTodoList(db *gorm.DB) ([]model.TodoList, error) {
// 	var todo_list []model.TodoParent
// 	err := db.Model(&model.TodoParent{}).Preload("todo_list").Find(&todo_list).Error
// 	// return todo_list, err
// }
