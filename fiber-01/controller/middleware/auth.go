package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/paung29/utils"
)

func JWTFilter(context *fiber.Ctx) error{

	authHeader :=  context.Get("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return  fiber.NewError(fiber.StatusUnauthorized, "missing or invalid Authorization header")
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	accountID, err := utils.ParseToken(tokenStr)

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid or expired token")
	}

	context.Locals("accountId",  accountID)

	return context.Next()
}