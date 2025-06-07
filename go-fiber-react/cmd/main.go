package main

import (
	"go-fiber-react/pkg/database"
	"go-fiber-react/pkg/routers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {

	database.ConnectDB()
}

func main() {

	godotenv.Load()
	app := fiber.New()

	app.Use(cors.New())
	postgresDb, err := database.DBConnection.DB()

	if err != nil {
		panic("Kernel Panic : Failed to connect to the data")
	}

	addr := os.Getenv("HTTP_LISTEN_ADDRESS")

	if addr == "" {
		panic(" Kernel Panic : FAILED TO LISTEN THE ADDRESS!")
	}

	defer postgresDb.Close()

	routers.SetupRoutes(app)

	app.Listen(addr)
}
