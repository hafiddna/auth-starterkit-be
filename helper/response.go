package helper

import (
	"encoding/xml"
	"github.com/gofiber/fiber/v2"
	"github.com/hafiddna/auth-starterkit-be/config"
)

type ResponseStruct struct {
	Ctx        *fiber.Ctx
	StatusCode int
	Message    string
	Error      interface{}
	Data       interface{}
}

type BaseResponse struct {
	XMLName    xml.Name    `json:"-" xml:"response"`
	StatusCode int         `json:"status_code" xml:"status_code"`
	Message    string      `json:"message,omitempty" xml:"message"`
	Error      interface{} `json:"error,omitempty" xml:"error"`
	Data       interface{} `json:"data,omitempty" xml:"data"`
}

func SendResponse(baseResponse ResponseStruct) error {
	newBaseResponse := BaseResponse{
		StatusCode: baseResponse.StatusCode,
	}

	if baseResponse.Message != "" {
		newBaseResponse.Message = baseResponse.Message
	}

	if baseResponse.Error != nil {
		newBaseResponse.Error = baseResponse.Error
	}

	if baseResponse.Data != nil {
		if config.Config.App.Environment == "development" {
			newBaseResponse.Data = baseResponse.Data
		} else {
			// TODO: Masking/encoded compression data
		}
	}

	accept := baseResponse.Ctx.Get("Accept")
	if accept == "application/xml" {
		baseResponse.Ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationXML)
		baseResponse.Ctx.Status(baseResponse.StatusCode)
		return baseResponse.Ctx.XML(newBaseResponse)
	} else {
		baseResponse.Ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		baseResponse.Ctx.Status(baseResponse.StatusCode)
		return baseResponse.Ctx.JSON(newBaseResponse)
	}
}
