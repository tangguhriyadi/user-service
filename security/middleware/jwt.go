package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v2"
)

func JWTProtect() func(*fiber.Ctx) error {
	// envconfig := infrastructure.New()
	config := jwtMiddleware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}
	return jwtMiddleware.New(config)
}

func jwtError(ctx *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unathorized"})
}
