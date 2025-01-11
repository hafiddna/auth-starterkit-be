package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hafiddna/auth-starterkit-be/helper"
	"github.com/hafiddna/auth-starterkit-be/service"
)

func AuthMiddleware(service service.JWTService) fiber.Handler {
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
		aToken, err := service.ValidateToken(token)
		if err != nil {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Unauthorized",
				Error:      err.Error(),
			})
		}

		mapStringClaims := make(map[string]interface{})
		for key, value := range aToken.Claims.(jwt.MapClaims) {
			mapStringClaims[key] = value
		}
		c.Locals("user", mapStringClaims)

		return c.Next()
	}
}

// TODO: Recheck this middleware
func PermissionAuthMiddleware(permission []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenPermission := c.Locals("user").(map[string]interface{})["permissions"].([]interface{})
		for _, v := range permission {
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

// TODO: Recheck this middleware
func RoleAuthMiddleware(role []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenRole := c.Locals("user").(map[string]interface{})["roles"].([]interface{})
		for _, v := range role {
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

// TODO: Recheck this middleware
func SystemRoleAuthMiddleware(role string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenRole := c.Locals("user").(map[string]interface{})["role"].(string)
		if tokenRole != role {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusForbidden,
				Message:    "Forbidden",
			})
		}
		return c.Next()
	}
}
