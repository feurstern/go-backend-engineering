package database

import (
	"go-fiber-react/pkg/model"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConnection *gorm.DB

func ConnectDB() {

	godotenv.Load()

	dsn := "host=localhost user=postgres dbname=go_fiber_quasar port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("Failed to connect the database")
	}

	log.Println("Connected")

	db.AutoMigrate(new(model.UserProfile))
	DBConnection = db

}
