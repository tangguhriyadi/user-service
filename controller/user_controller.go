package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/user-service/model"
	"github.com/tangguhriyadi/user-service/service"
)

type UserController interface {
	GetAll(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (uc UserControllerImpl) GetAll(ctx *fiber.Ctx) error {
	c := ctx.Context()
	result, err := uc.userService.GetAll(c)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (uc UserControllerImpl) Create(ctx *fiber.Ctx) error {
	c := ctx.Context()

	var user model.Users

	//body parsing
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	//run business logic
	if err := uc.userService.Create(c, &user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "succes create",
	})

}
