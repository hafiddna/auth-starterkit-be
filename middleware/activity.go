package middleware

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/hafiddna/auth-starterkit-be/helper"
	"github.com/hafiddna/auth-starterkit-be/model"
	"github.com/hafiddna/auth-starterkit-be/repository"
	"time"
)

func ActivityMiddleware(repository repository.SessionRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var oldSessionPayload model.SessionPayload
		var sessionPayload model.SessionPayload
		// TODO: Handle if the API accessed doesn't have Referer / is a functional API, not a page API
		sessionPayload.Previous.URL = c.Get("Referer")

		authorization := c.Get("Authorization")
		ipAddress := c.IP()
		userAgent := c.Get("User-Agent")
		// TODO: How to check a valid User-Agent?
		if len(userAgent) == 0 {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusBadRequest,
				Message:    "Bad Request",
				Error:      "User-Agent is required",
			})
		}

		appID := c.Get("X-App-Id")
		if len(appID) == 0 {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusBadRequest,
				Message:    "Bad Request",
				Error:      "X-App-ID is required",
			})
		} else if !helper.IsUUID(appID) {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusBadRequest,
				Message:    "Bad Request",
				Error:      "X-App-ID is not a valid",
			})
		}

		// TODO: How to check a valid X-Device-Category?
		deviceCategory := c.Get("X-Device-Category")
		if len(deviceCategory) == 0 {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusBadRequest,
				Message:    "Bad Request",
				Error:      "X-Device-Category is required, please provide either 'Web', 'Mobile', 'Desktop App', 'Smart Devices', 'Game Consoles', 'Bots and Automation', 'Virtual or Cloud', or 'Others'",
			})
		}

		// TODO: How to check a valid X-Device-Type?
		deviceType := c.Get("X-Device-Type")
		if len(deviceType) == 0 {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusBadRequest,
				Message:    "Bad Request",
				Error:      "X-Device-Type is required",
			})
		}

		if authorization == "" || len(authorization) < 7 {
			session := model.Session{
				IPAddress: sql.NullString{
					String: ipAddress,
					Valid:  true,
				},
				UserAgent: sql.NullString{
					String: userAgent,
					Valid:  true,
				},
				Payload:        sessionPayload.SessionEncode(),
				LastActivity:   time.Now().UnixNano() / int64(time.Millisecond),
				AppID:          appID,
				DeviceCategory: deviceCategory,
				DeviceType:     deviceType,
			}
			sessionData, err := repository.FindOneByAppID(appID)
			if err != nil {
				session.RememberToken = sql.NullString{
					String: helper.RandomString(10),
					Valid:  true,
				}
				go repository.Create(session)
			} else {
				oldSessionPayload.SessionDecode(sessionData.Payload)
				sessionPayload.Token = oldSessionPayload.Token
				go repository.Update(model.Session{
					Model: model.Model{
						ID:       sessionData.ID,
						Metadata: sessionData.Metadata,
					},
					UserID: sql.NullString{
						String: "",
						Valid:  false,
					},
					IPAddress:      session.IPAddress,
					UserAgent:      session.UserAgent,
					Payload:        sessionPayload.SessionEncode(),
					LastActivity:   time.Now().UnixNano() / int64(time.Millisecond),
					RememberToken:  sessionData.RememberToken,
					AppID:          sessionData.AppID,
					DeviceCategory: deviceCategory,
					DeviceType:     deviceType,
				})
			}
		} else {
			token := authorization[7:]
			aToken, err := helper.ValidateRS512Token(config.Config.App.JWT.PublicKey, token, false)
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

			var encryptedData helper.EncryptedData
			tokenData := helper.JSONMarshal(mapStringClaims["data"])
			helper.JSONUnmarshal([]byte(tokenData), &encryptedData)
			decryptedData, err := helper.DecryptAES256CBC(&encryptedData, []byte(config.Config.App.Secret.AuthKey))
			if err != nil {
				return helper.SendResponse(helper.ResponseStruct{
					Ctx:        c,
					StatusCode: fiber.StatusInternalServerError,
					Message:    "Internal Server Error",
					Error:      err.Error(),
				})
			}

			mapDecryptedData := make(map[string]interface{})
			mapDecryptedData["sub"] = mapStringClaims["sub"]
			helper.JSONUnmarshal([]byte(decryptedData), &mapDecryptedData)

			c.Locals("user", mapDecryptedData)

			sessionData, err := repository.FindOneByAppID(appID)
			if err == nil {
				oldSessionPayload.SessionDecode(sessionData.Payload)
				sessionPayload.Token = oldSessionPayload.Token
				go repository.Update(model.Session{
					Model: model.Model{
						ID:       sessionData.ID,
						Metadata: sessionData.Metadata,
					},
					UserID: sessionData.UserID,
					IPAddress: sql.NullString{
						String: ipAddress,
						Valid:  true,
					},
					UserAgent:      sessionData.UserAgent,
					Payload:        sessionPayload.SessionEncode(),
					LastActivity:   time.Now().UnixNano() / int64(time.Millisecond),
					RememberToken:  sessionData.RememberToken,
					AppID:          sessionData.AppID,
					DeviceCategory: deviceCategory,
					DeviceType:     deviceType,
				})
			} else {
				return helper.SendResponse(helper.ResponseStruct{
					Ctx:        c,
					StatusCode: fiber.StatusUnauthorized,
					Message:    "Unauthorized",
					Error:      err.Error(),
				})
			}
		}

		return c.Next()
	}
}
