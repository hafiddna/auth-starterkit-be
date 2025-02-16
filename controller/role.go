package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hafiddna/auth-starterkit-be/service"
)

type RoleController interface {
	GetAll(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	DeletePermanent(ctx *fiber.Ctx) error
}

type roleController struct {
	roleService service.RoleService
	validator   *validator.Validate
}

func NewRoleController(roleService service.RoleService, validator *validator.Validate) RoleController {
	return &roleController{
		roleService: roleService,
		validator:   validator,
	}
}

func (c *roleController) GetAll(ctx *fiber.Ctx) error {
	return nil
}

func (c *roleController) Create(ctx *fiber.Ctx) error {
	return nil
}

func (c *roleController) Get(ctx *fiber.Ctx) error {
	return nil
}

func (c *roleController) Update(ctx *fiber.Ctx) error {
	return nil
}

func (c *roleController) Delete(ctx *fiber.Ctx) error {
	return nil
}

func (c *roleController) DeletePermanent(ctx *fiber.Ctx) error {
	return nil
}
