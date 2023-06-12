package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"github.com/tangguhriyadi/user-service/controller"
	"github.com/tangguhriyadi/user-service/docs"
	"github.com/tangguhriyadi/user-service/infrastructure"
	"github.com/tangguhriyadi/user-service/repository"
	"github.com/tangguhriyadi/user-service/security/middleware"
	"github.com/tangguhriyadi/user-service/service"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	infrastructure.ConnectDB()
	infrastructure.RunGrpc()

	docs.SwaggerInfo.Title = "Content Service Dapur Santet"
	docs.SwaggerInfo.Description = "test"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "18.140.2.142:8000"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// config := infrastructure.New()
	app := fiber.New()
	var validator = validator.New()

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	userRepo := repository.NewUserRepository(infrastructure.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService, validator)

	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService, validator)

	app.Get("/documentation/*", fiberSwagger.WrapHandler)
	// v1 := app.Group("/v1")
	app.Get("/users", middleware.JWTProtect(), userController.GetAll)
	app.Post("/signup", userController.Create)
	app.Patch("/users/:id", middleware.JWTProtect(), userController.Update)
	app.Get("/users/:id", middleware.JWTProtect(), userController.GetById)
	app.Delete("/users/:id", middleware.JWTProtect(), userController.Delete)

	app.Post("/login", authController.Login)

	app.Listen(":8081")

	fmt.Println("running...")
}
