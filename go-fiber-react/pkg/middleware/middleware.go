package middleware

import (
	"go-fiber-react/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing authorization header",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization format",
			})
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return utils.SecretKey, nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token claims",
			})
		}

		c.Locals("userId", int64(claims["user_id"].(float64)))
		c.Locals("roleId", int(claims["role_id"].(float64)))

		return c.Next()
	}
}

func RoleAuthMiddleware(roleIdPayload int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleId := c.Locals("roleId")

		if roleId == nil || roleId != roleIdPayload {

			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Forbidden chamber!",
			})
		}

		return c.Next()
	}
}
