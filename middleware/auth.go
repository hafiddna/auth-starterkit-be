package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/hafiddna/auth-starterkit-be/helper"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		if authorization == "" || len(authorization) < 7 {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Unauthorized",
			})
		}

		token := authorization[7:]
		aToken, err := helper.ValidateRS512Token(config.Config.App.JWT.PublicKey, token)
		if err != nil {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Unauthorized",
				Error:      err.Error(),
			})
		}

		if !aToken.Valid {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Unauthorized",
			})
		}

		return c.Next()
	}
}

func PermissionAuthMiddleware(permissions []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: If the configuration is using team-based permission, also check the team permission
		tokenPermission := c.Locals("user").(map[string]interface{})["permissions"].([]interface{})
		for _, v := range permissions {
			if !helper.ArrayInterfaceContains(tokenPermission, v) {
				return helper.SendResponse(helper.ResponseStruct{
					Ctx:        c,
					StatusCode: fiber.StatusForbidden,
					Message:    "Forbidden",
				})
			}
		}

		return c.Next()
	}
}

func RoleAuthMiddleware(roles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: If the configuration is using team-based role, also check the team role
		tokenRole := c.Locals("user").(map[string]interface{})["roles"].([]interface{})
		for _, v := range roles {
			if !helper.ArrayInterfaceContains(tokenRole, v) {
				return helper.SendResponse(helper.ResponseStruct{
					Ctx:        c,
					StatusCode: fiber.StatusForbidden,
					Message:    "Forbidden",
				})
			}
		}

		return c.Next()
	}
}
