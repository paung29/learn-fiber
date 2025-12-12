package service

import (
	"github.com/paung29/controller/dto"
	"github.com/paung29/database"
	"github.com/paung29/model"
	"github.com/paung29/utils"
)


func CreateUser(form dto.CreateAccountForm) (model.Account, error) {


	account := model.Account {
		Name: form.Name,
		Email: form.Email,
		Password: form.Password,
	}

	hashedPassword, err := utils.HashedPassword(account.Password)

	if err != nil {
		return  model.Account{}, utils.SystemError("hashing password failed")
	}
	
	account.Password = hashedPassword

	if err := database.DB.Create(&account).Error; err != nil {
		return model.Account{}, utils.SystemError("failed to create Account")
	}
	return account, nil
}

func Login(form dto.LoginForm) (string, error) {
	var account model.Account

	if err := database.DB.Where("email = ?", form.Email).First(&account).Error;
	err != nil {
		return "", utils.SystemError("invalid email")
	}

	if err := utils.CheckPassword(account.Password, form.Password);
	err != nil {
		return  "", utils.SystemError("invalid password")
	}

	token, err := utils.GenerateToken(account.ID)

	if err != nil {
		return  "", utils.SystemError("failed to generate token")
	}

	return token, nil
}