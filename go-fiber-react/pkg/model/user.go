package model

type User struct {
	ID        int64  `gorm:"primaryKey"`
	Username  string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Role      int    `gorm:"not null"`
	Password  string `gorm:"password"`
	CreatedAt string `gorm:"created_at"`
	DeletedAt string `gorm:"deleted_at"`
}
