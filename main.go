package main

import (
	"context"
	validator2 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	config2 "github.com/hafiddna/auth-starterkit-be/config"
	"github.com/hafiddna/auth-starterkit-be/database"
	"github.com/hafiddna/auth-starterkit-be/helper"
	"github.com/hafiddna/auth-starterkit-be/tool"
	"github.com/hafiddna/auth-starterkit-be/util"
	"time"
)

var (
	config = config2.NewConfig().GetConfig()
	app    = fiber.New(fiber.Config{
		AppName: config.App.Name,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return helper.NewResponse(config).SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusBadRequest,
				Message:    "Bad Request",
				Error:      err.Error(),
			})
		},
	})

	ctx                       = context.Background()
	redisDB                   = database.NewRedis(config, ctx, 15)
	mongoClient, mongoErr     = database.NewMongoDB(config, ctx).Connect()
	mongoDB                   = mongoClient.Database(config.App.MongoDB.Database)
	sqlDB, gormDB, postgreErr = database.NewPostgreSQL(config).Connect(config.App.PostgreSQL.Database)
	minioClient               = tool.NewMinioTool(config).Connect()

	response  = helper.NewResponse(config)
	validator = validator2.New()
)

func main() {
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	time.Local = utc

	if err = util.NewLicensing(config).InitApp(); err != nil {
		panic(err)
	}

	// MongoDB
	defer database.NewMongoDB(config, ctx).Disconnect(mongoClient)

	// PostgreSQL
	defer database.NewPostgreSQL(config).Disconnect(sqlDB)

	// Minio
	if minioClient == nil {
		panic("Minio connection failed")
	}

	//setUpRoutes()
	app.Listen(":" + config.App.Server.Port)
}
