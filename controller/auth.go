package controller

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

	accessToken := a.authService.Login(user)
	if accessToken == nil {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Your account has been deactivated.",
		})
	}

	userAgent := c.Get("User-Agent")
	userId := user.ID
	ipAddress := c.IP()
	go a.sessionService.CreateOrUpdate(model.Session{
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
		Payload: "",
		// TODO: Update last activity shouldn't just be in here
		LastActivity: time.Now().Unix(),
	})

	return helper.SendResponse(helper.ResponseStruct{
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
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})

	}

	// TODO: Update session last activity here

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
