package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hafiddna/auth-starterkit-be/service"
)

type UserController interface {
	GetAll(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	DeletePermanent(ctx *fiber.Ctx) error
}

type userController struct {
	userService service.UserService
	validator   *validator.Validate
}

func NewUserController(userService service.UserService, validator *validator.Validate) UserController {
	return &userController{
		userService: userService,
		validator:   validator,
	}
}

func (c *userController) GetAll(ctx *fiber.Ctx) error {
	return nil
}

func (c *userController) Create(ctx *fiber.Ctx) error {
	return nil
}

func (c *userController) Get(ctx *fiber.Ctx) error {
	return nil
}

func (c *userController) Update(ctx *fiber.Ctx) error {
	return nil
}

func (c *userController) Delete(ctx *fiber.Ctx) error {
	return nil
}

func (c *userController) DeletePermanent(ctx *fiber.Ctx) error {
	return nil
}
