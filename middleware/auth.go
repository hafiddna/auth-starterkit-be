package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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

		mapStringClaims := make(map[string]interface{})
		for key, value := range aToken.Claims.(jwt.MapClaims) {
			mapStringClaims[key] = value
		}

		if !aToken.Valid {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Unauthorized",
			})
		}

		var encryptedData helper.EncryptedData
		tokenData := helper.JSONMarshal(mapStringClaims["data"])
		helper.JSONUnmarshal([]byte(tokenData), &encryptedData)
		decryptedData, err := helper.DecryptAES256CBC(&encryptedData, []byte(config.Config.App.Secret.AuthKey))
		if err != nil {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Unauthorized",
				Error:      err.Error(),
			})
		}

		mapDecryptedData := make(map[string]interface{})
		mapDecryptedData["sub"] = mapStringClaims["sub"]
		helper.JSONUnmarshal([]byte(decryptedData), &mapDecryptedData)

		c.Locals("user", mapDecryptedData)

		return c.Next()
	}
}

func PermissionAuthMiddleware(permissions []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
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
