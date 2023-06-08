package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tangguhriyadi/user-service/model"
	"github.com/tangguhriyadi/user-service/service"
)

type AuthController interface {
	Login(ctx *fiber.Ctx) error
}

type AuthControllerImpl struct {
	authService service.AuthService
	validate    *validator.Validate
}

func NewAuthController(authService service.AuthService, validate *validator.Validate) AuthController {
	return &AuthControllerImpl{
		authService: authService,
		validate:    validate,
	}
}

// ShowAccount godoc
// @Summary      Login
// @Description  Login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param payload body model.Login true "The input struct"
// @Success      200  {object}  model.LoginResponse
// @Router       /auth/login [post]
func (ac AuthControllerImpl) Login(ctx *fiber.Ctx) error {
	c := ctx.Context()

	var payload model.Login

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// request body validation
	if err := ac.validate.Struct(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	//run business logic
	result, err := ac.authService.Login(c, &payload)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(result)

}
