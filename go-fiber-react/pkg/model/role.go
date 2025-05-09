package model

import (
	"time"

	"gorm.io/gorm"
)

type Roles struct {
	ID        int            `gorm:"primaryKey"`
	Name      string         `gorm:"not null; column:name"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}
