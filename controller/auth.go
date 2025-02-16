package controller

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/hafiddna/auth-starterkit-be/dto"
	"github.com/hafiddna/auth-starterkit-be/helper"
	"github.com/hafiddna/auth-starterkit-be/model"
	"github.com/hafiddna/auth-starterkit-be/service"
	"reflect"
	"time"
)

type AuthController interface {
	Login(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
	GetProfile(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
}

type authController struct {
	authService    service.AuthService
	sessionService service.SessionService
	validator      *validator.Validate
}

func NewAuthController(authService service.AuthService, sessionService service.SessionService, validator *validator.Validate) AuthController {
	return &authController{
		authService:    authService,
		sessionService: sessionService,
		validator:      validator,
	}
}

func (a *authController) Login(c *fiber.Ctx) error {
	// TODO: Differentiate between login web and login mobile
	// TODO: Add rate limiter
	// TODO: Add MFA (Multi-Factor Authentication)
	// TODO: Add OAuth2 (https://chatgpt.com/c/67b15eba-82e4-8002-bebb-e4da309ddccc)
	var loginDto dto.LoginDTO
	if err := c.BodyParser(&loginDto); err != nil {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusBadRequest,
			Message:    "Bad Request",
			Error:      err.Error(),
		})
	}

	if err := a.validator.Struct(loginDto); err != nil {
		body := reflect.TypeOf(loginDto)
		errorMessages := helper.Validate(body, err)

		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusBadRequest,
			Message:    "Bad Request",
			Error:      errorMessages,
		})
	}

	user, err := a.authService.ValidateUser(loginDto)
	if err != nil {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}

	appID := c.Get("X-App-Id")
	sessionData, err := a.sessionService.FindOneByAppID(appID)
	if err != nil {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusNotFound,
			Message:    "Not Found",
		})
	}

	responseData := make(map[string]interface{})

	var oldSessionPayload model.SessionPayload
	var sessionPayload model.SessionPayload
	oldSessionPayload.SessionDecode(sessionData.Payload)
	accessToken := oldSessionPayload.Token.AccessToken
	refreshToken := oldSessionPayload.Token.RefreshToken

	aToken, err := helper.ValidateRS512Token(config.Config.App.JWT.PublicKey, accessToken, true)
	if accessToken != "" && err == nil && aToken.Valid {
		responseData["access_token"] = accessToken
		sessionPayload.Token.AccessToken = accessToken
	} else {
		newAccessToken, err := a.authService.Login(user)
		if newAccessToken == "" || err != nil {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Your account is not active",
			})
		}

		responseData["access_token"] = newAccessToken
		sessionPayload.Token.AccessToken = newAccessToken
	}

	if *loginDto.Remember {
		rToken, err := helper.ValidateRS512Token(config.Config.App.JWT.RememberTokenPublic, refreshToken, true)
		if refreshToken != "" && err == nil && rToken.Valid {
			responseData["refresh_token"] = refreshToken
			sessionPayload.Token.RefreshToken = refreshToken
		} else {
			rememberTokenDuration := time.Now().Add(time.Hour * 24)
			rememberData := helper.JwtRememberClaim{
				RememberToken: sessionData.RememberToken.String,
			}
			rememberAccessToken := helper.GenerateRS512Token(config.Config.App.JWT.RememberTokenPrivate, config.Config.App.Secret.RememberTokenKey, user.ID, rememberData, rememberTokenDuration)
			responseData["refresh_token"] = rememberAccessToken
			sessionPayload.Token.RefreshToken = rememberAccessToken
		}
	} else {
		sessionPayload.Token.RefreshToken = oldSessionPayload.Token.RefreshToken
	}

	sessionPayload.Previous = oldSessionPayload.Previous

	sessionData.Payload = sessionPayload.SessionEncode()
	sessionData.UserID = sql.NullString{
		String: user.ID,
		Valid:  true,
	}
	sessionData.LastActivity = time.Now().UnixNano() / int64(time.Millisecond)
	err = a.sessionService.Update(sessionData)
	if err != nil {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal Server Error",
			Error:      err.Error(),
		})
	}

	return helper.SendResponse(helper.ResponseStruct{
		Ctx:        c,
		StatusCode: fiber.StatusOK,
		Message:    "Success",
		Data:       responseData,
	})
}

