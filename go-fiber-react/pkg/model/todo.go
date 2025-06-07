package model

import (
	"time"

	"gorm.io/gorm"
)

type TodoParent struct {
	ID        int64          `gorm:"primaryKey"`
	Title     string         `gorm:"title"`
	List      []TodoList     `gorm:"foreignKey : TodoParentId"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"null;column_updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

type TodoPayload struct {
	Title string `json:"title"`
	List  []TodoListPayload
}

type TodoListPayload struct {
	TodoParentId int64  `json:"to_do_parent_id"`
	List         string `json:"list"`
}

type TodoList struct {
	ID           int64          `gorm:"id"`
	TodoParentId int64          `gorm:"foreignKey:TodoParent.id"`
	List         string         `gorm:"not null;column:list"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"null;column_updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index;column:deleted_at"`
}
