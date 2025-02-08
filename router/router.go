package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/hafiddna/auth-starterkit-be/config"
	"github.com/hafiddna/auth-starterkit-be/database"
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
	validator := validator.New()

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
		AllowMethods:     "GET, POST, PATCH, PUT, DELETE",
		AllowHeaders:     "Accept, Content-Type, Content-Length, Accept-Encoding, Accept-Language, X-CSRF-Token, Authorization, X-Requested-With, User-Agent, Connection, Host",
		AllowCredentials: true,
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

	// Service

	// Controller

	//Route::middleware('guest')->group(function () {
	//	//    Route::get('register', [RegisteredUserController::class, 'create'])->name('register');
	//
	//	//    Route::post('register', [RegisteredUserController::class, 'store']);
	//
	//Route::get('login', [AuthenticatedSessionController::class, 'create'])->name('login');
	//
	//Route::post('login', [AuthenticatedSessionController::class, 'store']);
	//
	////    Route::get('forgot-password', [PasswordResetLinkController::class, 'create'])->name('password.request');
	//
	////    Route::post('forgot-password', [PasswordResetLinkController::class, 'store'])->name('password.email');
	//
	////    Route::get('reset-password/{token}', [NewPasswordController::class, 'create'])->name('password.reset');
	//
	////    Route::post('reset-password', [NewPasswordController::class, 'store'])->name('password.update');
	//});
	//
	//Route::middleware('auth')->group(function () {
	////    Route::get('verify-email', [EmailVerificationPromptController::class, '__invoke'])->name('verification.notice');
	//
	////    Route::get('verify-email/{id}/{hash}', [VerifyEmailController::class, '__invoke'])->middleware(['signed', 'throttle:6,1'])->name('verification.verify');
	//
	////    Route::post('email/verification-notification', [EmailVerificationNotificationController::class, 'store'])->middleware('throttle:6,1')->name('verification.send');
	//
	////    Route::get('confirm-password', [ConfirmablePasswordController::class, 'show'])->name('password.confirm');
	//
	////    Route::post('confirm-password', [ConfirmablePasswordController::class, 'store']);
	//
	//Route::post('logout', [AuthenticatedSessionController::class, 'destroy'])->name('logout');
	//});
}
