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
