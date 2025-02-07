package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/hafiddna/auth-starterkit-be/helper"
	"github.com/hafiddna/auth-starterkit-be/router"
	"log"
	"time"
)

func main() {
	var err error

	// Timezone
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	time.Local = utc

	// Config
	config.Config, err = config.GetConfig()
	if err != nil {
		log.Fatalf("Error getting config: %v", err)
	}

	// Licensing
	if err = helper.InitApp(); err != nil {
		panic(err)
	}

	// Fiber
	app := fiber.New(fiber.Config{
		Prefork:       config.Config.App.Environment == "production",
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  config.Config.App.ServerName,
		AppName:       config.Config.App.Name,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusBadRequest,
				Message:    "Bad Request",
				Error:      err.Error(),
			})
		},
	})

	// Router
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":" + config.Config.App.Server.Port))
}
