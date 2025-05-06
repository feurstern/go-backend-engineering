package model

import "gorm.io/gorm"

type TodoParent struct {
	ID        int64      `gorm:"primaryKey"`
	Title     string     `gorm:"title"`
	List      []TodoList `gorm:"many2many:todo_lists"`
	CreatedAt string     `gorm:"autoCreateTime"`
	UpdatedAt string     `gorm:"autoUpdateTime"`
	DeletedAt string     `gorm:"autoDeleteTime"`
}

type TodoList struct {
	gorm.Model
	ID           int64  `gorm:"id"`
	TodoParentId int64  `gorm:"foreignKey:TodoParent.id"`
	ListName     string `gorm:"not null"`
	CreatedAt    string `gorm:"autoCreateTime"`
	UpdatedAt    string `gorm:"autoUpdateTime"`
	DeletedAt    string `gorm:"autoDeleteTime"`
}
