package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/paung29/controller/dto"
	"github.com/paung29/service"
	"github.com/paung29/utils"
)

var validate = validator.New()

func CreateAccount(c *fiber.Ctx) error{

	var form dto.CreateAccountForm

	if err := c.BodyParser(&form); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if err := validate.Struct(&form); err != nil {
		errors := utils.TranslateValidationError(err)
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	account, err := service.CreateUser(form)

	if err != nil {
		return  err
	}

	return c.Status(fiber.StatusCreated).JSON(account)

}

func Login(c *fiber.Ctx) error {
	var form dto.LoginForm

	if err := c.BodyParser(&form);
	err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid json format")
	}

	if err := validate.Struct(&form);
	err != nil {
		errors := utils.TranslateValidationError(err)
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	token, err := service.Login(form)
	if err != nil {
		return err
	}

	return c.JSON(dto.LoginResponse {
		Token: token,
	})
	
}