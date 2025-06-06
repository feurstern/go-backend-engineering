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
	Category_id int64          `gorm:"foreignKey:BookCategory.id"`
	Created_At  time.Time      `gorm:"column:created_at"`
	Deleted_at  gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

// type BookPayload struct {
// 	Title       string `json:"title"`
// 	Author      string `json:"author"`
// 	Description string `json:"description"`
// 	Category_id int64  `json:"category_id"`
// 	Image       []*File  `json"image"`
// }

type BookCategory struct {
	ID            int32          `gorm:"id"`
	Category_Name string         `gorm:"not null; column:category_name"`
	Created_At    time.Time      `gorm:"column:created_at"`
	Deleted_at    gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

type BookCover struct {
	ID     int64  `gorm:"id"`
	BookId int64  `gorm:"foreignKey:Book.id"`
	Name   string `gorm:"not null; column:name"`
	Path   string `gorm:"not null; column:path"`
}
