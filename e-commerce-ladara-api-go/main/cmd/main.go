package main

import (
	database "ladara-api/pkg/database/config"
	"ladara-api/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.ConnectionDB()
	r := gin.Default()

	r.GET("/users", handlers.GetUsers(db))
	r.GET("/")
	r.Run()
}
