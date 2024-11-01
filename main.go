package main

import (
	"context"
	validator2 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	config2 "github.com/hafiddna/auth-starterkit-be/config"
	"github.com/hafiddna/auth-starterkit-be/controller"
	"github.com/hafiddna/auth-starterkit-be/database"
	"github.com/hafiddna/auth-starterkit-be/helper"
	"github.com/hafiddna/auth-starterkit-be/middleware"
	"github.com/hafiddna/auth-starterkit-be/repository"
	"github.com/hafiddna/auth-starterkit-be/service"
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
	// Timezone
	utc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	time.Local = utc

	// Licensing
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

	// Start::Global Middleware
	app.Use(middleware.CORSMiddleware(config))
	app.Use(middleware.LoggerMiddleware)
	// End::Global Middleware

	// Start::Routes
	setUpGlobalRoutes()
	setUpPublicRoutes()
	setUpPrivateRoutes()
	// End::Routes

	app.Listen(":" + config.App.Server.Port)
}

var (
	// Start::Repository
	assetRepository = repository.NewAssetRepository(gormDB)
	// End::Repository

	// Start::Service
	jwtService     = service.NewJWTService(config)
	storageService = service.NewStorageService(minioClient, assetRepository)
	// End::Service

	// Start::Controller
	storageController = controller.NewStorageController(response, validator, storageService, minioClient, jwtService)
	// End::Controller
)

func setUpGlobalRoutes() {
	// Start::Not Found Handler
	app.Use(func(c *fiber.Ctx) error {
		return response.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusNotFound,
			Message:    "Not Found",
			Error:      "Cannot " + c.Method() + " " + c.Path(),
		})
	})
	// End::Not Found Handler
}

func setUpPublicRoutes() {
	public := app.Group("/api")

	// Start:Storage
	public.Get("/file", storageController.Get)
	// End:Storage
}

func setUpPrivateRoutes() {
	private := app.Group("/api")
	private.Use(middleware.AuthMiddleware(jwtService, response))

	// Start:Storage
	storage := private.Group("/file")
	storage.Post("/", storageController.Upload)
	storage.Delete("/", storageController.Delete)
	// End:Storage
}
