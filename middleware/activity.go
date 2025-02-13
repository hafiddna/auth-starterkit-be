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
		var sessionPayload model.SessionPayload
		sessionPayload.Previous.URL = c.Get("Referer")

		authorization := c.Get("Authorization")
		userAgent := c.Get("User-Agent")
		ipAddress := c.IP()
		appID := c.GetReqHeaders()["X-App-Id"]
		rememberToken := helper.RandomString(10)

		if len(appID) == 0 {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusBadRequest,
				Message:    "Bad Request",
				Error:      "X-App-ID is required",
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
				Payload:      sessionPayload.SessionEncode(),
				LastActivity: time.Now().UnixNano() / int64(time.Millisecond),
				AppID:        appID[0],
				RememberToken: sql.NullString{
					String: rememberToken,
					Valid:  true,
				},
			}
			sessionData, err := repository.FindOneByAppID(appID[0])
			if err != nil {
				err := repository.Create(session)
				if err != nil {
					return helper.SendResponse(helper.ResponseStruct{
						Ctx:        c,
						StatusCode: fiber.StatusInternalServerError,
						Message:    "Internal Server Error",
						Error:      err.Error(),
					})
				}
			} else {
				err := repository.Update(model.Session{
					Model: model.Model{
						ID:       sessionData.ID,
						Metadata: sessionData.Metadata,
					},
					UserID: sql.NullString{
						String: "",
						Valid:  false,
					},
					IPAddress:     session.IPAddress,
					UserAgent:     session.UserAgent,
					Payload:       sessionPayload.SessionEncode(),
					LastActivity:  time.Now().UnixNano() / int64(time.Millisecond),
					RememberToken: sessionData.RememberToken,
					AppID:         sessionData.AppID,
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

			sessionData, err := repository.FindOneByAppID(appID[0])
			if err == nil {
				err := repository.Update(model.Session{
					Model: model.Model{
						ID:       sessionData.ID,
						Metadata: sessionData.Metadata,
					},
					UserID: sessionData.UserID,
					IPAddress: sql.NullString{
						String: ipAddress,
						Valid:  true,
					},
					UserAgent:     sessionData.UserAgent,
					Payload:       sessionPayload.SessionEncode(),
					LastActivity:  time.Now().UnixNano() / int64(time.Millisecond),
					RememberToken: sessionData.RememberToken,
					AppID:         sessionData.AppID,
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

		c.Locals("app_id", appID[0])
		return c.Next()
	}
}
