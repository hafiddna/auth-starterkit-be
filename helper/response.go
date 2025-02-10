package helper

import (
	"encoding/xml"
	"github.com/gofiber/fiber/v2"
	"github.com/hafiddna/auth-starterkit-be/config"
	"log"
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

func SendResponse(baseResponse ResponseStruct) (err error) {
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
			marshalledData := JSONMarshal(baseResponse.Data)

			newBaseResponse.Data, err = EncryptAES256CBC([]byte(marshalledData), []byte(config.Config.App.Secret.DataEncryptionKey))
			if err != nil {
				log.Fatalf("Error encrypting data: %v", err)
			}
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
