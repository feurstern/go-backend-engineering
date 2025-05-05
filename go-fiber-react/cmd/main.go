package main

import (
	"go-fiber-react/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func init() {

	database.ConnectDB()
}

func main() {
	app := fiber.New()

	postgresDb, err := database.DBConnection.DB()

	if err != nil {
		panic("error in database connection")
	}

	defer postgresDb.Close()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Connected!")
	})

	app.Get("/welcome-message", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"Message": "Welcome to the app"})
	})

	app.Listen(":5353")
}
