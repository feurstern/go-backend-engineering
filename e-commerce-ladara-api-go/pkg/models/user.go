package model

type User struct {
	ID         uint `gorm:"primaryKey"`
	Username   string
	Email      string
	Role       int
	Deleted_at string
	Created_at string
	Updated_at string
	Google_id  string
}
