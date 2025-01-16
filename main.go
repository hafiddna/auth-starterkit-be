package main

import (
	validator2 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/hafiddna/auth-starterkit-be/controller"
	"github.com/hafiddna/auth-starterkit-be/database"
	"github.com/hafiddna/auth-starterkit-be/helper"
	"github.com/hafiddna/auth-starterkit-be/middleware"
	"github.com/hafiddna/auth-starterkit-be/repository"
	"github.com/hafiddna/auth-starterkit-be/service"
	"github.com/hafiddna/auth-starterkit-be/tool"
	"github.com/hafiddna/auth-starterkit-be/util"
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
	// TODO: Use Minio for file storage
	_, err = tool.ConnectToMinio()
	if err != nil {
		panic(err)
	}

	// Fiber
	app := fiber.New(fiber.Config{
		//Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       config.Config.App.Name,
		ServerHeader:  config.Config.App.ServerName,
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
	// TODO: Recheck this middleware
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: nil,
		AllowOrigins:     config.Config.App.Server.Cors,
		AllowMethods:     "GET, POST, PATCH, PUT, DELETE",
		AllowHeaders:     "Accept, Content-Type, Content-Length, Accept-Encoding, Accept-Language, X-CSRF-Token, Authorization, X-Requested-With, User-Agent, Connection, Host",
	}))
	app.Use(compress.New())
	app.Use(etag.New())
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
	app.Use(recover.New())
	app.Use(requestid.New())
	//app.Use(middleware.LoggerMiddleware)
	//file, err := os.OpenFile("./123.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalf("error opening file: %v", err)
	//}
	//defer file.Close()
	app.Use(logger.New(logger.Config{
		Format:     "Time: ${time} | Status: ${status} | PID: ${pid} | RequestID: ${locals:requestid} | Latency: ${latency} | IP: [${ip}]:${port} | Method: ${method} | Path: ${path} | Error: ${error}\n",
		TimeFormat: time.RFC3339Nano,
		TimeZone:   "UTC",
		Done: func(c *fiber.Ctx, logString []byte) {
			log.Println(string(logString))
			//if c.Response().StatusCode() != fiber.StatusOK {
			//reporter.SendToSlack(logString)
			//}
		},
		//CustomTags: map[string]logger.LogFunc{
		//	"custom_tag": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
		//		return output.WriteString("it is a custom tag")
		//	},
		//},
		//DisableColors: true,
		//Output: file,
	}))

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
	//public.Post("/register", authController.Register)
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
