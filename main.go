package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/user-service/controller"
	"github.com/tangguhriyadi/user-service/infrastructure"
	"github.com/tangguhriyadi/user-service/repository"
	"github.com/tangguhriyadi/user-service/service"
)

func main() {
	infrastructure.ConnectDB()
	app := fiber.New()
	var validator = validator.New()
	userRepo := repository.NewUserRepository(infrastructure.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService, validator)

	v1 := app.Group("/v1")
	v1.Get("/users", userController.GetAll)
	v1.Post("/users", userController.Create)
	v1.Patch("/users/:id", userController.Update)

	app.Listen(":8081")
}
