package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          int64          `gorm:"id"`
	Title       string         `gorm:"not null; column:title"`
	Author      string         `gorm:"not null; column:author"`
	Description string         `gorm:"not null; column:description"`
	Category_id int64          `gorm:"not null :colum:category_id"`
	Created_At  time.Time      `gorm:"column:created_at"`
	Deleted_at  gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

type BookCover struct {
	ID     int64  `gorm:"id"`
	BookId int64  `gorm:"foreignKey:Book.id"`
	Name   string `gorm:"not null; column:name"`
	Path   string `gorm:"not null; column:path"`
}
