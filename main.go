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
	config := infrastructure.New() //get env
	app := fiber.New()
	var validator = validator.New()

	userRepo := repository.NewUserRepository(infrastructure.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService, validator)

	authService := service.NewAuthService(userRepo, config)
	authController := controller.NewAuthController(authService, validator)

	v1 := app.Group("/v1")
	v1.Get("/users", middleware.JWTProtect(), userController.GetAll)
	v1.Post("/users", userController.Create)
	v1.Patch("/users/:id", userController.Update)
	v1.Get("/users/:id", userController.GetById)
	v1.Delete("/users/:id", userController.Delete)

	v1.Post("/auth/login", authController.Login)

	app.Listen(":8081")
}
