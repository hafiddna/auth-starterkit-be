package router

import (
	validator2 "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
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
	"log"
	"strings"
	"time"
)

func SetupRoutes(app *fiber.App) {
	var err error

	// Mongo
	mongo, err := database.ConnectToMongo()
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Postgres
	postgres, err := database.ConnectToPostgres()
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v", err)
	}

	// Minio
	minio, err := database.ConnectToMinio()
	if err != nil {
		log.Fatalf("Error connecting to Minio: %v", err)
	}

	// Validator
	validator := validator2.New()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			allowedOrigins := strings.Split(config.Config.App.Server.Cors, ",")
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					return true
				}
			}

			return false
		},
		AllowMethods: "GET, POST, PATCH, PUT, DELETE",
		AllowHeaders: "Content-Type, Content-Length, Host, User-Agent, Accept, Accept-Encoding, Connection, Authorization, Accept-Language, X-App-Id, X-Device-Category, X-Device-Type, Referer",
		// TODO: Limit ExposeHeaders to only the necessary headers
		//ExposeHeaders: "",
	}))
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(favicon.New(favicon.Config{
		File: "./asset/favicon.ico",
		URL:  "/favicon.ico",
	}))
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
	sessionRepository := repository.NewSessionRepository(postgres)
	roleUserRepository := repository.NewRoleUserRepository(postgres)
	userProfileRepository := repository.NewUserProfileRepository(mongo)
	userRepository := repository.NewUserRepository(postgres, minio)
	userSettingRepository := repository.NewUserSettingRepository(mongo)
	permissionRepository := repository.NewPermissionRepository(postgres)
	roleRepository := repository.NewRoleRepository(postgres)

	// Service
	sessionService := service.NewSessionService(sessionRepository)
	userService := service.NewUserService(userRepository, userProfileRepository, userSettingRepository, roleUserRepository)
	authService := service.NewAuthService(userService, sessionService)
	permissionService := service.NewPermissionService(permissionRepository)
	roleService := service.NewRoleService(roleRepository) // TODO: Check

	// Controller
	authController := controller.NewAuthController(authService, sessionService, validator)
	persmissionController := controller.NewPermissionController(permissionService, validator)
	roleController := controller.NewRoleController(roleService, validator) // TODO: Check
	userController := controller.NewUserController(userService, validator)

	// More Middleware
	app.Use(middleware.ActivityMiddleware(sessionRepository))

	// Routes
	// Start::Public Routes
	public := app.Group("/api")

	// Auth
	//public.Post("/register", authController.Register)
	public.Post("/login", authController.Login) // TODO: Check
	public.Post("/refresh-token", authController.RefreshToken)
	//public.Post("/forgot-password", authController.ForgotPassword)
	//public.Patch("/reset-password", authController.ResetPassword)
	// End::Public Routes

	// Start::Private Routes
	private := app.Group("/api")
	private.Use(middleware.AuthMiddleware())

	// Auth
	private.Get("/profile", authController.GetProfile) // TODO: Check
	//private.Get("/detail-profile", authController.GetDetailProfile)
	//private.Patch("/profile", authController.UpdateProfile)
	//private.Patch("/change-password", authController.ChangePassword)
	//private.Patch("/change-pin", authController.ChangePin)
	//private.Patch("/users/:id/account-activation", userController.AccountActivation)???
	//private.Get("/verify-email/{id}/{hash}", authController.VerifyEmail)???
	private.Post("/logout", authController.Logout)

	// Permission
	private.Get("/permissions", middleware.PermissionAuthMiddleware([]string{"read:permissions"}), persmissionController.GetAll)
	private.Post("/permissions", middleware.PermissionAuthMiddleware([]string{"write:permissions"}), persmissionController.Create)
	private.Get("/permissions/:id", middleware.PermissionAuthMiddleware([]string{"read:permissions"}), persmissionController.Get)
	private.Put("/permissions/:id", middleware.PermissionAuthMiddleware([]string{"write:permissions"}), persmissionController.Update)
	private.Delete("/permissions/:id", middleware.PermissionAuthMiddleware([]string{"delete:permissions"}), persmissionController.Delete)
	private.Delete("/permissions/:id/permanent", middleware.PermissionAuthMiddleware([]string{"delete:permissions"}), persmissionController.DeletePermanent)

	// Role
	private.Get("/roles", middleware.PermissionAuthMiddleware([]string{"read:roles"}), roleController.GetAll)
	private.Post("/roles", middleware.PermissionAuthMiddleware([]string{"write:roles"}), roleController.Create)
	private.Get("/roles/:id", middleware.PermissionAuthMiddleware([]string{"read:roles"}), roleController.Get)
	private.Put("/roles/:id", middleware.PermissionAuthMiddleware([]string{"write:roles"}), roleController.Update)
	private.Delete("/roles/:id", middleware.PermissionAuthMiddleware([]string{"delete:roles"}), roleController.Delete)
	private.Delete("/roles/:id/permanent", middleware.PermissionAuthMiddleware([]string{"delete:roles"}), roleController.DeletePermanent)

	// User
	private.Get("/users", middleware.PermissionAuthMiddleware([]string{"read:users"}), userController.GetAll)
	private.Post("/users", middleware.PermissionAuthMiddleware([]string{"write:users"}), userController.Create)
	private.Get("/users/:id", middleware.PermissionAuthMiddleware([]string{"read:users"}), userController.Get)
	private.Put("/users/:id", middleware.PermissionAuthMiddleware([]string{"write:users"}), userController.Update)
	private.Delete("/users/:id", middleware.PermissionAuthMiddleware([]string{"delete:users"}), userController.Delete)
	private.Delete("/users/:id/permanent", middleware.PermissionAuthMiddleware([]string{"delete:users"}), userController.DeletePermanent)
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
}
