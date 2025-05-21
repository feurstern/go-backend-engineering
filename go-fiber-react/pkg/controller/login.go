package controller

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	email := c.Params("email")
	password := c.Params("password")

}
