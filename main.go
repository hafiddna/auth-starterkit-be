package main

import (
	validator2 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/hafiddna/auth-starterkit-be/controller"
	"github.com/hafiddna/auth-starterkit-be/database"
	"github.com/hafiddna/auth-starterkit-be/helper"
	"github.com/hafiddna/auth-starterkit-be/middleware"
	"github.com/hafiddna/auth-starterkit-be/repository"
	"github.com/hafiddna/auth-starterkit-be/service"
	"github.com/hafiddna/auth-starterkit-be/util"
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
		panic(err)
	}

	// Licensing
	if err = util.InitApp(); err != nil {
		panic(err)
	}

	// MongoDB
	mongoDB, err := database.ConnectToMongoDB()
	if err != nil {
		panic(err)
	}

	// PostgreSQL
	postgreSQL, err := database.ConnectToPostgreSQL()
	if err != nil {
		panic(err)
	}

	// Minio
	// TODO: Enable Minio on Creating File Upload Feature
	//minio, err := tool.ConnectToMinio()
	//if err != nil {
	//	panic(err)
	//}

	// Fiber
	app := fiber.New(fiber.Config{
		//Prefork:       true,
		//CaseSensitive: true,
		//StrictRouting: true,
		AppName:      config.Config.App.Name,
		ServerHeader: config.Config.App.ServerName,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return helper.SendResponse(helper.ResponseStruct{
				Ctx:        c,
				StatusCode: fiber.StatusBadRequest,
				Message:    "Bad Request",
				Error:      err.Error(),
			})
		},
	})

	// Validator
	validator := validator2.New()

	// Global Middleware
	app.Use(middleware.CORSMiddleware())
	app.Use(middleware.LoggerMiddleware)
	//app.Use(cors.New())
	//app.Use(compress.New())
	//app.Use(etag.New())
	//app.Use(favicon.New())
	//app.Use(limiter.New(limiter.Config{
	//	Max: 100,
	//	LimitReached: func(c *fiber.Ctx) error {
	//		return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
	//			"status":  "fail",
	//			"message": "You have requested too many in a single time-frame! Please wait another minute!",
	//		})
	//	},
	//}))
	//app.Use(logger.New())
	//app.Use(recover.New())
	//app.Use(requestid.New())

	// Repository
	roleUserRepository := repository.NewRoleUserRepository(postgreSQL)
	sessionRepository := repository.NewSessionRepository(postgreSQL)
	userProfileRepository := repository.NewUserProfileRepository(mongoDB)
	userRepository := repository.NewUserRepository(postgreSQL)
	userSettingRepository := repository.NewUserSettingRepository(mongoDB)

	// Service
	jwtService := service.NewJWTService()
	sessionService := service.NewSessionService(sessionRepository)
	userService := service.NewUserService(userRepository, userProfileRepository, userSettingRepository, roleUserRepository)
	authService := service.NewAuthService(userService, jwtService)

	// Controller
	authController := controller.NewAuthController(authService, sessionService, validator)

	// Routes
	// Start::Public Routes
	public := app.Group("/api")

	// Auth
	public.Post("/login", authController.Login)
	public.Post("/refresh-token", authController.Refresh)
	// End::Public Routes

	// Start::Private Routes
	private := app.Group("/api")
	private.Use(middleware.AuthMiddleware(jwtService))

	// Auth
	private.Get("/profile", authController.GetProfile)
	private.Post("/logout", authController.Logout)
	//private.Patch("/users/:id/account-activation", userController.AccountActivation)
	//private.Patch("/users/:id/assign-roles", userController.AssignRoles)
	// End::Private Routes

	// Not Found
	app.Use(func(c *fiber.Ctx) error {
		return helper.SendResponse(helper.ResponseStruct{
			Ctx:        c,
			StatusCode: fiber.StatusNotFound,
			Message:    "Not Found",
			Error:      "Cannot " + c.Method() + " " + c.Path(),
		})
	})

	app.Listen(":" + config.Config.App.Server.Port)
}
