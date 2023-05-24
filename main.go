package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/user-service/controller"
	"github.com/tangguhriyadi/user-service/infrastructure"
	"github.com/tangguhriyadi/user-service/repository"
	"github.com/tangguhriyadi/user-service/service"
)

func main() {
	infrastructure.ConnectDB()
	app := fiber.New()

	userRepo := repository.NewUserRepository(infrastructure.DB)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	v1 := app.Group("/v1")
	v1.Get("/users", userController.GetAll)
	v1.Post("/Users", userController.Create)

	app.Listen(":8081")
}
