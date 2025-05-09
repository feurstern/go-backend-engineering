package model

import (
	"time"

	"gorm.io/gorm"
)

type TodoParent struct {
	ID        int64          `gorm:"primaryKey"`
	Title     string         `gorm:"title"`
	List      []TodoList     `gorm:"many2many:todo_lists"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"null;column_updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

type TodoList struct {
	gorm.Model
	ID           int64          `gorm:"id"`
	TodoParentId int64          `gorm:"foreignKey:TodoParent.id"`
	List         string         `gorm:"not null;column:list"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"null;column_updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index;column:deleted_at"`
}
