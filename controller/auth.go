package controller

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
	Refresh(c *fiber.Ctx) error
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

	aToken, err := helper.ValidateRS512Token(config.Config.App.JWT.PublicKey, accessToken)
	if accessToken != "" && err == nil && aToken.Valid {
		responseData["access_token"] = accessToken
		sessionPayload.Token.AccessToken = accessToken
	} else {
		loginTokens, err := a.authService.Login(user)
		if loginTokens == "" || err != nil {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Your account is not active",
			})
		}

		responseData["access_token"] = loginTokens
		sessionPayload.Token.AccessToken = loginTokens
	}

	if *loginDto.Remember {
		rToken, err := helper.ValidateRS512Token(config.Config.App.JWT.RememberTokenPublic, refreshToken)
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

func (a *authController) Refresh(c *fiber.Ctx) error {
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

	return nil
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
	// TODO: Remove session here
	return nil
}
