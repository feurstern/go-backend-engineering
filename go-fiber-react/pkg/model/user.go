package model

import (
	// "go-fiber-react/pkg/model"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int64          `gorm:"primaryKey"`
	Username  string         `gorm:"not null; column:username"`
	Email     string         `gorm:"not null;column:email"`
	RoleId    int            `gorm:"not null;column:role_id"`
	Role      UserRoles      `gorm:"foreignKey: RoleId"`
	Password  string         `gorm:"not null; column:password"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

type UserProfile struct {
	ID          int64          `gorm:"primaryKey"`
	User_Id     int64          `gorm:"foreignKey:User.id"`
	Profile_Pic string         `gorm:"column:profile_pic"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

type UserRoles struct {
	ID        int            `gorm:"primaryKey"`
	Name      string         `gorm:"not null; column:name"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}
