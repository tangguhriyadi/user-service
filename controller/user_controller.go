package controller

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/user-service/dto"
	"github.com/tangguhriyadi/user-service/model"
	"github.com/tangguhriyadi/user-service/security/token"
	"github.com/tangguhriyadi/user-service/service"
)

type UserController interface {
	GetAll(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	GetById(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type UserControllerImpl struct {
	userService service.UserService
	validate    *validator.Validate
}

func NewUserController(userService service.UserService, validate *validator.Validate) UserController {
	return &UserControllerImpl{
		userService: userService,
		validate:    validate,
	}
}

func (uc UserControllerImpl) GetAll(ctx *fiber.Ctx) error {
	c := ctx.Context()

	claims, err := token.ExtractTokenMetada(ctx)

	fmt.Println(claims)

	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "query page must be int",
		})
	}

	limit, err := strconv.Atoi(ctx.Query("limit", "10"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "query limit must be int",
		})
	}

	result, err := uc.userService.GetAll(c, page, limit)

	if result.Data == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed to retrieve data",
		})
	}

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
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

	// request body validation
	if err := uc.validate.Struct(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	//run business logic
	if err := uc.userService.Create(c, &user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "succes create",
	})

}

func (uc UserControllerImpl) Update(ctx *fiber.Ctx) error {
	c := ctx.Context()

	var user model.Users

	var userUpdate dto.UserUpdate

	userId := ctx.Params("id")

	//body parsing
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// request body validation
	if err := uc.validate.Struct(&userUpdate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	//run business logic
	if err := uc.userService.Update(c, userId, &user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "succes update",
	})
}

func (uc UserControllerImpl) GetById(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userId := ctx.Params("id")

	//param validator
	_, err := strconv.Atoi(userId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	//run business logic
	result, err := uc.userService.GetById(c, userId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
}

func (uc UserControllerImpl) Delete(ctx *fiber.Ctx) error {
	c := ctx.Context()

	var user model.Users

	var userUpdate dto.UserUpdate

	userId := ctx.Params("id")

	//body parsing
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// request body validation
	if err := uc.validate.Struct(&userUpdate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	//run business logic
	if err := uc.userService.Delete(c, userId); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "succes deleted",
	})
}
