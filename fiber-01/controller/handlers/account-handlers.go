package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paung29/controller/dto"
	"github.com/paung29/service"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateAccount(c *fiber.Ctx) error{

	var form dto.CreateAccountForm

	if err := c.BodyParser(&form); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if err := validate.Struct(&form); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	account, err := service.CreateUser(form)

	if err != nil {
		return  err
	}

	return c.Status(fiber.StatusCreated).JSON(account)

}