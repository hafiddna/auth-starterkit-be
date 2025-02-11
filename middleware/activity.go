package middleware

import "github.com/gofiber/fiber/v2"

// TODO: on creating session, using a middleware create remember_token using helper.RandomString(10)
func ActivityMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
