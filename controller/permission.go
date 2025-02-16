package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hafiddna/auth-starterkit-be/service"
)

type PermissionController interface {
	GetAll(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	DeletePermanent(ctx *fiber.Ctx) error
}

type permissionController struct {
	permissionService service.PermissionService
	validator         *validator.Validate
}

func NewPermissionController(permissionService service.PermissionService, validator *validator.Validate) PermissionController {
	return &permissionController{
		permissionService: permissionService,
		validator:         validator,
	}
}

func (c *permissionController) GetAll(ctx *fiber.Ctx) error {
	return nil
}

func (c *permissionController) Create(ctx *fiber.Ctx) error {
	return nil
}

func (c *permissionController) Get(ctx *fiber.Ctx) error {
	return nil
}

func (c *permissionController) Update(ctx *fiber.Ctx) error {
	return nil
}

func (c *permissionController) Delete(ctx *fiber.Ctx) error {
	return nil
}

func (c *permissionController) DeletePermanent(ctx *fiber.Ctx) error {
	return nil
}
