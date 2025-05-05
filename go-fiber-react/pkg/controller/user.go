package controller

import (
	"go-fiber-react/pkg/model"

	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) ([]model.User, error) {
	var users []model.User
	err := db.Model(&model.User{}).Find(&users).Error
	return users, err
}

func GetAllTodoList(db *gorm.DB) ([]model.TodoList, error) {
	var todo_list []model.TodoParent
	err := db.Model(&model.TodoParent{}).Preload("todo_list").Find(&todo_list).Error
	return todo_list, err
}
