package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/user-service/controller"
	"github.com/tangguhriyadi/user-service/infrastructure"
	"github.com/tangguhriyadi/user-service/repository"
	"github.com/tangguhriyadi/user-service/security/middleware"
	"github.com/tangguhriyadi/user-service/service"
)

func main() {
	infrastructure.ConnectDB()
	// config := infrastructure.New()
	app := fiber.New()
	var validator = validator.New()

	userRepo := repository.NewUserRepository(infrastructure.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService, validator)

	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService, validator)

	// v1 := app.Group("/v1")
	app.Get("/users", middleware.JWTProtect(), userController.GetAll)
	app.Post("/users", userController.Create)
	app.Patch("/users/:id", userController.Update)
	app.Get("/users/:id", userController.GetById)
	app.Delete("/users/:id", userController.Delete)

	app.Post("/login", authController.Login)

	app.Listen(":8081")
}
