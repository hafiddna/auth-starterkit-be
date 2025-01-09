package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hafiddna/auth-starterkit-be/config"
)

// TODO: Recheck this middleware
func CORSMiddleware(config config.CfgStruct) fiber.Handler {
	return cors.New(
		cors.Config{
			AllowOriginsFunc: nil,
			AllowOrigins:     config.App.Server.Cors,
			AllowMethods:     "GET, POST, PATCH, PUT, DELETE",
			AllowHeaders:     "Accept, Content-Type, Content-Length, Accept-Encoding, Accept-Language, X-CSRF-Token, Authorization, X-Requested-With, User-Agent, Connection, Host",
		},
	)
}