func (a *authController) RefreshToken(c *fiber.Ctx) error {
	var refreshDTO dto.RefreshDTO
	if err := c.BodyParser(&refreshDTO); err != nil {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusBadRequest,
			Message:    "Bad Request",
			Error:      err.Error(),
		})
	}

	if err := a.validator.Struct(refreshDTO); err != nil {
		body := reflect.TypeOf(refreshDTO)
		errorMessages := helper.Validate(body, err)

		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusBadRequest,
			Message:    "Bad Request",
			Error:      errorMessages,
		})
	}

	rToken, err := helper.ValidateRS512Token(config.Config.App.JWT.RememberTokenPublic, refreshDTO.RefreshToken, true)
	if err != nil {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}

	if !rToken.Valid {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}

	mapStringClaims := make(map[string]interface{})
	for key, value := range rToken.Claims.(jwt.MapClaims) {
		mapStringClaims[key] = value
	}

	appID := c.Get("X-App-Id")
	sessionData, err := a.sessionService.FindOneByAppID(appID)
	if err != nil {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusNotFound,
			Message:    "Not Found",
		})
	}

	var encryptedData helper.EncryptedData
	tokenData := helper.JSONMarshal(mapStringClaims["data"])
	helper.JSONUnmarshal([]byte(tokenData), &encryptedData)
	decryptedData, err := helper.DecryptAES256CBC(&encryptedData, []byte(config.Config.App.Secret.RememberTokenKey))
	if err != nil {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal Server Error",
			Error:      err.Error(),
		})
	}

	rememberToken := helper.RandomString(10)
	responseData := make(map[string]interface{})

	var oldSessionPayload model.SessionPayload
	var sessionPayload model.SessionPayload
	oldSessionPayload.SessionDecode(sessionData.Payload)

	var rememberTokenPayload helper.JwtRememberClaim
	helper.JSONUnmarshal([]byte(decryptedData), &rememberTokenPayload)
	if sessionData.RememberToken.String != rememberTokenPayload.RememberToken {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}

	user, err := a.authService.GetTokenData(sessionData.UserID.String)

	newAccessToken, err := a.authService.Login(user)
	if newAccessToken == "" || err != nil {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}
	responseData["access_token"] = newAccessToken
	sessionPayload.Token.AccessToken = newAccessToken

	rememberTokenDuration := time.Now().Add(time.Hour * 24)
	rememberData := helper.JwtRememberClaim{
		RememberToken: rememberToken,
	}
	rememberAccessToken := helper.GenerateRS512Token(config.Config.App.JWT.RememberTokenPrivate, config.Config.App.Secret.RememberTokenKey, sessionData.User.ID, rememberData, rememberTokenDuration)
	responseData["refresh_token"] = rememberAccessToken
	sessionPayload.Token.RefreshToken = rememberAccessToken

	sessionPayload.Previous = oldSessionPayload.Previous
	sessionData.Payload = sessionPayload.SessionEncode()
	sessionData.RememberToken = sql.NullString{
		String: rememberToken,
		Valid:  true,
	}
	sessionData.LastActivity = time.Now().UnixNano() / int64(time.Millisecond)
	err = a.sessionService.Update(sessionData)
	if err != nil {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal Server Error",
			Error:      err.Error(),
		})
	}

	return helper.SendResponse(helper.ResponseStruct{
		Ctx:        c,
		StatusCode: fiber.StatusOK,
		Message:    "Success",
		Data:       responseData,
	})
}

func (a *authController) GetProfile(c *fiber.Ctx) error {
	userId, isString := c.Locals("user").(map[string]interface{})["sub"].(string)
	if !isString {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})

	}

	user, err := a.authService.Profile(userId)
	if err != nil {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusNotFound,
			Message:    "Not Found",
		})
	}

	return helper.SendResponse(helper.ResponseStruct{
		Ctx:        c,
		StatusCode: fiber.StatusOK,
		Message:    "Success",
		Data:       user,
	})
}

func (a *authController) Logout(c *fiber.Ctx) error {
	appID := c.Get("X-App-Id")
	sessionData, err := a.sessionService.FindOneByAppID(appID)
	if err != nil {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusNotFound,
			Message:    "Not Found",
		})
	}

	var sessionPayload model.SessionPayload
	var oldSessionPayload model.SessionPayload
	oldSessionPayload.SessionDecode(sessionData.Payload)

	sessionPayload.Token.AccessToken = ""
	sessionPayload.Token.RefreshToken = ""
	sessionPayload.Previous = oldSessionPayload.Previous

	sessionData.Payload = sessionPayload.SessionEncode()
	sessionData.UserID = sql.NullString{
		String: "",
		Valid:  false,
	}
	go a.sessionService.Update(sessionData)

	return helper.SendResponse(helper.ResponseStruct{
		Ctx:        c,
		StatusCode: fiber.StatusOK,
		Message:    "Success",
	})
}
