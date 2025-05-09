package routers

import (
	"go-fiber-react/pkg/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/user", controller.Userlist)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Connected!")
	})

	app.Get("/user-role", controller.UserRoleList)

	app.Get("/welcome-message", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"Message": "Welcome to the app"})
	})

}
