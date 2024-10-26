package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	//TODO: Implement the token verification
	return c.Next()
}
