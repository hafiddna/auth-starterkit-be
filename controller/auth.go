package controller

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hafiddna/auth-starterkit-be/dto"
	"github.com/hafiddna/auth-starterkit-be/entity"
	"github.com/hafiddna/auth-starterkit-be/helper"
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
	response       helper.Response
	validator      *validator.Validate
	authService    service.AuthService
	sessionService service.SessionService
}

func NewAuthController(response helper.Response, validator *validator.Validate, authService service.AuthService, sessionService service.SessionService) AuthController {
	return &authController{
		response:       response,
		validator:      validator,
		authService:    authService,
		sessionService: sessionService,
	}
}

func (a *authController) Login(c *fiber.Ctx) error {
	var loginDto dto.LoginDto
	if err := c.BodyParser(&loginDto); err != nil {
		return a.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusBadRequest,
			Message:    "Bad Request",
			Error:      err.Error(),
		})
	}

	if err := a.validator.Struct(loginDto); err != nil {
		body := reflect.TypeOf(loginDto)
		errorMessages := helper.Validate(body, err)

		return a.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusBadRequest,
			Message:    "Bad Request",
			Error:      errorMessages,
		})
	}

	user := a.authService.ValidateUser(loginDto)
	if user == nil {
		return a.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}

	return a.response.SendResponse(helper.ResponseStruct{
		Ctx:        c,
		StatusCode: fiber.StatusUnauthorized,
		Message:    "Unauthorized",
	})

	accessToken := a.authService.Login(user.(entity.User))
	if accessToken == nil {
		return a.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Your account has been deactivated.",
		})
	}

	userAgent := c.Get("User-Agent")
	userId := user.(entity.User).ID
	ipAddress := c.IP()
	go a.sessionService.Create(entity.Session{
		UserID: sql.NullString{
			String: userId,
			Valid:  true,
		},
		IPAddress: sql.NullString{
			String: ipAddress,
			Valid:  true,
		},
		UserAgent: sql.NullString{
			String: userAgent,
			Valid:  true,
		},
		Payload:      "",
		LastActivity: time.Now().Unix(),
	})

	return a.response.SendResponse(helper.ResponseStruct{
		Ctx:        c,
		StatusCode: fiber.StatusOK,
		Message:    "Success",
		Data:       map[string]interface{}{"access_token": accessToken},
	})
}

func (a *authController) Refresh(c *fiber.Ctx) error {
	return nil
}

func (a *authController) GetProfile(c *fiber.Ctx) error {
	userId, isString := c.Locals("user").(map[string]interface{})["sub"].(string)
	if !isString {
		return a.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})

	}

	// TODO: Update session last activity here

	user := a.authService.Profile(userId)

	return a.response.SendResponse(helper.ResponseStruct{
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
