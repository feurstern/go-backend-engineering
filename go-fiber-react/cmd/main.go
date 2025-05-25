package main

import (
	"go-fiber-react/pkg/database"
	"go-fiber-react/pkg/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {

	database.ConnectDB()
}

func main() {
	app := fiber.New()

	app.Use(cors.New())
	postgresDb, err := database.DBConnection.DB()

	if err != nil {
		panic("error in database connection")
	}
	defer postgresDb.Close()

	// app.Use(logger.New({}))
	routers.SetupRoutes(app)
	app.Listen(":5353")
}
