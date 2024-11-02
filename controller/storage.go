package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hafiddna/auth-starterkit-be/dto"
	"github.com/hafiddna/auth-starterkit-be/entity"
	"github.com/hafiddna/auth-starterkit-be/helper"
	"github.com/hafiddna/auth-starterkit-be/service"
	"github.com/minio/minio-go/v7"
	"reflect"
	"time"
)

type StorageController interface {
	Get(c *fiber.Ctx) error
	Upload(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type storageController struct {
	response       helper.Response
	validator      *validator.Validate
	storageService service.StorageService
	minio          *minio.Client
	jwtService     service.JWTService
}

func NewStorageController(response helper.Response, validator *validator.Validate, storageService service.StorageService, minio *minio.Client, jwtService service.JWTService) StorageController {
	return &storageController{
		response:       response,
		validator:      validator,
		storageService: storageService,
		minio:          minio,
		jwtService:     jwtService,
	}
}

func (p storageController) Get(c *fiber.Ctx) error {
	var getOrDeleteStorageDto dto.GetOrDeleteStorageDto

	if err := c.QueryParser(&getOrDeleteStorageDto); err != nil {
		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusBadRequest,
			Message:    "Bad Request",
		})
	}

	if err := p.validator.Struct(getOrDeleteStorageDto); err != nil {
		body := reflect.TypeOf(getOrDeleteStorageDto)
		errorMessages := helper.Validate(body, err)

		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusBadRequest,
			Message:    "Bad Request",
			Error:      errorMessages,
		})
	}

	asset := p.storageService.FindByPath(getOrDeleteStorageDto)
	if asset == nil {
		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusNotFound,
			Message:    "Not Found",
		})
	} else {
		asset, isAsset := asset.(entity.Asset)
		if !isAsset {
			return p.response.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusInternalServerError,
				Message:    "Internal Server Error",
			})
		}

		if asset.BucketType == "public" {
			if asset.Access == "private" {
				authorization := c.Get("Authorization")
				if authorization == "" || len(authorization) < 7 {
					return p.response.SendResponse(helper.ResponseStruct{
						Ctx:        c,
						StatusCode: fiber.StatusUnauthorized,
						Message:    "Unauthorized",
					})
				}

				token := authorization[7:]
				aToken, err := p.jwtService.ValidateToken(token)
				if err != nil {
					return p.response.SendResponse(helper.ResponseStruct{
						Ctx:        c,
						StatusCode: fiber.StatusUnauthorized,
						Message:    "Unauthorized",
						Error:      err.Error(),
					})
				}

				claims := make(map[string]interface{})
				for key, value := range aToken.Claims.(jwt.MapClaims) {
					claims[key] = value
				}

				userId, isString := claims["sub"].(string)
				if !isString {
					// TODO: Implement owner type for organization/tenant
					return p.response.SendResponse(helper.ResponseStruct{
						Ctx:        c,
						StatusCode: fiber.StatusUnauthorized,
						Message:    "Unauthorized",
					})
				}

				if asset.OwnerID != "" && asset.OwnerID != userId {
					return p.response.SendResponse(helper.ResponseStruct{
						Ctx:        c,
						StatusCode: fiber.StatusForbidden,
						Message:    "Forbidden",
					})
				}
			}

			file, err := p.minio.GetObject(c.Context(), asset.BucketType, asset.Path, minio.GetObjectOptions{})
			if err != nil {
				return p.response.SendResponse(helper.ResponseStruct{
					Ctx:        c,
					StatusCode: fiber.StatusInternalServerError,
					Message:    "Internal Server Error",
				})
			}

			objectInfo, err := file.Stat()
			if err != nil {
				return p.response.SendResponse(helper.ResponseStruct{
					Ctx:        c,
					StatusCode: fiber.StatusInternalServerError,
					Message:    "Internal Server Error",
				})
			}

			c.Set("Content-Type", objectInfo.ContentType)
			c.Set("Content-Disposition", "attachment; filename="+asset.Name)
			c.Set("Connection", "close")
			c.Set("Vary", "Origin")
			c.Set("Cache-Control", "no-cache, private")
			return c.SendStream(file)
		} else {
			presignedURL, err := p.minio.PresignedGetObject(c.Context(), asset.BucketType, asset.Path, time.Second*60, nil)
			if err != nil {
				return p.response.SendResponse(helper.ResponseStruct{
					Ctx:        c,
					StatusCode: fiber.StatusInternalServerError,
					Message:    "Internal Server Error",
				})
			}

			return c.Redirect(presignedURL.String(), fiber.StatusTemporaryRedirect)
		}
	}
}

func (p storageController) Upload(c *fiber.Ctx) error {
	var createStorageDto dto.CreateStorageDto

	if err := c.BodyParser(&createStorageDto); err != nil {
		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusBadRequest,
			Message:    "Bad Request",
		})
	}

	if err := p.validator.Struct(createStorageDto); err != nil {
		body := reflect.TypeOf(createStorageDto)
		errorMessages := helper.Validate(body, err)

		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusBadRequest,
			Message:    "Bad Request",
			Error:      errorMessages,
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusBadRequest,
			Message:    "Bad Request",
		})
	}

	if asset, err := p.storageService.Upload(c, createStorageDto, file); err != nil {
		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal Server Error",
		})
	} else {
		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusOK,
			Message:    "Uploaded",
			Data:       map[string]interface{}{"object": asset.Path},
		})
	}
}

func (p storageController) Delete(c *fiber.Ctx) error {
	var getOrDeleteStorageDto dto.GetOrDeleteStorageDto

	if err := c.QueryParser(&getOrDeleteStorageDto); err != nil {
		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusBadRequest,
			Message:    "Bad Request",
		})
	}

	if err := p.validator.Struct(getOrDeleteStorageDto); err != nil {
		body := reflect.TypeOf(getOrDeleteStorageDto)
		errorMessages := helper.Validate(body, err)

		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusBadRequest,
			Message:    "Bad Request",
			Error:      errorMessages,
		})
	}

	asset := p.storageService.FindByPath(getOrDeleteStorageDto)
	if asset == nil {
		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusNotFound,
			Message:    "Not Found",
		})
	}

	assetData, isAsset := asset.(entity.Asset)
	if !isAsset {
		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal Server Error",
		})
	}

	authorization := c.Get("Authorization")
	if authorization == "" || len(authorization) < 7 {
		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}

	token := authorization[7:]
	aToken, err := p.jwtService.ValidateToken(token)
	if err != nil {
		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
			Error:      err.Error(),
		})
	}

	claims := make(map[string]interface{})
	for key, value := range aToken.Claims.(jwt.MapClaims) {
		claims[key] = value
	}

	userId, isString := claims["sub"].(string)
	if !isString {
		// TODO: Implement owner type for organization/tenant
		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusUnauthorized,
			Message:    "Unauthorized",
		})
	}

	if assetData.OwnerID != "" && assetData.OwnerID != userId {
		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusForbidden,
			Message:    "Forbidden",
		})
	}

	if err := p.storageService.Delete(assetData); err != nil {
		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal Server Error",
		})
	} else {
		return p.response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusOK,
			Message:    "Deleted",
		})
	}
}
