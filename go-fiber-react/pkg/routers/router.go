package routers

import (
	"go-fiber-react/pkg/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/user", controller.Userlist)
	// admin role only
	app.Get("/user-role", controller.UserRoleList)
	app.Post("/todo", controller.CreateTodo)
	app.Get("/todo", controller.TodoList)
	app.Post("/login", controller.Login)
	app.Post("/registration", controller.Registration)
	//end of admin role only
	app.Get("/welcome-message", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"Message": "Welcome to the app"})
	})

	// admin := app.Group("/", middleware.JWTMiddleware(), middleware.RoleAuthMiddleware(1))

	// admin.Post("/user-delete/:id", controller.DeleteUser)
	// admin.Post("/role-delete/:id", controller.DeleteTodoList)

	// global/ guess/ unauthorized

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Connected!")
	})

}
