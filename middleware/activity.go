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

		authorization := c.Get("Authorization")
		ipAddress := c.IP()
		userAgent := c.Get("User-Agent")
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

		deviceCategory := c.Get("X-Device-Category")
		if len(deviceCategory) == 0 {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusBadRequest,
				Message:    "Bad Request",
				Error:      "X-Device-Category is required",
			})
		} else if !helper.IsValidDeviceCategory(deviceCategory) {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusBadRequest,
				Message:    "Bad Request",
				Error:      "X-Device-Category is not a valid",
			})
		}

		deviceType := c.Get("X-Device-Type")
		if len(deviceType) == 0 {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusBadRequest,
				Message:    "Bad Request",
				Error:      "X-Device-Type is required",
			})
		} else if !helper.IsValidDeviceType(deviceType) {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusBadRequest,
				Message:    "Bad Request",
				Error:      "X-Device-Type is not a valid",
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
			oldSessionPayload.SessionDecode(sessionData.Payload)
			if err != nil {
				session.RememberToken = sql.NullString{
					String: helper.RandomString(10),
					Valid:  true,
				}
				if c.Get("Referer") != "" {
					sessionPayload.Previous.URL = c.Get("Referer")
				} else {
					sessionPayload.Previous.URL = oldSessionPayload.Previous.URL
				}
				err = repository.Create(session)
				if err != nil {
					return helper.SendResponse(helper.ResponseStruct{
						Ctx:        c,
						StatusCode: fiber.StatusInternalServerError,
						Message:    "Internal Server Error",
						Error:      err.Error(),
					})
				}
			} else {
				if c.Get("Referer") != "" {
					sessionPayload.Previous.URL = c.Get("Referer")
				} else {
					sessionPayload.Previous.URL = oldSessionPayload.Previous.URL
				}
				sessionPayload.Token = oldSessionPayload.Token
				err = repository.Update(model.Session{
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
				if err != nil {
					return helper.SendResponse(helper.ResponseStruct{
						Ctx:        c,
						StatusCode: fiber.StatusInternalServerError,
						Message:    "Internal Server Error",
						Error:      err.Error(),
					})
				}
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
			oldSessionPayload.SessionDecode(sessionData.Payload)
			if err == nil {
				if c.Get("Referer") != "" {
					sessionPayload.Previous.URL = c.Get("Referer")
				} else {
					sessionPayload.Previous.URL = oldSessionPayload.Previous.URL
				}
				sessionPayload.Token = oldSessionPayload.Token
				err = repository.Update(model.Session{
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
				if err != nil {
					return helper.SendResponse(helper.ResponseStruct{
						Ctx:        c,
						StatusCode: fiber.StatusInternalServerError,
						Message:    "Internal Server Error",
						Error:      err.Error(),
					})
				}
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
